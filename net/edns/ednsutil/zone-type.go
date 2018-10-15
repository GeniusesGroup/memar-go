// Copyright 2017 SabzCity

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ednsutil is Extended DNS package utility
// https://tools.ietf.org/html/rfc1034
// https://tools.ietf.org/html/rfc1035
// https://en.wikipedia.org/wiki/List_of_DNS_record_types
package ednsutil

// DNS is The standard struct of DNS zone.
// All subdomain must have full domain name even @ can't use as origin domain name e.g. "sabz.city." || "shop.sabz.city."
type DNS struct {
	Origin string
	TTL    uint32
	IN     map[string]RR // the Internet system is DNS class
	CH     map[string]RR // the Chaos system is DNS class
}

// RR is Resource Records of DNS zone.
// interface{} can be any of RDATA type by ID of records.
type RR struct {
	SOA SOA
	PTR
	NS
	A
	AAAA
	CNAME
	DNAME
	ALIAS
	TXT
	MX  []MX
	SRV []SRV
	DS  []DS
}

//						//
//		Standard RRs	//
//						//

// Container RDATA format
// Non Standard!!! just use by us!!!!
// It will send nearest server Address (A & AAAA) that have running container to user!
type Container struct {
	ContainerID string
}

// CNAME RDATA format.
// https://tools.ietf.org/html/rfc1035#section-3.3.1
// The IANA assigned value of the type is 5 (decimal).
type CNAME struct {
	CNAME string // Allias canonical name record.
}

// DNAME RDATA format.
// https://tools.ietf.org/html/rfc6672
// The IANA assigned value of the type is 39 (decimal).
type DNAME struct {
	DNAME string // Alias for a name and all its subnames, unlike CNAME, which is an alias for only the exact name.
}

// MX RDATA format.
// https://tools.ietf.org/html/rfc1035#section-3.3.9
// The IANA assigned value of the type is 15 (decimal).
type MX struct {
	PREFERENCE uint16
	EXCHANGE   string
}

// NS RDATA format
// https://tools.ietf.org/html/rfc1035#section-3.3.11
// The IANA assigned value of the type is 2 (decimal).
type NS struct {
	NSDNAME []string // NameServer of host. If set it, DNS server must query to these zone instead use of other record.
}

// PTR RDATA format
// https://tools.ietf.org/html/rfc1035#section-3.3.12
// The IANA assigned value of the type is 12 (decimal).
type PTR struct {
	PTRDNAME string // a pointer to another part of the domain name space. Unlike a CNAME, DNS processing stops and just the name is returned.
}

// SOA RDATA format
// https://tools.ietf.org/html/rfc1035#section-3.3.13
// The IANA assigned value of the type is 6 (decimal).
type SOA struct {
	MNAME   string // The <domain-name> of the name server that was the original or primary source of data for this zone.
	RNAME   string // A <domain-name> which specifies the mailbox of the person responsible for this zone.
	SERIAL  uint32 // The unsigned 32 bit version number of the original copy of the zone.  Zone transfers preserve this value. This value wraps and should be compared using sequence space arithmetic.
	REFRESH uint32 // A 32 bit time interval before the zone should be refreshed.
	RETRY   uint32 // A 32 bit time interval that should elapse before a failed refresh should be retried.
	EXPIRE  uint32 // A 32 bit time value that specifies the upper limit on the time interval that can elapse before the zone is no longer authoritative.
	MINIMUM uint32 // The unsigned 32 bit minimum TTL field that should be exported with any RR from this zone.
}

// TXT RDATA format
// https://tools.ietf.org/html/rfc1035#section-3.3.14
// The IANA assigned value of the type is 16 (decimal).
type TXT struct {
	TXTDATA []string // Text record
}

//							//
//	Internet specific RRs	//
//							//

// A RDATA format
// https://tools.ietf.org/html/rfc1035#section-3.4.1
// The IANA assigned value of the type is 1 (decimal).
type A struct {
	A []string // IP v4 of host.
}

// AAAA RDATA format
// https://tools.ietf.org/html/rfc3596
// The IANA assigned value of the type is 28 (decimal).
type AAAA struct {
	AAAA []string // IP v6 of host.
}

// ALIAS RDATA format
// Non Standard!!! just use by big company!!
// The ALIAS record will automatically resolve your domain to one or more A records at resolution time and
// thus resolvers see your domain simply as if it had A records.
type ALIAS struct {
	Alias string // Other domain name
}

// SRV RDATA format
// Replace it with URI   https://tools.ietf.org/html/rfc7553
type SRV struct {
	Target   string
	Priority int
	Weight   int
	Port     int
}

// DS RDATA format
type DS struct {
	KeyTag     int
	Algorithm  int
	DigestType int
	Digest     string
	MaxSigLife string
	Flags      string
	Protocol   string
	KeyDataALG string
	PublicKey  string
}
