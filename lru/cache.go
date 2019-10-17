package lru

// LRU is the cache for any particular type
type LRU struct {
	cache    map[string]Cache
	capacity int64
	size     int64
	head     *Cache
	tail     *Cache
}

// NewLRU is the constructor for the LRU
func NewLRU(capacity int64) *LRU {
	lru := &LRU{
		capacity: capacity,
		cache:    make(map[string]Cache),
		head:     NewCache("", 0),
		tail:     NewCache("", 0),
		size:     0,
	}

	lru.GetHead().SetNext(lru.GetTail())
	lru.GetTail().SetPre(lru.GetHead())
	lru.GetHead().SetPre(nil)
	lru.GetTail().SetNext(nil)

	return lru
}

// GetHead returns the head of the LRU
func (lru *LRU) GetHead() *Cache {
	return lru.head
}

// GetTail returns the tail of the LRU
func (lru *LRU) GetTail() *Cache {
	return lru.tail
}

// AddToHead adds a new node to the head of LRU
func (lru *LRU) AddToHead(cache *Cache) {
	cache.SetNext(lru.GetHead().GetNext())
	cache.GetNext().SetPre(cache)
	cache.SetPre(lru.GetHead())
	lru.GetHead().SetNext(cache)
}

// DeleteCache delets a cached element from the LRU
func DeleteCache(c *Cache) {
	c.GetPre().SetNext(c.GetNext())
	c.GetNext().SetPre(c.GetPre())
}
