/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// ProtocolsHandlers and its methods act as multiplexer and route income packet to desire protocol handler!
type ProtocolsHandlers struct {
	Handlers [65535]StreamHandler // TODO : It use 256||512 KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
}

// Init use to register all standard supported protocols!
// Even number for income connections (server - accept connection)
// Odd number for outcome connections (Peer - which start connection)
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
func (ph *ProtocolsHandlers) Init() {
	// SMTP
	// ph.Handlers[25] = smtpHandler
	// DNS
	// ph.Handlers[53] = dnsHandler
	// HTTP
	// ph.Handlers[80] = httpIncomeRequestHandler
	// ph.Handlers[81] = httpIncomeResponseHandler
	// sRPC
	ph.Handlers[0] = SrpcIncomeRequestHandler
	ph.Handlers[1] = SrpcIncomeResponseHandler
}

// SetSingleProtocolServer use to tell server use all port just for one type handler!
// If this handler set other handlers not work and all request will response just with this handler!
func (ph *ProtocolsHandlers) SetSingleProtocolServer(sh StreamHandler) {
	var i uint16
	for i = 0; i < 65535; i++ {
		ph.Handlers[i] = sh
	}
}

// SetProtocolHandler use to set or change specific handler!
func (ph *ProtocolsHandlers) SetProtocolHandler(protocolID uint16, sh StreamHandler) {
	if ph.Handlers[protocolID] != nil {
		Log("This protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
		panic("ChaparKhane occur panic situation due to ^")
	}
	ph.Handlers[protocolID] = sh
}

// GetProtocolHandler use to get specific protocol handler!
func (ph *ProtocolsHandlers) GetProtocolHandler(protocolID uint16) StreamHandler {
	return ph.Handlers[protocolID]
}
