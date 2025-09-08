/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	capsule_p "memar/computer/capsule/protocol"
	error_p "memar/error/protocol"
)

/*
**********************************************************************************
Transport (OSI Layer 5: Session)

https://en.wikipedia.org/wiki/Session_layer
https://en.wikipedia.org/wiki/Session_(computer_science)
**********************************************************************************
*/

// The session layer provides the mechanism for opening, closing and managing a session between end-user application processes,
type OSI_Session interface {
	capsule_p.LifeCycle
	// Init(dt OSI_Session_DialogueType)

	/* session data */
	Field_SessionID
	Field_DialogueType

	OSI_Session_LowLevelAPIs

	Framer
}

// Session_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// Multiplexing is the main service of the session layer in the OSI model. But it is part of the transport layer in the TCP/IP model.
type OSI_Session_LowLevelAPIs interface {
	FrameWriter

	Close() (err error_p.Error)  // Just once, must deregister the socket and notify peer in some proper way.
	Revoke() (err error_p.Error) // Just once
}
