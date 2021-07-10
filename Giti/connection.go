/* For license and copyright information please see LEGAL file in repository */

package giti

// ConnectionWeight indicate connection and stream state
type ConnectionState uint8

// Standrad Connection States
const (
	ConnectionStateUnset  ConnectionState = iota // State not set yet!
	ConnectionStateNew                           // means connection not saved yet to storage!
	ConnectionStateLoaded                        // means connection load from storage

	ConnectionStateOpening // connection||stream plan to open and not ready to accept stream!
	ConnectionStateOpen    // connection||stream is open and ready to use
	ConnectionStateClosing // connection||stream plan to close and not accept new stream
	ConnectionStateClosed  // connection||stream had been closed

	ConnectionStateNotResponse // peer not response to recently send request!
	ConnectionStateRateLimited // connection||stream limited due to higher usage than permitted!

	ConnectionStateBrokenPacket
	ConnectionStateReceivedCompletely
	ConnectionStateSentCompletely
	ConnectionStateEncrypted
	ConnectionStateDecrypted
	ConnectionStateReady
	ConnectionStateIdle

	ConnectionStateBlocked
	ConnectionStateBlockedByPeer
)

// ConnectionWeight indicate connection and stream weight
type ConnectionWeight uint8

// Standrad Connection Weights
const (
	ConnectionWeightUnset ConnectionWeight = iota
	
	ConnectionWeightNormal
	ConnectionWeightTimeSensitive        // If true must call related service in each received packet. VoIP, IPTV, Sensors data, ...
	// TODO::: 16 queue for priority weight of the connections exist.
)
