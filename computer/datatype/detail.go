/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

// Detail embed to provide some methods when not need to implement by others.
// Usually add in non language specific capsules and implement them in each desire language e.g. errors, services, ...
type Detail struct{}

//memar:impl memar/protocol.Detail
func (self *Detail) Domain() string   { return "" }
func (self *Detail) Summary() string  { return "" }
func (self *Detail) Overview() string { return "" }
func (self *Detail) UserNote() string { return "" }
func (self *Detail) DevNote() string  { return "" }
func (self *Detail) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (self *Detail) Name() string         { return "" }
func (self *Detail) Abbreviation() string { return "" }
func (self *Detail) Aliases() []string    { return []string{} }
