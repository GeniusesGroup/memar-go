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

// Developer must set dynamic routes in KVS microservice before use this filter!
// http://apis.sabz.city/kvs/v1/domains/{Domain}/directories/routing-rules/keys/530/

package filters

import (
	"time"

	cachedata "github.com/SabzCity/go-library/database/cache-data"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
)

var rpcRoutingCache = cachedata.NewObjectCache(time.Hour * 24 * 7)

// DynamicRPCRoute : All we need for process a request.
type DynamicRPCRoute struct {
	Type           string // Type can be "Files" || "Container" || "Proxy" || "R301"
	Destination    string // Destination for "Files":"" || "Container": "{NameOfContainer.UserName}" e.g. "html5-apps-engine.sabzcity" || "Proxy" : "Full-URL"
	ServiceName    string
	Authorization  bool
	GeoRestriction struct {
		Whitelist []string
		Blacklist []string
	}
}

// DynamicRPCRouter call as filter to specify Route for request in RPC standard
// Filter level (Server, Route): Server
// Needed filter: before(), after()
func DynamicRPCRouter(ctx *ehttp.Context) {

	// Find route for special Method.
	rpcRoute := getRPCRoute(ctx.Request.Host, ctx.Request.Header.Get("X-Service-Name"), ctx.Request.Header.Get("X-Service-Version"))
	if rpcRoute == nil {
		ctx.Error = errors.HostNotExist
		// Announce to server, response is ready!
		ctx.ResponseReady = true
		return
	}

	// Set for Authorization process.
	ctx.StringData["ServiceName"] = ctx.Request.Header.Get("X-Service-Name")
	ctx.StringData["MethodName"] = ctx.Request.Header.Get("X-Method-Name")
	ctx.StringData["ActiveUserID"] = ctx.Request.Header.Get("X-Active-User-ID")
	ctx.StringData["OwnerUserID"] = ctx.Request.Header.Get("X-Owner-User-ID")
	ctx.StringData["RecordID"] = ctx.Request.Header.Get("X-Record-ID")

	ctx.InterfaceData["DynamicRPCRoute"] = rpcRoute
}

func getRPCRoute(host, serviceName, serviceVersion string) *DynamicRPCRoute {

	DynamicRPCRouterCached, err := rpcRoutingCache.Get(host)
	if err != nil {
		// TODO : Get needed data from KVS ms!
		DynamicRPCRoutes := rpcRouteTest[host] //TEST
		//if err != nil {
		//	return nil, errors.HostNotExist
		//}

		dynamicRoutingCache.Set(host, &DynamicRPCRoutes, -1)

		serviceRoute := DynamicRPCRoutes[serviceName+"/"+serviceVersion]
		return &serviceRoute
	}

	serviceRoute := DynamicRPCRouterCached.(map[string]DynamicRPCRoute)[serviceName+"/"+serviceVersion]

	return &serviceRoute
}
