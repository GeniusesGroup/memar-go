/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// ProtocolsHandlers and its methods act as multiplexer and route income packet to desire protocol handler!
type ProtocolsHandlers struct {
	Handlers [65535]ProtocolHandler // TODO : It use 512||1024 KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
}

// ProtocolHandler store handler for specific protocol!
type ProtocolHandler struct {
	RequestHandler  StreamHandler
	ResponseHandler StreamHandler
}

// Init use to register all standard supported protocols!
// Odd number for income connections (server - accept connection)
// Even number for outcome connections (Peer - which start connection)
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
func (phs *ProtocolsHandlers) Init() {
	// SMTP
	// phs.Handlers[25] = smtpHandler
	// DNS
	// phs.Handlers[53] = dnsHandler
	// HTTP
	// phs.Handlers[443] = httpProtocolHandler
	// phs.Handlers[444] = httpProtocolHandler
	// sRPC
	phs.Handlers[0] = sRPCProtocolHandler
	phs.Handlers[1] = sRPCProtocolHandler
}

// SetSingleProtocolServer use to tell server use all port just for one type handler!
// If this handler set other handlers not work and all request will response just with this handler!
func (phs *ProtocolsHandlers) SetSingleProtocolServer(ph ProtocolHandler) {
	var i uint16
	for i = 0; i < 65535; i++ {
		phs.Handlers[i] = ph
	}
}

// SetProtocolHandler use to set or change specific handler!
func (phs *ProtocolsHandlers) SetProtocolHandler(protocolID uint16, ph ProtocolHandler) {
	if phs.Handlers[protocolID].RequestHandler != nil {
		Log("This protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
		panic("ChaparKhane occur panic situation due to ^")
	}
	phs.Handlers[protocolID] = ph
}

// GetProtocolHandler use to get specific protocol handler!
func (phs *ProtocolsHandlers) GetProtocolHandler(protocolID uint16) ProtocolHandler {
	return phs.Handlers[protocolID]
}
