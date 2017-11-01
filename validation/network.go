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

package validation

import (
	"net/http"
	"sort"
	"strings"
)

// CheckCookie : Check size of cookies.
func CheckCookie(cookie *http.Cookie) bool {

	if strings.Count(cookie.String(), "")-1 > 4096 {
		return false
	}

	return true
}

// IsValidOAuth2Platform : Check the name of oauth2 platfom is supported by sabzcity.
func IsValidOAuth2Platform(in interface{}) bool {

	if str, ok := in.(string); ok {
		if sort.SearchStrings(validOAuthPlatforms, str) != len(validOAuthPlatforms) {
			return true
		}
	}

	return false
}
