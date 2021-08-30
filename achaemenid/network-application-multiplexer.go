/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"../protocol"
	"../log"
)

// netAppMux (Network Application Multiplexer) and its methods act as multiplexer and route income packet to desire protocol handler!
type netAppMux struct {
	handlers [65536]protocol.NetworkApplicationHandler // TODO::: It use 256||512 KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
	lastUse  uint16
}

// SetNetworkApplicationHandler use to set or change specific handler!
func (nam *netAppMux) SetNetworkApplicationHandler(protocolID protocol.NetworkApplicationProtocolID, nah protocol.NetworkApplicationHandler) {
	if nam.handlers[protocolID] != nil {
		log.Warn("Protocol handler with ID: ", protocolID, ", register before, Double check for any mistake to prevent unexpected behavior")
	}
	nam.handlers[protocolID] = nah
	return
}

// GetNetworkApplicationHandler return the protocol handler for given ID!
func (nam *netAppMux) GetNetworkApplicationHandler(protocolID protocol.NetworkApplicationProtocolID) (nah protocol.NetworkApplicationHandler) {
	nah = nam.handlers[protocolID]
	if nah == nil {
		// TODO::: return not exist handler instead of nil handler
	}
	return
}

// DeleteNetworkApplicationHandler delete the handler from handlers.
func (nam *netAppMux) DeleteNetworkApplicationHandler(protocolID protocol.NetworkApplicationProtocolID) {
	nam.handlers[protocolID] = nil
}

// GetFreeProtocolID use to get free portID.
// usually use to start connection to other servers in random other than standards ports number.
func (nam *netAppMux) GetFreeProtocolID() (protocolID uint16) {
	for i := 0; i < 65536; i++ {
		nam.lastUse++
		if nam.handlers[nam.lastUse] == nil {
			return nam.lastUse
		}
	}
	return 0
}
