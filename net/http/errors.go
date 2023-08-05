/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	er "memar/error"
)

// Declare package errors
var (
	ErrNoConnection        er.Error
	ErrPacketTooShort      er.Error
	ErrPacketTooLong       er.Error
	ErrParseMethod         er.Error
	ErrParseVersion        er.Error
	ErrParseStatusCode     er.Error
	ErrParseReasonPhrase   er.Error
	ErrParseHeaderTooLarge er.Error

	ErrCookieBadName   er.Error
	ErrCookieBadValue  er.Error
	ErrCookieBadPath   er.Error
	ErrCookieBadDomain er.Error

	ErrNotFound             er.Error
	ErrUnsupportedMediaType er.Error

	ErrBodySizeMismatch er.Error
)

func init() {
	ErrNoConnection.Init("domain/http.wg.ietf.org; type=error; name=no-connection")
	ErrPacketTooShort.Init("domain/http.wg.ietf.org; type=error; name=packet-too-short")
	ErrPacketTooLong.Init("domain/http.wg.ietf.org; type=error; name=packet-too-long")
	ErrParseMethod.Init("domain/http.wg.ietf.org; type=error; name=parse-method")
	ErrParseVersion.Init("domain/http.wg.ietf.org; type=error; name=parse-version")
	ErrParseStatusCode.Init("domain/http.wg.ietf.org; type=error; name=parse-status-code")
	ErrParseReasonPhrase.Init("domain/http.wg.ietf.org; type=error; name=parse-reason-phrase")
	ErrParseHeaderTooLarge.Init("domain/http.wg.ietf.org; type=error; name=parse-header-too-large")

	ErrCookieBadName.Init("domain/http.wg.ietf.org; type=error; name=cookie-bad-name")
	ErrCookieBadValue.Init("domain/http.wg.ietf.org; type=error; name=cookie-bad-value")
	ErrCookieBadPath.Init("domain/http.wg.ietf.org; type=error; name=cookie-bad-path")
	ErrCookieBadDomain.Init("domain/http.wg.ietf.org; type=error; name=cookie-bad-domain")

	ErrNotFound.Init("domain/http.wg.ietf.org; type=error; name=not-found")
	ErrUnsupportedMediaType.Init("domain/http.wg.ietf.org; type=error; name=unsupported-media-type")

	ErrBodySizeMismatch.Init("domain/http.wg.ietf.org; type=error; name=body-size-mismatched")
}
