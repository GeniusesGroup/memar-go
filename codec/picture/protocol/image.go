/* For license and copyright information please see the LEGAL file in the code repository */

package picture_p

import (
	buffer_p "memar/buffer/protocol"
)

type Image interface {
	Format()
	Image() buffer_p.Buffer
}
