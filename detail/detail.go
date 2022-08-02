/* For license and copyright information please see LEGAL file in repository */

package detail

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// Detail store detail about an MediaType that implement protocol.Detail
type Detail struct {
	languageID protocol.LanguageID
	domain     string
	summary    string
	overview   string
	userNote   string
	devNote    string
	tags       []string
}

func (d *Detail) Language() protocol.LanguageID { return d.languageID }
func (d *Detail) Domain() string                { return d.domain }
func (d *Detail) Summary() string               { return d.summary }
func (d *Detail) Overview() string              { return d.overview }
func (d *Detail) UserNote() string              { return d.userNote }
func (d *Detail) DevNote() string               { return d.devNote }
func (d *Detail) TAGS() []string                { return d.tags }
