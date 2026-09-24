/* For license and copyright information please see the LEGAL file in the code repository */

package scheduler

import (
	"runtime"

	"sync/atomic"
)

// Causes the calling thread to yield execution to another thread that is ready to run on the current processor.
// The scheduler selects the next thread to be executed.
// After the yielding thread's time slice elapses, the scheduler reschedules execution for the yielding thread.
// depending on the priorities of other threads, the yielding thread can be scheduled back in one interrupt period later.
// in some ways like runtime.Gosched
func (t *Thread) Yield(wr Thread_WaitReason) {
	t.waitReason.Store(wr)
	// TODO::: below function is not what we want.
	runtime.Gosched()
}

// A Thread_WaitReason explains why a goroutine has been stopped.
// See gopark. Do not re-use reasons, add new ones.
type Thread_WaitReason uint32

const (
	Thread_WaitReason_Unset                 Thread_WaitReason = iota // ""
	Thread_WaitReason_GCAssistMarking                                // "GC assist marking"
	Thread_WaitReason_IOWait                                         // "IO wait"
	Thread_WaitReason_ChanReceiveNilChan                             // "chan receive (nil chan)"
	Thread_WaitReason_ChanSendNilChan                                // "chan send (nil chan)"
	Thread_WaitReason_DumpingHeap                                    // "dumping heap"
	Thread_WaitReason_GarbageCollection                              // "garbage collection"
	Thread_WaitReason_GarbageCollectionScan                          // "garbage collection scan"
	Thread_WaitReason_PanicWait                                      // "panic wait"
	Thread_WaitReason_Select                                         // "select"
	Thread_WaitReason_SelectNoCases                                  // "select (no cases)"
	Thread_WaitReason_GCAssistWait                                   // "GC assist wait"
	Thread_WaitReason_GCSweepWait                                    // "GC sweep wait"
	Thread_WaitReason_GCScavengeWait                                 // "GC scavenge wait"
	Thread_WaitReason_ChanReceive                                    // "chan receive"
	Thread_WaitReason_ChanSend                                       // "chan send"
	Thread_WaitReason_FinalizerWait                                  // "finalizer wait"
	Thread_WaitReason_ForceGCIdle                                    // "force gc (idle)"
	Thread_WaitReason_Semacquire                                     // "semacquire"
	Thread_WaitReason_Sleep                                          // "sleep"
	Thread_WaitReason_SyncCondWait                                   // "sync.Cond.Wait"
	Thread_WaitReason_TimerGoroutineIdle                             // "timer goroutine (idle)"
	Thread_WaitReason_TraceReaderBlocked                             // "trace reader (blocked)"
	Thread_WaitReason_WaitForGCCycle                                 // "wait for GC cycle"
	Thread_WaitReason_GCWorkerIdle                                   // "GC worker (idle)"
	Thread_WaitReason_Preempted                                      // "preempted"
	Thread_WaitReason_DebugCall                                      // "debug call"
)

func (s *Thread_WaitReason) Load() Thread_WaitReason {
	return Thread_WaitReason(atomic.LoadUint32((*uint32)(s)))
}
func (s *Thread_WaitReason) Store(status Thread_WaitReason) {
	atomic.StoreUint32((*uint32)(s), uint32(status))
}
func (s *Thread_WaitReason) CompareAndSwap(old, new Thread_WaitReason) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
