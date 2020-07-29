/* For license and copyright information please see LEGAL file in repository */

package assets

import (
	"bytes"
	"compress/gzip"
)

// File :
type File struct {
	FullName     string
	Name         string
	Extension    string
	MimeType     string
	Dep          *Folder
	State        uint8
	Data         []byte
	CompressData []byte
	CompressType string
}

// File||Folder State
const (
	StateUnChanged uint8 = iota
	StateChanged
)

// Copy returns a copy of the file.
func (f *File) Copy() *File {
	var file = File{
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		MimeType:  f.MimeType,
		Dep:       f.Dep,
		State:     f.State,
		Data:      f.Data,
	}
	return &file
}

// DeepCopy returns a deep copy of the file.
func (f *File) DeepCopy() *File {
	var file = File{
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		MimeType:  f.MimeType,
		Dep:       f.Dep,
		State:     f.State,
	}
	copy(file.Data, f.Data)
	return &file
}

// Supported compress types
const (
	CompressTypeGZIP = "gzip"
)

// Compress use to compress file data to mostly to use in serving file by servers.
func (f *File) Compress(compressType string) {
	// Check file type and compress just if it worth.
	switch f.Extension {
	case "png", "jpg", "gif", "jpeg", "mkv", "avi", "mp3", "mp4":
		f.CompressData = f.Data
		return
	}

	f.CompressType = compressType
	switch compressType {
	case "gzip":
		var b bytes.Buffer
		var gz = gzip.NewWriter(&b)
		gz.Write(f.Data)
		gz.Close()
		f.CompressData = b.Bytes()
	}
}
