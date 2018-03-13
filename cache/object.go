//Copyright 2018 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Use object cache to store golang data ready to use

package cache

import (
	"sync"
	"time"

	"github.com/SabzCity/go-library/errors"
)

// NewObjectCache : Create a new local cache storage.
func NewObjectCache(cleanDuration time.Duration) *ObjectCache {
	cache := &ObjectCache{
		items: map[string]cacheItem{}}

	if cleanDuration > 0 {
		go func() {
			for {
				select {
				case <-time.Tick(cleanDuration):
					cache.mutex.Lock()

					for key, item := range cache.items {
						if time.Now().UnixNano() > item.timeToLive {
							delete(cache.items, key)
						}
					}

					cache.mutex.Unlock()
				}
			}
		}()
	}

	return cache
}

// ObjectCache : This is use in everywhere that we need local cache storage.
type ObjectCache struct {
	mutex sync.RWMutex
	items map[string]cacheItem
}

// Set : Add or replace an Item in cache.
func (cache *ObjectCache) Set(key string, value interface{}, TTL time.Duration) {
	cache.mutex.Lock()

	cache.items[key] = cacheItem{
		object:     value,
		timeToLive: time.Now().Add(TTL).UnixNano()}

	cache.mutex.Unlock()
}

// Get : Get value from special key.
func (cache *ObjectCache) Get(key string) (interface{}, error) {
	cache.mutex.RLock()

	item, ok := cache.items[key]
	if !ok {
		cache.mutex.RUnlock()

		return nil, errors.ContentNotExist
	}

	cache.mutex.RUnlock()

	if time.Now().UnixNano() > item.timeToLive {
		cache.mutex.Lock()

		delete(cache.items, key)

		cache.mutex.Unlock()

		return nil, errors.ContentNotExist
	}

	return item.object, nil
}

type cacheItem struct {
	object     interface{}
	timeToLive int64
}
