/* For license and copyright information please see the LEGAL file in the code repository */

package detail

import (
	"libgo/protocol"
)

/*
func init() {
	Err__.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("").
		SetOverview("").
		SetUserNote("").
		SetDevNote("").
		SetTAGS([]string{})
	)
}
*/

func New(lang protocol.LanguageID, domain string) (d Detail) {
	d.languageID = lang
	d.domain = domain
	return
}

// Detail store detail about an MediaType that implement protocol.Detail
type Detail struct {
	languageID protocol.LanguageID
	Quiddity
	domain   string
	summary  string
	overview string
	userNote string
	devNote  string
	tags     []string
}

func (d *Detail) Language() protocol.LanguageID { return d.languageID }
func (d *Detail) Domain() string                { return d.domain }
func (d *Detail) Summary() string               { return d.summary }
func (d *Detail) Overview() string              { return d.overview }
func (d *Detail) UserNote() string              { return d.userNote }
func (d *Detail) DevNote() string               { return d.devNote }
func (d *Detail) TAGS() []string                { return d.tags }

// Helper methods to set details more easily
func (d Detail) SetName(v string) Detail         { d.Quiddity.SetName(v); return d }
func (d Detail) SetAbbreviation(v string) Detail { d.Quiddity.SetAbbreviation(v); return d }
func (d Detail) SetAliases(v []string) Detail    { d.Quiddity.SetAliases(v); return d }

func (d Detail) SetSummary(v string) Detail  { d.summary = v; return d }
func (d Detail) SetOverview(v string) Detail { d.overview = v; return d }
func (d Detail) SetUserNote(v string) Detail { d.userNote = v; return d }
func (d Detail) SetDevNote(v string) Detail  { d.devNote = v; return d }
func (d Detail) SetTAGS(v []string) Detail   { d.tags = v; return d }
