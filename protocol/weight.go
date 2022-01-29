/* For license and copyright information please see LEGAL file in repository */

package protocol

// Weight indicate connection and stream weight
type Weight uint8

// Standrad  Weights
const (
	Weight_Unset Weight = iota
	Weight_LowestPriority
	Weight_LowerPriority
	Weight_LowPriority
	Weight_Normal
	// TODO::: add more queue for priority weights
	Weight_TimeSensitive Weight = 255 // Call related service in each received packet. VoIP, IPTV, Sensors data, ...
)
