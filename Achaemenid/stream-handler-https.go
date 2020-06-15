/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net/url"
	"strconv"
	"strings"

	"../http"
)

const (
	// ProtocolPortHTTPS indicate standard port number register for https protocol
	ProtocolPortHTTPS uint16 = 443
	// ProtocolPortHTTPDev use for http protocol in development phase without TLS
	ProtocolPortHTTPDev uint16 = 8080
)

// HTTPHandler use to standard HTTP handlers in any layer!
type HTTPHandler func(*Server, *Stream, *http.Request, *http.Response)

// HTTPSIncomeRequestHandler handle incoming HTTP request streams!
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func HTTPSIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	var service *Service
	var uri *url.URL
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()

	err = req.UnMarshal(st.Payload)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		goto End
	}

	// Encode URL to route request
	uri, err = req.GetURI()
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		goto End
	}

	// HTTP may transmit over TCP||UDP and we can't make full detail connection in that protocols!!
	if st.Connection.UserID == [16]byte{} {

	}

	// Find related services
	if uri.Path == "/apis" {
		var id uint64
		id, err = strconv.ParseUint(uri.RawQuery, 10, 32)
		if err == nil {
			st.ServiceID = uint32(id)
			service = s.Services.GetServiceHandlerByID(st.ServiceID)
		}
		// Add some header for /apis like not index by SE(google, ...), ...
		res.Header.SetValue("X-Robots-Tag", "noindex")
	} else {
		// First check of request send over non www subdomain
		var host = req.Header.GetHost(uri)
		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
		if !strings.HasPrefix(host, "www.") && !('0' <= host[0] && host[0] <= '9') {
			host = "www." + host
			var target = "https://" + host + uri.Path
			if len(uri.RawQuery) > 0 {
				target += "?" + uri.RawQuery // + "&rd=tls" // TODO::: add rd query for analysis purpose??
			}
			res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
			res.Header.SetValue(http.HeaderKeyLocation, target)
			res.Header.SetValue(http.HeaderKeyCacheControl, "max-age=2592000")
			goto End
		}

		// Route by URL
		service = s.Services.GetServiceHandlerByURI(uri.Path)

		// Route by assets
		if service == nil {
			var path = strings.Split(uri.Path, "/")
			var lastPath = path[len(path)-1]

			var file = s.Assets.GetFile(lastPath)
			if file == nil {
				file = s.Assets.GetFile("main.html")
			}

			// serve-to-robots
			// serve-index
			// serve-assets

			res.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
			res.Header.SetValue(http.HeaderKeyContentType, file.MimeType)
			res.Body = file.Data
			goto End
		}
	}
	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
	if service == nil {
		err = ErrHTTPServiceNotFound
		st.Connection.ServiceCallCount++
		st.Connection.FailedServiceCall++
		// handle err
		// Attack??

		res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
	} else {
		// Handle request stream
		service.HTTPHandler(s, st, req, res)
	}

End:
	// Do some global assignment to response
	res.Version = req.Version
	// res.Header.SetValue(http.HeaderKeyAccessControlAllowOrigin, "*")
	res.Header.SetValue(http.HeaderKeyContentLength, strconv.FormatUint(uint64(len(res.Body)), 10))
	res.Header.SetValue(http.HeaderKeyConnection, http.HeaderValueKeepAlive)
	res.Header.SetValue(http.HeaderKeyStrictTransportSecurity, "max-age=31536000")
	// Add Server Header to response : "Achaemenid"
	res.Header.SetValue(http.HeaderKeyServer, http.DefaultServer)

	st.ReqRes.Payload = res.Marshal()
}

// HTTPIncomeResponseHandler use to handle incoming HTTP response streams!
func HTTPIncomeResponseHandler(s *Server, st *Stream) {
	// tell request stream that response stream ready to use!
	st.ReqRes.StateChannel <- StreamStateReady
}

// HTTPOutcomeRequestHandler use to handle outcoming HTTP request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func HTTPOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	// send stream to outcome pool by weight
	err = s.SendStream(st)

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus streamState = <-st.StateChannel
	if responseStatus == StreamStateReady {

	} else {

	}

	return
}

// TODO::: Have default error pages and can get customizes!
// Send beauty HTML response in http error situation like 500, 404, ...
