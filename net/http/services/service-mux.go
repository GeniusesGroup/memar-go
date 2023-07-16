/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"libgo/detail"
	"libgo/http"
	"libgo/mediatype"
	"libgo/protocol"
	"libgo/service"
	uuid "libgo/uuid/32byte"
)

const MuxService_Path = "/s"

var MuxService muxService

type muxService struct {
	detail.DS
	mediatype.MT
	service.Service
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (s *muxService) Init() (err protocol.Error) {
	err = s.MT.Init("domain/httpwg.org; type=service; name=multiplexer")

	return
}

//libgo:impl libgo/protocol.MediaType
func (s *muxService) FileExtension() string               { return "" }
func (s *muxService) Status() protocol.SoftwareStatus     { return protocol.Software_PreAlpha }
func (s *muxService) ReferenceURI() string                { return "" }
func (s *muxService) IssueDate() protocol.Time            { return nil } // 1587282740
func (s *muxService) ExpiryDate() protocol.Time           { return nil }
func (s *muxService) ExpireInFavorOf() protocol.MediaType { return nil }

//libgo:impl libgo/protocol.Object
func (s *muxService) Fields() []protocol.Object_Member_Field   { return nil }
func (s *muxService) Methods() []protocol.Object_Member_Method { return nil }

//libgo:impl libgo/protocol.Service
func (s *muxService) URI() string                 { return MuxService_Path }
func (s *muxService) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *muxService) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *muxService) CRUDType() protocol.CRUD     { return protocol.CRUD_All }
func (s *muxService) UserType() protocol.UserType { return protocol.UserType_All }

func (s *muxService) ServeHTTP(st protocol.Stream, httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	if httpReq.Method() != http.MethodPOST {
		// err =
		httpRes.SetStatus(http.StatusMethodNotAllowedCode, http.StatusMethodNotAllowedPhrase)
		return
	}

	var serviceID uint64
	var service protocol.Service
	serviceID, err = uuid.IDfromString(httpReq.URI().Query())
	if err == nil {
		service, err = protocol.App.GetServiceByID(protocol.ID(serviceID))
		if err != nil {
			httpRes.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
			httpRes.SetError(&ErrNotFound)
			// err = &ErrNotFound
			return
		}
	} else {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
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

func (s *muxService) doHTTP(httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	return
}
