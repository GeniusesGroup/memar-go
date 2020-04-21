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
	var protocolHandler ProtocolHandler = s.ProtocolsHandlers.GetProtocolHandler(protocolID)
	if protocolHandler.RequestHandler == nil {
		// Send response or just ignore packet
		// TODO : DDOS!!??
		return
	}

	var ok bool
	var err error
	var conn *Connection
	var st *Stream
	var peerUIP = uip.GetSourceUIP(osIncomePacket)
	// Find Connection from ConnectionPoolByPeerUIP by requester UIP
	conn, ok = s.Connections.GetConnectionByPeerUIP(peerUIP)
	// If it is first time that ConnectionID used
	if !ok {
		conn, err = s.Connections.MakeNewConnectionByPeerUIP(peerUIP)
		if err != nil {
			// Send response or just ignore packet
			// TODO : DDOS!!??
			return
		}
	}

	conn.PacketPayloadSize = uip.GetPayloadSize(osIncomePacket)
	conn.PacketsReceived++
	conn.BytesReceived = conn.BytesReceived + uint64(uip.GetPayloadSize(osIncomePacket))

	// Decrypt data part of packet!
	err = uip.DecryptDataPart(osIncomePacket, conn.FrameSize, conn.EncryptionKey, conn.CipherSuite)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		// Send response or just ignore packet
		// TODO : DDOS!!??
		return
	}

	var streamID uint32 = uip.GetStreamID(osIncomePacket)
	st, ok = conn.GetStreamByID(streamID)
	if !ok {
		st = conn.MakeUnidirectionalStream(streamID)
		st.ProtocolID = protocolID
		// Increment request count for rate limiting
		conn.ServiceCallCount++
	}

	var packetID uint32 = uip.GetPacketID(osIncomePacket)
	// add payload to Stream payload!
	err = st.AddNewUIPPacket(uip.GetPayload(osIncomePacket), packetID)
	if err != nil {
		// TODO::: Can't just ignore err on not TimeSensitive stream and send adversative stream to handler!!
	}

	// Check TimeSensitive or last packet of stream here.
	// Client said stream had been finished in request(4294967294)||response(4294967295) and server must continue process stream, call app protocol requested!
	if (st.TimeSensitive && packetID&1 == 0 && err != ErrPacketArrivedPosterior) || packetID == 4294967294 {
		protocolHandler.RequestHandler(s, st)
	} else if (st.TimeSensitive && packetID&1 == 1 && err != ErrPacketArrivedPosterior) || packetID == 4294967295 {
		protocolHandler.ResponseHandler(s, st)
	}
}
