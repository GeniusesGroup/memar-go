/* For license and copyright information please see the LEGAL file in the code repository */

package errs

func init() {
	ErrServiceNotFound.Init()
	ErrServiceCallByAlias.Init()
	ErrServiceNotAcceptCLI.Init()

	ErrFlagNotFound.Init()
	ErrFlagBadSyntax.Init()
	ErrFlagNeedsAnArgument.Init()
}
