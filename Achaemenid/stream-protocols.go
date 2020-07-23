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
// Not rule but suggest to use even port number to listen||receive||response||server and odd for send||request||client
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
// https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml
func (sp *streamProtocols) Init() {
	// sRPC
	sp.handlers[ProtocolPortSRPCReceive] = SrpcIncomeRequestHandler
	sp.handlers[ProtocolPortSRPCSend] = SrpcIncomeResponseHandler

	// SMTP
	// sp.handlers[25] = smtpHandler

	// DNS
	sp.handlers[ProtocolPortDNSReceive] = DNSIncomeRequestHandler
	sp.handlers[ProtocolPortDNSSend] = DNSIncomeResponseHandler

	// HTTP & HTTPS
	sp.handlers[ProtocolPortHTTPReceive] = HTTPIncomeRequestHandler
	sp.handlers[ProtocolPortHTTPSend] = HTTPIncomeResponseHandler
}

// SetSingleProtocol use to tell server use all port just for one type handler!
// If this handler set other handlers not work and all request will response just with this handler!
func (sp *streamProtocols) SetSingleProtocol(sh StreamHandler) {
	for i := 0; i < 65536; i++ {
		sp.handlers[i] = sh
	}
}

// SetProtocolHandler use to set or change specific handler!
func (sp *streamProtocols) SetProtocolHandler(protocolID uint16, sh StreamHandler) {
	if sp.handlers[protocolID] != nil {
		log.Warn("Protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
	}
	sp.handlers[protocolID] = sh
	return
}

// GetProtocolHandler use to get specific protocol handler!
func (sp *streamProtocols) GetProtocolHandler(protocolID uint16) StreamHandler {
	return sp.handlers[protocolID]
}

// DeleteService use to delete specific service in services list.
func (sp *streamProtocols) DeleteHandler(protocolID uint16) {
	sp.handlers[protocolID] = nil
}

// GetFreePortID use to get free portID.
// usually use to start connection to other servers in random other than standards ports number.
func (sp *streamProtocols) GetFreePortID() (protocolID uint16) {
	for i := 0; i < 65536; i++ {
		sp.lastUse++
		if sp.handlers[sp.lastUse] == nil {
			return sp.lastUse
		}
	}
	return 0
}
