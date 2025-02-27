// from https://github.com/libp2p/go-libp2p-rendezvous/blob/implement-spec/discovery.go
// @NOTE(gfanton): we need to dedup this file to add unregister method

package tinder

import (
	"context"
	"math"
	mrand "math/rand"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	p2p_rp "github.com/libp2p/go-libp2p-rendezvous"
	"go.uber.org/zap"
)

type rendezvousDiscovery struct {
	logger       *zap.Logger
	rp           p2p_rp.RendezvousPoint
	peerCache    map[string]*rpCache
	peerCacheMux sync.RWMutex
	rng          *mrand.Rand
	rngMux       sync.Mutex
	selfID       peer.ID
}

type rpCache struct {
	recs   map[peer.ID]*rpRecord
	cookie []byte
	mux    sync.Mutex
}

type rpRecord struct {
	peer   peer.AddrInfo
	expire int64
}

func NewRendezvousDiscovery(logger *zap.Logger, host host.Host, rdvPeer peer.ID, rng *mrand.Rand) UnregisterDiscovery {
	rp := p2p_rp.NewRendezvousPoint(host, rdvPeer)
	return &rendezvousDiscovery{
		logger:    logger.Named("tinder/rdvp"),
		rp:        rp,
		rng:       rng,
		peerCache: make(map[string]*rpCache),
		selfID:    host.ID(),
	}
}

func (c *rendezvousDiscovery) Advertise(ctx context.Context, ns string, opts ...discovery.Option) (time.Duration, error) {
	// Get options
	var options discovery.Options
	err := options.Apply(opts...)
	if err != nil {
		return 0, err
	}

	ttl := options.Ttl
	var ttlSeconds int

	if ttl == 0 {
		ttlSeconds = 7200
	} else {
		ttlSeconds = int(math.Round(ttl.Seconds()))
	}

	rttl, err := c.rp.Register(ctx, ns, ttlSeconds)
	if err != nil {
		return 0, err
	}

	return rttl, nil
}

func (c *rendezvousDiscovery) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	// Get options
	var options discovery.Options
	err := options.Apply(opts...)
	if err != nil {
		return nil, err
	}

	const maxLimit = 1000
	limit := options.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}

	/*
		FIXME: enable async behavior

		if enabled, ok := options.Other[optionRendezvousUseDiscoverAsync]; ok {
			enabled, ok := enabled.(bool)
			if ok && enabled {
				return c.asyncFindPeers(ctx, ns, limit)
			}
		}
	*/

	return c.syncFindPeers(ctx, ns, limit)
}

func (c *rendezvousDiscovery) syncFindPeers(ctx context.Context, ns string, limit int) (<-chan peer.AddrInfo, error) {
	// Get cached peers
	var cache *rpCache
	var err error

	c.peerCacheMux.RLock()
	cache, ok := c.peerCache[ns]
	c.peerCacheMux.RUnlock()
	if !ok {
		c.peerCacheMux.Lock()
		cache, ok = c.peerCache[ns]
		if !ok {
			cache = &rpCache{recs: make(map[peer.ID]*rpRecord)}
			c.peerCache[ns] = cache
		}
		c.peerCacheMux.Unlock()
	}

	cache.mux.Lock()
	defer cache.mux.Unlock()

	// Remove all expired entries from cache
	currentTime := time.Now().Unix()
	newCacheSize := len(cache.recs)

	for p := range cache.recs {
		rec := cache.recs[p]
		if rec.expire < currentTime {
			newCacheSize--
			delete(cache.recs, p)
		}
	}

	cookie := cache.cookie

	// Discover new records if we don't have enough
	if newCacheSize < limit {
		// TODO: Should we return error even if we have valid cached results?
		var regs []p2p_rp.Registration
		var newCookie []byte

		if regs, newCookie, err = c.rp.Discover(ctx, ns, limit, cookie); err == nil {
			for _, reg := range regs {
				rec := &rpRecord{peer: reg.Peer, expire: int64(reg.Ttl) + currentTime}
				cache.recs[rec.peer.ID] = rec
			}
			cache.cookie = newCookie
		}
	}

	// Randomize and fill channel with available records
	count := len(cache.recs)
	if limit < count {
		count = limit
	}

	// c.logger.Debug("found peers", logutil.PrivateString("key", ns), zap.Int("count", count))

	// fmt.Printf("rdvp findpeers found [%s]: %d peers found\n", ns, count)
	chPeer := make(chan peer.AddrInfo, count)

	c.rngMux.Lock()
	perm := c.rng.Perm(len(cache.recs))[0:count]
	c.rngMux.Unlock()

	permSet := make(map[int]int)
	for i, v := range perm {
		permSet[v] = i
	}

	sendLst := make([]*peer.AddrInfo, count)
	iter := 0
	for k := range cache.recs {
		if sendIndex, ok := permSet[iter]; ok {
			sendLst[sendIndex] = &cache.recs[k].peer
		}
		iter++
	}

	for _, send := range sendLst {
		chPeer <- *send
	}

	close(chPeer)
	return chPeer, err
}

// nolint:unused
func (c *rendezvousDiscovery) asyncFindPeers(ctx context.Context, ns string, limit int) (<-chan peer.AddrInfo, error) {
	// Get cached peers
	var cache *rpCache

	c.peerCacheMux.RLock()
	cache, ok := c.peerCache[ns]
	c.peerCacheMux.RUnlock()
	if !ok {
		c.peerCacheMux.Lock()
		cache, ok = c.peerCache[ns]
		if !ok {
			cache = &rpCache{recs: make(map[peer.ID]*rpRecord)}
			c.peerCache[ns] = cache
		}
		c.peerCacheMux.Unlock()
	}

	cache.mux.Lock()
	// Remove all expired entries from cache
	currentTime := time.Now().Unix()
	newCacheSize := len(cache.recs)

	for p := range cache.recs {
		rec := cache.recs[p]
		if rec.expire < currentTime {
			newCacheSize--
			delete(cache.recs, p)
		}
	}

	count := len(cache.recs)
	if limit < count {
		count = limit
	}

	ctx, cancel := context.WithCancel(ctx)

	ch, err := c.rp.DiscoverAsync(ctx, ns)
	if err != nil {
		cancel()
		cache.mux.Unlock()
		return nil, err
	}

	chPeer := make(chan peer.AddrInfo, count)

	c.rngMux.Lock()
	perm := c.rng.Perm(len(cache.recs))[0:count]
	c.rngMux.Unlock()

	permSet := make(map[int]int)
	for i, v := range perm {
		permSet[v] = i
	}

	sendLst := make([]peer.ID, count)
	iter := 0
	for k := range cache.recs {
		if sendIndex, ok := permSet[iter]; ok {
			sendLst[sendIndex] = k
		}
		iter++
	}

	for _, pid := range sendLst {
		chPeer <- cache.recs[pid].peer
	}

	go func() {
		defer cancel()
		defer cache.mux.Unlock()

		for reg := range ch {
			currentTime := time.Now().Unix()

			rec := &rpRecord{peer: reg.Peer, expire: int64(reg.Ttl) + currentTime}
			_, alreadySent := cache.recs[rec.peer.ID]
			cache.recs[rec.peer.ID] = rec

			if !alreadySent {
				chPeer <- rec.peer
			}

			limit--

			if limit == 0 {
				cancel()
			}
		}
	}()

	return chPeer, nil
}

func (c *rendezvousDiscovery) Unregister(ctx context.Context, ns string) error {
	c.peerCacheMux.RLock()
	cache, ok := c.peerCache[ns]
	c.peerCacheMux.RUnlock()
	if ok {
		cache.mux.Lock()
		delete(cache.recs, c.selfID)
		cache.mux.Unlock()
	}
	return c.rp.Unregister(ctx, ns)
}

// nolint:gochecknoinits
func init() {
	p2p_rp.DiscoverAsyncInterval = time.Second * 10
}
