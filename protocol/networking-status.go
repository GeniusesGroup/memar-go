/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Network_Status interface {
	Status() NetworkStatus     // return last socket status
	State() chan NetworkStatus // return status channel to listen to new socket status. for more than one listener use channel hub(repeater)

	// SetStatus is low level API that must use very carefully and usually in not services layer.
	// It is Non-Blocking operation. Change status of socket and send notification on socket StateChannel.
	SetStatus(ns NetworkStatus)
}

// NetworkStatus indicate socket state
type NetworkStatus uint32

// Connection States
const (
	NetworkStatus_Unset NetworkStatus = iota // State not set yet

	NetworkStatus_New          // means socket session not saved yet to storage
	NetworkStatus_Loaded       // means socket session load from storage
	NetworkStatus_Unregistered //

	NetworkStatus_Opening // socket plan to open and not ready to accept actions
	NetworkStatus_Open    // socket is open and ready to use
	NetworkStatus_Closing // socket plan to close and not accept new action
	NetworkStatus_Closed  // socket had been closed

	NetworkStatus_NotResponse // peer not response to recently send request
	NetworkStatus_RateLimited // socket limited due to higher usage than permitted

	NetworkStatus_Timeout_Read  // socket timeout(DeadlineExceeded) and must close
	NetworkStatus_Timeout_Write // socket timeout(DeadlineExceeded) and must close

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
)
