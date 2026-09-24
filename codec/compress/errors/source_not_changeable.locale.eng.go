//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (self *sourceNotChangeable) Domain() string  { return domainEnglish }
func (self *sourceNotChangeable) Summary() string { return "Source not Changeable" }
func (self *sourceNotChangeable) Overview() string {
	return "Can't read from other source than source given in compression||decompression creation"
}
func (self *sourceNotChangeable) UserActionNote() string { return "" }
func (self *sourceNotChangeable) DevActionNote() string  { return "" }
func (self *sourceNotChangeable) TAGS() []string         { return []string{} }

//memar:impl memar/protocol.Quiddity
func (self *sourceNotChangeable) Name() string         { return "" }
func (self *sourceNotChangeable) Abbreviation() string { return "" }
func (self *sourceNotChangeable) Aliases() []string    { return []string{} }
