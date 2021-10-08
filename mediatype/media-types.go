/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"../protocol"
)

// The currently registered types
const (
	MimeSubTypeApplication = "application"
	MimeSubTypeAudio       = "audio"
	MimeSubTypeFont        = "font"
	MimeSubTypeExample     = "example"
	MimeSubTypeImage       = "image"
	MimeSubTypeMessage     = "message"
	MimeSubTypeModel       = "model"
	MimeSubTypeMultipart   = "multipart"
	MimeSubTypeText        = "text"
	MimeSubTypeVideo       = "video"
)

// MediaTypes store all data structure details
type MediaTypes struct{}

// RegisterMediaType register given datastructure
func (dss *MediaTypes) RegisterMediaType(mt protocol.MediaType) {
	if mt.URN().URI() == "" {
		protocol.App.LogFatal("Data structure '", mt.URN().ID, "' doesn't has an URN! It is rule to add more detail about data structure before register it!")
	}
	if mt.URN().ID() == 0 {
		protocol.App.LogWarn("Data structure '", mt.URN, "', give 0 as data structure ID! it won't register to use by ID! legal ID must hash of data structure URN")
		return
	}

	poolByID[mt.URN().ID()] = mt
	poolByURN[mt.URN().URI()] = mt
	poolByMediaType[mt.MediaType()] = mt
	poolByFileExtension[mt.FileExtension()] = mt
}

func (dss *MediaTypes) GetMediaTypeByID(id uint64) protocol.MediaType   { return ByID(id) }
func (dss *MediaTypes) GetMediaTypeByURN(urn string) protocol.MediaType { return ByURN(urn) }
func (dss *MediaTypes) GetMediaTypeByFileExtension(ex string) protocol.MediaType {
	return ByFileExtension(ex)
}
func (dss *MediaTypes) GetMediaTypeByType(mediaType string) protocol.MediaType {
	return ByType(mediaType)
}

var (
	poolByID            = map[uint64]protocol.MediaType{}
	poolByURN           = map[string]protocol.MediaType{}
	poolByFileExtension = map[string]protocol.MediaType{}
	poolByMediaType     = map[string]protocol.MediaType{}
)

func ByID(id uint64) protocol.MediaType            { return poolByID[id] }
func ByURN(urn string) protocol.MediaType          { return poolByURN[urn] }
func ByFileExtension(ex string) protocol.MediaType { return poolByFileExtension[ex] }
func ByType(mediaType string) protocol.MediaType   { return poolByMediaType[mediaType] }
