/* For license and copyright information please see the LEGAL file in the code repository */

package error

// package errors
var (
	ErrNotFound Error
	ErrNotExist Error
)

func init() {
	ErrNotFound.Init("domain/libgo.scm.geniuses.group; package=error; type=error; name=not_found")
	ErrNotExist.Init("domain/libgo.scm.geniuses.group; package=error; type=error; name=not_exist")
}
