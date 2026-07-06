/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/protocol"
)

// ContentType read all value about content in header
func (h *Header) ContentType() (c ContentType) {
	var contentType = h.Header_Get(HeaderKey_ContentType)
	c.FromString(contentType)
	return
}
func (h *Header) AddContentType(ct string) {
	h.Header_Add(HeaderKey_ContentType, ct)
}

// ContentType store
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
type ContentType struct {
	Type     mimeType
	SubType  mimeSubType
	Charset  string
	Boundary string
}

//memar:impl memar/protocol.Stringer_From
func (ct *ContentType) FromString(contentType string) (err protocol.Error) {
	var mediaTypeFirst, mediaTypeSecond string
	var index int
	for i := 0; i < len(contentType); i++ {
		switch contentType[i] {
		case '/':
			mediaTypeFirst = contentType[:i]
			index = i + 1
		case ';':
			mediaTypeSecond = contentType[index:i]
			contentType = contentType[i+2:] // +2 due to have space after semicolon
		}
	}

	switch contentType[0] {
	case 'c': // charset=
		ct.Charset = contentType[8:]
	case 'b': // boundary=
		ct.Boundary = contentType[9:]
	}

	switch mediaTypeFirst {
	case "text":
		ct.Type = ContentTypeMimeTypeText
	case "application":
		ct.Type = ContentTypeMimeTypeApplication
	}

	switch mediaTypeSecond {
	case "html":
		ct.SubType = ContentTypeMimeSubTypeHTML
	case "json":
		ct.SubType = ContentTypeMimeSubTypeJSON
	}

	return
}

type mimeType uint16

// Standard HTTP content type
const (
	ContentTypeMimeTypeUnset mimeType = iota
	ContentTypeMimeTypeText
	ContentTypeMimeTypeApplication
)

type mimeSubType uint16

// Standard HTTP content type
const (
	ContentTypeMimeSubTypeUnset mimeSubType = iota
	ContentTypeMimeSubTypeHTML
	ContentTypeMimeSubTypeJSON
)
