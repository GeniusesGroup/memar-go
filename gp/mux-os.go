/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	"../protocol"
)

type OSMultiplexer struct {
	users          [8]uint16
	appMultiplexer map[uint16]protocol.NetworkTransportMultiplexer // TODO::: map performance penalty!??
}

// MakeGPNetwork register app to OS GP router and start handle income GP packets.
func (osMux *OSMultiplexer) Init() {
	// send user PublicKey to router and get GP if user granted. otherwise log error.

	osMux.appMultiplexer = make(map[uint16]protocol.NetworkTransportMultiplexer, 512)
}

func (osMux *OSMultiplexer) HeaderID() protocol.NetworkLinkNextHeaderID {
	return protocol.NetworkLinkNextHeaderGP
}

// Receive handle GP packet with any application protocol and response just some basic data!
// Protocol Standard : https://github.com/GeniusesGroup/RFCs/blob/master/Giti-Network.md
func (osMux *OSMultiplexer) Receive(conn protocol.NetworkLinkConnection, packet []byte) {
	// ChaparKhane or OS must always check and penalty other routers or societies
	var err = CheckPacket(packet)
	if err != nil {
		// TODO::: tell router and other society on repeated attack
		return
	}

	var desAddr Addr = GetDestinationAddr(packet)
	var appMux = osMux.getAppMultiplexer(desAddr.AppID())
	if appMux == nil {
		// TODO::: tell router and other society on repeated attack
		return
	}
	appMux.Receive(conn, packet)
}

// RegisterAppMultiplexer will register multiplexer only if it is GP multiplexer.
func (osMux *OSMultiplexer) RegisterAppMultiplexer(appMux protocol.NetworkTransportMultiplexer) {
	if appMux.HeaderID() != protocol.NetworkLinkNextHeaderGP {
		return
	}

	// TODO::: check application have GP access from user!

	var appID uint16 = protocol.OS.AppManifest().AppID()
	osMux.appMultiplexer[appID] = appMux
}
func (osMux *OSMultiplexer) UnRegisterAppMultiplexer(appMux protocol.NetworkTransportMultiplexer) {
	if appMux.HeaderID() != protocol.NetworkLinkNextHeaderGP {
		return
	}
	var appID uint16 = protocol.OS.AppManifest().AppID()
	osMux.appMultiplexer[appID] = nil
}
func (osMux *OSMultiplexer) getAppMultiplexer(appID uint16) protocol.NetworkTransportMultiplexer {
	return osMux.appMultiplexer[appID]
}

// Shutdown the listener when the application closes or force to closes by not recovered panic.
func (osMux *OSMultiplexer) Shutdown() {
	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
}
