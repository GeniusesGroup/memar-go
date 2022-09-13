/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../log"
	"../mediatype"
	"../protocol"
	"../service"
)

var HostSupportedService = hostSupportedService{
	Service: service.New("", mediatype.New("domain/http.protocol.service; name=host-supported").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Host Supported",
		"Service to check if requested host is valid or not",
		"",
		"",
		nil).SetInfo(protocol.Software_PreAlpha, 1587282740, "")).
		SetAuthorization(protocol.CRUDAll, protocol.UserType_All),
}

type hostSupportedService struct {
	*service.Service
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
		if protocol.LogMode_DeepDebug {
			protocol.App.Log(log.DebugEvent(domainEnglish, "Host Check - IP host: "+host))
		}

		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.H.Set(HeaderKeyLocation, target)
		httpRes.H.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if len(host) > 4 && host[:4] == "www." {
		if host[4:] != domainName {
			if protocol.LogMode_DeepDebug {
				protocol.App.Log(log.DebugEvent(domainEnglish, "Host Check - Unknown WWW host: "+host))
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return false
		}

		if protocol.LogMode_DeepDebug {
			protocol.App.Log(log.DebugEvent(domainEnglish, "Host Check - WWW host: "+host))
		}

		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www
		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.H.Set(HeaderKeyLocation, target)
		httpRes.H.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if host != domainName {
		if protocol.LogMode_DeepDebug {
			protocol.App.Log(log.DebugEvent(domainEnglish, "Host Check - Unknown host: "+host))
		}
		// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
		return false
	}
	return true
}

func (ser *hostSupportedService) DoHTTP(httpReq *Request, httpRes *Response) (err protocol.Error) {
	return
}
