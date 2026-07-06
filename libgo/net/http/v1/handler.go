/* For license and copyright information please see the LEGAL file in the code repository */

package hh

import (
	"github.com/GeniusesGroup/libgo/convert"
	"github.com/GeniusesGroup/libgo/http"
	hs "github.com/GeniusesGroup/libgo/http/services"
	"github.com/GeniusesGroup/libgo/log"
	"github.com/GeniusesGroup/libgo/protocol"
)

// Protocol Standard - HTTP/1 : https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol
type V1 struct{}

// HandleIncomeRequest handle incoming HTTP request streams.
// It can use for architectures like restful, ...
func (h *V1) HandleIncomeRequest(st protocol.Stream) (err protocol.Error) {
	var httpReq http.Request
	var httpRes http.Response
	httpReq.Init()
	httpRes.Init()

	_, err = httpReq.Decode(st)
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		h.HandleOutcomeResponse(st, &httpReq, &httpRes)
		return
	}

	err = h.ServeHTTP(st, &httpReq, &httpRes)
	return
}

// ServeHTTP handle incoming HTTP request.
// Developers can develope new one with its desire logic to more respect protocols like restful, ...
func (h *V1) ServeHTTP(st protocol.Stream, httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	if !protocol.AppMode_Dev && !hs.HostSupportedService.ServeHTTP(st, httpReq, httpRes) {
		h.HandleOutcomeResponse(st, httpReq, httpRes)
		return
	}

	var path = httpReq.URI().Path()
	switch path {
	case hs.MuxService_Path:
		err = hs.MuxService.ServeHTTP(st, httpReq, httpRes)
	case hs.LandingService_Path:
		// TODO:::
	default:
		// serve by URL
		var service protocol.Service
		service, err = protocol.App.GetServiceByURI(path)
		if service == nil {
			// If project don't have any logic that support data on e.g. HTTP (restful, ...) we send platform GUI app for web
			err = hs.ServeWWWService.ServeHTTP(st, httpReq, httpRes)
		} else {
			err = service.ServeHTTP(st, httpReq, httpRes)
		}
	}
	h.HandleOutcomeResponse(st, httpReq, httpRes)
	return
}

// HandleOutcomeResponse use to handle outcome HTTP response stream.
func (h *V1) HandleOutcomeResponse(st protocol.Stream, httpReq *http.Request, httpRes *http.Response) {
	// Do some global assignment to response
	httpRes.SetVersion(httpReq.Version())
	if httpRes.Body() != nil {
		var mediaType = httpRes.Body().MediaType()
		if mediaType != nil {
			httpRes.H.Set(http.HeaderKeyContentType, mediaType.ToString())
		}
		var compressType = httpRes.Body().CompressType()
		if compressType != nil {
			httpRes.H.Set(http.HeaderKeyContentEncoding, compressType.ContentEncoding())
		}

		if httpRes.Body().Len() > 0 {
			httpRes.SetContentLength()
		} else {
			httpRes.H.SetTransferEncoding(http.HeaderValueChunked)
		}
	} else {
		httpRes.H.SetZeroContentLength()
	}

	// httpRes.H.Set(HeaderKeyAccessControlAllowOrigin, "*")

	st.Encode(httpRes)

	if protocol.AppMode_Dev {
		// TODO::: req||res not serialized yet to log it! any idea to have better performance below??
		var req, _ = httpReq.Marshal()
		var res, _ = httpRes.Marshal()
		protocol.App.Log(log.ConfEvent(domainEnglish, convert.UnsafeByteSliceToString(req)))
		protocol.App.Log(log.ConfEvent(domainEnglish, convert.UnsafeByteSliceToString(res)))
	}
}

// SendBidirectionalRequest use to handle outcoming HTTP request stream
// It block caller until get response or error
func SendBidirectionalRequest(conn protocol.Connection, service protocol.Service, httpReq *http.Request) (httpRes *http.Response, err protocol.Error) {
	var st protocol.Stream
	st, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	_, err = st.Encode(httpReq)
	if err != nil {
		return
	}

	var streamStatus = st.State()
	for {
		var status = <-streamStatus
		switch status {
		case protocol.NetworkStatus_Timeout:
			// err =
		case protocol.NetworkStatus_ReceivedCompletely:
			var res http.Response
			res.Init()
			_, err = res.Decode(st)
			if err == nil {
				err = res.GetError()
			}
			httpRes = &res
		default:
			continue
		}
		st.Close()
		break
	}
	return
}

// SendUnidirectionalRequest use to send outcome HTTP request and don't expect any response.
// It block caller until request send successfully or return error
func SendUnidirectionalRequest(conn protocol.Connection, service protocol.Service, httpReq *http.Request) (err protocol.Error) {
	var st protocol.Stream
	st, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	_, err = st.Encode(httpReq)
	if err != nil {
		return
	}

	var streamStatus = st.State()
	for {
		var status = <-streamStatus
		switch status {
		case protocol.NetworkStatus_Timeout:
			// err =
		case protocol.NetworkStatus_SentCompletely:
			// Nothing to do. Just let execution go to st.Close() and break the loop
		default:
			continue
		}
		st.Close()
		break
	}
	return
}
