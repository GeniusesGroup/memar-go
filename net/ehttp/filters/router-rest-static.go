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

	"github.com/SabzCity/go-library/documentation/openapi"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
	"github.com/SabzCity/go-library/net/ehttp"
)

var staticRestRoutes StaticRestRoutesTree

// StaticRestRoute :
type StaticRestRoute struct {
	Path           string                  // Full Path URL
	PathParts      []string                // Auto vlaue and cached for dispatching
	PathItem       openapi.PathItem        //
	MethodsFilters map[string]ehttp.Filter // *, GET, POST, HEAD, PUT, DELETE, OPTIONS, CONNECT, TRACE, PATCH HTTP method filters for this URL
	CORS           Cors                    // Cors setting
	Metadata       map[string]interface{}  // Extra information used to store custom information about the route.
}

// StaticRestRouter call as filter to specify.
// Filter level:(Server, Route)  Server
// Needed filter: before(), after()
func StaticRestRouter(ctx *ehttp.Context) {
	requestURLParts := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")

	// Find route for special URL.
	StaticRestRoute := staticRestRoutes.get(requestURLParts)
	if StaticRestRoute == nil {
		ctx.Error = errors.RouterNotFound
		// Announce to server, response is ready!
		ctx.ResponseReady = true
		return
	}

	for i, part := range StaticRestRoute.PathParts {
		// Check the URL parameter.
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			// parameterName without enclosing {}
			parameterName := part[1 : len(part)-1]
			ctx.StringData[parameterName] = requestURLParts[i]
		}
	}

	switch ctx.Request.Method {
	case "GET":
		if StaticRestRoute.PathItem.GET != nil {
			ctx.StringData["MethodName"] = StaticRestRoute.PathItem.GET.OperationID
		}
	case "PUT":
		if StaticRestRoute.PathItem.PUT != nil {
			ctx.StringData["MethodName"] = StaticRestRoute.PathItem.PUT.OperationID
		}
	case "POST":
		if StaticRestRoute.PathItem.POST != nil {
			ctx.StringData["MethodName"] = StaticRestRoute.PathItem.POST.OperationID
		}
	case "DELETE":
		if StaticRestRoute.PathItem.DELETE != nil {
			ctx.StringData["MethodName"] = StaticRestRoute.PathItem.DELETE.OperationID
		}
	case "PATCH":
		if StaticRestRoute.PathItem.PATCH != nil {
			ctx.StringData["MethodName"] = StaticRestRoute.PathItem.PATCH.OperationID
		}
	}

	ctx.InterfaceData["StaticRestRoute"] = StaticRestRoute
	ctx.StringData["ServiceName"] = StaticRestRoute.PathParts[0] //// ??
}

// NewStaticRestRoute creates a new StaticRestRoute instance and add it to StaticRestRoutesTree
// It expects an ID as parameter but not passing an ID is also valid.
func NewStaticRestRoute(path string) *StaticRestRoute {
	sr := StaticRestRoute{}
	sr.Path = path
	sr.PathParts = strings.Split(strings.Trim(sr.Path, "/"), "/")

	microservice.MSDetails.OpenAPI.Paths[path] = &sr.PathItem

	staticRestRoutes.set(sr.PathParts, &sr)

	return &sr
}

// StaticRestRoutesTree struct, represent a name tree
type StaticRestRoutesTree struct {
	urlPart         string
	StaticRestRoute *StaticRestRoute
	children        []*StaticRestRoutesTree
	parameterChild  *StaticRestRoutesTree
}

// Set : adds a url to the tree with the given DynamicRestRoutec address.
func (srt *StaticRestRoutesTree) set(urlParts []string, StaticRestRoute *StaticRestRoute) {
	innerTree := srt

L1:
	for i, urlPart := range urlParts {
		if strings.HasPrefix(urlPart, "{") && strings.HasSuffix(urlPart, "}") {
			if innerTree.parameterChild == nil {
				innerTree.parameterChild = &StaticRestRoutesTree{}
			}

			innerTree = innerTree.parameterChild

			if i == len(urlParts)-1 {
				innerTree.StaticRestRoute = StaticRestRoute
			}
		} else {
			for _, child := range innerTree.children {
				if urlPart == child.urlPart {
					if i == len(urlParts)-1 {
						child.StaticRestRoute = StaticRestRoute
					}
					innerTree = child
					continue L1
				}
			}

			newNode := &StaticRestRoutesTree{urlPart: urlPart}
			if i == len(urlParts)-1 {
				newNode.StaticRestRoute = StaticRestRoute
			}

			innerTree.children = append(innerTree.children, newNode)
			innerTree = newNode
		}
	}
}

// Get finds the best match for a given url
func (srt *StaticRestRoutesTree) get(urlParts []string) *StaticRestRoute {
	lastNode := srt.walkTree(urlParts)
	if lastNode != nil {
		return lastNode.StaticRestRoute
	}
	return nil
}

func (srt *StaticRestRoutesTree) walkTree(urlParts []string) *StaticRestRoutesTree {
	if len(urlParts) == 0 {
		return srt
	}

	if len(srt.children) > 0 {
		for _, child := range srt.children {
			if child.urlPart == urlParts[0] {
				return child.walkTree(urlParts[1:])
			}
		}
	}

	if srt.parameterChild != nil {
		return srt.parameterChild.walkTree(urlParts[1:])
	}

	return nil
}
