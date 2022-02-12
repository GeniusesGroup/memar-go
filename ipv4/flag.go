/* For license and copyright information please see LEGAL file in repository */

package ipv4

const (
	Flag_Reserved byte = 0b10000000
	Flag_DF       byte = 0b01000000
	Flag_MF       byte = 0b00100000

	// 00 – Non ECN-Capable Transport, Non-ECT
	Flag_ECT0 byte = 0b00000010 // ECN Capable Transport
	Flag_ECT1 byte = 0b00000001 // ECN Capable Transport
	Flag_CE   byte = 0b00000011 // Congestion Encountered
)
