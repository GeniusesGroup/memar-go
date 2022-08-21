/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Priority indicate to use by any queue systems like connection and stream to implement priority/weight mechanism.
type Priority uint8

// Priorities
const (
	Priority_Unset         Priority = iota
	Priority_TimeSensitive          // Call related service in each received packet. VoIP, IPTV, Sensors data, ...
	Priority_Normal
	Priority_Low
	Priority_Lower
	// TODO::: add more priorities
	Priority_Lowest Priority = 255
)
