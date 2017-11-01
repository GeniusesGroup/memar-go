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

// Set a route in specific URL and add this to its filters. e.g.
/*
	// Register OpenAPI dynamic API document creator.
	openAPI := filters.NewStaticRestRoute("/document/openapi/")
	openAPI.MethodsFilters = map[string]ehttp.Filter{"GET": filters.OpenAPI}
	openAPI.CORS = filters.Cors{Origins: []string{"*"}, Methods: []string{"GET"}}
	openAPI.PathItem.GET = openapi.NewOperation("OpenAPI").WithDescription("Return OpenAPI document for service")
*/

package filters

import (
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
	"github.com/SabzCity/go-library/net/ehttp"
)

// OpenAPI : Install document of service in ".../document/openapi/" URL.
// Filter level: just Routes
// Needed filter: before(StaticRestRouter/), after()
func OpenAPI(ctx *ehttp.Context) {
	if microservice.MSDetails.OpenAPI != nil {
		ctx.ContentEncoder(microservice.MSDetails.OpenAPI)
		return
	}

	ctx.Error = errors.ContentNotExist
}
