/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

type flag byte

const (
	Flag_Reserved1 flag = 0b00001000
	Flag_Reserved2 flag = 0b00000100
	Flag_Reserved3 flag = 0b00000010
	Flag_NS        flag = 0b00000001
	Flag_CWR       flag = 0b10000000
	Flag_ECE       flag = 0b01000000
	Flag_URG       flag = 0b00100000
	Flag_ACK       flag = 0b00010000
	Flag_PSH       flag = 0b00001000
	Flag_RST       flag = 0b00000100
	Flag_SYN       flag = 0b00000010
	Flag_FIN       flag = 0b00000001
)
