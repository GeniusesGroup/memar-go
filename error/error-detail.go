/* For license and copyright information please see LEGAL file in repository */

package error

import "../protocol"

// Detail store detail about an error that impelement protocol.ErrorDetail
type Detail struct {
	languageID protocol.LanguageID
	domain     string
	summary    string
	overview   string
	userAction string
	devAction  string
}

func (d *Detail) Language() protocol.LanguageID { return d.languageID }
func (d *Detail) Domain() string                { return d.domain }
func (d *Detail) Summary() string               { return d.summary }
func (d *Detail) Overview() string              { return d.overview }
func (d *Detail) UserAction() string            { return d.userAction }
func (d *Detail) DevAction() string             { return d.devAction }
