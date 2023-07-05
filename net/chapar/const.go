/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

const (
	// MinFrameLen is minimum Chapar frame length
	MinFrameLen = int(frameFixedLength + minHopCount)
	// MaxFrameLen is maximum Chapar frame length
	MaxFrameLen = int(frameFixedLength) + int(maxHopCount)

	// AcceptLastHop indicate that package must accept frames in last hop or not.
	AcceptLastHop = true

	// 256 is max ports that Chapar protocol support directly in one hop.
	defaultPortNumber = 256

	frameFixedLength  byte = 3 // without path part
	minHopCount       byte = 1
	maxHopCount       byte = 255
	broadcastHopCount byte = 0
)
