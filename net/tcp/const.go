/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/timer"
)

// ATTENTION:::: Don't changed below settings without any good reason
// https://man7.org/linux/man-pages/man7/tcp.7.html
// https://man7.org/linux/man-pages/man7/socket.7.html

// SYN config values
const (
	// The maximum number of times initial SYNs for an active TCP
	// connection attempt will be retransmitted. The default value is 6,
	// which corresponds to retrying for up to approximately 127
	// seconds. Before Linux 3.7, the default value was 5, which
	// (in conjunction with calculation based on other kernel
	// parameters) corresponded to approximately 180 seconds.
	CNF_SynRetries uint8 = 6

	// The maximum number of times a SYN/ACK segment for a
	// passive TCP connection will be retransmitted.
	CNF_SynAck_Retries uint8 = 5

	// Enable TCP syncookies.  The kernel must be compiled with
	// CONFIG_SYN_COOKIES.  The syncookies feature attempts to
	// protect a socket from a SYN flood attack.  This should be
	// used as a last resort, if at all.  This is a violation of
	// the TCP protocol, and conflicts with other areas of TCP
	// such as TCP extensions.  It can cause problems for clients
	// and relays.  It is not recommended as a tuning mechanism
	// for heavily loaded servers to help with overloaded or
	// misconfigured conditions. For recommended alternatives
	// see CNF_MaxSynBacklog, CNF_SynAck_Retries, and
	// tcp_abort_on_overflow.  Set to one of the following
	// values:
	// 0  Disable TCP syncookies.
	// 1  Send out syncookies when the syn backlog queue of a
	//    socket overflows.
	// 2  (since Linux 3.12) Send out syncookies unconditionally.
	//    This can be useful for network testing.
	CNF_SynCookies = 1

	// The maximum number of queued connection requests which
	// have still not received an acknowledgement from the
	// connecting client. If this number is exceeded, the kernel
	// will begin dropping requests. The default value of 256 is
	// increased to 1024 when the memory present in the system is
	// adequate or greater (>= 128 MB), and reduced to 128 for
	// those systems with very low memory (<= 32 MB).
	//
	// Prior to Linux 2.6.20, it was recommended that if this
	// needed to be increased above 1024, the size of the SYNACK
	// hash table (TCP_SYNQ_HSIZE) in include/net/tcp.h should be
	// modified to keep
	//
	// 	TCP_SYNQ_HSIZE * 16 <= CNF_MaxSynBacklog
	//
	// and the kernel should be recompiled.  In Linux 2.6.20, the
	// fixed sized TCP_SYNQ_HSIZE was removed in favor of dynamic
	// sizing.
	CNF_MaxSynBacklog = 256
)

// congestion-control algorithm config values
const (
	// default congestion-control algorithm to be used for new tcp sockets
	CNF_CongestionControlAlgorithm = CongestionControlAlgorithm_Reno

	// When enabled, connectivity to some destinations could be
	// affected due to older, misbehaving middle boxes along the
	// path, causing connections to be dropped.  However, to
	// facilitate and encourage deployment with option 1, and to
	// work around such buggy equipment, the tcp_ecn_fallback
	// option has been introduced.
	CNF_ExplicitCongestionNotification = ExplicitCongestionNotification_EnabledOnRequested
	// Enable RFC 3168, Section 6.1.1.1. fallback. When enabled,
	// outgoing ECN-setup SYNs that time out within the normal
	// SYN retransmission timeout will be resent with CWR and ECE cleared.
	CNF_ExplicitCongestionNotification_fallback = true

	// NoDelay controls whether TCP should delay packet transmission in hopes of sending fewer packets (Nagle's algorithm).
	// The default is true (no delay), meaning that data is sent as soon as possible after a Write.
	CNF_NoDelay = true // TODO::: need it in user-space??
)

// TCP Selective Acknowledgements config values
const (
	// Enable RFC 2018 TCP Selective Acknowledgements.
	CNF_Sack = true
)

// Delayed acknowledgment(acknowledgements) config values
const (
	// negative of TCP_QUICKACK in linux which will disable the "Nagle" algorithm
	CNF_DelayedAcknowledgment = true
	// CNF_DelayedAcknowledgment_PerStream indicate that in init phase of each stream package must enable DelayedAcknowledgment.
	CNF_DelayedAcknowledgment_PerStream = true
	CNF_DelayedAcknowledgment_Timeout   = 500 * timer.Millisecond
)

// Keep alive config values
// https://www.rfc-editor.org/rfc/rfc1122#page-101
const (
	// keep-alive message means send ACK to peer to force OS or
	// any NAT mechanism in network layer to keep open the tcp stream in their routing tables.
	CNF_KeepAlive = true
	// CNF_KeepAlive_PerStream indicate that in init phase of each stream package must enable KeepAlive.
	CNF_KeepAlive_PerStream = false

	// The time (in seconds) the connection needs to remain idle before TCP starts sending keepalive probes.
	// terminated after approximately an additional CNF_KeepAlive_Probes*CNF_KeepAlive_Interval e.g.
	// 9 probes an interval of 75 seconds apart be 11 minutes.
	CNF_KeepAlive_Idle = 7200 * timer.Second // (2 hours)
	// The number of seconds between TCP keep-alive probes.
	CNF_KeepAlive_Interval = 75 * timer.Second
	// The maximum number of TCP keep-alive probes to send before giving up and
	// killing the socket connection if no response is obtained from the other end.
	CNF_KeepAlive_Probes = 9
)

// User timeout config values
// https://www.rfc-editor.org/rfc/rfc0793
// https://www.rfc-editor.org/rfc/rfc1122
const (
	// The TCP user timeout controls how long transmitted data may remain unacknowledged,
	// or buffered data may remain untransmitted (due to zero window size),
	// before TCP will forcibly close the corresponding connection
	CNF_UserTimeout = true
	// CNF_UserTimeout_PerStream indicate that in init phase of each stream package must enable UserTimeout.
	CNF_UserTimeout_PerStream      = true
	CNF_UserTimeout_Retransmission = 3
	CNF_UserTimeout_Idle           = 100 * timer.Second
	CNF_UserTimeout_SynIdle        = 180 * timer.Second // (3 minutes)
)

// timeout config values
// https://www.rfc-editor.org/rfc/rfc1337
const (
	// TIME_WAIT indicates that local endpoint (this side) has closed the connection.
	// The connection is being kept around so that any delayed packets
	// can be matched to the connection and handled appropriately.
	// The connections will be removed when they time out within four minutes.
	CNF_Timeout           = true
	CNF_Timeout_PerStream = true
	// Enable TCP behavior conformant with RFC 1337.  When
	// disabled, if a RST is received in TIME_WAIT state, we
	// close the socket immediately without waiting for the end
	// of the TIME_WAIT period.
	CNF_Timeout_RFC1337 = false

	// The maximum number of sockets in TIME_WAIT state allowed
	// in the system.  This limit exists only to prevent simple
	// denial-of-service attacks.  The default value of NR_FILE*2
	// is adjusted depending on the memory in the system.  If
	// this number is exceeded, the socket is closed and a
	// warning is printed.
	CNF_Timeout_MaxBuckets = 256

	// Enable fast recycling of TIME_WAIT sockets.  Enabling this
	// option is not recommended as the remote IP may not use
	// monotonically increasing timestamps (devices behind NAT,
	// devices with per-connection timestamp offsets).  See RFC
	// 1323 (PAWS) and RFC 6191.
	CNF_Timeout_Recycle = false

	// Allow to reuse TIME_WAIT sockets for new connections when
	// it is safe from protocol viewpoint.  It should not be
	// changed without advice/request of technical experts.
	CNF_Timeout_Reuse = false

	// Linger sets the behavior of Close on a connection which still
	// has data waiting to be sent or to be acknowledged.
	//
	// - If sec < 0 (the default), TCP finishes sending the data in the background.
	// - If sec == 0, TCP discards any unsent or unacknowledged data.
	// - If sec > 0, the data is sent in the background as with sec < 0.
	// 		On some implementation after seconds have elapsed any remaining unsent data may be discarded.
	CNF_Timeout_Linger = -1
)

// FIN config values
const (
	// This specifies how many seconds to wait for a final FIN
	// packet (The lifetime of orphaned FIN_WAIT2 state sockets)
	// before the socket is forcibly closed. (Also know as Linger2)
	// This is strictly a violation of the TCP specification, but
	// required to prevent denial-of-service attacks.  In Linux
	// 2.2, the default value was 180.
	//  This option should not be used in code intended to be portable.
	CNF_FinTimeout = 60 * timer.Second
)

// segment config values
const (
	CNF_Segment_MinSize = 20 // 5words * 4bit
	// mss or maximum segment size greater than the (eventual) interface MTU have no effect.
	CNF_Segment_MaxSize = 536
)
