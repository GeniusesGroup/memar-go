/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

type SOAP_Handler interface {
	Action() string
	ServeSOAP(sk Socket, req SOAP_Message, res SOAP_Message) (err Error)
}

type SOAP_Message interface {
	Header() SOAP_Message_Header
	Body() SOAP_Message_Body
}

type SOAP_Message_Header interface {
}

type SOAP_Message_Body interface {
	// XMLName xml.Name    `xml:"soap:Body"`
	Payload() interface{}
	Action() string
	Fault() SOAP_Message_Fault
}

type SOAP_Message_Fault interface {
	// XMLName xml.Name `xml:"soap:Fault"`
	// Code    string   `xml:"Code"`
	// Reason  string   `xml:"Reason"`
	// Node    string   `xml:"Node,omitempty"`
	// Role    string   `xml:"Role,omitempty"`
	// Detail  string   `xml:"Detail,omitempty"`
}
