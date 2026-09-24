/* For license and copyright information please see the LEGAL file in the code repository */

package compress_p

// Field_ContentEncoding is standard shape of http compress coding type string
type Field_ContentEncoding interface {
	// https://en.wikipedia.org/wiki/HTTP_compression
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
	// https://www.iana.org/assignments/http-parameters/http-parameters.xml#content-coding
	ContentEncoding() string
}
