/* For license and copyright information please see LEGAL file in repository */

package tcp

type flag byte

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
