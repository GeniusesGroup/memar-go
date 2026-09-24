/* For license and copyright information please see the LEGAL file in the code repository */

package file

// Metadata : The standard struct of Metadata for object
// TODO : Add validator for these fields.
const (
	// Name store name of data
	Name uint32 = iota
)

type metaData struct {
	Title    string
	Comment  string
	MimeType uint16 // Use media types.
	MD5      string
	ReadOnly bool // No one can delete readonly objects! Even in user delete account situation!
	Trash    bool
}
