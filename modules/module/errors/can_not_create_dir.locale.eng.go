//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	mp "memar/modules/module/protocol"
)

//memar:impl errCanNotCreateDir memar/protocol.Detail
func (d *errCanNotCreateDir) Domain() string  { return mp.DomainEnglish }
func (d *errCanNotCreateDir) Summary() string { return "Can Not Create Directory" }
func (d *errCanNotCreateDir) Overview() string {
	return ""
}
func (d *errCanNotCreateDir) UserActionNote() string {
	return ""
}
func (d *errCanNotCreateDir) DevActionNote() string {
	return ""
}
func (d *errCanNotCreateDir) TAGS() []string { return []string{} }

//memar:impl errCanNotCreateDir memar/protocol.Quiddity
func (d *errCanNotCreateDir) Name() string         { return "" }
func (d *errCanNotCreateDir) Abbreviation() string { return "" }
func (d *errCanNotCreateDir) Aliases() []string    { return []string{} }
