package cache

type CachePolicy interface {
	set(key string, val int)
	get(key string) (int, error)
}

type Cache struct {
	policy CachePolicy
}

func NewCache(policy CachePolicy) Cache {
	return Cache{policy: policy}
}

func (c *Cache) set(key string, val int) bool {
	c.policy.set(key, val)
	return true
}

func (c *Cache) get(key string) (int, error) {
	return c.policy.get(key)
}
