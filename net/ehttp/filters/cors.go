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

// Cross Origin Resource Sharing (CORS)

package filters

import (
	"sort"
	"strconv"
	"strings"

	"github.com/SabzCity/go-library/convert"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// CORS is a filter that manage CORS.
// Filter level (Server, Route): Server
// Needed filter: before(StaticRestRouter/), after()
func CORS(ctx *ehttp.Context) {

	if ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Origin) != "" {

		route := ctx.InterfaceData["StaticRestRoute"].(*StaticRestRoute)

		// Check request can be process
		for _, o := range route.CORS.Origins {
			if o == "*" || strings.HasSuffix(o, ctx.Request.Host) {
				goto CheckMethod
			}
		}
		ctx.Error = errors.AccessForbidden
		ctx.ResponseReady = true

	CheckMethod:
		if sort.SearchStrings(route.CORS.Methods, ctx.Request.Method) == len(route.CORS.Methods) {
			if _, ok := route.MethodsFilters[ctx.Request.Method]; ok {
				ctx.Error = errors.AccessForbidden
				ctx.ResponseReady = true
			}
		}

		// Set "AllowMethods".
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.AccessControl.AllowMethods, strings.Join(route.CORS.Methods, ","))

		// Set "AllowOrgins".
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.AccessControl.AllowOrigin,
			strings.Join(route.CORS.Origins, ","))

		// Set "AllowHeaders".
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.AccessControl.AllowHeaders,
			strings.Join(route.CORS.Headers, ","))

		// Set "MaxAge".
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.AccessControl.MaxAge,
			convert.IntToString(route.CORS.MaxAge))

		// Set "AllowCredentials".
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.AccessControl.AllowCredentials,
			strconv.FormatBool(route.CORS.Credentials))

		// Add needed vary header
		ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Vary, ehttputil.RequestStandards.Headers.Origin)

		// Check it is Preflight request or Actual request
		if ctx.Request.Method == "OPTIONS" {
			// Announce to server, response is ready!
			ctx.ResponseReady = true
		}
	}
}

// Cors saves some data to handle CORS by "CORS.Handle()".
type Cors struct {
	Origins     []string
	Headers     []string
	Methods     []string
	MaxAge      int
	Credentials bool
}
