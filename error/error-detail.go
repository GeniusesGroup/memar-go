/* For license and copyright information please see LEGAL file in repository */

package error

// Detail store detail about an error
type Detail struct {
	domain     string // Locale domain name that error belongs to it!
	short      string // Locale general short error detail
	long       string // Locale general long error detail
	userAction string // Locale user action that user do when face this error
	devAction  string // Locale technical advice for developers
}

// Domain return locale domain name that error belongs to it!
func (d Detail) Domain() string {
	return d.domain
}

// Short return locale general short error detail
func (d Detail) Short() string {
	return d.short
}

// Long return locale general long error detail
func (d Detail) Long() string {
	return d.long
}

// UserAction return locale user action that user do when face this error
func (d Detail) UserAction() string {
	return d.userAction
}

// DevAction return locale technical advice for developers
func (d Detail) DevAction() string {
	return d.devAction
}
