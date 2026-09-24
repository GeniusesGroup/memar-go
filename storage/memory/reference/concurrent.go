/* For license and copyright information please see the LEGAL file in the code repository */

package reference

// Concurrent reference is an mutable reference that is safe to pass between threads.
// It will allocated on global heap not current thread.
// Immutable references doesn't need concurrent safely by introducing new capsule like this one.
type Concurrent[T any] struct{}
