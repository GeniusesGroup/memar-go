//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package flate

const domainEnglish = "Flate Compress"

//memar:impl memar/protocol.Detail
func (d *deflate) Domain() string         { return domainEnglish }
func (d *deflate) Summary() string        { return "" }
func (d *deflate) Overview() string       { return "" }
func (d *deflate) UserActionNote() string { return "" }
func (d *deflate) DevActionNote() string  { return "" }
func (d *deflate) TAGS() []string         { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *deflate) Name() string         { return "" }
func (d *deflate) Abbreviation() string { return "" }
func (d *deflate) Aliases() []string    { return []string{} }
