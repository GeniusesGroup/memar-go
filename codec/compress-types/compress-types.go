/* For license and copyright information please see the LEGAL file in the code repository */

package cts

import (
	errs "memar/compress-types/errors"
	"memar/protocol"
)

var cts = compressTypes{
	poolByID:              map[protocol.ID]protocol.CompressType{},
	poolByMediaType:       map[string]protocol.CompressType{},
	poolByFileExtension:   map[string]protocol.CompressType{},
	poolByContentEncoding: map[string]protocol.CompressType{},
}

func Register(ct protocol.CompressType) (err protocol.Error)                { return cts.Register(ct) }
func GetByID(id protocol.ID) (ct protocol.CompressType, err protocol.Error) { return cts.GetByID(id) }
func GetByMediaType(mt string) (ct protocol.CompressType, err protocol.Error) {
	return cts.GetByMediaType(mt)
}
func GetByFileExtension(ex string) (ct protocol.CompressType, err protocol.Error) {
	return cts.GetByFileExtension(ex)
}
func GetByContentEncoding(ce string) (ct protocol.CompressType, err protocol.Error) {
	return cts.GetByContentEncoding(ce)
}
func ContentEncodings() []string { return cts.ContentEncodings() }

type compressTypes struct {
	poolByID              map[protocol.ID]protocol.CompressType
	poolByMediaType       map[string]protocol.CompressType
	poolByFileExtension   map[string]protocol.CompressType
	poolByContentEncoding map[string]protocol.CompressType
	contentEncodings      []string
}

// RegisterCompressType register given CompressType
func (cts *compressTypes) Register(ct protocol.CompressType) (err protocol.Error) {
	// TODO::: change panics to error
	if ct.MediaType() == "" {
		panic("CompressType doesn't has a valid MediaType. Can't register it.")
	}
	if ct.ContentEncoding() == "" {
		panic("CompressType doesn't has a valid ContentEncoding. Can't register it.")
	}

	// TODO::: lock??
	cts.poolByID[ct.ID()] = ct
	cts.poolByMediaType[ct.ToString()] = ct
	cts.poolByFileExtension[ct.FileExtension()] = ct
	var ce = ct.ContentEncoding()
	cts.poolByContentEncoding[ce] = ct
	cts.contentEncodings = append(cts.contentEncodings, ce)
	return
}

func (cts *compressTypes) GetByID(id protocol.ID) (ct protocol.CompressType, err protocol.Error) {
	ct = cts.poolByID[id]
	if ct == nil {
		err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByMediaType(mt string) (ct protocol.CompressType, err protocol.Error) {
	ct = cts.poolByMediaType[mt]
	if ct == nil {
		err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByFileExtension(ex string) (ct protocol.CompressType, err protocol.Error) {
	ct = cts.poolByFileExtension[ex]
	if ct == nil {
		err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) GetByContentEncoding(ce string) (ct protocol.CompressType, err protocol.Error) {
	ct = cts.poolByContentEncoding[ce]
	if ct == nil {
		err = &errs.ErrNotFound
	}
	return
}
func (cts *compressTypes) ContentEncodings() []string { return cts.contentEncodings }
