/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"../protocol"
)

// StandardMimeTypes : All standard mime types.
// The currently registered types are: application, audio, example, font, image, message, model, multipart, text and video.
const (
	MimeSubTypeApplication = "application"
	MimeSubTypeAudio       = "audio"
	MimeSubTypeFont        = "font"
	MimeSubTypeImage       = "image"
	MimeSubTypeMessage     = "message"
	MimeSubTypeText        = "text"
	MimeSubTypeVideo       = "video"
)

var (
	mediaTypeByID            = map[uint64]*mediaType{}
	mediaTypeByFileExtension = map[string]*mediaType{}
	mediaTypeByType          = map[string]*mediaType{}
)

func MediaTypeByID(id uint64) protocol.MediaType            { return mediaTypeByID[id] }
func MediaTypeByFileExtension(ex string) protocol.MediaType { return mediaTypeByFileExtension[ex] }
func MediaTypeByType(mediatype string) protocol.MediaType   { return mediaTypeByType[mediatype] }
