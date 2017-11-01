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

// Package ednsutil is Extended DNS package utility
package ednsutil

import "strings"

// NormalizeDNS : Get DNS and origin and normalize a DNS.
func NormalizeDNS(DNS, origin string) string {
	normalized := strings.Replace(DNS, "@", origin, -1)

	if normalized != DNS {
		return normalized
	}

	if DNS[strings.Count(DNS, "")-1] != '.' {
		return DNS + origin
	}

	return DNS
}
