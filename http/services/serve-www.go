/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/http"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/service"
	"github.com/GeniusesGroup/libgo/www"
)

var ServeWWWService = serveWWWService{}

func init() {
	ServeWWWService.MT.Init("domain/http.protocol.service; name=serve-www")
	ServeWWWService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Serve WWW",
		"",
		"",
		"",
		nil)
}

type serveWWWService struct {
	detail.DS
	mediatype.MT
	service.Service
	WWW www.Assets
}

//libgo:impl protocol.MediaType
func (s *serveWWWService) FileExtension() string               { return "" }
func (s *serveWWWService) Status() protocol.SoftwareStatus     { return protocol.Software_PreAlpha }
func (s *serveWWWService) ReferenceURI() string                { return "" }
func (s *serveWWWService) IssueDate() protocol.Time            { return nil } // 1587282740
func (s *serveWWWService) ExpiryDate() protocol.Time           { return nil }
func (s *serveWWWService) ExpireInFavorOf() protocol.MediaType { return nil }
func (s *serveWWWService) Fields() []protocol.Field            { return nil }

//libgo:impl protocol.Service
func (s *serveWWWService) URI() string                 { return "" }
func (s *serveWWWService) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *serveWWWService) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *serveWWWService) CRUDType() protocol.CRUD     { return protocol.CRUD_All }
func (s *serveWWWService) UserType() protocol.UserType { return protocol.UserType_All }

// ServeWWW will serve WWW assets to request
func (s *serveWWWService) ServeHTTP(stream protocol.Stream, httpReq *http.Request, httpRes *http.Response) (err protocol.Error) {
	var reqFile, _ = s.WWW.GUI.FileByPath(httpReq.URI().Path())
	if reqFile == nil {
		// TODO::: SSR to serve-to-robots
		// TODO::: Have default error pages and can get customizes!
		// Send beauty HTML response in http error situation like 500, 404, ...

		const supportedLang = "en" // TODO::: get from header
		reqFile, err = s.WWW.MainHTMLDir.File(supportedLang)
		// if err != nil {
		// TODO::: check other user language and at the end send better error
		// }
	}
	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.H.Set(http.HeaderKeyCacheControl, "max-age=31536000, immutable")
	httpRes.SetBody(reqFile.Data())
	return
}
