//Copyright 2017 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package filters

import (
	"net"

	"github.com/miekg/dns"

	"github.com/SabzCity/go-library/net/edns"
	"github.com/SabzCity/go-library/net/edns/ednsutil"
)

var responseHandlers = map[uint16]ResponseHandler{
	dns.TypeA:     AHandler,
	dns.TypeAAAA:  AAAAHandler,
	dns.TypeCNAME: CNameHandler,
	dns.TypeMX:    MXHandler,
	dns.TypeNS:    NSHandler,
	dns.TypeSOA:   SOAHandler,
	dns.TypeTXT:   TXTHandler}

// ResponseHandler definition how developer must declare handler!
type ResponseHandler func(*edns.Context, *ednsutil.MiniDNS, string)

// SOAHandler handle A question.
func SOAHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	ctx.Response.Answer = append(ctx.Response.Answer, &dns.SOA{
		Hdr:     dns.RR_Header{Name: question, Rrtype: dns.TypeSOA, Class: dns.ClassINET, Ttl: zone.TTL},
		Ns:      zone.IN.SOA.RNAME,
		Mbox:    zone.IN.SOA.MNAME,
		Serial:  zone.IN.SOA.SERIAL,
		Refresh: zone.IN.SOA.REFRESH,
		Retry:   zone.IN.SOA.RETRY,
		Expire:  zone.IN.SOA.EXPIRE,
		Minttl:  zone.IN.SOA.MINIMUM})
}

// NSHandler handle NS question.
func NSHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	for _, value := range zone.IN.NSDNAME {
		ctx.Response.Answer = append(ctx.Response.Answer, &dns.NS{
			Hdr: dns.RR_Header{Name: question, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: zone.TTL},
			Ns:  value})
	}
}

// AHandler handle A question.
func AHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	CNameHandler(ctx, zone, question)

	for _, value := range zone.IN.A.A {
		ctx.Response.Answer = append(ctx.Response.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: question, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: zone.TTL},
			A:   net.ParseIP(value)})
	}
}

// AAAAHandler handle AAAA question.
func AAAAHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	CNameHandler(ctx, zone, question)

	for _, value := range zone.IN.AAAA.AAAA {
		ctx.Response.Answer = append(ctx.Response.Answer, &dns.AAAA{
			Hdr:  dns.RR_Header{Name: question, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: zone.TTL},
			AAAA: net.ParseIP(value)})
	}
}

// CNameHandler handle CName question.
func CNameHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	if zone.IN.CNAME.CNAME != "" {
		ctx.Response.Answer = append(ctx.Response.Answer, &dns.CNAME{
			Hdr:    dns.RR_Header{Name: question, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: zone.TTL},
			Target: zone.IN.CNAME.CNAME})
	}
}

// MXHandler handle MX question.
func MXHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	for _, value := range zone.IN.MX {
		ctx.Response.Answer = append(ctx.Response.Answer, &dns.MX{
			Hdr:        dns.RR_Header{Name: question, Rrtype: dns.TypeMX, Class: dns.ClassINET, Ttl: zone.TTL},
			Preference: uint16(value.PREFERENCE),
			Mx:         value.EXCHANGE})
	}
}

// TXTHandler handle TXT question.
func TXTHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	ctx.Response.Answer = append(ctx.Response.Answer, &dns.TXT{
		Hdr: dns.RR_Header{Name: question, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: zone.TTL},
		Txt: zone.IN.TXT.TXTDATA})
}

// ANYHandler handle ANY question.
func ANYHandler(ctx *edns.Context, zone *ednsutil.MiniDNS, question string) {
	SOAHandler(ctx, zone, question)
	NSHandler(ctx, zone, question)
	AHandler(ctx, zone, question)
	AAAAHandler(ctx, zone, question)
	MXHandler(ctx, zone, question)
	TXTHandler(ctx, zone, question)
}
