/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype

import (
	"libgo/protocol"
)

func RegisterMediaType(mt protocol.MediaType) {
	// Check due to os can be nil almost in tests and benchmarks build
	if protocol.OS != nil {
		protocol.OS.RegisterMediaType(mt)
	}
}

// MediaTypes store all data structure details
type MediaTypes struct{}

// RegisterMediaType register given MediaType to the application
func (dss *MediaTypes) RegisterMediaType(mt protocol.MediaType) {
	if mt.MainType() == "" && mt.SubType() == "" {
		panic("Mediatype doesn't has main and sub type. Can't register empty mediatype.")
	}
	register(mt)
}

func (dss *MediaTypes) GetMediaTypeByID(id protocol.MediaTypeID) protocol.MediaType { return ByID(id) }
func (dss *MediaTypes) GetMediaTypeByFileExtension(ex string) protocol.MediaType {
	return ByFileExtension(ex)
}
func (dss *MediaTypes) GetMediaType(mediaType string) protocol.MediaType {
	return ByMediaType(mediaType)
}

var (
	poolByMediaType     = map[string]protocol.MediaType{}
	poolByID            = map[protocol.MediaTypeID]protocol.MediaType{}
	poolByFileExtension = map[string]protocol.MediaType{}
)

func ByMediaType(mediaType string) protocol.MediaType { return poolByMediaType[mediaType] }
func ByID(id protocol.MediaTypeID) protocol.MediaType { return poolByID[id] }
func ByFileExtension(ex string) protocol.MediaType    { return poolByFileExtension[ex] }

func register(mt protocol.MediaType) {
	// TODO::: lock??
	poolByMediaType[mt.ToString()] = mt
	poolByID[mt.ID()] = mt
	var fe = mt.FileExtension()
	if fe != "" {
		poolByFileExtension[fe] = mt
	}
}
