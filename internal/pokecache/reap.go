package pokecache

import "time"

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) //sends a signal every time the interval passes
	defer ticker.Stop()

	for range ticker.C { //waits for each signal and runs the cleanup
		c.mu.Lock()

		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval { //asks whether the entry is older than the allowed cache lifetime
				delete(c.entries, key) //removed expired entry from the map
			}
		}
		c.mu.Unlock()
	}
}
