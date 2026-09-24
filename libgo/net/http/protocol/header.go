/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	array_p "memar/computer/adt/array/protocol"
)

type Field_Header interface {
	Header() Header
}

// Header indicate HTTP header semantic.
// STR MUST just include ASCII characters.
//
// Other frameworks:
// https://developer.mozilla.org/en-US/docs/Web/API/Headers
type Header interface {
	Header_Get(key string) (value string)
	Header_Add(key, value string)
	// Header_Set is same as Header_Del() >> Header_Add()
	Header_Set(key, value string)
	Header_Del(key string)

	// some header fields such as "Set-Cookie", "WWW-Authenticate", "Proxy-Authenticate" break multiple values
	// separate by comma and use multi line same key! implementations MUST provide iteration mechanism over all header fields.
	Header_Iteration() array_p.Iteration_KV[string, string]
}
