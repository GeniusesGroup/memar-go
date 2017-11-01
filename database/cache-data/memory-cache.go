//Copyright 2017 SabzCity
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

// Use binary cache to store data for often access and response to requests

package cachedata

import (
	"time"

	"github.com/SabzCity/go-library/errors"
	"github.com/allegro/bigcache"
)

// NewMemoryCache : Create a new binary cache storage.
func NewMemoryCache(evictionTime time.Duration, maxSize int) *MemoryCache {

	config := bigcache.DefaultConfig(evictionTime)
	config.HardMaxCacheSize = maxSize

	manager, _ := bigcache.NewBigCache(config)

	return &MemoryCache{manager: manager}
}

// MemoryCache : This is use in everywhere that we need public cache storage.
type MemoryCache struct {
	manager *bigcache.BigCache
}

// Set : Add or replace an Item in cache.
func (cache *MemoryCache) Set(key string, val []byte) error {

	if err := cache.manager.Set(key, val); err != nil {
		return errors.SomeThingIsWrong
	}

	return nil
}

// Get : Get value from special key.
func (cache *MemoryCache) Get(key string) ([]byte, error) {

	data, err := cache.manager.Get(key)
	if err != nil {
		return nil, errors.ContentNotExist
	}

	return data, nil
}
