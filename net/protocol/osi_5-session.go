/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Transport (OSI Layer 5: Session)

https://en.wikipedia.org/wiki/Session_layer
https://en.wikipedia.org/wiki/Session_(computer_science)
**********************************************************************************
*/

// The session layer provides the mechanism for opening, closing and managing a session between end-user application processes,
type OSI_Session interface {
	ObjectLifeCycle
	Init(dt OSI_Session_DialogueType)

	/* session data */
	DialogueType() OSI_Session_DialogueType

	/* Peer data */
	DomainName() string // if exist
	UserID() UserID
	DelegateUserID() UserID // Persons can delegate to things(as a user type)

	Close() (err Error)  // Just once, must deregister the socket and notify peer in some proper way.
	Revoke() (err Error) // Just once
	// Authorize request by data in the session for many access control like service, time, location, ...
	// Dev must extend this method in each service by it uses.
	Authorize() (err Error)

	OSI_Session_LowLevelAPIs

	Network_Framer
}

// Session_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// Multiplexing is the main service of the session layer in the OSI model. But it is part of the transport layer in the TCP/IP model.
type OSI_Session_LowLevelAPIs interface {
	Network_FrameWriter
}

// Dialogue is a discussion intended to produce an agreement
// https://en.wikipedia.org/wiki/Session_layer#Dialogue_control
type OSI_Session_DialogueType uint8

const (
	OSI_Session_DialogueType_Unset      OSI_Session_DialogueType = iota
	OSI_Session_DialogueType_FullDuplex                          // allowing communication in opposite directions simultaneously
	OSI_Session_DialogueType_HalfDuplex                          // information can be sent in only one direction at a time (two way alternate)
	OSI_Session_DialogueType_Simplex                             // one way (Monolog)
)
