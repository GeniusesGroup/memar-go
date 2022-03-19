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

func (cts *CompressTypes) GetCompressTypeByID(id uint64) (ct protocol.CompressType, err protocol.Error) {
	ct = ByID(id)
	if ct == nil {
		err = ErrNotFound
	}
	return
}
func (cts *CompressTypes) GetCompressTypeByMediaType(mt string) (ct protocol.CompressType, err protocol.Error) {
	ct = ByMediaType(mt)
	if ct == nil {
		err = ErrNotFound
	}
	return
}
func (cts *CompressTypes) GetCompressTypeByFileExtension(ex string) (ct protocol.CompressType, err protocol.Error) {
	ct = ByFileExtension(ex)
	if ct == nil {
		err = ErrNotFound
	}
	return
}
func (cts *CompressTypes) GetCompressTypeByContentEncoding(ce string) (ct protocol.CompressType, err protocol.Error) {
	ct = ByContentEncoding(ce)
	if ct == nil {
		err = ErrNotFound
	}
	return
}
func (cts *CompressTypes) ContentEncodings() []string { return contentEncodings }

var (
	poolByID              = map[uint64]protocol.CompressType{}
	poolByMediaType       = map[string]protocol.CompressType{}
	poolByFileExtension   = map[string]protocol.CompressType{}
	poolByContentEncoding = map[string]protocol.CompressType{}
	contentEncodings      []string
)

func ByID(id uint64) protocol.CompressType              { return poolByID[id] }
func ByMediaType(mt string) protocol.CompressType       { return poolByMediaType[mt] }
func ByFileExtension(ex string) protocol.CompressType   { return poolByFileExtension[ex] }
func ByContentEncoding(ce string) protocol.CompressType { return poolByContentEncoding[ce] }
func ContentEncodings() []string                        { return contentEncodings }

func register(ct protocol.CompressType) {
	// TODO::: lock??
	poolByID[ct.MediaType().ID()] = ct
	poolByMediaType[ct.MediaType().MediaType()] = ct
	poolByFileExtension[ct.FileExtension()] = ct
	var ce = ct.ContentEncoding()
	poolByContentEncoding[ce] = ct
	contentEncodings = append(contentEncodings, ce)
}
