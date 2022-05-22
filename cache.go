package cache

import (
	"time"
)

type Time struct {
	IsExpired bool
	Date      time.Time
}

type Item struct {
	Value        string
	ExpireAtTime Time
}

type Cache struct {
	CacheItems map[string]Item
}

func NewCache() Cache {
	i := make(map[string]Item)
	return Cache{CacheItems: i}
}

func (cache Cache) Get(key string) (string, bool) {
	cache.CheckExpiredTime(key)
	value, ok := cache.CacheItems[key]
	return value.Value, ok
}

func (cache Cache) CheckExpiredTime(key string) {
	now := time.Now()
	if cache.CacheItems[key].ExpireAtTime.IsExpired && !cache.CacheItems[key].ExpireAtTime.Date.After(now) {
		cache.Delete(key)
	}
}

func (cache Cache) Delete(key string) {
	delete(cache.CacheItems, key)
}

func (cache Cache) Put(key, value string) {
	cache.CacheItems[key] = Item{Value: value, ExpireAtTime: Time{IsExpired: false}}
}

func (cache Cache) Keys() []string {
	now := time.Now()
	var keys []string
	for key, value := range cache.CacheItems {
		if !value.ExpireAtTime.IsExpired || value.ExpireAtTime.Date.After(now) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
	now := time.Now()
	if deadline.After(now) {
		cache.CacheItems[key] = Item{Value: value, ExpireAtTime: Time{IsExpired: true, Date: deadline}}
	}
}
