/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type CompressTypes interface {
	Register(ct CompressType) (err Error)
	GetByID(id ID) (ct CompressType, err Error)
	GetByMediaType(mt string) (ct CompressType, err Error)
	GetByFileExtension(ex string) (ct CompressType, err Error)
	GetByContentEncoding(ce string) (ct CompressType, err Error)

	ContentEncodings() []string
}
