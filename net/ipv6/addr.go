/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

import (
	"memar/protocol"
)

// Well-known IPv6 addresses
var (
	AddrZero                   = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	AddrUnspecified            = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // "::"
	AddrLoopback               = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	AddrInterfaceLocalAllNodes = Addr{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	AddrLinkLocalAllnodes      = Addr{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	AddrLinkLocalAllRouters    = Addr{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)

// An Addr is an IP address version 6.
type Addr [AddrLen]byte

func (addr Addr) IsUnspecified() bool             { return addr == AddrUnspecified }
func (addr Addr) IsLoopback() bool                { return addr == AddrLoopback }
func (addr Addr) IsMulticast() bool               { return addr[0] == 0xff }
func (addr Addr) IsInterfaceLocalMulticast() bool { return addr[0] == 0xff && addr[1]&0x0f == 0x01 }
func (addr Addr) IsLinkLocalMulticast() bool      { return addr[0] == 0xff && addr[1]&0x0f == 0x02 }
func (addr Addr) IsLinkLocalUnicast() bool        { return addr[0] == 0xfe && addr[1]&0xc0 == 0x80 }

// IsGlobalUnicast reports whether ip is a global unicast address.
// It returns true even if ip is local IPv6 unicast address space.
// https://tools.ietf.org/html/rfc1122
// https://tools.ietf.org/html/rfc4632
// https://tools.ietf.org/html/rfc4291
func (addr *Addr) IsGlobalUnicast() bool {
	return !addr.IsUnspecified() &&
		!addr.IsLoopback() &&
		!addr.IsMulticast() &&
		!addr.IsLinkLocalUnicast()
}

// FromIPv4 set given the IPv4 address in 16-byte form
func (addr *Addr) FromIPv4(v4 [4]byte) {
	var v4InV6Prefix = [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

	copy(addr[:], v4InV6Prefix[:])
	copy(addr[12:], v4[:])
	return
}

// ToString returns canonical string representation of IPv6.
//
//memar:impl memar/protocol.Stringer
func (addr Addr) ToString() string {
	// Find longest run of zeros.
	var e0 = -1
	var e1 = -1
	for i := 0; i < AddrLen; i += 2 {
		j := i
		for j < AddrLen && addr[j] == 0 && addr[j+1] == 0 {
			j += 2
		}
		if j > i && j-i > e1-e0 {
			e0 = i
			e1 = j
			i = j
		}
	}
	// The symbol "::" MUST NOT be used to shorten just one 16 bit 0 field.
	if e1-e0 <= 2 {
		e0 = -1
		e1 = -1
	}

	const (
		hexTable     = "0123456789abcdef"
		maxStringLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	)

	var b = make([]byte, 0, maxStringLen)

	// Print with possible :: in place of run of zeros
	for i := 0; i < AddrLen; i += 2 {
		if i == e0 {
			b = append(b, ':', ':')
			i = e1
			if i >= AddrLen {
				break
			}
		} else if i > 0 {
			b = append(b, ':')
		}
		b[i*2] = hexTable[addr[i]>>4]
		b[i*2+1] = hexTable[addr[i+1]&0x0f]
	}
	return string(b)
}

// FromString parses ip as a literal IPv6 address described in RFC 4291 and RFC 5952.
//
//memar:impl memar/protocol.Stringer
func (addr *Addr) FromString(ip string) (err protocol.Error) {
	var ellipsis = -1 // position of ellipsis in ip

	// Might have leading ellipsis
	if len(ip) >= 2 && ip[0] == ':' && ip[1] == ':' {
		ellipsis = 0
		ip = ip[2:]
		// Might be only ellipsis
		if len(ip) == 0 {
			return
		}
	}

	// Loop, parsing hex numbers followed by colon.
	var i = 0
	for i < AddrLen {
		// Hex number.
		n, c, ok := xtoi(ip)
		if !ok || n > 0xFFFF {
			return
		}

		// Save this 16-bit chunk.
		addr[i] = byte(n >> 8)
		addr[i+1] = byte(n)
		i += 2

		// Stop at end of string.
		ip = ip[c:]
		if len(ip) == 0 {
			break
		}

		// Otherwise must be followed by colon and more.
		if ip[0] != ':' || len(ip) == 1 {
			return
		}
		ip = ip[1:]

		// Look for ellipsis.
		if ip[0] == ':' {
			if ellipsis >= 0 { // already have one
				return
			}
			ellipsis = i
			ip = ip[1:]
			if len(ip) == 0 { // can be at end
				break
			}
		}
	}

	// Must have used entire string.
	if len(ip) != 0 {
		return
	}

	// If didn't parse enough, expand ellipsis.
	if i < AddrLen {
		if ellipsis < 0 {
			return
		}
		var n = AddrLen - i
		for j := i - 1; j >= ellipsis; j-- {
			addr[j+n] = addr[j]
		}
		for j := ellipsis + n - 1; j >= ellipsis; j-- {
			addr[j] = 0
		}
	} else if ellipsis >= 0 {
		// Ellipsis must represent at least one 0 group.
		return
	}

	return
}

// Hexadecimal to integer.
// Returns number, characters consumed, success.
func xtoi(s string) (n int, i int, ok bool) {
	const (
		// Bigger than we need, not too big to worry about overflow
		big = 0xFFFFFF
	)

	n = 0
	for i = 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			n *= 16
			n += int(s[i] - '0')
		} else if 'a' <= s[i] && s[i] <= 'f' {
			n *= 16
			n += int(s[i]-'a') + 10
		} else if 'A' <= s[i] && s[i] <= 'F' {
			n *= 16
			n += int(s[i]-'A') + 10
		} else {
			break
		}
		if n >= big {
			return 0, i, false
		}
	}
	if i == 0 {
		return 0, i, false
	}
	return n, i, true
}
