/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/timer"
)

// ATTENTION:::: Don't changed below settings without any good reason
// https://man7.org/linux/man-pages/man7/tcp.7.html
const (
	// default congestion-control algorithm to be used for new tcp sockets
	CongestionControlAlgorithm = CongestionControlAlgorithm_Reno

	// When enabled, connectivity to some destinations could be
	// affected due to older, misbehaving middle boxes along the
	// path, causing connections to be dropped.  However, to
	// facilitate and encourage deployment with option 1, and to
	// work around such buggy equipment, the tcp_ecn_fallback
	// option has been introduced.
	ExplicitCongestionNotification = ExplicitCongestionNotification_EnabledOnRequested
    // Enable RFC 3168, Section 6.1.1.1. fallback. When enabled,
    // outgoing ECN-setup SYNs that time out within the normal
    // SYN retransmission timeout will be resent with CWR and ECE cleared.
	ExplicitCongestionNotification_fallback = true

	// negative of TCP_QUICKACK in linux which will disable the "Nagle" algorithm
	DelayedAcknowledgment         = true
	DelayedAcknowledgment_Timeout = 500 * timer.Millisecond
	// NoDelay controls whether the operating system should delay
	// packet transmission in hopes of sending fewer packets (Nagle's algorithm).
	// The default is true (no delay), meaning that data is
	// sent as soon as possible after a Write.
	NoDelay = true // TODO::: need it in user-space??

	// keep-alive message means send ACK to peer to force OS and
	// any NAT mechanism in network layer to keep open the tcp stream in their routing tables.
	KeepAlive_Message = true

	// The number of seconds between TCP keep-alive probes.
	KeepAlive_Interval = 75 * timer.Second
	// The maximum number of TCP keep-alive probes to send before giving up and
	// killing the socket connection if no response is obtained from the other end.
	KeepAlive_Probes = 9
	// The time (in seconds) the connection needs to remain idle before TCP starts sending keepalive probes.
	// terminated after approximately an additional 11 minutes (9 probes an interval of 75 seconds apart)
	KeepAlive_Idle = 7200 * timer.Second // (2 hours)

	// Linger sets the behavior of Close on a connection which still
	// has data waiting to be sent or to be acknowledged.
	//
	// If sec < 0 (the default), the operating system finishes sending the
	// data in the background.
	//
	// If sec == 0, the operating system discards any unsent or
	// unacknowledged data.
	//
	// If sec > 0, the data is sent in the background as with sec < 0. On
	// some operating systems after sec seconds have elapsed any remaining
	// unsent data may be discarded.
	Linger = -1

	// The lifetime of orphaned FIN_WAIT2 state sockets. This
	// option can be used to override the system-wide setting in
	// the file /proc/sys/net/ipv4/tcp_fin_timeout for this
	// socket. This is not to be confused with the socket(7)
	// level option SO_LINGER.  This option should not be used in
	// code intended to be portable.
	Linger2 = 0
)

// ATTENTION:::: Don't changed below settings without any good reason
const (
	MinPacketLen      = 20 // 5words * 4bit
	OptionDefault_MSS = 536
)
