/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../mediatype"
	"../protocol"
	"../service"
)

const serviceMuxPath = "/m"

var MuxService = muxService{
	Service: service.New("", mediatype.New("domain/http.protocol.service; name=service-multiplexer").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Multiplexer",
		"Multiplex services by its ID with impressive performance",
		"",
		"",
		nil).SetInfo(protocol.Software_PreAlpha, 1587282740, "")).
		SetAuthorization(protocol.CRUDAll, protocol.UserType_All),
}

type muxService struct {
	*service.Service
}

func (ser *muxService) ServeHTTP(st protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	if httpReq.method != MethodPOST {
		// err =
		httpRes.SetStatus(StatusMethodNotAllowedCode, StatusMethodNotAllowedPhrase)
		return
	}

	var serviceID uint64
	var service protocol.Service
	serviceID, err = mediatype.IDfromString(httpReq.uri.query)
	if err == nil {
		service, err = protocol.App.GetServiceByID(serviceID)
		if err != nil {
			httpRes.SetStatus(StatusNotFoundCode, StatusNotFoundPhrase)
			httpRes.SetError(ErrNotFound)
			// err = ErrNotFound
			return
		}
	} else {
		httpRes.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		httpRes.SetError(err)
		return
	}

	// Add some header for dynamically services like not index by SE(google, ...), ...
	httpRes.H.Set("X-Robots-Tag", "noindex")
	// httpRes.H.Set(HeaderKeyCacheControl, "no-store")

	st.SetService(service)
	// TODO::: can't easily call service and must schedule it by its weight.
	err = service.ServeHTTP(st, httpReq, httpRes)
	if err != nil {
		httpRes.SetError(err)
	}
	return
}

func (ser *muxService) DoHTTP(httpReq *Request, httpRes *Response) (err protocol.Error) {
	return
}
