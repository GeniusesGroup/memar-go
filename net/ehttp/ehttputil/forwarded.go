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
	"net/http"
	"strings"
)

// Forwarded : Implement of "Forwarded" header.
type Forwarded struct {
	// The interface where the request came in to the proxy server.
	By string
	// The client that initiated the request and subsequent proxies in a chain of proxies.
	For []string
	// The Host request header field as received by the proxy.
	Host string
	// Indicates which protocol was used to make the request (typically "http" or "https").
	Protocol string
}

// GetForwarded : Parse "Forwarded" header.
func GetForwarded(req *http.Request) Forwarded {

	forwarded := req.Header.Get(RequestStandards.Headers.Forwarded)
	if forwarded == "" {
		return Forwarded{For: []string{req.RemoteAddr},
			Host:     req.Host,
			Protocol: req.URL.Scheme}
	}

	// TODO : Parse "Forwarded" values by regex.
	return Forwarded{For: []string{req.RemoteAddr},
		Host:     req.Host,
		Protocol: req.URL.Scheme}
}

// SetForwarded : Set the "Forwarded" header.
func SetForwarded(header http.Header, forwarded Forwarded) {

	var headerValue []string

	if len(forwarded.For) != 0 {
		var forValue []string

		for _, value := range headerValue {
			forValue = append(forValue, "for="+value)
		}

		headerValue = append(headerValue, strings.Join(forValue, ", "))
	}
	if forwarded.By != "" {
		headerValue = append(headerValue, "by="+forwarded.By)
	}
	if forwarded.Host != "" {
		headerValue = append(headerValue, "host="+forwarded.Host)
	}
	if forwarded.Protocol != "" {
		headerValue = append(headerValue, "proto="+forwarded.Protocol)
	}

	header.Set(RequestStandards.Headers.Forwarded, strings.Join(headerValue, "; "))
}
