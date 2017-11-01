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

package oauth2

import (
	"sort"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/instagram"
)

// OAuth2Platforms : List of all supported platforms of "oauth2".
var OAuth2Platforms []string

var key []byte

func init() {

	key = []byte{12, 55, 89, 13, 147, 252, 8, 9}

	store := sessions.NewCookieStore(key)
	gothic.Store = store

	goth.UseProviders(
		instagram.New("a377919d5e07466fab7be0153c51ce88", "e0d312b57f894b8b8dc95e9b4e6b92c1", "http://apis.sabz.city/auth/v1/users/ouath2/instagram/"),
	)

	OAuth2Platforms := append(OAuth2Platforms, "instagram")

	sort.Strings(OAuth2Platforms)
}
