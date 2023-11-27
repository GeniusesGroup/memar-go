/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
)

// Errors
var (
	ErrParse             er.Error
	ErrQueryBadKey       er.Error
	ErrInvalidURLEscape  er.Error
	ErrInvalidHostEscape er.Error
)

func init() {
	ErrParse.Init("domain/uri.wg.ietf.org; type=error; name=parse-uri")
	ErrQueryBadKey.Init("domain/uri.wg.ietf.org; type=error; name=query-bad-key")
	ErrInvalidURLEscape.Init("domain/uri.wg.ietf.org; type=error; name=invalid-url-escape ")
	ErrInvalidHostEscape.Init("domain/uri.wg.ietf.org; type=error; name=invalid-host-escape ")
}
