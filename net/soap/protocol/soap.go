/* For license and copyright information please see the LEGAL file in the code repository */

package soap_p

import (
	error_p "memar/error/protocol"
	net_p "memar/net/protocol"
)

type Handler interface {
	Action() string
	ServeSOAP(sk net_p.Socket, req Message, res Message) (err error_p.Error)
}

type Message interface {
	Header() Message_Header
	Body() Message_Body
}

type Message_Header interface {
}

type Message_Body interface {
	// XMLName xml.Name    `xml:"soap:Body"`
	Payload() interface{}
	Action() string
	Fault() Message_Fault
}

type Message_Fault interface {
	// XMLName xml.Name `xml:"soap:Fault"`
	// Code    string   `xml:"Code"`
	// Reason  string   `xml:"Reason"`
	// Node    string   `xml:"Node,omitempty"`
	// Role    string   `xml:"Role,omitempty"`
	// Detail  string   `xml:"Detail,omitempty"`
}
