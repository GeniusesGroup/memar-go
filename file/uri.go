/* For license and copyright information please see the LEGAL file in the code repository */

package file

import (
	"path/filepath"

	"memar/convert"
	"memar/protocol"
)

// URI implement protocol.URI and protocol.FileURI interface
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

func (uri *URI) URI() string                  { return uri.uri }
func (uri *URI) Scheme() string               { return "file" }
func (uri *URI) Domain() string               { return uri.domain }
func (uri *URI) Path() string                 { return uri.path }
func (uri *URI) Name() string                 { return uri.name }
func (uri *URI) NameWithoutExtension() string { return uri.name[:len(uri.name)-len(uri.extension)] }
func (uri *URI) Extension() string            { return uri.extension }
func (uri *URI) IsDirectory() bool            { return IsPathDirectory(uri.path) }
func (uri *URI) PathParts() []string          { return FilePathParts(uri.path) }

// Rename just rename the file name
func (uri *URI) Rename(newName string) {
	uri.name = newName
	uri.extension = Extension(newName)
}

//memar:impl memar/protocol.Codec
func (uri *URI) MediaType() protocol.MediaType       { return &MediaType }
func (uri *URI) CompressType() protocol.CompressType { return nil }

// Marshal encodes and return whole file data.
func (uri *URI) Marshal() (data []byte, err protocol.Error) {
	var ln = 5
	data = make([]byte, 0, ln)
	return
}

// MarshalTo encodes whole file data to given data and return it with new len.
func (uri *URI) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	return data, nil
}

// Unmarshal save given data to the file. It will overwritten exiting data.
func (uri *URI) Unmarshal(data []byte) (n int, err protocol.Error) {
	var dataAsString = convert.UnsafeByteSliceToString(data)
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
