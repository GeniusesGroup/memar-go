/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

const (
	FrameType_Length = 1 // byte get on byte space
)

// FrameType is Frame type ID like service ID but fixed ID with 8bit length. Just some few services get one byte length service ID
// Common services must register by 64bit unsigned integer.
type FrameType byte

// https://github.com/GeniusesGroup/memar/blob/main/networking.md#frames-number
const (
	FrameType_Unset FrameType = iota

	// e.g. A way to process all old protocols e.g. Ethernet, ATM, IPv4, IPv6,
	FrameType_OldProtocols

	FrameType_Asb
	FrameType_Chapar
	FrameType_GP

	FrameType_Padding // FrameType = 128 + iota
	FrameType_CallService
	FrameType_OpenStream
	FrameType_CloseStream
	FrameType_Data
	FrameType_Error
	FrameType_Signature
)
