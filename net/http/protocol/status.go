/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Status /*[STR string_p.String]*/ interface {
	StatusCode() string_p.String

	// The Reason-Phrase is indeed optional.
	// HTTP/2 and HTTP/3 even dropped it entirely.
	ReasonPhrase() string_p.String
}

type Method_Status /*[STR string_p.String]*/ interface {
	// TODO::: How can we made reasonPhrase optional? below is good?
	// If pass just one status, it means status code.
	// If pass two args, means code and phrase.
	// It will ignore more than two args.
	// SetStatus(status ...STR)
	SetStatus(statusCode, reasonPhrase string_p.String)
}
