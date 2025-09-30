/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	cts "memar/compress-types"
)

func init() {
	Deflate.Init()
	cts.Register(&Deflate)
}
