/* For license and copyright information please see the LEGAL file in the code repository */

package ap

import (
	container_p "memar/adt/container/protocol"
	function_p "memar/computer/function/protocol"
	error_p "memar/process/error/protocol"
)

type Domain string

//memar:impl computer/capsule/protocol.Accessor
func (f *Domain) Get() string      { return string(*f) }
//memar:impl computer/capsule/protocol.Assignment
func (f *Domain) Set(value string) { *f = Domain(value) }

func (f *Domain) Name() string         { return "location" }
func (f *Domain) Abbreviation() string { return "loc" }

func (f *Domain) Access() function_p.Access { return function_p.Access_Unset }

func (f *Domain) Optional() bool          { return false }
func (f *Domain) Validate() error_p.Error { return nil }
func (f *Domain) SetDefault()             {}

//memar:impl memar/protocol.Len
func (f *Domain) Length() container_p.NumberOfElement { return -1 }

//memar:impl memar/codec/string/protocol.Stringer
func (f *Domain) ToString() (str string, err error_p.Error) {
	str = string(*f)
	return
}
func (f *Domain) FromString(str string) (err error_p.Error) {
	*f = Domain(str)
	return
}

// flagSet.StringVar((*string)(&req.location), "d", "", "module domain name e.g. user-name")
