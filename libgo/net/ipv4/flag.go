/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

const (
	flag_Reserved byte = 0b10000000
	flag_DF       byte = 0b01000000
	flag_MF       byte = 0b00100000

	// 00 – Non ECN-Capable Transport, Non-ECT
	flag_ECT0 byte = 0b00000010 // ECN Capable Transport
	flag_ECT1 byte = 0b00000001 // ECN Capable Transport
	flag_CE   byte = 0b00000011 // Congestion Encountered
)
