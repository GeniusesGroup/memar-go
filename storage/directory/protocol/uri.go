/* For license and copyright information please see the LEGAL file in the code repository */

package directory_p

// Helper functions
type Method_URI_Directory interface {
	// URI path end with "/" if it is a file directory
	// or has `text/directory` mimetype
	IsDirectory() bool 
}
