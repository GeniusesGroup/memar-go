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
	"github.com/SabzCity/go-library/authorization"
	"github.com/SabzCity/go-library/net/ehttp"
)

// Authorization check user access.
// Filter level (Server, Route): both
// Needed filter: before(CookieAuthentication||AuthorizationHeader && DynamicRestRouter||StaticRestRouter||DynamicRPCRouter), after()
func Authorization(ctx *ehttp.Context) {

	// Check DynamicRestRouter want authorization
	restRoute, ok := ctx.InterfaceData["DynamicRestRoute"].(*DynamicRestRoute)
	if ok {
		if !restRoute.Authorization {
			return
		}
	}

	// Check DynamicRPCRouter want authorization
	rpcRoute, ok := ctx.InterfaceData["DynamicRPCRoute"].(*DynamicRPCRoute)
	if ok {
		if !rpcRoute.Authorization {
			return
		}
	}

	// if this filter add in StaticRestRouter means developer want this filter!

	// Check method name be available, otherwise filter can't process!
	if ctx.StringData["MethodName"] == "" {
		return
	}

	request := authorization.Request{
		RequestedUserID: ctx.StringData["ActiveUserID"],
		OwnerUserID:     ctx.StringData["OwnerUserID"],
		RecordID:        ctx.StringData["RecordID"],
		MethodName:      ctx.StringData["ServiceName"] + "." + ctx.StringData["MethodName"]}

	ctx.Error = authorization.Authorization(request)
	if ctx.Error != nil {
		ctx.ResponseReady = true
	}
}
