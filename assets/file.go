/* For license and copyright information please see LEGAL file in repository */

package assets

import "unsafe"

// File :
type File struct {
	FullName   string
	Name       string
	Extension  string
	MimeType   string
	Dep        *Folder
	State      uint8
	Data       []byte
	DataString string
}

// File||Folder State
const (
	StateUnChanged uint8 = iota
	StateChanged
)

// Copy returns a copy of the file.
func (f *File) Copy() *File {
	var file = File{
		FullName:   f.FullName,
		Name:       f.Name,
		Extension:  f.Extension,
		MimeType:   f.MimeType,
		Dep:        f.Dep,
		State:      f.State,
		Data:       f.Data,
		DataString: f.DataString,
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
	file.DataString = *(*string)(unsafe.Pointer(&file.Data))
	return &file
}
