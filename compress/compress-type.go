/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../protocol"
	"../urn"
)

type compressType struct {
	urn             urn.Giti
	contentEncoding string
	extension       string // Use as file extension usually in windows os
	compression     compression
	decompression   decompression
}

type compression func(raw protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec)
type decompression func(compress protocol.Codec) (raw protocol.Codec)

func (ct *compressType) URN() protocol.GitiURN   { return &ct.urn }
func (ct *compressType) ContentEncoding() string { return ct.contentEncoding }
func (ct *compressType) Extension() string       { return ct.extension }
func (ct *compressType) Compression(raw protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec) {
	return ct.compression(raw, compressLevel)
}
func (ct *compressType) Decompression(compress protocol.Codec) (raw protocol.Codec) {
	return ct.decompression(raw)
}

func NewCompressType(urn, contentEncoding, extension string, compression compression, decompression decompression) (ct *compressType) {
	ct = &compressType{
		contentEncoding: contentEncoding,
		extension:       extension,
		compression:     compression,
		decompression:   decompression,
	}
	ct.urn.Init(urn)

	compressTypeByID[ct.urn.ID()] = ct
	compressTypeByContentEncoding[ct.contentEncoding] = ct
	compressTypeByFileExtension[ct.extension] = ct
	return
}
