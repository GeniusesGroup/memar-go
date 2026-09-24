/* For license and copyright information please see the LEGAL file in the code repository */

package file

import (
	"path/filepath"

	container_p "memar/computer/adt/container/protocol"
	"memar/computer/buffer/byteslice/convert"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
)

// URI implement uri_p.URI and file_p.URI interface
type URI struct {
	uri       string
	domain    string
	path      string
	name      string
	extension string
}

func (uri *URI) Init(u string) {
	var dir, fileName = filepath.Split(u)
	uri.path = dir
	uri.Rename(fileName)
	// TODO::: decode more data
}

//memar:impl memar/storage/file/protocol.URI
func (uri *URI) URI() string                  { return uri.uri }
func (uri *URI) Scheme() string               { return "file" }
func (uri *URI) Domain() string               { return uri.domain }
func (uri *URI) Path() string                 { return uri.path }
func (uri *URI) Name() string                 { return uri.name }
func (uri *URI) NameWithoutExtension() string { return uri.name[:len(uri.name)-len(uri.extension)] }
func (uri *URI) IsDirectory() bool            { return IsPathDirectory(uri.path) }
func (uri *URI) PathParts() []string          { return FilePathParts(uri.path) }

//memar:impl memar/storage/protocol.FileExtension
func (uri *URI) FilExtension() string { return uri.extension }

// Rename just rename the file name
func (uri *URI) Rename(newName string) {
	uri.name = newName
	uri.extension = Extension(newName)
}

//memar:impl memar/codec/protocol.Codec
func (uri *URI) MediaType() mediatype_p.MediaType { return MediaType.MediaType() }

// Marshal encodes and return whole file data.
func (uri *URI) Marshal() (data []byte, err error_p.Error) {
	var ln = 5
	data = make([]byte, 0, ln)
	return
}

// MarshalTo encodes whole file data to given data and return it with new len.
func (uri *URI) MarshalTo(data []byte) (added []byte, err error_p.Error) {
	return data, nil
}

// Unmarshal save given data to the file. It will overwritten exiting data.
func (uri *URI) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	var dataAsString = convert.UnsafeByteSliceToString(source)
	var dir, fileName = filepath.Split(dataAsString)
	uri.path = dir
	uri.name = fileName
	// TODO::: decode more data
	return
}

func Extension(name string) (ex string) {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			ex = name[i+1:]
			break
		}
	}
	return
}

func FilePathParts(uriPath string) (parts []string) {
	parts = make([]string, 0, 4)
	for i := 0; i < len(uriPath); i++ {
		if uriPath[i] == '/' {
			parts = append(parts, uriPath[:i])
			uriPath = uriPath[i+1:]
			i = 0
		}
	}
	parts = append(parts, uriPath)
	return
}

func IsPathDirectory(uriPath string) bool { return uriPath[len(uriPath)-1] == '/' }
