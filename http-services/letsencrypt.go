/* For license and copyright information please see LEGAL file in repository */

package services

import (
	"net/http"
	"strings"
)

// LetsEncrypt :
// Filter level: just Server
// Needed filter: before(), after()
func LetsEncrypt(ctx *ehttp.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, "/.well-known/acme-challenge/") {
		// Get needed certificate to connect a container.
		ctx.Request.URL.Scheme = "http" // Change for TEST
		// Set "DomainMS:Port" instead of requested Host.
		ctx.Request.URL.Host = "v0.domains.sabzcity.containers.sabz.city:50030"
		ctx.Request.Host = "v0.domains.sabzcity.containers.sabz.city:50030"

		// Use default transport.
		ctx.InterfaceData["Proxy.Transport"] = http.DefaultTransport

		ctx.ResponseReady = true
	}
}
