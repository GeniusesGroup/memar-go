/* For license and copyright information please see LEGAL file in repository */

package assets

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
