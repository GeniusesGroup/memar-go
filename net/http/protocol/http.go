/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/string/protocol"
)

// Other languages:
// - https://www.php-fig.org/psr/psr-7/
// - https://nodejs.org/api/http.html#requestgetheaders

// Request indicate HTTP request semantic.
type Request /*[STR String]*/ interface {
	PseudoHeader_Request[string_p.String]
	Header[string_p.String]
	Body
}

// Response indicate HTTP response semantic.
type Response /*[STR String]*/ interface {
	PseudoHeader_Response[string_p.String]
	Header[string_p.String]
	Body
}
