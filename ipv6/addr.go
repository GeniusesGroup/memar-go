/* For license and copyright information please see LEGAL file in repository */

package ipv6

// An Addr is a IPv6 IP address.
type Addr [IPv6len]byte

const (
	// IPv6len address lengths 128 bit || 16 byte.
	IPv6len = 16

	hextable = "0123456789abcdef"
)

// Well-known IPv6 addresses
var (
	IPv6zero                   = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = Addr{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = Addr{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = Addr{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}

	IPv6MaxStringLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
)

// IsUnspecified reports whether addr is an unspecified address "::".
func (addr *Addr) IsUnspecified() bool {
	return *addr == IPv6unspecified
}

// IsLoopback reports whether ip is a loopback address.
func (addr *Addr) IsLoopback() bool {
	return *addr == IPv6loopback
}

// IsMulticast reports whether ip is a multicast address.
func (addr *Addr) IsMulticast() bool {
	return addr[0] == 0xff
}

// IsInterfaceLocalMulticast reports whether ip is
// an interface-local multicast address.
func (addr *Addr) IsInterfaceLocalMulticast() bool {
	return addr[0] == 0xff && addr[1]&0x0f == 0x01
}

// IsLinkLocalMulticast reports whether ip is a link-local multicast address.
func (addr *Addr) IsLinkLocalMulticast() bool {
	return addr[0] == 0xff && addr[1]&0x0f == 0x02
}

// IsLinkLocalUnicast reports whether ip is a link-local unicast address.
func (addr *Addr) IsLinkLocalUnicast() bool {
	return addr[0] == 0xfe && addr[1]&0xc0 == 0x80
}

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

// ToString returns canonical string representation of IPv6.
func (addr *Addr) ToString() string {
	p := *addr
	// Find longest run of zeros.
	e0 := -1
	e1 := -1
	for i := 0; i < IPv6len; i += 2 {
		j := i
		for j < IPv6len && p[j] == 0 && p[j+1] == 0 {
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

	b := make([]byte, 0, IPv6MaxStringLen)

	// Print with possible :: in place of run of zeros
	for i := 0; i < IPv6len; i += 2 {
		if i == e0 {
			b = append(b, ':', ':')
			i = e1
			if i >= IPv6len {
				break
			}
		} else if i > 0 {
			b = append(b, ':')
		}
		b[i*2] = hextable[p[i]>>4]
		b[i*2+1] = hextable[p[i+1]&0x0f]
	}
	return string(b)
}

// FromString parses ip as a literal IPv6 address described in RFC 4291 and RFC 5952.
func (addr *Addr) FromString(ip string) {
	ellipsis := -1 // position of ellipsis in ip

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
	i := 0
	for i < IPv6len {
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
	if i < IPv6len {
		if ellipsis < 0 {
			return
		}
		n := IPv6len - i
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
}
