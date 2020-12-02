/* For license and copyright information please see LEGAL file in repository */

package achaemenid

type state uint8

// State
const (
	StateUnset  state = iota // State not set yet!
	StateNew                 // means connection not saved yet to storage!
	StateLoaded              // means connection load from storage

	StateOpening // connection||stream plan to open and not ready to accept stream!
	StateOpen    // connection||stream is open and ready to use
	StateClosing // connection||stream plan to close and not accept new stream
	StateClosed  // connection||stream had been closed

	StateNotResponse // peer not response to recently send request!
	StateRateLimited // connection||stream limited due to higher usage than permitted!

	StateBrokenPacket
	StateReady
)

// Weight indicate connection and stream weight
type Weight uint8

// Weight
const (
	WeightNormal        Weight = iota
	WeightTimeSensitive        // If true must call related service in each received packet. VoIP, IPTV, Sensors data, ...
	// 16 queue for priority weight of the connections exist.
)
