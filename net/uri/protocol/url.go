/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/string/protocol"
)

// URL indicate "Uniform Resource Locators".
// https://datatracker.ietf.org/doc/html/rfc1738
type URL[STR string_p.String] URI[STR]
