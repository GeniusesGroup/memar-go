/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../protocol"
)

var (
	compressTypeByID              = map[uint64]*compressType{}
	compressTypeByContentEncoding = map[string]*compressType{}
	compressTypeByFileExtension   = map[string]*compressType{}
)

func CompressTypeByID(id uint64) protocol.CompressType { return compressTypeByID[id] }
func CompressTypeByContentEncoding(ce string) protocol.CompressType {
	return compressTypeByContentEncoding[ce]
}
func CompressTypeByFileExtension(fe string) protocol.CompressType {
	return compressTypeByFileExtension[fe]
}
