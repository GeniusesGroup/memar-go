/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../protocol"
)

// CompressTypes store all compress types to use by anyone want protocol.CompressTypes
type CompressTypes struct{}

// RegisterCompressType register given CompressType
func (cts *CompressTypes) RegisterCompressType(ct protocol.CompressType) {
	if ct.MediaType() == nil {
		panic("CompressType doesn't has a valid MediaType. Can't register it.")
	}
	if ct.ContentEncoding() == "" {
		panic("CompressType doesn't has a valid ContentEncoding. Can't register it.")
	}
	register(ct)
}

func (cts *CompressTypes) GetCompressTypeByID(id uint64) protocol.CompressType { return ByID(id) }
func (cts *CompressTypes) GetCompressTypeByMediaType(mt string) protocol.CompressType {
	return ByMediaType(mt)
}
func (cts *CompressTypes) GetCompressTypeByFileExtension(ex string) protocol.CompressType {
	return ByFileExtension(ex)
}
func (cts *CompressTypes) GetCompressTypeByContentEncoding(ce string) protocol.CompressType {
	return ByContentEncoding(ce)
}

var (
	poolByID              = map[uint64]protocol.CompressType{}
	poolByMediaType       = map[string]protocol.CompressType{}
	poolByFileExtension   = map[string]protocol.CompressType{}
	poolByContentEncoding = map[string]protocol.CompressType{}
)

func ByID(id uint64) protocol.CompressType              { return poolByID[id] }
func ByMediaType(mt string) protocol.CompressType       { return poolByMediaType[mt] }
func ByFileExtension(ex string) protocol.CompressType   { return poolByFileExtension[ex] }
func ByContentEncoding(ce string) protocol.CompressType { return poolByContentEncoding[ce] }

func register(ct protocol.CompressType) {
	// TODO::: lock??
	poolByID[ct.MediaType().ID()] = ct
	poolByMediaType[ct.MediaType().MediaType()] = ct
	poolByFileExtension[ct.FileExtension()] = ct
	poolByContentEncoding[ct.ContentEncoding()] = ct
}
