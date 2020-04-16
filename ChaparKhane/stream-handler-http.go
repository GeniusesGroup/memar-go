/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// httpHandler use to handle HTTP protocol! It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func httpHandler(s *Server, st *Stream) {
	// Ready data for logics & do some logic
	// - Route by URL
	// - Encode||Decode body by mime type header

	// Add Server Header to response : "ChaparKhane"

	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
}

// httpOptionsHandler is a filter function that inspects the HTTP Request for the OPTIONS method when Origin header not exist
// and provides the response with a set of allowed methods for the request URL Path.
func httpOptionsHandler() {
	// if ctx.Request.Method == "OPTIONS" && ctx.Request.Header.Get(http.RequestStandards.Headers.Origin) == "" {
	// 	var allowMethods []string
	// 	staticRestRoute := ctx.InterfaceData["StaticRestRoute"].(*StaticRestRoute)
	// 	for m := range staticRestRoute.MethodsFilters {
	// 		allowMethods = append(allowMethods, m)
	// 	}

	// 	ctx.Response.Header.Set(http.ResponseStandards.Headers.Allow, strings.Join(allowMethods, ","))
	// }
}
