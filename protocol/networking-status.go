/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Network_Status interface {
	Status() NetworkStatus     // return last connection||stream status
	State() chan NetworkStatus // return status channel to listen to new connection||stream status. for more than one listener use channel hub(repeater)

	// SetStatus is low level API that must use very carefully and usually in not services layer.
	// It is Non-Blocking operation. Change status of stream and send notification on stream StateChannel.
	SetStatus(ns NetworkStatus)
}

// NetworkStatus indicate connection and stream state
type NetworkStatus uint32

// Connection States
const (
	NetworkStatus_Unset NetworkStatus = iota // State not set yet

	NetworkStatus_New          // means connection not saved yet to storage
	NetworkStatus_Loaded       // means connection load from storage
	NetworkStatus_Unregistered //

	NetworkStatus_Opening // connection||stream plan to open and not ready to accept stream
	NetworkStatus_Open    // connection||stream is open and ready to use
	NetworkStatus_Closing // connection||stream plan to close and not accept new stream
	NetworkStatus_Closed  // connection||stream had been closed

	NetworkStatus_NotResponse // peer not response to recently send request
	NetworkStatus_RateLimited // connection||stream limited due to higher usage than permitted
	NetworkStatus_Timeout     // connection||stream timeout(DeadlineExceeded) and must close

	NetworkStatus_BrokenPacket
	NetworkStatus_NeedMoreData
	NetworkStatus_Sending
	NetworkStatus_Receiving
	NetworkStatus_ReceivedCompletely
	NetworkStatus_SentCompletely

	NetworkStatus_Encrypted
	NetworkStatus_Decrypted
	NetworkStatus_Ready
	NetworkStatus_Idle

	NetworkStatus_Blocked
	NetworkStatus_BlockedByPeer

	// Each package can declare it's own state with NetworkStatus > 127 e.g. tcp: TCP_SYN_SENT, TCP_SYN_RECV, ...
)
