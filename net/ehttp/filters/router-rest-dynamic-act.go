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
	"net/url"
	"strings"

	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// DynamicRestRouteAct : Do what dynamic route type want! Use DynamicRestRoute.Type & Destionion to proccess request.
// Filter level: just Server
// Needed filter: before(DynamicRestRouter), after()
func DynamicRestRouteAct(ctx *ehttp.Context) {

	route := ctx.InterfaceData["DynamicRestRoute"].(*DynamicRestRoute)

	// Process the client request with special parameters.
	switch route.Type {
	case "Container":

		// Get needed certificate to connect a container.
		ctx.Request.URL.Scheme = "http" ///Change for TEST

		// Set "DNS:Port" instead of requested Host. e.g. auth.sabzcity.containers.sabz.city:50000
		ctx.Request.URL.Host = route.Destination
		ctx.Request.Host = route.Destination

		// Use default transport.
		ctx.InterfaceData["Proxy.Transport"] = http.DefaultTransport
	case "Proxy":
		dstURL, _ := url.Parse(route.Destination)

		ctx.Request.URL.Scheme = dstURL.Scheme
		ctx.Request.URL.Host = dstURL.Host
		ctx.Request.URL.Path = strings.TrimRight(dstURL.Path, "/") + "/" + strings.TrimLeft(ctx.Request.URL.Path, "/")

		ctx.Request.Host = dstURL.Host

		// Use default transport.
		ctx.InterfaceData["Proxy.Transport"] = http.DefaultTransport
	case "Redirect":
		ctx.Response.Status = http.StatusPermanentRedirect

		dstURL, _ := url.Parse(route.Destination)

		ctx.Request.URL.Scheme = dstURL.Scheme
		ctx.Request.URL.Host = dstURL.Host
		ctx.Request.URL.Path = strings.TrimRight(dstURL.Path, "/") + "/" + strings.TrimLeft(ctx.Request.URL.Path, "/")

		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Location, ctx.Request.URL.String())
	case "File":
	}
}
