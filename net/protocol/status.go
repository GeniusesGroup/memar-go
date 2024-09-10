/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

type Socket_Status interface {
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
	Status_Unset NetworkStatus = iota // State not set yet

	Status_New          // means socket session not saved yet to storage
	Status_Loaded       // means socket session load from storage
	Status_Unregistered //

	Status_Opening // socket plan to open and not ready to accept actions
	Status_Open    // socket is open and ready to use
	Status_Closing // socket plan to close and not accept new action
	Status_Closed  // socket had been closed

	Status_NotResponse // peer not response to recently send request
	Status_RateLimited // socket limited due to higher usage than permitted

	Status_Timeout_Read  // socket timeout(DeadlineExceeded) and must close
	Status_Timeout_Write // socket timeout(DeadlineExceeded) and must close

	Status_BrokenPacket
	Status_NeedMoreData
	Status_Sending
	Status_Receiving
	Status_ReceivedCompletely
	Status_SentCompletely

	Status_Encrypted
	Status_Decrypted
	Status_Ready
	Status_Idle

	Status_Blocked
	Status_BlockedByPeer
)
