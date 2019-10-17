package lru

// Cache is a node in the LRU Caching mechanism
type Cache struct {
	key   string
	value interface{}
	pre   *Cache
	next  *Cache
}

// NewCache is a contructor function for the LRU
func NewCache(key string, value interface{}) *Cache {
	return &Cache{
		key:   key,
		value: value,
	}
}

func (c *Cache) SetNext(n *Cache) {
	c.next = n
}

func (c *Cache) SetPre(p *Cache) {
	c.pre = p
}

func (c *Cache) GetNext() *Cache {
	return c.next
}

func (c *Cache) GetPre() *Cache {
	return c.pre
}
