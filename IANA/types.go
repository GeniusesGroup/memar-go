/* For license and copyright information please see LEGAL file in repository */

package iana

// MetaData : The standard struct of MetaData for object
// TODO : Add validator for these fields.
const (
	// Name store name of data
	Name uint32 = iota
)

type metaData struct {
	Title    string
	Comment  string
	MimeType uint16 // Use PersiaDB media types.
	MD5      string
	ReadOnly bool // No one can delete readonly objects! Even in user delete account situation!
	Trash    bool
}

// Service use to store some information about APIs services
type Service struct {
	ServiceID   uint32
	ServiceName string
	Tags        []string
	ExpireDate  uint64
	Status      uint8 // Alfa, Beta, PreStable, Stable,
}
