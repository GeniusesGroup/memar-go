/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

type GUI interface {
	// Notify error to user by graphic, sound and vibration (Haptic Feedback)
	Notify()
}
