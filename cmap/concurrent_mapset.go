package cmap

import (
	"sync"
)

// ConcurrentMapset A "thread" safe mapset of type uint64:uint64.
// To avoid lock bottlenecks this map is dived to several (MapsetShardCount) map shards.
type ConcurrentMapset struct {
	mapsetShardCount uint64
	maxCardinality   uint64
	shard            []*ConcurrentMapsetShared
}

// ConcurrentMapsetShared A "thread" safe uint64 to uint64 map.
type ConcurrentMapsetShared struct {
	items        map[uint64]uint64
	fifoQueue    []uint64
	sequence     uint64 // circular fifoQueue header position
	sync.RWMutex        // Read Write mutex, guards access to internal map.
}

// NewMapset Creates a new concurrent map. 'maxCardinality' is for each shard, the total is 'maxCardinality'*MapsetShardCount.
func NewMapset(shardCount uint64, maxCardinality uint64) ConcurrentMapset {
	shard := make([]*ConcurrentMapsetShared, shardCount)
	for i := uint64(0); i < shardCount; i++ {
		shard[i] = &ConcurrentMapsetShared{items: make(map[uint64]uint64, maxCardinality), fifoQueue: make([]uint64, maxCardinality)}
	}

	return ConcurrentMapset{shardCount, maxCardinality, shard}
}

func (m ConcurrentMapset) IsEmpty() bool {
	if m.shard == nil || len(m.shard) == 0 {
		return true
	} else {
		return false
	}
}

// Set add an element to the set
// - value is updated anyway when set is called;
// - but fifiQueue is only updated when key is new.
// - The oldest Set key will be deleted (kick out) firstly when mapset is full.
func (m ConcurrentMapset) Set(key uint64, value uint64) bool {
	shard := m.GetShard(key)
	shard.Lock()
	_, exist := shard.items[key]

	if uint64(len(shard.items)) >= m.maxCardinality {
		// release / kick out the oldest Set key
		delete(shard.items, shard.fifoQueue[shard.sequence%m.maxCardinality])
	}
	// update value anyway no matter exist or not
	shard.items[key] = value
	if !exist {
		// but only update fifiQueue for new key
		shard.fifoQueue[shard.sequence%m.maxCardinality] = key
		shard.sequence++
	}

	shard.Unlock()
	return true
}

// GetShard returns shard under given key
func (m ConcurrentMapset) GetShard(key uint64) *ConcurrentMapsetShared {
	return m.shard[key%m.mapsetShardCount]
}

// Cardinality returns the number of elements within the map.
func (m ConcurrentMapset) Cardinality() int {
	count := 0
	for i := uint64(0); i < m.mapsetShardCount; i++ {
		shard := m.shard[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has Looks up an item under specified key
func (m ConcurrentMapset) Has(key uint64) bool {
	shard := m.GetShard(key)
	shard.RLock()
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Get Looks up an item under specified key
func (m ConcurrentMapset) Get(key uint64) (uint64, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	ts, ok := shard.items[key]
	shard.RUnlock()
	return ts, ok
}
