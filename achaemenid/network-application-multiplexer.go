/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"strconv"

	"libgo/log"
	"libgo/protocol"
)

// netAppMux (Network Application Multiplexer) and its methods act as multiplexer and route income packet to desire protocol handler.
type netAppMux struct {
	handlers [65536]protocol.NetworkApplication_Handler // TODO::: It use 256||512 KB of RAM on 32||64bit! Other alternative? map use more than this simple array.
	lastUse  uint16
}

// SetNetworkApplicationHandler use to set or change specific handler.
func (nam *netAppMux) SetNetworkApplicationHandler(nah protocol.NetworkApplication_Handler) {
	var protocolID = nah.ProtocolID()
	if nam.handlers[protocolID] != nil {
		var eventMsg = "Protocol handler with this ID: " + strconv.FormatUint(uint64(protocolID), 10) + ", register before, Double check for any mistake to prevent unexpected behavior"
		protocol.App.Log(log.WarnEvent(&DefaultEvent_MediaType, eventMsg))
	}
	nam.handlers[protocolID] = nah
}

// GetNetworkApplicationHandler return the protocol handler for given ID.
func (nam *netAppMux) GetNetworkApplicationHandler(protocolID protocol.NetworkApplication_ProtocolID) (nah protocol.NetworkApplication_Handler) {
	nah = nam.handlers[protocolID]
	if nah == nil {
		// TODO::: return not exist handler instead of nil handler??
	}
	return
}

// DeleteNetworkApplicationHandler delete the handler from handlers.
func (nam *netAppMux) DeleteNetworkApplicationHandler(protocolID protocol.NetworkApplication_ProtocolID) {
	nam.handlers[protocolID] = nil
}

// GetFreeProtocolID use to get free portID.
// usually use to start connection to other application in random other than standards ports number.
func (nam *netAppMux) GetFreeProtocolID() (protocolID uint16) {
	for i := 0; i < 65536; i++ {
		nam.lastUse++
		if nam.handlers[nam.lastUse] == nil {
			return nam.lastUse
		}
	}
	return 0
}
