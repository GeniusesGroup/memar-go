/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../urn"
	"../protocol"
	"../service"
)

const serviceMuxPath = "/m"

var MuxService = muxService{
	Service: service.New("urn:giti:http.protocol:service:service-multiplexer", serviceMuxPath, protocol.Software_PreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Multiplexer", 
			"It use giti urn mechanism to call requested service",
			``,
			[]string{}).
		SetAuthorization(protocol.CRUDAll, protocol.UserTypeAll).Expired(0, ""),
}

type muxService struct {
	service.Service
}

func (ser *muxService) ServeHTTP(st protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	var connection = st.Connection()

	if httpReq.method != MethodPOST {
		connection.StreamFailed()
		httpRes.SetStatus(StatusMethodNotAllowedCode, StatusMethodNotAllowedPhrase)
		return
	}

	var serviceID uint64
	var service protocol.Service
	serviceID, err = urn.IDfromString(httpReq.uri.query)
	if err == nil {
		service, err = protocol.App.GetServiceByID(serviceID)
		if err != nil {
			connection.StreamFailed()
			httpRes.SetStatus(StatusNotFoundCode, StatusNotFoundPhrase)
			httpRes.SetError(ErrNotFound)
			// err = ErrNotFound
			return
		}
	} else {
		connection.StreamFailed()
		httpRes.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		httpRes.SetError(err)
		return
	}

	// Add some header for dynamically services like not index by SE(google, ...), ...
	httpRes.header.Set("X-Robots-Tag", "noindex")
	// httpRes.header.Set(HeaderKeyCacheControl, "no-store")

	st.SetService(service)
	err = service.ServeHTTP(st, httpReq, httpRes)
	if err != nil {
		connection.StreamFailed()
		httpRes.SetError(err)
	} else {
		connection.StreamSucceed()
	}
	return
}

func (ser *muxService) DoHTTP(httpReq *Request, httpRes *Response) (err protocol.Error) {
	return
}
