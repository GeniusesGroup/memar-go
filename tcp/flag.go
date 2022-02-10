/* For license and copyright information please see LEGAL file in repository */

package tcp

const (
	Flag_Reserved1 byte = 0b00001000
	Flag_Reserved2 byte = 0b00000100
	Flag_Reserved3 byte = 0b00000010
	Flag_NS        byte = 0b00000001
	Flag_CWR       byte = 0b10000000
	Flag_ECE       byte = 0b01000000
	Flag_URG       byte = 0b00100000
	Flag_ACK       byte = 0b00010000
	Flag_PSH       byte = 0b00001000
	Flag_RST       byte = 0b00000100
	Flag_SYN       byte = 0b00000010
	Flag_FIN       byte = 0b00000001
)

// Each flag in TCP is an individual bit representing On or Off—to manage data flow in specific situations.
// Reserved flags in TCP headers always has a value of zero.
type Flags struct {
	// The first 3 bits (rsvd) are not used.
	Reserved1, Reserved2, Reserved3 bool

	// ECN-nonce - concealment protection
	NS bool

	// Congestion window reduced (CWR) flag is set by the sending host
	// to indicate that it received a TCP segment with the ECE flag set and had
	// responded in congestion control mechanism.
	// Used for informing that the sender reduced its sending rate.
	CWR bool

	// ECN-Echo has a dual role, depending on the value of the SYN flag.
	// It indicates:
	// - If the SYN flag is set (1), that the TCP peer is ECN capable.
	// - If the SYN flag is clear (0), that a packet with Congestion Experienced flag
	//   set (ECN=11) in the IP header was received during normal transmission.
	//   This serves as an indication of network congestion (or impending congestion) to the TCP sender.
	ECE bool

	// Urgent Pointer (U) indicates that the segment contains prioritized data.
	URG bool

	// ACK Indicates that the Acknowledgment field is significant.
	// All packets after the initial SYN packet sent by the client should have this flag set.
	ACK bool

	// Push function is used to indicate that the receiver should “push” the data to the service logic as soon as possible.
	PSH bool

	// RST resets the TCP connection.
	RST bool

	// SYN (S) is used to synchronize sequence numbers in the initial handshake.
	// So Only the first packet sent from each end should have this flag set.
	// Some other flags and fields change meaning based on this flag,
	// and some are only valid when it is set, and others when it is clear.
	SYN bool

	// FIN indicates that the sender has finished sending data.
	FIN bool
}
