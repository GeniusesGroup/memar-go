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
// http://apis.sabz.city/kvs/v1/domains/{Domain}/directories/routing-rules/keys/443/

package filters

import (
	"strings"
	"time"

	cachedata "github.com/SabzCity/go-library/database/cache-data"
	"github.com/SabzCity/go-library/documentation/openapi"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
)

var dynamicRoutingCache = cachedata.NewObjectCache(time.Hour * 24 * 7)

// DynamicRestRoute : All we need for process a request.
type DynamicRestRoute struct {
	Type           string           // Type can be "Files" || "Container" || "Proxy" || "R301"
	Destination    string           // Destination for "Files":"" || "Container": "{NameOfContainer.UserName}" e.g. "html5-apps-engine.sabzcity" || "Proxy" : "Full-URL"
	Path           string           // Full Path URL. Can be regex and must start with % e.g. %{RequestURI} !.*\.(svg|jpg|png|css|js|html|php)
	PathParts      []string         // Auto vlaue and cached for dispatching
	PathItem       openapi.PathItem //
	Authorization  bool
	GeoRestriction struct {
		Whitelist []string
		Blacklist []string
	}
}

// DynamicRestRouter call as filter to specify DynamicRestRoute for request.
// Filter level (Server, Route): Server
// Needed filter: before(), after()
func DynamicRestRouter(ctx *ehttp.Context) {
	requestURLParts := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")

	// Find RouteTree for requested host.
	DynamicRestRoutesTree := getDynamicRestRoutesTree(ctx.Request.Host)
	if DynamicRestRoutesTree == nil {
		ctx.Error = errors.HostNotExist
		ctx.ResponseReady = true
		return
	}

	// Find route for special URL.
	DynamicRestRoute := DynamicRestRoutesTree.get(ctx.Request.URL.Path)
	if DynamicRestRoute == nil {
		ctx.Error = errors.URLIsWrong
		ctx.ResponseReady = true
		return
	}

	if DynamicRestRoute.Type == "Container" {
		for i, part := range DynamicRestRoute.PathParts {
			// Check the URL parameter.
			if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				// parameterName without enclosing {}
				parameterName := part[1 : len(part)-1]
				ctx.StringData[parameterName] = requestURLParts[i]
				ctx.Request.Header.Set("X-"+parameterName, requestURLParts[i])
			}
		}

		switch ctx.Request.Method {
		case "GET":
			if DynamicRestRoute.PathItem.GET != nil {
				ctx.StringData["MethodName"] = DynamicRestRoute.PathItem.GET.OperationID
				ctx.Request.Header.Set("X-Method-Name", DynamicRestRoute.PathItem.GET.OperationID)
			}
		case "PUT":
			if DynamicRestRoute.PathItem.PUT != nil {
				ctx.StringData["MethodName"] = DynamicRestRoute.PathItem.PUT.OperationID
				ctx.Request.Header.Set("X-Method-Name", DynamicRestRoute.PathItem.PUT.OperationID)
			}
		case "POST":
			if DynamicRestRoute.PathItem.POST != nil {
				ctx.StringData["MethodName"] = DynamicRestRoute.PathItem.POST.OperationID
				ctx.Request.Header.Set("X-Method-Name", DynamicRestRoute.PathItem.POST.OperationID)
			}
		case "DELETE":
			if DynamicRestRoute.PathItem.DELETE != nil {
				ctx.StringData["MethodName"] = DynamicRestRoute.PathItem.DELETE.OperationID
				ctx.Request.Header.Set("X-Method-Name", DynamicRestRoute.PathItem.DELETE.OperationID)
			}
		case "PATCH":
			if DynamicRestRoute.PathItem.PATCH != nil {
				ctx.StringData["MethodName"] = DynamicRestRoute.PathItem.PATCH.OperationID
				ctx.Request.Header.Set("X-Method-Name", DynamicRestRoute.PathItem.PATCH.OperationID)
			}
		}

		ctx.StringData["ServiceName"] = requestURLParts[0]           //// ??
		ctx.Request.Header.Set("X-Service-Name", requestURLParts[0]) //// ??
	}

	ctx.InterfaceData["DynamicRestRoute"] = DynamicRestRoute
}

func getDynamicRestRoutesTree(host string) *DynamicRestRoutesTree {
	DynamicRestRoutesTreeCached, err := dynamicRoutingCache.Get(host)
	if err != nil {
		// TODO : Get needed data from KVS ms!
		DynamicRestRoutes := dynaimcRouteTest[host] //TEST
		//if err != nil {
		//	return nil, errors.HostNotExist
		//}

		// If routing exist in KVS MS convert it to "Routing" type.
		DynamicRestRoutesTree := DynamicRestRoutesTree{}
		for i := range DynamicRestRoutes {
			//Check container type and get extra info (PathItem:spec.PathItem) from destinition.

			DynamicRestRoutesTree.set(DynamicRestRoutes[i].Path, &DynamicRestRoutes[i])
		}
		dynamicRoutingCache.Set(host, &DynamicRestRoutesTree, -1)

		return &DynamicRestRoutesTree
	}

	return DynamicRestRoutesTreeCached.(*DynamicRestRoutesTree)
}

// DynamicRestRoutesTree struct, represent a name tree
type DynamicRestRoutesTree struct {
	urlPart          string
	dynamicRestRoute *DynamicRestRoute
	children         []*DynamicRestRoutesTree
	parameterChild   *DynamicRestRoutesTree
	regexChild       *DynamicRestRoutesTree
	wildCardChild    *DynamicRestRoutesTree
}

// Set : adds a url to the tree with the given DynamicRestRoutec address.
func (drt *DynamicRestRoutesTree) set(url string, DynamicRestRoute *DynamicRestRoute) {
	urlParts := strings.Split(strings.Trim(url, "/"), "/")
	innerTree := drt

Loop:
	for i, urlPart := range urlParts {
		if strings.HasPrefix(urlPart, "%") {
			innerTree.regexChild = &DynamicRestRoutesTree{dynamicRestRoute: DynamicRestRoute}
			continue Loop
		} else if strings.HasPrefix(urlPart, "*") {
			innerTree.wildCardChild = &DynamicRestRoutesTree{dynamicRestRoute: DynamicRestRoute}
			continue Loop
		} else if strings.HasPrefix(urlPart, "{") && strings.HasSuffix(urlPart, "}") {
			if innerTree.parameterChild == nil {
				innerTree.parameterChild = &DynamicRestRoutesTree{}
			}

			innerTree = innerTree.parameterChild

			if i == len(urlParts)-1 {
				innerTree.dynamicRestRoute = DynamicRestRoute
			}
			continue Loop
		} else {
			for _, child := range innerTree.children {
				if urlPart == child.urlPart {
					if i == len(urlParts)-1 {
						child.dynamicRestRoute = DynamicRestRoute
					}
					innerTree = child
					continue Loop
				}
			}

			newNode := &DynamicRestRoutesTree{urlPart: urlPart}
			if i == len(urlParts)-1 {
				newNode.dynamicRestRoute = DynamicRestRoute
			}

			innerTree.children = append(innerTree.children, newNode)
			innerTree = newNode
		}
	}
}

// Get finds the best match for a given url
func (drt *DynamicRestRoutesTree) get(url string) *DynamicRestRoute {
	urlParts := strings.Split(strings.Trim(url, "/"), "/")
	lastNode := drt.walkTree(urlParts)
	if lastNode != nil {
		return lastNode.dynamicRestRoute
	}
	return nil
}

func (drt *DynamicRestRoutesTree) walkTree(urlParts []string) *DynamicRestRoutesTree {
	if len(urlParts) == 0 {
		return drt
	}

	if len(drt.children) > 0 {
		for _, child := range drt.children {
			if child.urlPart == urlParts[0] {
				return child.walkTree(urlParts[1:])
			}
		}
	}

	// Check parameterChild and check if requestURL match!
	if drt.parameterChild != nil {
		return drt.parameterChild.walkTree(urlParts[1:])
	}

	// Check regexChild and check if requestURL match! Regex routes start with %.
	if drt.regexChild != nil {
		//// TODO : Check url with regex
		if false {
			return drt.regexChild.walkTree(urlParts[1:])
		}
	}

	// check if wildCardChild exist! WildCard routes start with *.
	if drt.wildCardChild != nil {
		return drt.wildCardChild.walkTree(urlParts[1:])
	}

	// return drt means we need most match not exactly one
	return drt
}
