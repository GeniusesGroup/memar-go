/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../convert"
	"../log"
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
	var streamSocket = stream.Socket()

	var maybeBody []byte
	maybeBody, err = httpReq.UnmarshalFrom(streamSocket.Marshal())
	if err != nil {
		httpRes.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return
	}
	httpReq.body.checkAndSetCodecAsIncomeBody(maybeBody, streamSocket, &httpReq.H)

	err = handler.ServeHTTP(stream, httpReq, httpRes)
	return
}

// ServeHTTP handle incoming HTTP request.
// Developers can develope new one with its desire logic to more respect protocols like restful, ...
func (handler *Handler) ServeHTTP(stream protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	if !protocol.AppMode_Dev && !HostSupportedService.ServeHTTP(stream, httpReq, httpRes) {
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return
	}

	switch httpReq.uri.path {
	case serviceMuxPath:
		err = MuxService.ServeHTTP(stream, httpReq, httpRes)
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
			err = ServeWWWService.ServeHTTP(stream, httpReq, httpRes)
		}
	}
	handler.HandleOutcomeResponse(stream, httpReq, httpRes)
	return
}

// HandleOutcomeResponse use to handle outcoming HTTP response stream!
func (handler *Handler) HandleOutcomeResponse(stream protocol.Stream, httpReq *Request, httpRes *Response) {
	// Do some global assignment to response
	httpRes.version = httpReq.version
	if httpRes.Body() != nil {
		var mediaType = httpRes.body.MediaType()
		if mediaType != nil {
			httpRes.H.Set(HeaderKeyContentType, mediaType.MediaType())
		}
		var compressType = httpRes.body.CompressType()
		if compressType != nil {
			httpRes.H.Set(HeaderKeyContentEncoding, compressType.ContentEncoding())
		}

		if httpRes.body.Len() > 0 {
			httpRes.SetContentLength()
		} else {
			httpRes.H.SetTransferEncoding(HeaderValueChunked)
		}
	} else {
		httpRes.H.SetZeroContentLength()
	}

	// httpRes.H.Set(HeaderKeyAccessControlAllowOrigin, "*")

	stream.SendResponse(httpRes)

	if protocol.LogMode_DeepDebug {
		// TODO::: req||res not serialized yet to log it!! any idea to have better performance below??
		protocol.App.Log(log.ConfEvent(domainEnglish, convert.UnsafeByteSliceToString(httpReq.Marshal())))
		protocol.App.Log(log.ConfEvent(domainEnglish, convert.UnsafeByteSliceToString(httpRes.Marshal())))
	}
}

// SendBidirectionalRequest use to handle outcoming HTTP request stream
// It block caller until get response or error
func SendBidirectionalRequest(conn protocol.Connection, service protocol.Service, httpReq *Request) (httpRes *Response, err protocol.Error) {
	var stream protocol.Stream
	stream, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	err = stream.SendRequest(httpReq)
	if err != nil {
		return
	}

	for {
		var status = <-stream.State()
		switch status {
		case protocol.ConnectionState_Timeout:
			// err =
		case protocol.ConnectionState_ReceivedCompletely:
			httpRes = NewResponse()
			err = httpRes.Unmarshal(stream.Socket().Marshal())
			if err == nil {
				err = httpRes.GetError()
			}
		default:
			continue
		}
		stream.Close()
		break
	}
	return
}

// SendUnidirectionalRequest use to send outcoming HTTP request and don't expect any response.
// It block caller until request send successfully or return error
func SendUnidirectionalRequest(conn protocol.Connection, service protocol.Service, httpReq *Request) (err protocol.Error) {
	var stream protocol.Stream
	stream, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	err = stream.SendRequest(httpReq)
	if err != nil {
		return
	}

	for {
		var status = <-stream.State()
		switch status {
		case protocol.ConnectionState_Timeout:
			// err =
		case protocol.ConnectionState_SentCompletely:
			// Nothing to do. Just let execution go to stream.Close() and break the loop
		default:
			continue
		}
		stream.Close()
		break
	}
	return
}
