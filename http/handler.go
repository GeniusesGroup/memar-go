/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../protocol"
	"../www"
)

// TODO::: Have default error pages and can get customizes!
// Send beauty HTML response in http error situation like 500, 404, ...

type Handler struct {
	WWW www.Assets
}

// HandleIncomeRequest handle incoming HTTP request streams.
// It can use for architectures like restful, ...
// Protocol Standard - HTTP/1 : https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol
// Protocol Standard - HTTP/2 : https://httpwg.org/specs/rfc7540.html
// Protocol Standard - HTTP/3 : https://quicwg.org/base-drafts/draft-ietf-quic-http.html
func (handler *Handler) HandleIncomeRequest(stream protocol.Stream) (err protocol.Error) {
	var httpReq = NewRequest()
	var httpRes = NewResponse()

	err = httpReq.Unmarshal(stream.IncomeData().Marshal())
	if err != nil {
		httpRes.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return
	}
	err = handler.ServeHTTP(stream, httpReq, httpRes)
	return
}

// ServeHTTP handle incoming HTTP request.
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func (handler *Handler) ServeHTTP(stream protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	if !protocol.AppDevMode && handler.HostCheck(stream, httpReq, httpRes) {
		return
	}

	var connection = stream.Connection()
	var service protocol.Service

	switch httpReq.uri.path {
	case muxPath:
		MuxService.ServeHTTP(stream, httpReq, httpRes)
	case shortenerPath:
		// TODO:::
	default:
		// serve by URL
		service = protocol.App.GetServiceByURI(httpReq.uri.host, httpReq.uri.path)
		if service == nil {
			// If project don't have any logic that support data on e.g. HTTP (restful, ...) we send platform GUI app for web
			handler.ServeWWW(stream, httpReq, httpRes)
			return
		}
	}

	handler.HandleOutcomeResponse(stream, httpReq, httpRes)
	return
}

// ServeWWW will serve WWW assets to request
func (handler *Handler) ServeWWW(stream protocol.Stream, httpReq *Request, httpRes *Response) {
	var reqFile, _ = handler.WWW.GUI.FileByPath(httpReq.uri.path)
	if reqFile == nil {
		// TODO::: SSR to serve-to-robots

		const supportedLang = "en" // TODO::: get from header
		reqFile, _ = handler.WWW.MainHTML.File(supportedLang)
	}
	connection.ServiceCalled()
	httpRes.SetStatus(StatusOKCode, StatusOKPhrase)
	httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
	httpRes.SetBody(reqFile.Data())

	handler.HandleOutcomeResponse(stream, httpReq, httpRes)
}

// HandleOutcomeResponse use to handle outcoming HTTP response stream!
func (handler *Handler) HandleOutcomeResponse(stream protocol.Stream, httpReq *Request, httpRes *Response) {
	stream.Close()

	// Do some global assignment to response
	httpRes.version = httpReq.version
	if httpRes.Body() != nil {
		var mediaType = httpRes.body.MediaType()
		if mediaType != nil {
			httpRes.header.Set(HeaderKeyContentType, mediaType.MediaType())
		}
		var compressType = httpRes.body.CompressType()
		if compressType != nil {
			httpRes.header.Set(HeaderKeyContentEncoding, compressType.ContentEncoding())
		}
	}
	// httpRes.header.Set(HeaderKeyAccessControlAllowOrigin, "*")
	httpRes.SetContentLength()

	// TODO::: Have default error pages and can get customizes!
	// Send beauty HTML response in http error situation like 500, 404, ...

	stream.SetOutcomeData(httpRes)

	if protocol.AppDeepDebugMode {
		// TODO::: body not serialized yet to log it!! any idea to have better performance below??
		protocol.App.LogConfidential("HTTP - Request:::", httpReq.uri.uri, httpReq.header, string(httpReq.body.Marshal()))
		protocol.App.LogConfidential("HTTP - Response:::", httpRes.ReasonPhrase, httpRes.header, string(httpRes.body.Marshal()))
	}
}

func (handler *Handler) HostCheck(stream protocol.Stream, httpReq *Request, httpRes *Response) (redirect bool) {
	var domainName = protocol.OS.AppManifest().DomainName()
	var host = httpReq.uri.host
	var path = httpReq.uri.path
	var query = httpReq.uri.query

	if host == "" {
		// TODO::: noting to do or reject request??
	} else if '0' <= host[0] && host[0] <= '9' {
		// check of request send over IP
		if protocol.AppDeepDebugMode {
			protocol.App.LogDebug("HTTP - Host Check - IP host:", host)
		}

		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return true
	} else if len(host) > 4 && host[:4] == "www." {
		if host[4:] != domainName {
			if protocol.AppDeepDebugMode {
				protocol.App.LogDebug("HTTP - Host Check - Unknown WWW host:", host)
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return true
		}

		if protocol.AppDeepDebugMode {
			protocol.App.LogDebug("HTTP - Host Check - WWW host:", host)
		}

		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return true
	} else if host != domainName {
		if protocol.AppDeepDebugMode {
			protocol.App.LogDebug("HTTP - Host Check - Unknown host:", host)
		}
		// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
		return true
	}
	return
}

// HandleOutcomeRequest use to handle outcoming HTTP request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func HandleOutcomeRequest(conn protocol.Connection, service protocol.Service, httpReq *Request) (httpRes *Response, err protocol.Error) {
	var stream protocol.Stream
	stream, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	stream.SetOutcomeData(httpReq)

	err = conn.Send(stream)
	if err != nil {
		return
	}

	httpRes = NewResponse()
	err = httpRes.Unmarshal(stream.IncomeData().Marshal())
	return
}
