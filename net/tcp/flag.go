/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

type flag byte

const (
	flag_Reserved1 flag = 0b00001000
	flag_Reserved2 flag = 0b00000100
	flag_Reserved3 flag = 0b00000010
	flag_NS        flag = 0b00000001
	flag_CWR       flag = 0b10000000
	flag_ECE       flag = 0b01000000
	flag_URG       flag = 0b00100000
	flag_ACK       flag = 0b00010000
	flag_PSH       flag = 0b00001000
	flag_RST       flag = 0b00000100
	flag_SYN       flag = 0b00000010
	flag_FIN       flag = 0b00000001
)
