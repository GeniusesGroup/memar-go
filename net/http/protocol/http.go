/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

// Other frameworks:
// - https://www.php-fig.org/psr/psr-7/
// - https://nodejs.org/api/http.html#requestgetheaders

// Request indicate HTTP request semantic.
// 
// Other frameworks:
// https://developer.mozilla.org/en-US/docs/Web/API/Request
type Request /*[STR String]*/ interface {
	PseudoHeader_Request
	Header

	Field_Body
	Method_Body
}

// Response indicate HTTP response semantic.
// 
// Other frameworks:
// https://developer.mozilla.org/en-US/docs/Web/API/Response
type Response /*[STR String]*/ interface {
	PseudoHeader_Response
	Header

	Field_Body
	Method_Body
}
