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

	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// Proxy will tunneling the request to destination.
// Filter level (Server, Route): just Server
func Proxy(ctx *ehttp.Context) {

	if ctx.InterfaceData["Proxy.Transport"] != nil {
		transport := ctx.InterfaceData["Proxy.Transport"].(*http.Transport)

		res, err := transport.RoundTrip(ctx.Request)
		if err != nil {
			ctx.Error = errors.CanNotConnectToResource
			return
		}

		ctx.Response.Status = res.StatusCode

		ehttputil.CopyHeaders(ctx.Response.Header, res.Header)

		ctx.Response.Body.ReadFrom(res.Body)
	}
}
