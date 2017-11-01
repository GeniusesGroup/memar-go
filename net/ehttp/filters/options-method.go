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
	"strings"

	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// OPTIONS is a filter function that inspects the HTTP Request for the OPTIONS method when Origin header not exist
// and provides the response with a set of allowed methods for the request URL Path.
// Filter level: Server
// Needed filter: before(StaticRestRouter/), after()
func OPTIONS(ctx *ehttp.Context) {

	if ctx.Request.Method == "OPTIONS" && ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Origin) == "" {
		var allowMethods []string
		staticRestRoute := ctx.InterfaceData["StaticRestRoute"].(*StaticRestRoute)
		for m := range staticRestRoute.MethodsFilters {
			allowMethods = append(allowMethods, m)
		}

		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Allow, strings.Join(allowMethods, ","))

		// Announce to server, response is ready!
		ctx.ResponseReady = true
	}
}
