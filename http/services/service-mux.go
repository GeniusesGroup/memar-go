/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/http"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/service"
	uuid "github.com/GeniusesGroup/libgo/uuid/32byte"
)

const MuxService_Path = "/m"

var MuxService = muxService{}

func init() {
	MuxService.MT.Init("domain/http.protocol.service; name=service-multiplexer")
	MuxService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Multiplexer",
		"Multiplex services by its ID with impressive performance",
		"",
		"",
		nil)
}

type muxService struct {
	detail.DS
	mediatype.MT
	service.Service
}

//libgo:impl protocol.MediaType
func (s *muxService) FileExtension() string               { return "" }
func (s *muxService) Status() protocol.SoftwareStatus     { return protocol.Software_PreAlpha }
func (s *muxService) ReferenceURI() string                { return "" }
func (s *muxService) IssueDate() protocol.Time            { return nil } // 1587282740
func (s *muxService) ExpiryDate() protocol.Time           { return nil }
func (s *muxService) ExpireInFavorOf() protocol.MediaType { return nil }
func (s *muxService) Fields() []protocol.Field            { return nil }

//libgo:impl protocol.Service
func (s *muxService) URI() string                 { return MuxService_Path }
func (s *muxService) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *muxService) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *muxService) CRUDType() protocol.CRUD     { return protocol.CRUD_All }
func (s *muxService) UserType() protocol.UserType { return protocol.UserType_All }

func (ser *muxService) ServeHTTP(st protocol.Stream, httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
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

func (ser *muxService) doHTTP(httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	return
}
