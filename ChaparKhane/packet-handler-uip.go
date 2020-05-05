/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"../uip"
)

// HandleUIP use to handle UIP with any application protocol and response just some basic data!
func (s *Server) HandleUIP(osIncomePacket []byte) {
	// Don't need to check packet here due Router layer must always check and penalty other router or XP
	// But it can panic server due to memory overflow, so decide to check or not!
	// uip.CheckPacket()

	// Check server supported requested protocol
	var protocolID uint16 = uip.GetDestinationAppProtocol(osIncomePacket)
	var protocolHandler StreamHandler = s.ProtocolsHandlers.GetProtocolHandler(protocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	var ok bool
	var err error
	var conn *Connection
	var peerUIP = uip.GetSourceUIP(osIncomePacket)
	// Find Connection from ConnectionPoolByPeerUIP by requester UIP
	conn, ok = s.Connections.GetConnectionByPeerUIP(peerUIP)
	// If it is first time that ConnectionID used
	if !ok {
		conn, err = s.Connections.MakeNewConnectionByPeerUIP(peerUIP)
		if err != nil {
			// Send response or just ignore packet
			// TODO::: DDOS!!??
			return
		}
		conn.PacketPayloadSize = uip.GetPayloadLength(osIncomePacket)
	}

	conn.PacketsReceived++
	conn.BytesReceived += uint64(uip.GetPayloadLength(osIncomePacket))

	// Decrypt packet!
	err = uip.Decrypt(osIncomePacket, conn.Cipher)
	if err != nil {
		conn.FailedPacketsReceived++
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	var st *Stream
	var streamID uint32 = uip.GetStreamID(osIncomePacket)

	st, ok = conn.GetStreamByID(streamID)
	if !ok {
		st, err = conn.MakeUnidirectionalStream(streamID)
		if err != nil {
			conn.FailedServiceCall++
			conn.FailedPacketsReceived++
			// Send response or just ignore stream
			// TODO::: DDOS!!??
			return
		}
		st.ProtocolID = protocolID
	}

	var packetID uint32 = uip.GetPacketID(osIncomePacket)

	// add payload to Stream payload!
	err = st.AddNewUIPPacket(uip.GetPayload(osIncomePacket), packetID)

	// Check TimeSensitive or stream ready to call requested app protocol to process stream.
	if (st.TimeSensitive && err != ErrPacketArrivedPosterior) || (st.Status == StreamStateReady) {
		protocolHandler(s, st)
	}
}
