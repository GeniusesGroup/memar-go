/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"strings"

	"../protocol"
	"../urn"
)

// MediaType implement protocol.MediaType interface
type mediaType struct {
	urn            urn.Giti // Use ID instead of other data to improve efficiency
	mediaType      string
	mainType       string
	subType        string
	extension      string // Use as file extension usually in windows os
	description    string
	reference      string
	registeredDate int64
	approvedDate   int64
}

func (mt *mediaType) URN() protocol.GitiURN { return &mt.urn }
func (mt *mediaType) MediaType() string     { return mt.mediaType }
func (mt *mediaType) MainType() string      { return mt.mainType }
func (mt *mediaType) SubType() string       { return mt.subType }
func (mt *mediaType) FileExtension() string { return mt.extension }

func newMediaType(urn, media, fileExtension, description string) (mt *mediaType) {
	mt = &mediaType{
		mediaType:   media,
		extension:   fileExtension,
		description: description,
	}
	mt.mainType, mt.subType = splitMediaType(media)
	mt.urn.Init(urn)

	mediaTypeByID[mt.urn.ID()] = mt
	mediaTypeByType[mt.mediaType] = mt
	mediaTypeByFileExtension[mt.extension] = mt
	return
}

func splitMediaType(mt string) (mainType, subType string) {
	var i = strings.IndexByte(mt, '/')
	mainType = mt[:i]
	subType = mt[i+1:]
	return
}
