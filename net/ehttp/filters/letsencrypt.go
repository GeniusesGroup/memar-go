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

	"github.com/SabzCity/go-library/net/ehttp"
)

// LetsEncrypt :
// Filter level: just Server
// Needed filter: before(), after()
func LetsEncrypt(ctx *ehttp.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, "/.well-known/acme-challenge/") {
		// Get needed certificate to connect a container.
		ctx.Request.URL.Scheme = "http" // Change for TEST
		// Set "DomainMS:Port" instead of requested Host.
		ctx.Request.URL.Host = "v0.domains.sabzcity.containers.sabz.city:50030"
		ctx.Request.Host = "v0.domains.sabzcity.containers.sabz.city:50030"

		// Use default transport.
		ctx.InterfaceData["Proxy.Transport"] = http.DefaultTransport

		ctx.ResponseReady = true
	}
}
