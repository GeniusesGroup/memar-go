/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	http_p "memar/libgo/net/http/protocol"
)

var _ http_p.Request = &Request{}
