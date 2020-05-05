/* For license and copyright information please see LEGAL file in repository */

package assets

// File :
type File struct {
	FullName   string
	Name       string
	Extension  string
	MimeType   string
	Dep        *Folder
	Data       []byte
	DataString string
}
