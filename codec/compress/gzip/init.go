/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	cts "memar/codec/compress-types"
)

func init() {
	GZIP.Init()
	cts.Register(&GZIP)
}
