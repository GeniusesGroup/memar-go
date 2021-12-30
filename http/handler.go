/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../protocol"
)

type Handler struct{}

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
	if !protocol.AppDevMode && !HostSupportedService.ServeHTTP(stream, httpReq, httpRes) {
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return
	}

	switch httpReq.uri.path {
	case serviceMuxPath:
		MuxService.ServeHTTP(stream, httpReq, httpRes)
	case shortenerPath:
		// TODO:::
	case landingPath:
		// TODO:::
	default:
		// serve by URL
		var service protocol.Service
		service, err = protocol.App.GetServiceByURI(httpReq.uri.path)
		if service == nil {
			// If project don't have any logic that support data on e.g. HTTP (restful, ...) we send platform GUI app for web
			ServeWWWService.ServeHTTP(stream, httpReq, httpRes)
		}
	}

	handler.HandleOutcomeResponse(stream, httpReq, httpRes)
	return
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

		if httpRes.body.Len() > 0 {
			httpRes.SetContentLength()
		} else {
			httpRes.header.SetTransferEncoding(HeaderValueChunked)
		}
	} else {
		httpRes.header.SetZeroContentLength()
	}

	// httpRes.header.Set(HeaderKeyAccessControlAllowOrigin, "*")

	stream.SetOutcomeData(httpRes)

	if protocol.AppDeepDebugMode {
		// TODO::: body not serialized yet to log it!! any idea to have better performance below??
		protocol.App.Log(protocol.Log_Confidential, "HTTP - Request:::", httpReq.uri.uri, httpReq.header, string(httpReq.body.Marshal()))
		protocol.App.Log(protocol.Log_Confidential, "HTTP - Response:::", httpRes.ReasonPhrase, httpRes.header, string(httpRes.body.Marshal()))
	}
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

	err = stream.SendRequest()
	if err != nil {
		return
	}

	httpRes = NewResponse()
	err = httpRes.Unmarshal(stream.IncomeData().Marshal())
	return
}
