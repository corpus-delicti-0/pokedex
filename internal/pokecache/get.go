package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key] //value stored under the key
	if !ok {                    //whether the key exists or not
		return nil, false
	}

	return entry.val, true

}
