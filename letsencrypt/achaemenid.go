/* For license and copyright information please see LEGAL file in repository */

package letsencrypt

import (
	"../achaemenid"
	as "../assets"
	"../http"
	"../log"
)

// CheckByAchaemenid use to check to make||renew server certificate for TLS protocol.
// make and returns new certificate for given domain in s.Manifest.domain + "www." subdomain
// If don't want to reuse any exiting data like ECKey, delete it before run app!
func CheckByAchaemenid(s *achaemenid.Server) (err error) {
	var letsEncrypt = LetsEncrypt{
		SecretFolder:     s.Assets.Secret,
		CommonName:       s.Manifest.Domain,
		EmailAddresses:   []string{s.Manifest.Email},
		Organization:     s.Manifest.Organization,
		OrganizationUnit: []string{"ICT"},
		Domains:          []string{s.Manifest.Domain, "www." + s.Manifest.Domain},
	}

	// First check exiting files if exist!
	err = letsEncrypt.init()
	if err != nil || letsEncrypt.CertificateRequest != nil {
		return
	}

	// Second get certificate from letsencrypt
	// first register http:80 protocol handler
	s.StreamProtocols.SetProtocolHandler(80, httpIncomeRequestHandler)

	// Get exiting LetsEncrypt account or create one
	var acc *as.File
	acc = s.Assets.Secret.GetFile(s.Manifest.Domain + "-letsencrypt-account.json")
	if acc == nil {
		// TODO::: complete proccess and remove below return
		return
	}
	// } else {
	// 	// json.UnMarshal(acc.Data, )
	// }

	log.Warn("Can't get new certificate from LetsEncrypt with this err: ", err)

	// Save to Secret assets
	err = letsEncrypt.saveToSecretAssets()
	// Delete hand;er due to don't need it any more until next 90 days!
	s.StreamProtocols.DeleteHandler(80)

	return
}

// HTTPIncomeRequestHandler handle incoming HTTP request streams use to handle certificate generation proccess!
func httpIncomeRequestHandler(s *achaemenid.Server, st *achaemenid.Stream) {
	var err error
	var req = http.MakeNewRequest()
	var res = http.MakeNewResponse()
	err = req.UnMarshal(st.Payload)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		res.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		goto End
	}

	// handle income request in "/.well-known/acme-challenge/" url

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
