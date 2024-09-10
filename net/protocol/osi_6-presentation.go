/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

/*
**********************************************************************************
Transport (OSI Layer 6: Presentation)

https://en.wikipedia.org/wiki/Presentation_layer
**********************************************************************************
*/

// Serialization of complex data structures into flat byte-strings (using mechanisms such as TLV, XML or JSON)
// can be thought of as the key functionality of the presentation layer.
// 
// Encryption and Decryption are typically done at this level too,
// although it can be done on the application, session, transport, or network layers, each having its own advantages and disadvantages.
// For example, when logging on to bank account sites the presentation layer will decrypt the data as it is received.
type OSI_Presentation interface {
	Framer

	OSI_Presentation_LowLevelAPIs
}

// Session_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// Multiplexing is the main service of the session layer in the OSI model. But it is part of the transport layer in the TCP/IP model.
type OSI_Presentation_LowLevelAPIs interface {
}