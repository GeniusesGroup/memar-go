/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"encoding/hex"
	"strconv"
	"strings"

	"../http"
)

// Indicate standard listen and send port number register for http protocol
const (
	ProtocolPortHTTPReceive uint16 = 80
	ProtocolPortHTTPSend    uint16 = 81
)

// HTTPHandler use to standard HTTP handlers in any layer!
type HTTPHandler func(*Server, *Stream, *http.Request, *http.Response)

// HTTPIncomeRequestHandler handle incoming HTTP request streams!
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func HTTPIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()

	err = req.UnMarshal(st.Payload)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		res.SetError(err)
		HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
		return
	}

	// TODO::: due to IPv4&&IPv6 support we must handle some functionality here! Remove them when remove those support.
	// HTTP may transmit over TCP||UDP and we can't make full detail connection in that protocols!!
	if st.Connection == nil {
		// First try to find cookie on http header!
		var cookies = req.Header.GetCookies()
		if len(cookies) > 0 {
			// Default Achaemenid HTTPS handler just support one cookie!
			var cookie = cookies[0]
			var acid, err = hex.DecodeString(cookie.Value)
			if err == nil {
				var acidArray [16]byte
				copy(acidArray[:], acid)
				st.Connection = s.Connections.GetConnectionByID(acidArray)
			}
		}
		if st.Connection == nil {
			// If can't find or get exiting connection from local and remote, make new one if server allow!
			err = makeNewGuestConnection(s, st)
			if err != nil {
				// Can't make new guest connection almost due to lack of enough resources, So simply return to close the connection
				return
			}

			res.Header.SetSetCookies([]http.SetCookie{
				http.SetCookie{
					Name:     "ACID", // "Achaemenid Connection ID"
					Value:    hex.EncodeToString(st.Connection.ID[:]),
					MaxAge:   strconv.FormatUint(365*24*60*60, 10),
					Secure:   true,
					HTTPOnly: true,
					SameSite: "Lax",
				},
			})
		}

		// This header just need in IP connections, so add here not globally in HTTPOutcomeResponseHandler()
		res.Header.SetValue(http.HeaderKeyConnection, http.HeaderValueKeepAlive)
		res.Header.SetValue(http.HeaderKeyKeepAlive, "timeout=60")
		res.Header.SetValue(http.HeaderKeyStrictTransportSecurity, "max-age=31536000")

		st.Connection.RegisterStream(st)
	}

	var service *Service
	// Find related services
	if req.URI.Path == "/apis" {
		if req.Method != http.MethodPOST {
			res.SetStatus(http.StatusMethodNotAllowedCode, http.StatusMethodNotAllowedPhrase)
			HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
			return
		}

		var id uint64
		id, err = strconv.ParseUint(req.URI.Query, 10, 32)
		if err == nil {
			st.ServiceID = uint32(id)
			service = s.Services.GetServiceHandlerByID(st.ServiceID)
		}
		// Add some header for /apis like not index by SE(google, ...), ...
		res.Header.SetValue("X-Robots-Tag", "noindex")
	} else if req.URI.Path == "/objects" {
		var file = s.Assets.Objects.GetFile(req.URI.Query)
		if file == nil {
			file = s.Assets.WWW.GetFile("404.html")
			res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
		} else {
			res.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
		}
		res.Header.SetValue(http.HeaderKeyContentType, file.MimeType)
		res.Body = file.Data
		HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
		return
	} else {
		// First check of request send over non www subdomain
		var host = req.GetHost()
		// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
		if len(host) > 4 && !(host[:4] == "www.") && !('0' <= host[0] && host[0] <= '9') {
			host = "www." + host
			var target = "https://" + host + req.URI.Path
			if len(req.URI.Query) > 0 {
				target += "?" + req.URI.Query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
			}
			res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
			res.Header.SetValue(http.HeaderKeyLocation, target)
			res.Header.SetValue(http.HeaderKeyCacheControl, "max-age=2592000")
			HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
			return
		}

		// Route by URL
		service = s.Services.GetServiceHandlerByURI(req.URI.Path)

		// Route by WWW assets
		if service == nil {
			var path = strings.Split(req.URI.Path, "/")
			var lastPath = path[len(path)-1]

			var file = s.Assets.WWW.GetFile(lastPath)
			if file == nil {
				file = s.Assets.WWW.GetFile("main.html")
			}

			// TODO::: serve-to-robots

			res.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
			res.Header.SetValue(http.HeaderKeyContentType, file.MimeType)
			res.Body = file.Data
			HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
			return
		}
	}
	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
	if service == nil {
		st.Connection.ServiceCallCount++
		st.Connection.FailedServiceCall++
		// Attack??
		res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
		res.SetError(ErrHTTPServiceNotFound)
	} else {
		// Handle request stream
		service.HTTPHandler(s, st, req, res)
		if st.ReqRes.Err != nil {
			st.Connection.FailedPacketsReceived++
			st.Connection.FailedServiceCall++
			res.SetError(st.ReqRes.Err)
			// Attack??
		}
	}
	HTTPOutcomeResponseHandler(s, st.ReqRes, req, res)
}

// HTTPIncomeResponseHandler use to handle incoming HTTP response streams!
func HTTPIncomeResponseHandler(s *Server, st *Stream) {
	st.SetState(StreamStateReady)
}

// HTTPOutcomeRequestHandler use to handle outcoming HTTP request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func HTTPOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	err = st.SendReq()
	return
}

// HTTPOutcomeResponseHandler use to handle outcoming HTTP response stream!
func HTTPOutcomeResponseHandler(s *Server, st *Stream, req *http.Request, res *http.Response) (err error) {
	// Close request stream!
	st.Connection.CloseStream(st.ReqRes)

	// Do some global assignment to response
	res.Version = req.Version
	// res.Header.SetValue(http.HeaderKeyAccessControlAllowOrigin, "*")
	res.Header.SetValue(http.HeaderKeyContentLength, strconv.FormatUint(uint64(len(res.Body)), 10))
	// Add Server Header to response : "Achaemenid"
	res.Header.SetValue(http.HeaderKeyServer, http.DefaultServer)

	st.Payload = res.Marshal()

	// send stream to outcome pool by weight
	err = st.Send()
	// TODO::: handle server error almost due to no network available or connection closed!

	return
}

// TODO::: Have default error pages and can get customizes!
// Send beauty HTML response in http error situation like 500, 404, ...

// HTTPToHTTPSHandler handle incoming HTTP request streams to redirect to HTTPS!
// TODO::: remove this when remove support of IP!
func HTTPToHTTPSHandler(s *Server, st *Stream) {
	var err error
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()
	err = req.UnMarshal(st.Payload)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
	} else {
		var host = req.GetHost()
		// Add www to domain due to we just support http on www server app!
		if len(host) >= 4 && host[0:4] != "www." {
			host = "www." + host
		}

		// redirect http to https
		// remove/add not default ports from req.Host
		var target = "https://" + host + req.URI.Path
		if len(req.URI.Query) > 0 {
			target += "?" + req.URI.Query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
		res.Header.SetValue(http.HeaderKeyLocation, target)
	}

	// Do some global assignment to response
	res.Version = req.Version
	res.Header.SetValue(http.HeaderKeyContentLength, "0")
	// Add cache to decrease server load
	res.Header.SetValue(http.HeaderKeyCacheControl, "public, max-age=2592000")
	// Add Server Header to response : "Achaemenid"
	res.Header.SetValue(http.HeaderKeyServer, http.DefaultServer)

	st.ReqRes.Payload = res.Marshal()
}
