/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

const (
	// MinFrameLen is minimum Chapar frame length
	// 4 Byte Chapar header + 8 Byte min payload
	MinFrameLen = 12
	// MaxFrameLen is maximum Chapar frame length
	MaxFrameLen = 8192

	// AcceptLastHop indicate that package must accept frames in last hop or not.
	AcceptLastHop = true

	// 256 is max ports that Chapar protocol support directly in one hop.
	defaultPortNumber = 256

	// 256 is max next header ID that Chapar protocol support.
	maxHeaderID = 256

	fixedHeaderLength      byte = 3 // without path part
	maxHopCount            byte = 255
	broadcastHopCount      byte = 0
	maxBroadcastPayloadLen int  = MaxFrameLen - (int(fixedHeaderLength) + int(maxHopCount))
)
