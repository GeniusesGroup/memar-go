/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	"../protocol"
)

type AppMultiplexer struct{}

// MakeGPNetwork register app to OS GP router and start handle income GP packets.
func (appMux *AppMultiplexer) Init() (err protocol.Error) {
	protocol.App.LogInfo("GP network begin listening...")
	return
}

func (appMux *AppMultiplexer) HeaderID() protocol.NetworkLinkNextHeaderID {
	return protocol.NetworkLinkNextHeaderGP
}

// Receive handle GP packet with any application protocol and response just some basic data!
// Protocol Standard : https://github.com/GeniusesGroup/RFCs/blob/master/Giti-Network.md
func (appMux *AppMultiplexer) Receive(linkConn protocol.NetworkLinkConnection, packet []byte) {
	var err protocol.Error
	var gpAddr [16]byte = GetSourceAddr(packet)
	// Find Connection from ConnectionPoolByPeerAdd by requester GP
	var conn protocol.Connection = protocol.App.GetConnectionByPeerAddr(gpAddr)
	// If it is first time that user want to connect or longer than server GC old unused connections!
	if conn == nil {
		// conn, err = MakeNewConnectionByPeerAdd(gpAddr, linkConn)
		if err != nil {
			// Send response or just ignore packet
			// TODO::: DDOS!!??
			return
		}
		protocol.App.RegisterConnection(conn)
	}
	conn.(*Connection).Receive(packet)
	return
}

// Shutdown the listener when the application closes or force to closes by not recovered panic!
func (appMux *AppMultiplexer) Shutdown() {
	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
}
