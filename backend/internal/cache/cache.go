package cache

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

type LRUCache struct {
	Cache      map[string]CacheItem
	Keys       []string
	MaxSize    int
	mutex      sync.Mutex
	db         *sql.DB
}

func NewLRUCache(maxSize int, db *sql.DB) *LRUCache {
	return &LRUCache{
		Cache:   make(map[string]CacheItem),
		Keys:    make([]string, 0, maxSize),
		MaxSize: maxSize,
		mutex:   sync.Mutex{},
		db:      db,
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, ok := c.Cache[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(item.Expiration) {
		delete(c.Cache, key)
		c.removeFromDB(key)
		return nil, false
	}

	return item.Value, true
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.Keys) >= c.MaxSize {
		delete(c.Cache, c.Keys[0])
		c.removeFromDB(c.Keys[0])
		c.Keys = c.Keys[1:]
	}

	c.Keys = append(c.Keys, key)
	c.Cache[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}

	c.addToDB(key, value, expiration)
}

func (c *LRUCache) addToDB(key string, value interface{}, expiration time.Time) {
	_, err := c.db.Exec("INSERT INTO cache (key, value, expiration) VALUES ($1, $2, $3)", key, value, expiration)
	if err != nil {
		log.Printf("Error adding item to DB: %v", err)
	}
}

func (c *LRUCache) removeFromDB(key string) {
	_, err := c.db.Exec("DELETE FROM cache WHERE key = $1", key)
	if err != nil {
		log.Printf("Error removing item from DB: %v", err)
	}
}

func (c *LRUCache) InitDB() error {
	_, err := c.db.Exec(`CREATE TABLE IF NOT EXISTS cache (
		key TEXT PRIMARY KEY,
		value JSONB,
		expiration TIMESTAMP
	)`)
	if err != nil {
		return err
	}
	return nil
}

func (c *LRUCache) CleanupExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, item := range c.Cache {
		if time.Now().After(item.Expiration) {
			delete(c.Cache, key)
			c.removeFromDB(key)
		}
	}
}

func (c *LRUCache) StartCleanupTask(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			c.CleanupExpired()
		}
	}()
}


func (c *LRUCache) CloseDB() error {
	return c.db.Close()
}

func (c *LRUCache) HandleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key not provided", http.StatusBadRequest)
			return
		}

		value, ok := c.Get(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(value)
	}
}

func (c *LRUCache) HandleSet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		c.Set(data.Key, data.Value, time.Now().Add(5*time.Second))
		w.WriteHeader(http.StatusCreated)
	}
}
