/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"fmt"
)

// ProtocolsHandlers and its methods act as multiplexer and route income packet to desire protocol handler!
type ProtocolsHandlers struct {
	SingleProtocolServer StreamHandler
	Handlers             [65535]StreamHandler // TODO : It use 256||512KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
}

// SetSingleProtocolServer use to tell server use all port just for one type handler!
// If this handler set other handlers not work and all request will response just with this handler!
func (ph *ProtocolsHandlers) SetSingleProtocolServer(sh StreamHandler) {
	ph.SingleProtocolServer = sh
}

// SetProtocolHandler use to set or change specific handler!
func (ph *ProtocolsHandlers) SetProtocolHandler(protocolID uint16, sh StreamHandler) {
	if ph.Handlers[protocolID] != nil {
		fmt.Print("This protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
	}
	ph.Handlers[protocolID] = sh
}

// GetProtocolHandler use to get specific protocol handler!
func (ph *ProtocolsHandlers) GetProtocolHandler(protocolID uint16) StreamHandler {
	if ph.SingleProtocolServer != nil {
		return ph.SingleProtocolServer
	}
	return ph.Handlers[protocolID]
}
