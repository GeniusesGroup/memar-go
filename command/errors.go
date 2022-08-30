/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	er "github.com/GeniusesGroup/libgo/error"
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
	ErrServiceNotAcceptCLI.Init("domain/geniuses.group; type=error; package=command; name=service-not-accept-cli")

	ErrFlagNotFound.Init("domain/geniuses.group; type=error; package=command; name=flag-not_found")
	ErrFlagBadSyntax.Init("domain/geniuses.group; type=error; package=command; name=flag-bad_syntax")
	ErrFlagNeedsAnArgument.Init("domain/geniuses.group; type=error; package=command; name=flag-needs_an_arguments")
}
