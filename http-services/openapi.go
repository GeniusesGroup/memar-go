/* For license and copyright information please see LEGAL file in repository */

package services

// Set a route in specific URL and add this to its filters. e.g.
/*
	// Register OpenAPI dynamic API document creator.
	openAPI := filters.NewStaticRestRoute("/document/openapi/")
	openAPI.MethodsFilters = map[string]ehttp.Filter{"GET": filters.OpenAPI}
	openAPI.CORS = filters.Cors{Origins: []string{"*"}, Methods: []string{"GET"}}
	openAPI.PathItem.GET = openapi.NewOperation("OpenAPI").WithDescription("Return OpenAPI document for service")
*/

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
