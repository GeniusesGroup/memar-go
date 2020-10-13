/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Indicate standard listen and send port number register for DNS protocol.
const (
	ProtocolPortDNSReceive uint16 = 50
	ProtocolPortDNSSend    uint16 = 51
)

// DNSHandler use to standard services handlers in any layer!
type DNSHandler func(*Stream)

// DNSIncomeRequestHandler handle incoming DNS request streams!
func DNSIncomeRequestHandler(s *Server, st *Stream) {

	// Handle response stream
	DNSOutcomeResponseHandler(s, st)
}

// DNSIncomeResponseHandler use to handle incoming DNS response streams!
func DNSIncomeResponseHandler(s *Server, st *Stream) {

}

// DNSOutcomeRequestHandler use to handle outcoming DNS request stream!
func DNSOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	return
}

// DNSOutcomeResponseHandler use to handle outcoming DNS response stream!
func DNSOutcomeResponseHandler(s *Server, st *Stream) (err error) {
	return
}
