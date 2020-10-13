/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"encoding/hex"
	"strconv"
	"strings"

	"../http"
	"../log"
)

// Indicate standard listen and send port number register for http protocol
const (
	ProtocolPortHTTPReceive uint16 = 80
	ProtocolPortHTTPSend    uint16 = 81
)

// HTTPHandler use to standard HTTP handlers in any layer!
type HTTPHandler func(*Stream, *http.Request, *http.Response)

// HTTPIncomeRequestHandler handle incoming HTTP request streams!
// It can use for architectures like restful, ...
// Protocol Standard - http2 : https://httpwg.org/specs/rfc7540.html
func HTTPIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()

	err = req.UnMarshal(st.IncomePayload)
	if err != nil {
		if st.Connection == nil {
			// TODO::: due to IPv4&&IPv6 support we must handle some functionality here! Remove it when remove those support.
			// TODO::: attack??
			return
		}
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		res.SetError(err)
		st.Connection.FailedPacketsReceived++
		HTTPOutcomeResponseHandler(s, st, req, res)
		return
	}

	// TODO::: due to IPv4&&IPv6 support we must handle some functionality here! Remove them when remove those support.
	if st.Connection == nil {
		// HTTP may transmit over TCP||UDP and we can't make full detail connection in that protocols!!
		err = makeNewConnectionByHTTP(s, st, req, res)
		if err != nil {
			return
		}
	}

	if hostCheck(s, st, req, res) {
		return
	}

	// Find related services
	if req.URI.Path == "/apis" {
		if req.Method != http.MethodPOST {
			res.SetStatus(http.StatusMethodNotAllowedCode, http.StatusMethodNotAllowedPhrase)
			HTTPOutcomeResponseHandler(s, st, req, res)
			return
		}

		var id uint64
		id, err = strconv.ParseUint(req.URI.Query, 10, 32)
		if err == nil {
			st.Service = s.Services.GetServiceHandlerByID(uint32(id))
		}
		// Add some header for /apis like not index by SE(google, ...), ...
		res.Header.Set("X-Robots-Tag", "noindex")

		// res.Header.Set(http.HeaderKeyContentEncoding, "gzip")
		// var b bytes.Buffer
		// var gz = gzip.NewWriter(&b)
		// gz.Write(res.Body)
		// gz.Close()
		// res.Body = b.Bytes()
	} else if req.URI.Path == "/objects" {
		var file = s.Assets.Objects.GetFile(req.URI.Query)
		if file == nil {
			file = s.Assets.WWW.GetFile("404.html")
			res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
		} else {
			res.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
			res.Header.Set(http.HeaderKeyCacheControl, "max-age=31536000")
			res.Header.Set(http.HeaderKeyContentType, file.MimeType)
			res.Header.Set(http.HeaderKeyContentEncoding, file.CompressType)
			res.Body = file.CompressData
		}
		HTTPOutcomeResponseHandler(s, st, req, res)
		return
	} else {
		// Route by URL
		st.Service = s.Services.GetServiceHandlerByURI(req.URI.Path)

		// Route by WWW assets
		if st.Service == nil {
			var path = strings.Split(req.URI.Path, "/")
			var lastPath = path[len(path)-1]

			var file = s.Assets.WWW.GetFile(lastPath)
			if file == nil && strings.IndexByte(lastPath, '.') == -1 {
				// TODO::: serve-to-robots
				file = s.Assets.WWWMain
			}

			if file == nil {
				res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
			} else {
				res.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
				res.Header.Set(http.HeaderKeyContentType, file.MimeType)
				res.Header.Set(http.HeaderKeyCacheControl, "max-age=31536000")
				res.Header.Set(http.HeaderKeyContentEncoding, file.CompressType)
				res.Body = file.CompressData
			}

			HTTPOutcomeResponseHandler(s, st, req, res)
			return
		}
	}

	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
	if st.Service == nil {
		st.Connection.ServiceCallCount++
		st.Connection.FailedServiceCall++
		// Attack??
		res.SetStatus(http.StatusNotFoundCode, http.StatusNotFoundPhrase)
		res.SetError(http.ErrHTTPNotFound)
	} else {
		// Handle request stream
		st.Service.HTTPHandler(st, req, res)
		if st.Err != nil {
			st.Connection.ServiceCallCount++
			st.Connection.FailedServiceCall++
			res.SetError(st.Err)
			// Attack??
		}
	}
	HTTPOutcomeResponseHandler(s, st, req, res)
}

// HTTPIncomeResponseHandler use to handle incoming HTTP response streams!
func HTTPIncomeResponseHandler(s *Server, st *Stream) {
	st.SetState(StateReady)
}

// HTTPOutcomeRequestHandler use to handle outcoming HTTP request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func HTTPOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	err = st.SendAndWait()
	return
}

// HTTPOutcomeResponseHandler use to handle outcoming HTTP response stream!
func HTTPOutcomeResponseHandler(s *Server, st *Stream, req *http.Request, res *http.Response) (err error) {
	st.Connection.StreamPool.CloseStream(st)

	// Do some global assignment to response
	res.Version = req.Version
	// res.Header.Set(http.HeaderKeyAccessControlAllowOrigin, "*")
	res.Header.Set(http.HeaderKeyContentLength, strconv.FormatUint(uint64(len(res.Body)), 10))
	// Add Server Header to response : "Achaemenid"
	res.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	st.OutcomePayload = res.Marshal()

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
	err = req.UnMarshal(st.IncomePayload)
	if err != nil {
		// TODO::: attack??
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
	} else {
		// redirect http to https
		// remove/add not default ports from req.Host
		var target = "https://" + req.GetHost() + req.URI.Path
		if len(req.URI.Query) > 0 {
			target += "?" + req.URI.Query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
		res.Header.Set(http.HeaderKeyLocation, target)
		res.Header.Set(http.HeaderKeyConnection, http.HeaderValueClose)
	}

	// Do some global assignment to response
	res.Version = req.Version
	res.Header.Set(http.HeaderKeyContentLength, "0")
	// Add cache to decrease server load
	res.Header.Set(http.HeaderKeyCacheControl, "public, max-age=2592000")
	// Add Server Header to response : "Achaemenid"
	res.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	st.OutcomePayload = res.Marshal()
}

// TODO::: due to IPv4&&IPv6 support need this func! Remove them when remove those support.
func makeNewConnectionByHTTP(s *Server, st *Stream, req *http.Request, res *http.Response) (err error) {
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
		st.Connection, err = s.Connections.MakeNewGuestConnection()
		if err != nil {
			// Can't make new guest connection almost due to lack of enough resources, So simply return to close the connection
			return
		}
		s.Connections.RegisterConnection(st.Connection)

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

		// It is first time user reach platform, So tell the browser always reach by https!
		res.Header.Set(http.HeaderKeyStrictTransportSecurity, "max-age=31536000; includeSubDomains; preload")

	}

	// This header just need in IP connections, so add here not globally in HTTPOutcomeResponseHandler()
	res.Header.Set(http.HeaderKeyConnection, http.HeaderValueKeepAlive)
	res.Header.Set(http.HeaderKeyKeepAlive, "timeout=120")

	st.Connection.StreamPool.RegisterStream(st)

	return
}

// TODO::: due to IPv4&&IPv6 support need this func! Remove them when remove those support.
func hostCheck(s *Server, st *Stream, req *http.Request, res *http.Response) (redirect bool) {
	if !log.DevMode {
		var host = req.GetHost()

		// First check of request send over IP
		if '0' <= host[0] && host[0] <= '9' {
			if log.DebugMode {
				log.Debug("HTTP - IP host:", host)
			}

			// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
			var target = "https://" + s.Manifest.DomainName + req.URI.Path
			if len(req.URI.Query) > 0 {
				target += "?" + req.URI.Query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
			}
			res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
			res.Header.Set(http.HeaderKeyLocation, target)
			res.Header.Set(http.HeaderKeyCacheControl, "max-age=31536000")
			HTTPOutcomeResponseHandler(s, st, req, res)
			return true
		} else if host != s.Manifest.DomainName {
			if log.DebugMode {
				log.Debug("HTTP - Unknown host:", host)
			}
			// TODO::: Silently ignoring a request might not be a good idea and perhaps breaks the RFC's for HTTP.
			return true
		} else if len(host) > 4 && host[:4] == "www." {
			if log.DebugMode {
				log.Debug("HTTP - WWW host:", host)
			}

			// Add www to domain. Just support http on www server app due to SE duplicate content both on www && non-www!
			var target = "https://" + s.Manifest.DomainName + req.URI.Path
			if len(req.URI.Query) > 0 {
				target += "?" + req.URI.Query // + "&rd=tls" // TODO::: add rd query for analysis purpose??
			}
			res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
			res.Header.Set(http.HeaderKeyLocation, target)
			res.Header.Set(http.HeaderKeyCacheControl, "max-age=31536000")
			HTTPOutcomeResponseHandler(s, st, req, res)
			return true
		}
	}
	return
}
