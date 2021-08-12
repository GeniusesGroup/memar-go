/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strings"

	"../convert"
	"../giti"
	"../log"
	"../www"
)

// TODO::: Have default error pages and can get customizes!
// Send beauty HTML response in http error situation like 500, 404, ...

type Handler struct {
	Application giti.Application
	WWW         www.Asset
}

// HandleIncomeRequest handle incoming HTTP request streams!
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func (handler *Handler) HandleIncomeRequest(stream giti.Stream) {
	var err giti.Error
	var connection = stream.Connection()
	var httpReq = NewRequest()
	var httpRes = NewResponse()

	err = httpReq.UnMarshal(stream.IncomeData().Marshal())
	if err != nil {
		connection.ServiceCallFail()
		httpRes.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		httpRes.SetError(err)
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return
	}
	err = handler.ServeHTTP(stream, httpReq, httpRes)
}

// HandleIncomeRequest handle incoming HTTP request streams!
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func (handler *Handler) ServeHTTP(stream giti.Stream, httpReq *Request, httpRes *Response) (err giti.Error) {
	if !giti.AppDevMode && handler.HostCheck(stream, httpReq, httpRes) {
		return
	}

	var connection = stream.Connection()
	var service giti.Service

	// Find related services
	if httpReq.uri.path == "/apis" {
		if httpReq.method != MethodPOST {
			connection.ServiceCallFail()
			httpRes.SetStatus(StatusMethodNotAllowedCode, StatusMethodNotAllowedPhrase)
			handler.HandleOutcomeResponse(stream, httpReq, httpRes)
			return
		}

		var serviceID uint64
		serviceID, err = convert.Base10StringToUint64(httpReq.uri.query)
		if err == nil {
			service = handler.Application.GetServiceByID(serviceID)
		}
		// Add some header for /apis like not index by SE(google, ...), ...
		httpRes.header.Set("X-Robots-Tag", "noindex")
		// httpRes.header.Set(HeaderKeyCacheControl, "no-store")
	} else {
		// Route by URL
		service = handler.Application.GetServiceByURI(httpReq.uri.path)

		// Route by WWW assets
		if service == nil {
			var path = strings.Split(httpReq.uri.path, "/")
			var lastPath = path[len(path)-1]

			var file = handler.WWW.Folder.GetFile(lastPath)
			if file == nil && strings.IndexByte(lastPath, '.') == -1 {
				// TODO::: SSR to serve-to-robots
				file = handler.WWW.Main
			}

			if file == nil {
				connection.ServiceCallFail()
				httpRes.SetStatus(StatusNotFoundCode, StatusNotFoundPhrase)
			} else {
				connection.ServiceCalled()
				httpRes.SetStatus(StatusOKCode, StatusOKPhrase)
				httpRes.header.Set(HeaderKeyContentType, file.MimeType)
				httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
				httpRes.header.Set(HeaderKeyContentEncoding, file.CompressType)
				httpRes.Body = file.CompressData
			}

			handler.HandleOutcomeResponse(stream, httpReq, httpRes)
			return
		}
	}

	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
	if service == nil {
		connection.ServiceCallFail()
		httpRes.SetStatus(StatusNotFoundCode, StatusNotFoundPhrase)
		httpRes.SetError(ErrNotFound)
	} else {
		stream.SetService(service)
		err = service.ServeHTTP(stream, httpReq, httpRes)
		if err != nil {
			connection.ServiceCallFail()
			stream.SetError(err)
			httpRes.SetError(err)
		}

		connection.ServiceCalled()
	}

	handler.HandleOutcomeResponse(stream, httpReq, httpRes)
	return
}

// HandleIncomeResponse use to handle incoming HTTP response streams!
func (handler *Handler) HandleIncomeResponse(stream giti.Stream) {
	stream.SetState(giti.ConnectionStateReady)
}

// HandleOutcomeRequest use to handle outcoming HTTP request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func (handler *Handler) HandleOutcomeRequest(stream giti.Stream) (err giti.Error) {
	// err = stream.Send()
	return
}

// HandleOutcomeResponse use to handle outcoming HTTP response stream!
func (handler *Handler) HandleOutcomeResponse(stream giti.Stream, httpReq *Request, httpRes *Response) {
	stream.Connection().CloseStream(stream)

	// Do some global assignment to response
	httpRes.version = httpReq.version
	// httpRes.header.Set(HeaderKeyAccessControlAllowOrigin, "*")
	httpRes.SetContentLength()
	// Add Server Header to response : "Achaemenid"
	httpRes.header.Set(HeaderKeyServer, DefaultServer)

	stream.SetOutcomeData(httpRes)

	if giti.AppDeepDebugMode {
		// TODO::: body not serialized yet to log it!! any idea to have better performance below??
		log.DeepDebug("HTTP - Request:::", httpReq.uri.raw, httpReq.header, string(httpReq.body.Marshal()))
		log.DeepDebug("HTTP - Response:::", httpRes.ReasonPhrase, httpRes.header, string(httpRes.body.Marshal()))
	}
}

// HTTPtoHTTPS handle incoming HTTP request to redirect to HTTPS!
func (handler *Handler) HTTPtoHTTPS(stream giti.Stream, httpReq *Request, httpRes *Response) {
	// redirect http to https
	// remove/add not default ports from httpReq.Host
	var target = "https://" + httpReq.URI().Host() + httpReq.URI().Path()
	var targetQuery = httpReq.URI().Query()
	if len(targetQuery) > 0 {
		target += "?" + targetQuery // + "&rd=tls" // TODO::: add rd query for analysis purpose??
	}
	httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
	httpRes.Header().Set(HeaderKeyLocation, target)
	httpRes.Header().Set(HeaderKeyConnection, HeaderValueClose)
	// Add cache to decrease server load
	httpRes.Header().Set(HeaderKeyCacheControl, "public, max-age=2592000")

	// Do some global assignment to response
	httpRes.version = httpReq.version
	httpRes.Header().Set(HeaderKeyContentLength, "0")
	// Add Server Header to response : "Achaemenid"
	httpRes.Header().Set(HeaderKeyServer, DefaultServer)
}

func (handler *Handler) HostCheck(stream giti.Stream, httpReq *Request, httpRes *Response) (redirect bool) {
	var domainName = handler.Application.Manifest().DomainName()
	var host = httpReq.uri.host
	var path = httpReq.uri.path
	var query = httpReq.uri.query

	if host == "" {
		// TODO::: noting to do or reject request??
	} else if '0' <= host[0] && host[0] <= '9' {
		// check of request send over IP
		if giti.AppDeepDebugMode {
			log.Debug("HTTP - Host Check - IP host:", host)
		}

		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return true
	} else if len(host) > 4 && host[:4] == "www." {
		if host[4:] != domainName {
			if giti.AppDeepDebugMode {
				log.Debug("HTTP - Host Check - Unknown WWW host:", host)
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return true
		}

		if giti.AppDeepDebugMode {
			log.Debug("HTTP - Host Check - WWW host:", host)
		}

		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
		var target = "https://" + domainName + path
		if len(query) > 0 {
			target += "?" + query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		httpRes.SetStatus(StatusMovedPermanentlyCode, StatusMovedPermanentlyPhrase)
		httpRes.header.Set(HeaderKeyLocation, target)
		httpRes.header.Set(HeaderKeyCacheControl, "max-age=31536000, immutable")
		handler.HandleOutcomeResponse(stream, httpReq, httpRes)
		return true
	} else if host != domainName {
		if giti.AppDeepDebugMode {
			log.Debug("HTTP - Host Check - Unknown host:", host)
		}
		// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
		return true
	}
	return
}
