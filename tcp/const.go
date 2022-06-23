/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"../timer"
)

// ATTENTION:::: Don't changed below settings without any good reason
// https://man7.org/linux/man-pages/man7/tcp.7.html
const (
	// default congestion-control algorithm to be used for new tcp sockets
	CongestionControlAlgorithm = "reno"

	// negative of TCP_QUICKACK in linux which will disable the "Nagle" algorithm
	DelayedAcknowledgment        = true
	DelayedAcknowledgmentTimeout = 500 * timer.Millisecond
	// NoDelay controls whether the operating system should delay
	// packet transmission in hopes of sending fewer packets (Nagle's
	// algorithm).  The default is true (no delay), meaning that data is
	// sent as soon as possible after a Write.
	NoDelay = true // TODO::: need it in user-space??

	// keep-alive message means send ACK to peer to force OS and
	// any NAT mechanism in network layer to keep open the tcp stream in their routing tables.
	KeepAlive_Message = true

	// The number of seconds between TCP keep-alive probes.
	KeepAlive_Interval = 75
	// The maximum number of TCP keep-alive probes to send before giving up and
	// killing the connection if no response is obtained from the other end.
	KeepAlive_Probes = 9
	// The time (in seconds) the connection needs to remain idle before TCP starts sending keepalive probes.
	// terminated after approximately an additional 11 minutes (9 probes an interval of 75 seconds apart)
	KeepAlive_Idle = 7200 // seconds (2 hours)

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
)

// ATTENTION:::: Don't changed below settings without any good reason
const (
	MinPacketLen      = 20 // 5words * 4bit
	OptionDefault_MSS = 536
)
