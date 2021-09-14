/* For license and copyright information please see LEGAL file in repository */

package error

import "../protocol"

// Detail store detail about an error
type Detail struct {
	languageID protocol.LanguageID
	domain     string // Locale domain name that error belongs to it!
	short      string // Locale general short error detail
	long       string // Locale general long error detail
	userAction string // Locale user action that user do when face this error
	devAction  string // Locale technical advice for developers
}

func (d *Detail) Language() protocol.LanguageID { return d.languageID }
func (d *Detail) Domain() string                { return d.domain }
func (d *Detail) Short() string                 { return d.short }
func (d *Detail) Long() string                  { return d.long }
func (d *Detail) UserAction() string            { return d.userAction }
func (d *Detail) DevAction() string             { return d.devAction }
