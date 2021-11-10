/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../convert"
	"../protocol"
	"../service"
)

const muxPath = "/m"

var MuxService = muxService{
	Service: service.New("urn:giti:http.protocol:service:service-multiplexer", muxPath, protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Service Multiplexer",
			``,
			[]string{}).
		SetAuthorization(protocol.CRUDAll, protocol.UserTypeAll).Expired(0, ""),
}

type muxService struct {
	service.Service
}

func (ser *muxService) ServeHTTP(st protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	if httpReq.method != MethodPOST {
		st.Connection().ServiceCallFail()
		httpRes.SetStatus(StatusMethodNotAllowedCode, StatusMethodNotAllowedPhrase)
		return
	}

	var serviceID uint64
	serviceID, err = convert.Base10StringToUint64(httpReq.uri.query)
	if err == nil {
		service = protocol.App.GetServiceByID(serviceID)
	}
	// Add some header for /apis like not index by SE(google, ...), ...
	httpRes.header.Set("X-Robots-Tag", "noindex")
	// httpRes.header.Set(HeaderKeyCacheControl, "no-store")

	if service == nil {
		st.Connection().ServiceCallFail()
		httpRes.SetStatus(StatusNotFoundCode, StatusNotFoundPhrase)
		httpRes.SetError(ErrNotFound)
		// err = ErrNotFound
		return
	}

	st.SetService(service)
	err = service.ServeHTTP(st, httpReq, httpRes)
	if err != nil {
		st.Connection().ServiceCallFail()
		httpRes.SetError(err)
	} else {
		st.Connection().ServiceCalled()
	}
	return
}

func (ser *muxService) DoHTTP(httpReq *Request, httpRes *Response) (err protocol.Error) {
	return
}
