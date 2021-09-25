/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"path/filepath"

	"../convert"
	"../protocol"
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

func IsPathDirectory(uriPath string) bool { return uriPath[len(uriPath)-1] == '/' }

// Rename just rename the file name
func (uri *URI) Rename(newName string) {
	uri.name = newName
	for i := len(newName) - 1; i >= 0; i-- {
		if newName[i] == '.' {
			uri.extension = newName[i+1:]
		}
	}
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

/*
********** protocol.Codec interface **********
 */

func (uri *URI) MediaType() string    { return "application/uri" }
func (uri *URI) CompressType() string { return "" }

// Marshal enecodes and return whole file data!
func (uri *URI) Marshal() (data []byte) {
	var ln = 5
	data = make([]byte, 0, ln)
	return
}

// MarshalTo enecodes whole file data to given data and return it with new len!
func (uri *URI) MarshalTo(data []byte) []byte {
	return data
}

// Unmarshal save given data to the file. It will overwritten exiting data.
func (uri *URI) Unmarshal(data []byte) (err protocol.Error) {
	var dataAsString = convert.UnsafeByteSliceToString(data)
	var dir, fileName = filepath.Split(dataAsString)
	uri.path = dir
	uri.name = fileName
	// TODO::: decode more data
	return
}
