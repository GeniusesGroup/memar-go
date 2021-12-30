/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../compress"
	"../protocol"
	"../service"
	"../www"
)

var ServeWWWService = serveWWWService{
	Service: service.New("urn:giti:http.protocol:service:serve-www", "", protocol.Software_PreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Multiplexer",
			"",
			``,
			[]string{}).
		SetAuthorization(protocol.CRUDAll, protocol.UserTypeAll).Expired(0, ""),
	WWW: www.Assets{
		ContentEncodings: []string{compress.DeflateContentEncoding, compress.GZIPContentEncoding, compress.BrotliContentEncoding},
	},
}

type serveWWWService struct {
	service.Service
	WWW www.Assets
}

// ServeWWW will serve WWW assets to request
func (ser *serveWWWService) ServeHTTP(stream protocol.Stream, httpReq *Request, httpRes *Response) {
	var connection = stream.Connection()

	var reqFile, _ = ser.WWW.GUI.FileByPath(httpReq.uri.path)
	if reqFile == nil {
		// TODO::: SSR to serve-to-robots
		// TODO::: Have default error pages and can get customizes!
		// Send beauty HTML response in http error situation like 500, 404, ...

		const supportedLang = "en" // TODO::: get from header
		reqFile, _ = ser.WWW.MainHTMLDir.File(supportedLang)
	}
	connection.StreamSucceed()
	httpRes.SetStatus(StatusOKCode, StatusOKPhrase)
	httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
	httpRes.SetBody(reqFile.Data())
}
