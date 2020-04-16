/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"../uip"
)

// HandleUIP use to handle UIP with all application protocol and response just some basic data!
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
// HTTP(80,443), DNS(53), SMTP(25,465,587), ...
func (s *Server) HandleUIP(osIncomePacket []byte) {
	// Don't need to check packet here due Router layer must always check and penalty other router or XP
	// But it can panic server due to memory overflow, so decide to check or not!
	// uip.CheckPacket()

	// Check server supported requested protocol
	var protocolID uint16 = uip.GetDestinationAppProtocol(osIncomePacket)
	var protocolHandler StreamHandler = s.ProtocolsHandlers.GetProtocolHandler(protocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO : DDOS!!??
		return
	}

	var ok bool
	var conn *Connection
	var st *Stream
	var peerUIP = uip.GetSourceUIP(osIncomePacket)
	// Find Connection from ConnectionPoolByPeerUIP by requester UIP
	conn, ok = s.Connections.GetConnectionByPeerUIP(peerUIP)
	// If it is first time that ConnectionID used
	if !ok {
		conn = &Connection{
			PeerUIPAddress: peerUIP,
		}
		st = &Stream{
			Connection: conn,
			ProtocolID: protocolID,
			Payload:    uip.GetPayload(osIncomePacket),
		}
		// Data is enough in one frame, Lets check and create connection!
		// Get sRPC handler and ask it to handle request
		protocolHandler = s.ProtocolsHandlers.GetProtocolHandler(0) // TODO: Is it efficient enough to get sRPC handler in each request!?
		protocolHandler(s, st)
		return
	}

	conn.PacketsReceived++
	conn.BytesReceived = conn.BytesReceived + uint64(uip.GetPayloadLength(osIncomePacket))

	// Decrypt data part of packet!
	var err error
	err = uip.DecryptDataPart(osIncomePacket, conn.FrameSize, conn.EncryptionKey, conn.CipherSuite)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		// Send response or just ignore packet
		// TODO : DDOS!!??
		return
	}

	var streamID uint32 = uip.GetStreamID(osIncomePacket)
	st, ok = conn.StreamPool[streamID]
	if !ok {
		// TODO : Check user can open new stream first
		st = &Stream{
			StreamID:   streamID,
			ProtocolID: protocolID,
		}
		conn.RegisterStream(st)
		// Increment request count for rate limiting
		conn.ServiceCallCount++
	}

	// Check TimeSensitive or last packet of stream here.
	// Check PacketID == 4294967295 in sdk for client
	// Send last part of that stream!
	// Client said stream had been finished and server must continue process.
	// call app protocol requested!
	var packetID uint32 = uip.GetPacketID(osIncomePacket)
	if st.TimeSensitive {
		st.LastPacketID = packetID
		st.Payload = uip.GetPayload(osIncomePacket)
		protocolHandler(s, st)
	} else {
		// add payload to StreamPool but it doesn't work due packet may not receive by order!
		copy(st.Payload, uip.GetPayload(osIncomePacket))
		if packetID == 4294967294 {
			protocolHandler(s, st)
			conn.CloseStream(st)
		}
	}
}
