/* For license and copyright information please see the LEGAL file in the code repository */

package errors

func init() {
	ErrNotFound.Init()
	ErrNotExist.Init()

	// This conditions must be true just in the dev phase.
	ErrNotProvideIdentifier.Init()
	ErrDuplicateIdentifier.Init()
}
