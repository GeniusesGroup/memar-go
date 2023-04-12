/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

// UpperHandlerNonExist use for default and empty switch port due to non of ports can be nil!
type UpperHandlerNonExist struct {
	headerID protocol.NetworkLink_NextHeaderID
}

// RegisterMultiplexer register given multiplexer to the port for further usage!
func (h *UpperHandlerNonExist) HeaderID() (id protocol.NetworkLink_NextHeaderID) { return h.headerID }

// Receive get packet to route it to its path!
func (h *UpperHandlerNonExist) Receive(conn protocol.NetworkLink_Connection, packet []byte) {}

func (h *UpperHandlerNonExist) RegisterAppMultiplexer(tranMux protocol.NetworkTransport_Multiplexer) {
}
func (h *UpperHandlerNonExist) UnRegisterAppMultiplexer(tranMux protocol.NetworkTransport_Multiplexer) {
}

func (h *UpperHandlerNonExist) Shutdown() {}
