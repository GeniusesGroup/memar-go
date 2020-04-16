/* For license and copyright information please see LEGAL file in repository */

package services

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
