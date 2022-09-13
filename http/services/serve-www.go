/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../mediatype"
	"../protocol"
	"../service"
	"../www"
)

var ServeWWWService = serveWWWService{
	Service: service.New("", mediatype.New("domain/http.protocol.service; name=serve-www").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Serve WWW",
		"",
		"",
		"",
		nil).SetInfo(protocol.Software_PreAlpha, 1587282740, "")).
		SetAuthorization(protocol.CRUDAll, protocol.UserType_All),
}

type serveWWWService struct {
	*service.Service
	WWW www.Assets
}

// ServeWWW will serve WWW assets to request
func (ser *serveWWWService) ServeHTTP(stream protocol.Stream, httpReq *Request, httpRes *Response) (err protocol.Error) {
	var reqFile, _ = ser.WWW.GUI.FileByPath(httpReq.uri.path)
	if reqFile == nil {
		// TODO::: SSR to serve-to-robots
		// TODO::: Have default error pages and can get customizes!
		// Send beauty HTML response in http error situation like 500, 404, ...

		const supportedLang = "en" // TODO::: get from header
		reqFile, err = ser.WWW.MainHTMLDir.File(supportedLang)
		// if err != nil {
		// TODO::: check other user language and at the end send better error
		// }
	}
	httpRes.SetStatus(StatusOKCode, StatusOKPhrase)
	httpRes.H.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
	httpRes.SetBody(reqFile.Data())
	return
}
