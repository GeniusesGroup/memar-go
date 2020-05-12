/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	gp "../GP"
)

// HandleGP use to handle GP with any application protocol and response just some basic data!
func (s *Server) HandleGP(osIncomePacket []byte) {
	// Don't need to check packet here due Router layer must always check and penalty other router or XP
	// But it can panic server due to memory overflow, so decide to check or not!
	// gp.CheckPacket()

	// Check server supported requested protocol
	var protocolID uint16 = gp.GetDestinationAppProtocol(osIncomePacket)
	var protocolHandler StreamHandler = s.ProtocolsHandlers.GetProtocolHandler(protocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	var ok bool
	var err error
	var conn *Connection
	var peerGP = gp.GetSourceGP(osIncomePacket)
	// Find Connection from ConnectionPoolByPeerGP by requester GP
	conn, ok = s.Connections.GetConnectionByPeerGP(peerGP)
	// If it is first time that ConnectionID used
	if !ok {
		conn, err = s.Connections.MakeNewConnectionByPeerGP(peerGP)
		if err != nil {
			// Send response or just ignore packet
			// TODO::: DDOS!!??
			return
		}
		conn.PacketPayloadSize = gp.GetPayloadLength(osIncomePacket)
	}

	conn.PacketsReceived++
	conn.BytesReceived += uint64(gp.GetPayloadLength(osIncomePacket))

	// Decrypt packet!
	err = gp.Decrypt(osIncomePacket, conn.Cipher)
	if err != nil {
		conn.FailedPacketsReceived++
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	var st *Stream
	var streamID uint32 = gp.GetStreamID(osIncomePacket)

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

	var packetID uint32 = gp.GetPacketID(osIncomePacket)

	// add payload to Stream payload!
	err = st.AddNewGPPacket(gp.GetPayload(osIncomePacket), packetID)

	// Check TimeSensitive or stream ready to call requested app protocol to process stream.
	if (st.TimeSensitive && err != ErrPacketArrivedPosterior) || (st.Status == StreamStateReady) {
		protocolHandler(s, st)
	}
}
