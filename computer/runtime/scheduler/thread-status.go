/* For license and copyright information please see the LEGAL file in the code repository */

package scheduler

import (
	"sync/atomic"
)

type ThreadStatus uint32

const (
	ThreadState_Unset ThreadStatus = iota
	ThreadState_Ready // ThreadState_Runnable
	ThreadState_Running
	ThreadState_Waiting
	ThreadState_Sleeping // Idle
	ThreadState_Blocked  // Usually due to a syscall
	ThreadState_Terminated
	ThreadState_Dead
)

func (s *ThreadStatus) Load() ThreadStatus {
	return ThreadStatus(atomic.LoadUint32((*uint32)(s)))
}
func (s *ThreadStatus) Store(status ThreadStatus) {
	atomic.StoreUint32((*uint32)(s), uint32(status))
}
func (s *ThreadStatus) CompareAndSwap(old, new ThreadStatus) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
