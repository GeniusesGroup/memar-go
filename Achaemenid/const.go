/* For license and copyright information please see LEGAL file in repository */

package achaemenid

type state uint8

// State
const (
	// StateUnset indicate state not set yet!
	StateUnset state = iota

	// StateClosed indicate connection had been closed
	StateClosed
	// StateClosing indicate connection plan to close and not accept new stream
	StateClosing
	// StateNotResponse indicate peer not response to recently send request!
	StateNotResponse
	// StateOpen indicate connection is open and ready to use
	StateOpen
	// StateOpening indicate connection plan to open and not ready to accept stream!
	StateOpening
	// StateRateLimited indicate connection limited due to higher usage than permitted!
	StateRateLimited

	StateBrokenPacket
	StateReady
)

type weight uint8

// weight
const (
	WeightNormal        weight = iota
	WeightTimeSensitive        // If true must call related service in each received packet. VoIP, IPTV, Sensors data, ...
	// 16 queue for priority weight of the connections exist.
)
