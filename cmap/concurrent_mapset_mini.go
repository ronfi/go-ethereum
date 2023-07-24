package cmap

import (
	"sync"
)

// ConcurrentMapsetMini A "thread" safe mapset of type uint64:struct{}.
// To avoid lock bottlenecks this map is dived to several (MapsetMiniShardCount) map shards.
type ConcurrentMapsetMini struct {
	MapsetMiniShardCount uint64
	maxCardinality       uint64
	shard                []*ConcurrentMapsetMiniShared
}

// ConcurrentMapsetMiniShared A "thread" safe uint64 to uint64 map.
type ConcurrentMapsetMiniShared struct {
	items        map[uint64]struct{}
	fifoQueue    []uint64
	sequence     uint64 // circular fifoQueue header position
	sync.RWMutex        // Read Write mutex, guards access to internal map.
}

// NewMapsetMini Creates a new concurrent map. 'maxCardinality' is for each shard, the total is 'maxCardinality'*MapsetMiniShardCount.
func NewMapsetMini(shardCount uint64, maxCardinality uint64) ConcurrentMapsetMini {
	shard := make([]*ConcurrentMapsetMiniShared, shardCount)
	for i := uint64(0); i < shardCount; i++ {
		shard[i] = &ConcurrentMapsetMiniShared{items: make(map[uint64]struct{}, maxCardinality), fifoQueue: make([]uint64, maxCardinality)}
	}

	return ConcurrentMapsetMini{shardCount, maxCardinality, shard}
}

// Add an element to the set. Returns 'true' means add success, 'false' means already exist.
func (m ConcurrentMapsetMini) Add(key uint64) bool {
	shard := m.GetShard(key)
	shard.Lock()
	if _, exist := shard.items[key]; exist {
		shard.Unlock()
		return false
	}

	if uint64(len(shard.items)) >= m.maxCardinality { // only check shard length for efficiency
		delete(shard.items, shard.fifoQueue[shard.sequence%m.maxCardinality]) // remove an oldest item
	}

	shard.items[key] = struct{}{}
	shard.fifoQueue[shard.sequence%m.maxCardinality] = key
	shard.sequence++
	shard.Unlock()
	return true
}

// GetShard returns shard under given key
func (m ConcurrentMapsetMini) GetShard(key uint64) *ConcurrentMapsetMiniShared {
	return m.shard[key%m.MapsetMiniShardCount]
}

// Cardinality returns the number of elements within the map.
func (m ConcurrentMapsetMini) Cardinality() int {
	count := 0
	for i := uint64(0); i < m.MapsetMiniShardCount; i++ {
		shard := m.shard[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// GetAll returns all elements within the map.
func (m ConcurrentMapsetMini) GetAll() (all []uint64) {
	count := m.Cardinality()
	all = make([]uint64, 0, count)
	for i := uint64(0); i < m.MapsetMiniShardCount; i++ {
		shard := m.shard[i]
		shard.RLock()
		for k := range shard.items {
			all = append(all, k)
		}
		shard.RUnlock()
	}
	return
}

// Has Looks up an item under specified key
func (m ConcurrentMapsetMini) Has(key uint64) bool {
	shard := m.GetShard(key)
	shard.RLock()
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}
