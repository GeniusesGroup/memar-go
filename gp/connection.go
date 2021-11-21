/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	"sync/atomic"
	"time"

	"../authorization"
	"../connection"
	"../protocol"
	"../uuid"
)

// Connection indicate the
// https://github.com/balacode/udpt/blob/main/sender.go
type Connection struct {
	writeTime etime.Time

	/* Connection data */
	ID         [32]byte
	StreamPool StreamPool
	State      protocol.ConnectionState
	Weight     protocol.ConnectionWeight
	linkConn   protocol.NetworkLinkConnection
	mtu        int // Maximum Transmission Unit. max GP payload size 

	/* Peer data */
	addr       Addr
	thingID    [32]byte // Use as ConnectionID too!
	domainName string
	// Peer Identifiers
	userID           [32]byte // Can't change on StateLoaded!
	userType         protocol.UserType
	delegateUserID   [32]byte // Can't change on StateLoaded!
	delegateUserType protocol.UserType

	/* Security data */
	AccessControl authorization.AccessControl
	cipher        protocol.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite

	connection.Metric
}

func (conn *Connection) ID() uint32                        { return conn.id }
func (conn *Connection) MTU() int                          { return conn.mtu }
func (conn *Connection) Addr() [16]byte                    { return conn.addr }
func (conn *Connection) AddrType() NetworkLinkNextHeaderID { return protocol.NetworkLinkNextHeaderGP }
func (conn *Connection) ThingID() [32]byte                 { return conn.thingID }
func (conn *Connection) DomainName() string                { return conn.domainName }
func (conn *Connection) UserID() [32]byte                  { return conn.userID }
func (conn *Connection) UserType() UserType                { return conn.userType }
func (conn *Connection) DelegateUserID() [32]byte          { return conn.delegateUserID }
func (conn *Connection) DelegateUserType() UserType        { return conn.delegateUserType }
func (conn *Connection) Cipher() Cipher                    { return conn.cipher }

// SetThingID set thingID only if it is not set before
func (conn *Connection) SetThingID(thingID [32]byte) {
	if conn.ThingID == [32]byte{} {
		conn.ThingID = thingID
	}
}

// Receive use for default and empty switch port due to non of ports can be nil!
func (conn *Connection) Receive(packet []byte) {
	var err protocol.Error

	// TODO::: check packet signature and decrypt it
	// Decrypt packet!
	var frames []byte
	frames, err = Decrypt(GetPayload(packet), conn.Cipher())
	if err != nil {
		conn.FailedPacketsReceived()
		// Send NACK or store and send later
		// TODO::: DDOS!!??
		return
	}

	// TODO::: check packet number and send ACK||NACK frame

	// Metrics data
	// conn.PacketReceived(uint64(len(packet)))
	// conn.PacketPayloadSize = GetPayloadLength(packet) // It's not working due to packet not encrypted yet!

	err = srpconn.HandleFrames(c, frames)
	if err != nil {
		// TODO:::
	}
	return
}

func handleDataFrame(st protocol.Stream, dataFrame []byte, packetID uint32) (err protocol.Error) {
	// Handle packet received not by order
	if packetID < st.LastPacketID {
		st.State = StateBrokenPacket
		err = ErrPacketArrivedPosterior
	} else if packetID > st.LastPacketID+1 {
		st.State = StateBrokenPacket
		err = ErrPacketArrivedAnterior
		// TODO::: send request to sender about not received packets!!
	} else if packetID+1 == st.LastPacketID {
		st.LastPacketID = packetID
	}
	// TODO::: non of above cover for packet 0||1 drop situation!

	// Use PacketID 0||1 for request||response to set stream settings!
	if packetID < 2 {
		// as.SetStreamSettings(st, p)
	} else {
		// TODO::: can't easily copy this way!!
		// copy(st.IncomePayload, p)
	}

	return
}

// MakeIncomeStream make and return the new stream with income ID!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeIncomeStream(streamID uint32) (st protocol.Stream, err protocol.Error) {
	// TODO::: Check user can open new stream first as stream policy!

	// if given streamID is 0, return new incremental streamID from pool
	if streamID == 0 {
		streamID = conn.StreamPool.freeIncomeStreamID
		conn.StreamPool.freeIncomeStreamID += 2
	}

	st = &Stream{
		id:         streamID,
		connection: conn,
		status:     protocol.ConnectionStateOpen,
		state:      make(chan protocol.ConnectionState),
	}
	conn.StreamPool.RegisterStream(st)
	return
}

// MakeOutcomeStream make and return the new stream with outcome ID!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeOutcomeStream(streamID uint32) (st *Stream, err protocol.Error) {
	// TODO::: Check user can open new stream first as stream policy!

	// if given streamID is 0, return new incremental streamID from pool
	if streamID == 0 {
		streamID = conn.StreamPool.freeOutcomeStreamID
		conn.StreamPool.freeOutcomeStreamID += 2
	}

	st = &Stream{
		id:         streamID,
		connection: conn,
		status:     protocol.ConnectionStateOpen,
		state:      make(chan protocol.ConnectionState),
	}
	conn.StreamPool.RegisterStream(st)
	return
}

// MakeSubscriberStream make new Publish–Subscribe stream!
func (conn *Connection) MakeSubscriberStream() (st *Stream) {
	// TODO:::
	return
}

// EstablishNewConnectionByDomainID make new connection by peer domain ID and initialize it!
func EstablishNewConnectionByDomainID(domainID [32]byte) (conn *Connection, err protocol.Error) {
	// TODO::: Get closest domain GP add
	var domainGPAddr = Addr{}
	conn, err = EstablishNewConnectionByPeerAdd(domainGPAddr)
	if err != nil {
		return
	}
	conn.UserID = domainID
	conn.UserType = protocol.UserTypeApp
	return
}

// EstablishNewConnectionByPeerAdd make new connection by peer GP and initialize it!
func EstablishNewConnectionByPeerAdd(gpAddr Addr) (conn *Connection, err protocol.Error) {
	// var userID, thingID [32]byte
	// TODO::: Get peer publickey & userID & thingID from peer GP router!

	// if userID != [32]byte{} {
	// conn = protocol.App.GetConnectionByUserIDThingID(userID, thingID)
	// }

	// If conn not exist means guest connection.
	if conn == nil {
		conn, err = MakeNewGuestConnection()
		if err == nil {
			conn.Addr = gpAddr
			// conn.Cipher = crypto.NewGCM(crypto.NewAES256([32]byte{}))
		}
	}
	return
}

// MakeNewGuestConnection make new connection and register on given stream due to it is first attempt connect to server!
func MakeNewGuestConnection() (conn *Connection, err protocol.Error) {
	// if Server.Manifest.TechnicalInfo.GuestMaxConnections == 0 {
	// 	return nil, ErrGuestConnectionNotAllow
	// } else if Server.Manifest.TechnicalInfo.GuestMaxConnections > 0 && Server.Connections.GuestConnectionCount > Server.Manifest.TechnicalInfo.GuestMaxConnections {
	// 	return nil, ErrGuestConnectionMaxReached
	// }

	conn = &Connection{
		ID:       uuid.Random32Byte(),
		State:    protocol.ConnectionStateNew,
		UserType: protocol.UserTypeGuest,
	}
	conn.AccessControl.GiveFullAccess()
	conn.StreamPool.Init()
	return
}
