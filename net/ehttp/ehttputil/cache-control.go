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

// Package ehttputil is Extended HTTP package utility
package ehttputil

import (
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/SabzCity/go-library/convert"
)

var cacheControlRegex = regexp.MustCompile(`([a-zA-Z][a-zA-Z_-]*)\s*(?:=(?:"([^"]*)"|([^ \t",;]*)))?`)

// GetCacheControl : Parses the cache directive of "Cache-Control" header.
func GetCacheControl(header http.Header) CacheControl {

	directives := map[string]string{}
	control := CacheControl{}
	matches := cacheControlRegex.FindAllString(header.Get("Cache-Control"), -1)

	for _, match := range matches {
		var key, value string
		key = match

		if index := strings.Index(match, "="); index != -1 {
			key, value = match[:index], match[index+1:]
		}

		directives[strings.ToLower(key)] = strings.TrimSpace(value)
	}

	_, control.Public = directives["public"]
	_, control.NoStore = directives["no-store"]
	_, control.NoTransform = directives["no-transform"]
	_, control.OnlyIfCached = directives["only-if-cached"]
	_, control.Private = directives["private"]
	_, control.NoCache = directives["no-cache"]
	_, control.MustRevalidate = directives["must-revalidate"]
	_, control.ProxyRevalidate = directives["proxy-revalidate"]
	control.MaxAge = timedDirective(directives, "max-age")
	control.SMaxAge = timedDirective(directives, "s-maxage")
	control.MinFresh = timedDirective(directives, "min-fresh")
	control.MaxStale = timedDirective(directives, "max-stale")

	return control
}

// SetCacheControl : Marshal "control" and set it to "header".
func SetCacheControl(header http.Header, control CacheControl) {

	var headerValue []string

	if control.Public {
		headerValue = append(headerValue, "pubilc")
	}
	if control.NoStore {
		headerValue = append(headerValue, "no-store")
	}
	if control.NoTransform {
		headerValue = append(headerValue, "no-transform")
	}
	if control.OnlyIfCached {
		headerValue = append(headerValue, "only-if-cached")
	}
	if control.Private {
		headerValue = append(headerValue, "private")
	}
	if control.NoCache {
		headerValue = append(headerValue, "no-cache")
	}
	if control.MustRevalidate {
		headerValue = append(headerValue, "must-revalidate")
	}
	if control.ProxyRevalidate {
		headerValue = append(headerValue, "proxy-revalidate")
	}
	if control.MaxAge.Seconds() > -1 {
		headerValue = append(headerValue, "max-age="+convert.IntToString(int(control.MaxAge.Seconds())))
	}
	if control.SMaxAge.Seconds() > -1 {
		headerValue = append(headerValue, "s-maxage="+convert.IntToString(int(control.SMaxAge.Seconds())))
	}
	if control.MinFresh.Seconds() > -1 {
		headerValue = append(headerValue, "min-fresh="+convert.IntToString(int(control.MinFresh.Seconds())))
	}
	if control.MaxStale.Seconds() > -1 {
		headerValue = append(headerValue, "max-stale="+convert.IntToString(int(control.MaxStale.Seconds())))
	}
	header.Set(RequestStandards.Headers.CacheControl, strings.Join(headerValue, ", "))
}

func timedDirective(directives map[string]string, key string) time.Duration {

	t, ok := directives[key]
	if !ok {
		return -1
	}

	i, err := convert.StringToInt(t)
	if err != nil {
		return -1
	}

	return time.Duration(i) * time.Second
}

// Returns -1 if the directive wasn't present or if an error happened
// during parsing the value. It returns math.MaxInt64 if it was present but if
// no value was provided. Otherwise, it returns the provided duration.
func getMaxStale(directives map[string]string) time.Duration {

	t, ok := directives["max-stale"]
	if !ok {
		return -1
	}

	if t == "" {
		return math.MaxInt64
	}

	i, err := convert.StringToInt(t)
	if err != nil {
		return -1
	}

	return time.Duration(i) * time.Second
}

// CacheControl : holds "Cache-Control" header directives and gives a few utility methods to quickly deal
// with directives values.
type CacheControl struct {
	Public          bool
	NoStore         bool
	NoTransform     bool
	OnlyIfCached    bool
	Private         bool
	NoCache         bool
	MustRevalidate  bool
	ProxyRevalidate bool
	MaxAge          time.Duration
	SMaxAge         time.Duration
	MinFresh        time.Duration
	MaxStale        time.Duration
}
