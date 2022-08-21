/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// A StreamID in many protocols that suggest use this rules:
// - The least significant bit (0x01) of the stream ID identifies the initiator of the stream.
// Client-initiated streams have even-numbered stream IDs (with the bit set to 0),
// and server-initiated streams have odd-numbered stream IDs (with the bit set to 1).
// - The second least significant bit (0x02) of the stream ID distinguishes between
// bidirectional streams (with the bit set to 0) and unidirectional streams (with the bit set to 1).
type StreamID uint64

// Stream_ID indicate a minimum networking stream functionality.
type Stream_ID interface {
	StreamID() StreamID
	PeerInitiated() bool
	Bidirectional() bool
}
