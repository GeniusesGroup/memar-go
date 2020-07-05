/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../log"

// StreamHandler use to standard stream handler in any layer!
type StreamHandler func(*Server, *Stream)

// streamProtocols and its methods act as multiplexer and route income packet to desire protocol handler!
type streamProtocols struct {
	handlers [65536]StreamHandler // TODO::: It use 256||512 KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
	lastUse  uint16
}

// Init use to register all standard supported protocols!
// Usually Dev must register needed stream protocol not use this method to register all available protocols!
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
// https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml
func (sp *streamProtocols) Init() {
	// sRPC
	sp.handlers[ProtocolPortSRPC] = SrpcIncomeRequestHandler
	// SMTP
	// sp.handlers[25] = smtpHandler
	// DNS
	// sp.handlers[53] = dnsHandler
	// HTTP & HTTPS
	sp.handlers[ProtocolPortHTTP] = HTTPIncomeRequestHandler
	sp.handlers[ProtocolPortHTTPS] = HTTPSIncomeRequestHandler
	sp.handlers[ProtocolPortHTTPDev] = HTTPSIncomeRequestHandler
}

// SetSingleProtocol use to tell server use all port just for one type handler!
// If this handler set other handlers not work and all request will response just with this handler!
func (sp *streamProtocols) SetSingleProtocol(sh StreamHandler) {
	for i := 0; i < 65536; i++ {
		sp.handlers[i] = sh
	}
}

// SetProtocolHandler use to set or change specific handler!
// If given protocolID==0, register given handler in free ID and return protocol ID, use to start connection to other servers.
func (sp *streamProtocols) SetProtocolHandler(protocolID uint16, sh StreamHandler) uint16 {
	if protocolID == 0 {
		protocolID = sp.lastUse
		sp.lastUse++
	}
	if sp.handlers[protocolID] != nil {
		log.Warn("Protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
	}
	sp.handlers[protocolID] = sh
	return protocolID
}

// GetProtocolHandler use to get specific protocol handler!
func (sp *streamProtocols) GetProtocolHandler(protocolID uint16) StreamHandler {
	return sp.handlers[protocolID]
}

// DeleteService use to delete specific service in services list.
func (sp *streamProtocols) DeleteHandler(protocolID uint16) {
	sp.handlers[protocolID] = nil
}
