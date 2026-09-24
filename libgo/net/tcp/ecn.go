/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

// ECN is RFC 3168 Explicit Congestion Notification
type ECN uint8

const (
	ExplicitCongestionNotification_Unset ECN = iota

	// Disable ECN Neither initiate nor accept ECN.
	// This was the default up to and including Linux
	// 2.6.30.
	ExplicitCongestionNotification_Disabled
	// Enable ECN when requested by incoming connections
	// and also request ECN on outgoing connection
	// attempts.
	ExplicitCongestionNotification_Enabled
	// Enable ECN when requested by incoming connections,
	// but do not request ECN on outgoing connections.
	// This value is supported, and is the default, since
	// Linux 2.6.31.
	ExplicitCongestionNotification_EnabledOnRequested
)
