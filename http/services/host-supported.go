/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/http"
	"github.com/GeniusesGroup/libgo/log"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/service"
)

var HostSupportedService = hostSupportedService{}

func init() {
	HostSupportedService.MT.Init("domain/http.protocol.service; name=host-supported")
	HostSupportedService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Host Supported",
		"Service to check if requested host is valid or not",
		"",
		"",
		nil)
}

type hostSupportedService struct {
	detail.DS
	mediatype.MT
	service.Service
}

//libgo:impl protocol.MediaType
func (s *hostSupportedService) FileExtension() string               { return "" }
func (s *hostSupportedService) Status() protocol.SoftwareStatus     { return protocol.Software_PreAlpha }
func (s *hostSupportedService) ReferenceURI() string                { return "" }
func (s *hostSupportedService) IssueDate() protocol.Time            { return nil } // 1587282740
func (s *hostSupportedService) ExpiryDate() protocol.Time           { return nil }
func (s *hostSupportedService) ExpireInFavorOf() protocol.MediaType { return nil }
func (s *hostSupportedService) Fields() []protocol.Field            { return nil }

//libgo:impl protocol.Service
func (s *hostSupportedService) URI() string                 { return "" }
func (s *hostSupportedService) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *hostSupportedService) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *hostSupportedService) CRUDType() protocol.CRUD     { return protocol.CRUD_All }
func (s *hostSupportedService) UserType() protocol.UserType { return protocol.UserType_All }

func (ser *hostSupportedService) ServeHTTP(stream protocol.Stream, httpReq *http.Request, httpRes *http.Response) (supported bool) {
	var domainName = protocol.OS.AppManifest().DomainName()
	var host = httpReq.URI().Host()
	var path = httpReq.URI().Path()
	var query = httpReq.URI().Query()

	if host == "" {
		// TODO::: noting to do or reject request??
	} else if '0' <= host[0] && host[0] <= '9' {
		// check of request send over IP
		if protocol.AppMode_Dev {
			protocol.App.Log(log.DeepDebugEvent(domainEnglish, "Host Check - IP host: "+host))
		}

		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
		httpRes.H.Set(http.HeaderKeyLocation, target)
		httpRes.H.Set(http.HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if len(host) > 4 && host[:4] == "www." {
		if host[4:] != domainName {
			if protocol.AppMode_Dev {
				protocol.App.Log(log.DeepDebugEvent(domainEnglish, "Host Check - Unknown WWW host: "+host))
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return false
		}

		if protocol.AppMode_Dev {
			protocol.App.Log(log.DeepDebugEvent(domainEnglish, "Host Check - WWW host: "+host))
		}

		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www
		// TODO::: target alloc occur multiple, improve it.
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
		httpRes.H.Set(http.HeaderKeyLocation, target)
		httpRes.H.Set(http.HeaderKeyCacheControl, "max-age=31536000, immutable")
		return false
	} else if host != domainName {
		if protocol.AppMode_Dev {
			protocol.App.Log(log.DeepDebugEvent(domainEnglish, "Host Check - Unknown host: "+host))
		}
		// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
		return false
	}
	return true
}

func (ser *hostSupportedService) doHTTP(httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	return
}
