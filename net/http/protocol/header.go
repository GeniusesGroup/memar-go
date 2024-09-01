/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/string/protocol"
)

// Header indicate HTTP header semantic.
// STR MUST just include ASCII characters.
type Header[STR string_p.String] interface {
	Header_Get(key STR) (value STR)
	Header_Add(key, value STR)
	// Header_Set is same as Header_Del() >> Header_Add()
	Header_Set(key, value STR)
	Header_Del(key STR)
}

// some header fields such as "Set-Cookie", "WWW-Authenticate", "Proxy-Authenticate" break multiple values
// separate by comma and use multi line same key! implementations MUST provide iteration mechanism over all header fields.
// type Header_Iteration[STR String] adt_p.Iteration_KV[STR, STR]
