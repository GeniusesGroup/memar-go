/* For license and copyright information please see LEGAL file in repository */

package timer

// The Timer type represents a single event.
// When the Timer expires, a signal will be sent on Signal() channel,
// unless the Timer was created by AfterFunc.
// A Timer must be created with Init, After or AfterFunc.
//
// An active timer (one that has been called Start()) may be
// passed to deltimer (time.stopTimer), after which it is no longer an
// active timer. It is an inactive timer.
// In an inactive timer the period, f, arg, and seq fields may be modified,
// but not the when field.
// It's OK to just drop an inactive timer and let the GC collect it.
// It's not OK to pass an inactive timer to addtimer.
// Only newly allocated timer values may be passed to addtimer.
//
// We don't permit calling addtimer/deltimer/modtimer/resettimer simultaneously,
// but adjusttimers and runtimer can be called at the same time as any of those.

// Init initialize the Timer with given callback function or make the channel and send signal on it
// Be aware that given function must not be closure and must not block the caller.

// Stop prevents the Timer from firing (no more ticks will be sent).
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped (It reports whether t was stopped before being run.).
// Stop does not close the channel, to prevent a read from the channel succeeding
// incorrectly (seeing an erroneous "tick").
//
// To ensure the channel is empty after a call to Stop, check the
// return value and drain the channel.
// For example, assuming the program has not received from t.Signal() already:
//
// 	if !t.Stop() {
// 		<-t.Signal()
// 	}
//
// This cannot be done concurrent to other receives from the Timer's
// channel or other calls to the Timer's Stop method.
//
// For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
// has already expired and the function f has been started in its own goroutine;
// Stop does not wait for f to complete before returning.
// If the caller needs to know whether f is completed, it must coordinate
// with f explicitly.

// Reset changes the timer to expire after duration d.
// It returns true if the timer had been active, false if the timer had
// expired or been stopped.
//
// For a Timer created with NewTimer, Reset should be invoked only on
// stopped or expired timers with drained channels.
//
// If a program has already received a value from t.Signal(), the timer is known
// to have expired and the channel drained, so t.Reset can be used directly.
// If a program has not yet received a value from t.Signal(), however,
// the timer must be stopped and—if Stop reports that the timer expired
// before being stopped—the channel explicitly drained:
//
// 	if !t.Stop() {
// 		<-t.Signal()
// 	}
// 	t.Reset(d)
//
// This should not be done concurrent to other receives from the Timer's
// channel.
//
// Note that it is not possible to use Reset's return value correctly, as there
// is a race condition between draining the channel and the new timer expiring.
// Reset should always be invoked on stopped or expired channels, as described above.
// The return value exists to preserve compatibility with existing programs.
//
// For a Timer created with AfterFunc(d, f), Reset either reschedules
// when f will run, in which case Reset returns true, or schedules f
// to run again, in which case it returns false.
// When Reset returns false, Reset neither waits for the prior f to
// complete before returning nor does it guarantee that the subsequent
// goroutine running f does not run concurrently with the prior
// one. If the caller needs to know whether the prior execution of
// f is completed, it must coordinate with f explicitly.
//
// Reset stops a ticker and resets its period to the specified duration.
// The next tick will arrive after the new period elapses. The duration d
// must be greater than zero; if not, Reset will panic.

// Modify modifies an existing timer.
// Reports whether the timer was modified before it was run.
// An inactive timer may be passed to Modify to turn into an
// active timer with an updated when, period fields.
// It's OK to call Modify() on a newly allocated Timer.
// An active timer may call Modify(). No fields may be touched. It remains an active timer.

// After waits for the duration to elapse and then sends signal on the returned channel.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, copy the body
// instead and call timer.Stop() if the timer is no longer needed.

// AfterFunc waits for the duration to elapse and then calls f
// in its own goroutine. It returns a Timer that can
// be used to cancel the call using its Stop method.
