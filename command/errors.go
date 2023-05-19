/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	er "libgo/error"
)

// Errors
var (
	ErrServiceNotFound     er.Error
	ErrServiceNotAcceptCLI er.Error

	ErrFlagNotFound        er.Error
	ErrFlagBadSyntax       er.Error
	ErrFlagNeedsAnArgument er.Error
)

func init() {
	ErrServiceNotFound.Init("domain/libgo.scm.geniuses.group; type=error; package=command; name=service-not-found")
	ErrServiceNotAcceptCLI.Init("domain/libgo.scm.geniuses.group; type=error; package=command; name=service-not-accept-cli")

	ErrFlagNotFound.Init("domain/libgo.scm.geniuses.group; type=error; package=command; name=flag-not_found")
	ErrFlagBadSyntax.Init("domain/libgo.scm.geniuses.group; type=error; package=command; name=flag-bad_syntax")
	ErrFlagNeedsAnArgument.Init("domain/libgo.scm.geniuses.group; type=error; package=command; name=flag-needs_an_arguments")
}
