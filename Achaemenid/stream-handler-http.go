/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net/url"
	"strings"

	"../http"
)

// ProtocolPortHTTP indicate standard port number register for http protocol
const (
	ProtocolPortHTTP uint16 = 80
)

// HTTPIncomeRequestHandler handle incoming HTTP request streams!
// We just support http to https redirect!
func HTTPIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()
	err = req.UnMarshal(st.Payload)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		goto End
	}

	{
		// Encode URL to route request
		var url *url.URL
		url, err = req.GetURI()
		if err != nil {
			st.Connection.FailedPacketsReceived++
			res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
			goto End
		}

		var host = req.Header.GetHost(url)
		// Add www to domain due to we just support http on www server app!
		if !strings.HasPrefix(host, "www.") {
			host = "www." + host
		}

		// redirect http to https
		// remove/add not default ports from req.Host
		var target = "https://" + host + url.Path
		if len(url.RawQuery) > 0 {
			target += "?" + url.RawQuery // + "&rd=tls" // TODO::: add rd query for analysis purpose??
		}
		res.SetStatus(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase)
		res.Header.SetValue(http.HeaderKeyLocation, target)
	}

End:
	// Do some global assignment to response
	res.Version = req.Version
	res.Header.SetValue(http.HeaderKeyContentLength, "0")
	// Add cache to decrease server load
	res.Header.SetValue(http.HeaderKeyCacheControl, "public, max-age=2592000")
	// Add Server Header to response : "Achaemenid"
	res.Header.SetValue(http.HeaderKeyServer, http.DefaultServer)

	st.ReqRes.Payload = res.Marshal()
}
