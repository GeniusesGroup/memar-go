/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

// MediaTypeDetail store detail about an MediaType that impelement protocol.MediaTypeDetail
type MediaTypeDetail struct {
	languageID protocol.LanguageID
	domain     string
	summary    string
	overview   string
	userNote   string
	devNote    string
	tags       []string
}

func (d *MediaTypeDetail) Language() protocol.LanguageID { return d.languageID }
func (d *MediaTypeDetail) Domain() string                { return d.domain }
func (d *MediaTypeDetail) Summary() string               { return d.summary }
func (d *MediaTypeDetail) Overview() string              { return d.overview }
func (d *MediaTypeDetail) UserNote() string              { return d.userNote }
func (d *MediaTypeDetail) DevNote() string               { return d.devNote }
func (d *MediaTypeDetail) TAGS() []string                { return d.tags }
