/* For license and copyright information please see LEGAL file in repository */

package chapar

import "../giti"

// UpperHandlerNonExist use for default and empty switch port due to non of ports can be nil!
type UpperHandlerNonExist struct {
	headerID byte
}

// RegisterMultiplexer register given multiplexer to the port for further usage!
func (h *UpperHandlerNonExist) HeaderID() (id byte) { return h.headerID }

// Receive get packet to route it to its path!
func (h *UpperHandlerNonExist) Receive(conn giti.NetworkLinkConnection, packet []byte) {}

func (h *UpperHandlerNonExist) RegisterAppMultiplexer(appMux NetworkTransportAppMultiplexer) {}
func (h *UpperHandlerNonExist) UnRegisterAppMultiplexer(appMux NetworkTransportAppMultiplexer) {}

func (h *UpperHandlerNonExist) Shutdown() {}
