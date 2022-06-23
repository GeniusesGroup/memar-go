/* For license and copyright information please see LEGAL file in repository */

package tcp

const (
	// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
	tcpProtocolNumberOverIP byte = 0x06
)

// TODO::: impelemnet checksums over IPv4, IPv6, standalone
// https://github.com/google/gopacket/blob/master/layers/tcpip.go
// https://github.com/tass-belgium/picotcp/blob/master/modules/pico_tcp.c
