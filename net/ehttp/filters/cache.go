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

package filters

import (
	"net/http"
	"strings"
	"time"

	"github.com/vmihailenco/msgpack"

	persiadb "github.com/SabzCity/PersiaDB"
	"github.com/SabzCity/go-library/convert"
	"github.com/SabzCity/go-library/cryptography/md5"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// NewHTTPCache : Create new "HTTPCache" object.
func NewHTTPCache() *HTTPCache {
	return &HTTPCache{
		value: persiadb.NewMemoryCache(time.Hour*24*365, 512),
		vary:  persiadb.NewMemoryCache(time.Hour*24*365*10, 128)}
}

// HTTPCache : Manage HTTP request and response cache.
type HTTPCache struct {
	value *persiadb.MemoryCache
	vary  *persiadb.MemoryCache
}

// InputCache : This is a filter that check if requested data is in the cache
// then we can send response from cache.
// Filter level (Server, Route): just Server
//// TODO : Handle :
//// "Cache-Control: max-age=<seconds>"
//// "Cache-Control: max-stale[=<seconds>]"
//// "Cache-Control: min-fresh=<seconds>"
func (cache *HTTPCache) InputCache(ctx *ehttp.Context) {

	// First we need check last layer error.
	// Or this "Cache" flag is on.
	if ctx.Error != nil {
		return
	}

	control := ehttputil.GetCacheControl(ctx.Request.Header)

	if control.NoStore || ctx.Request.Method != ehttputil.RequestStandards.Methods.GET || // HTTP methods (POST, PUT, DELETE, and PATCH) are not cached and simply proxied to the origin
		// Conditional requests should send to the origin server.
		ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.If.NoneMatch) != "" ||
		ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.If.Match) != "" ||
		ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.If.ModifiedSince) != "" ||
		ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.If.UnmodifiedSince) != "" ||
		ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.If.Range) != "" {
		// We do the request without any cache.

		return
	}

	cachedResponse, err := cache.Get(cache.CreateCacheKey(ctx.Request.URL.String(), ctx.Request.Header))
	if err != nil {
		if control.OnlyIfCached {
			ctx.Error = errors.ItemNotExistInCache
		}

		return
	}

	if control.OnlyIfCached {
		cachedResponse.CopyToResponse(ctx.Response)
		ctx.ResponseReady = true
		return
	}

	// Cache item is expired or we have to revalidate it before using.
	if cachedResponse.NeedRevalidate() || control.NoCache {
		// Revalidate with "ETag" way.
		if eTag := cachedResponse.Header.Get(ehttputil.ResponseStandards.Headers.ETag); eTag != "" {
			ctx.Request.Header.Set(ehttputil.RequestStandards.Headers.If.NoneMatch, eTag)
		}
		// Revalidate with "LastModified" way.
		if lastModified := cachedResponse.Header.Get(ehttputil.ResponseStandards.Headers.LastModified); lastModified != "" {
			ctx.Request.Header.Set(ehttputil.RequestStandards.Headers.If.ModifiedSince, lastModified)
		}

		ctx.InterfaceData["CachedResponse"] = cachedResponse

		return
	}

	cachedResponse.CopyToResponse(ctx.Response)

	// Set Age header.
	ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Age, convert.IntToString(cachedResponse.Age()))

	ctx.ResponseReady = true
}

// OutputCache : This is a filter that save the response data in cache.
// Filter level (Server, Route): just Server
func (cache *HTTPCache) OutputCache(ctx *ehttp.Context) {

	// First we need check last layer error.
	// Or this "Cache" flag is on.
	if ctx.Error != nil {
		return
	}

	control := ehttputil.GetCacheControl(ctx.Response.Header)

	// TODO : Handle none caching states.
	// HTTP methods (POST, PUT, DELETE, and PATCH) are not cached.
	if ctx.Request.Method != ehttputil.RequestStandards.Methods.GET ||
		ctx.Response.Status == http.StatusPartialContent ||
		control.Private || control.NoStore {

		return
	}

	// The cache value not modified.
	if ctx.Response.Status == http.StatusNotModified {
		if value, ok := ctx.InterfaceData["CachedResponse"]; ok {
			cachedResponse := value.(CachedResponse)

			// Reset cache.
			ehttputil.ReplaceHeaders(cachedResponse.Header, ctx.Response.Header)
			cache.Set(cache.CreateCacheKey(ctx.Request.URL.String(), ctx.Request.Header), cachedResponse)

			cachedResponse.CopyToResponse(ctx.Response)

			// Set Age header.
			ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Age, convert.IntToString(cachedResponse.Age()))

			ctx.ResponseReady = true
		}

		return
	}

	// Save "Vary" header in cache.
	cache.SaveVary(ctx.Request.URL.String(), ctx.Response.Header)

	// Save response in cache.
	cache.Set(cache.CreateCacheKey(ctx.Request.URL.String(), ctx.Request.Header),
		CachedResponse{
			Status: ctx.Response.Status,
			Header: ctx.Response.Header,
			Body:   ctx.Response.Body.Bytes()})
}

// Get caches responses for special key.
// e.g. hashed value of URL.
func (cache *HTTPCache) Get(key string) (CachedResponse, error) {

	value, err := cache.value.Get(key)
	if err != nil {
		return CachedResponse{}, err
	}

	res := unmarshalResponse(value)

	// Prepare "Cache-Control" header.
	res.control = ehttputil.GetCacheControl(res.Header)
	res.age = int(time.Duration(time.Now().Unix()) - res.StartTime)

	return res, nil
}

// Set the response in cache.
func (cache *HTTPCache) Set(key string, res CachedResponse) error {

	control := ehttputil.GetCacheControl(res.Header)

	// TODO : Handle "Age" header if exist.
	res.StartTime = time.Duration(time.Now().Unix())
	res.ExpireTime = res.StartTime

	if control.MaxAge > 0 {
		res.ExpireTime = res.StartTime + control.MaxAge
	} else if control.MaxAge == -1 {
		// Get expire time from "Expires" header
		expTime, err := http.ParseTime(res.Header.Get(ehttputil.ResponseStandards.Headers.Expires))
		if err == nil {
			res.ExpireTime = time.Duration(expTime.Unix())
		}
	}

	err := cache.value.Set(key, marshalResponse(res))
	if err != nil {
		return err
	}

	return nil
}

func marshalResponse(res CachedResponse) []byte {

	data, _ := msgpack.Marshal(res)

	return data
}

func unmarshalResponse(data []byte) CachedResponse {

	var res CachedResponse

	msgpack.Unmarshal(data, &res)

	return res
}

// CreateCacheKey : Create a hashed key by URL and vary headers.
func (cache *HTTPCache) CreateCacheKey(requestURL string, header http.Header) string {

	data := requestURL

	cacheExist, err := cache.vary.Get(requestURL)
	if err != nil {
		return md5.SumString(data)
	}

	vary := strings.Split(string(cacheExist), ",")
	for _, value := range vary {
		data += strings.TrimSpace(value)
	}

	return md5.SumString(data)
}

// SaveVary Add "Vary" value by URL in cache.
func (cache *HTTPCache) SaveVary(requestURL string, header http.Header) error {

	vary := header.Get(ehttputil.ResponseStandards.Headers.Vary)

	if err := cache.vary.Set(requestURL, []byte(vary)); err != nil {
		return err
	}

	return nil
}

// CachedResponse : This struct is a cache storage for each shared response.
// NOTE : "msgpack" tag is written by initial words because it has direct effect on performance.
type CachedResponse struct {
	age        int
	Status     int           `msgpack:"S"`
	Header     http.Header   `msgpack:"H"`
	Body       []byte        `msgpack:"B"`
	ExpireTime time.Duration `msgpack:"E"`
	StartTime  time.Duration `msgpack:"T"`
	control    ehttputil.CacheControl
}

// CopyToResponse copy cached data to "ehttp.Response".
func (res CachedResponse) CopyToResponse(response *ehttp.Response) {

	response.Status = res.Status

	ehttputil.CopyHeaders(response.Header, res.Header)

	response.Body.Write(res.Body)
}

// CacheControl : Returns the "res.control".
func (res CachedResponse) CacheControl() ehttputil.CacheControl {

	return res.control
}

// Expired : Check the expiration time of cache.
func (res CachedResponse) Expired() bool {

	return time.Now().Unix() >= int64(res.ExpireTime)
}

// NeedRevalidate : Check validation state of cache.
func (res CachedResponse) NeedRevalidate() bool {

	return res.Expired() || res.control.NoCache || res.control.MustRevalidate || res.control.ProxyRevalidate
}

// Age : Returns remaining time of cache.
func (res CachedResponse) Age() int {

	return res.age
}
