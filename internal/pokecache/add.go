package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) { //*Cache means the methods operate on existing cache
	c.mu.Lock()         //prevents other goroutines from accessing the map while it changes
	defer c.mu.Unlock() //promises to unlock when Add() finishes

	c.entries[key] = cacheEntry{ //key associated with a new entry
		createdAt: time.Now(), //record when the value entered cache
		val:       val,
	}
}
