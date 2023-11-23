/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type MediaTypes interface {
	Register(mt MediaType) (err Error) 
	GetByMediaType(mediaType string) (mt MediaType, err Error)
	GetByID(id MediaTypeID) (mt MediaType, err Error)
	GetByFileExtension(ex string) (mt MediaType, err Error)
}
