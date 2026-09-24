/* For license and copyright information please see the LEGAL file in the code repository */

package cts

import (
	compress_p "memar/codec/compress/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

var cts = compressTypes{
	poolByID:              map[datatype_p.ID]compress_p.CompressType{},
	poolByMediaType:       map[string]compress_p.CompressType{},
	poolByFileExtension:   map[string]compress_p.CompressType{},
	poolByContentEncoding: map[string]compress_p.CompressType{},
}

func Register(ct compress_p.CompressType) (err error_p.Error) { return cts.Register(ct) }
func GetByID(id datatype_p.ID) (ct compress_p.CompressType, err error_p.Error) {
	return cts.GetByID(id)
}
func GetByMediaType(mt string) (ct compress_p.CompressType, err error_p.Error) {
	return cts.GetByMediaType(mt)
}
func GetByFileExtension(ex string) (ct compress_p.CompressType, err error_p.Error) {
	return cts.GetByFileExtension(ex)
}
func GetByContentEncoding(ce string) (ct compress_p.CompressType, err error_p.Error) {
	return cts.GetByContentEncoding(ce)
}
func ContentEncodings() []string { return cts.ContentEncodings() }

type compressTypes struct {
	poolByID              map[datatype_p.ID]compress_p.CompressType
	poolByMediaType       map[string]compress_p.CompressType
	poolByFileExtension   map[string]compress_p.CompressType
	poolByContentEncoding map[string]compress_p.CompressType
	contentEncodings      []string
}

// RegisterCompressType register given CompressType
func (cts *compressTypes) Register(ct compress_p.CompressType) (err error_p.Error) {
	// TODO::: change panics to error
	if ct.MediaType() == "" {
		panic("CompressType doesn't has a valid MediaType. Can't register it.")
	}
	if ct.ContentEncoding() == "" {
		panic("CompressType doesn't has a valid ContentEncoding. Can't register it.")
	}

	// TODO::: lock??
	cts.poolByID[ct.DataTypeID()] = ct
	cts.poolByMediaType[ct.MediaType()] = ct
	cts.poolByFileExtension[ct.FileExtension()] = ct
	var ce = ct.ContentEncoding()
	cts.poolByContentEncoding[ce] = ct
	cts.contentEncodings = append(cts.contentEncodings, ce)
	return
}

func (cts *compressTypes) GetByID(id datatype_p.ID) (ct compress_p.CompressType, err error_p.Error) {
	ct = cts.poolByID[id]
	if ct == nil {
		// err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByMediaType(mt string) (ct compress_p.CompressType, err error_p.Error) {
	ct = cts.poolByMediaType[mt]
	if ct == nil {
		// err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByFileExtension(ex string) (ct compress_p.CompressType, err error_p.Error) {
	ct = cts.poolByFileExtension[ex]
	if ct == nil {
		// err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByContentEncoding(ce string) (ct compress_p.CompressType, err error_p.Error) {
	ct = cts.poolByContentEncoding[ce]
	if ct == nil {
		// err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) ContentEncodings() []string { return cts.contentEncodings }
