/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

// GUI indicate how Error capsule must act in GUI applications.
//
// GUI can use also in CLI apps by e.g. use `//go:build gui` tag,
// But strongly suggest DON'T think in this way and just use `Log` concept
type GUI interface {
	// Notify error to user by graphic, sound and vibration (Haptic Feedback)
	Notify()
}
