/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../protocol"
	"../service"
)

var HostSupportedService = hostSupportedService{
	Service: service.New("urn:giti:http.protocol:service:host-supported", "", protocol.Software_PreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Multiplexer",
			"",
			``,
			[]string{}).
		SetAuthorization(protocol.CRUDAll, protocol.UserTypeAll).Expired(0, ""),
}

type hostSupportedService struct {
	service.Service
}

func (ser *hostSupportedService) ServeHTTP(stream protocol.Stream, httpReq *Request, httpRes *Response) (supported bool) {
	var domainName = protocol.OS.AppManifest().DomainName()
	var host = httpReq.uri.host
	var path = httpReq.uri.path
	var query = httpReq.uri.query

	if host == "" {
		// TODO::: noting to do or reject request??
	} else if '0' <= host[0] && host[0] <= '9' {
		// check of request send over IP
		if protocol.AppDeepDebugMode {
			protocol.App.Log(protocol.Log_Debug, "HTTP - Host Check - IP host:", host)
		}

		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if len(host) > 4 && host[:4] == "www." {
		if host[4:] != domainName {
			if protocol.AppDeepDebugMode {
				protocol.App.Log(protocol.Log_Debug, "HTTP - Host Check - Unknown WWW host:", host)
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return false
		}

		if protocol.AppDeepDebugMode {
			protocol.App.Log(protocol.Log_Debug, "HTTP - Host Check - WWW host:", host)
		}

		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www
		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if host != domainName {
		if protocol.AppDeepDebugMode {
			protocol.App.Log(protocol.Log_Debug, "HTTP - Host Check - Unknown host:", host)
		}
		// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
		return false
	}
	return true
}

func (ser *hostSupportedService) DoHTTP(httpReq *Request, httpRes *Response) (err protocol.Error) {
	return
}
