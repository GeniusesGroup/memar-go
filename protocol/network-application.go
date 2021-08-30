/* For license and copyright information please see LEGAL file in repository */

package protocol

import "net"

/*
**********************************************************************************
Application (OSI Layer 7: Application)
**********************************************************************************
*/

type NetworkApplicationProtocolID uint16

// Indicate standard listen and send port number register for application layer protocols.
// Usually Dev must register needed stream protocol not use this method to register all available protocols!
// Not rule but suggest to use even port number to listen||receive||response||server and odd for send||request||client
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
// https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml
const (
	NetworkApplicationSRPCSyllab NetworkApplicationProtocolID = 4
	NetworkApplicationDNS        NetworkApplicationProtocolID = 50
	NetworkApplicationHTTP       NetworkApplicationProtocolID = 80
)

// NetworkApplicationMultiplexer
type NetworkApplicationMultiplexer interface {
	GetNetworkApplicationHandler(protocolID NetworkApplicationProtocolID) NetworkApplicationHandler
	SetNetworkApplicationHandler(protocolID NetworkApplicationProtocolID, nah NetworkApplicationHandler)
	DeleteNetworkApplicationHandler(protocolID NetworkApplicationProtocolID)
}

// NetworkApplicationHandler
type NetworkApplicationHandler interface {
	HandleIncomeRequest(stream Stream)
	HandleIncomeResponse(stream Stream)
	HandleStreamConnection(stream Stream, conn net.Conn)
}
