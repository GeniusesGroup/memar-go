/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	gp "../GP"
)

type gpNetwork struct {
	InterfaceID uint8
	MTU         uint16 // Maximum Transmission Unit. GP packet size
	RouterID    uint32
	GPAddr      [14]byte
}

// MakeGPNetwork register app to OS GP router and start handle income GP packets.
func MakeGPNetwork(s *Server) (err error) {
	// send PublicKey to router and get GP if user granted. otherwise log error.
	// n.gp.GPRange = [14]byte{}

	// Get MTU from router
	// n.gp.MTU = 8192 || 1500 || ...

	// Because Achaemenid is server based application must have GP access.

	return
}

// handleGP handle GP packet with any application protocol and response just some basic data!
// Protocol Standard : https://github.com/SabzCity/RFCs/blob/master/GP.md
func handleGP(s *Server, packet []byte) (conn *Connection, st *Stream) {
	// Don't need to check packet here due to ChaparKhane or OS must always check and penalty other routers or societies
	// But it can panic server due to memory overflow, so decide to check or not!
	// gp.CheckPacket()

	// Check server supported requested protocol
	var protocolID uint16 = gp.GetDestinationProtocol(packet)
	var protocolHandler StreamHandler = s.StreamProtocols.GetProtocolHandler(protocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	var err error
	var gpAddr [14]byte = gp.GetSourceAddr(packet)
	// Find Connection from ConnectionPoolByPeerAdd by requester GP
	conn = s.Connections.GetConnectionByPeerGPAdd(gpAddr)
	// If it is first time that user want to connect or longer than server GC old unused connections!
	if conn == nil {
		conn, err = s.Connections.MakeNewConnectionByPeerAdd(gpAddr)
		if err != nil {
			// Send response or just ignore packet
			// TODO::: DDOS!!??
			return
		}
		s.Connections.RegisterConnection(conn)
		conn.PacketPayloadSize = gp.GetPayloadLength(packet) // It's not working due to packet not encrypted yet!
	}

	conn.PacketsReceived++

	// Decrypt packet!
	err = gp.Decrypt(packet, conn.Cipher)
	if err != nil {
		conn.FailedPacketsReceived++
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	conn.SocietyID = gp.GetSourceSociety(packet)
	conn.RouterID = gp.GetSourceRouter(packet)
	conn.BytesReceived += uint64(gp.GetPayloadLength(packet))
	var streamID uint32 = gp.GetStreamID(packet)

	st = conn.StreamPool.GetStreamByID(streamID)
	if st == nil {
		st, err = conn.MakeIncomeStream(streamID)
		if err != nil {
			conn.FailedServiceCall++
			conn.FailedPacketsReceived++
			// Send response or just ignore stream
			// TODO::: DDOS!!??
			return
		}
	}

	var packetID uint32 = gp.GetPacketID(packet)

	// add payload to Stream payload!
	err = addNewGPPacket(st, gp.GetPayload(packet), packetID)

	// Check TimeSensitive or stream ready to call requested app protocol to process stream.
	if (st.Weight == WeightTimeSensitive && err != gp.ErrGPPacketArrivedPosterior) || (st.State == StateReady) {
		protocolHandler(s, st)
		conn.StreamPool.CloseStream(st)
	}

	return
}

// AddNewGPPacket add new GP packet payload to the stream!
func addNewGPPacket(st *Stream, p []byte, packetID uint32) (err error) {
	// Handle packet received not by order
	if packetID < st.LastPacketID {
		st.State = StateBrokenPacket
		err = gp.ErrGPPacketArrivedPosterior
	} else if packetID > st.LastPacketID+1 {
		st.State = StateBrokenPacket
		err = gp.ErrGPPacketArrivedAnterior
		// TODO::: send request to sender about not received packets!!
	} else if packetID+1 == st.LastPacketID {
		st.LastPacketID = packetID
	}
	// TODO::: non of above cover for packet 0||1 drop situation!

	// Use PacketID 0||1 for request||response to set stream settings!
	if packetID < 2 {
		setStreamSettings(st, p)
	} else {
		// TODO::: can't easily copy this way!!
		copy(st.IncomePayload, p)
	}

	// Check stream ready situation!
	if st.TotalPacket == st.PacketReceived {
		st.State = StateReady
	}

	return
}

// Just to show transfer data for setStreamSettings()! We never use this type!
type setStreamSettingsReq struct {
	TotalPacket uint32 // Expected packets count that send over this stream!
	PayloadSize uint64
	Weight      weight
}

// setStreamSettings set stream settings like time sensitive use in VoIP, IPTV, ...
func setStreamSettings(st *Stream, p []byte) {
	// TODO::: allow multiple settings set??

	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
}
