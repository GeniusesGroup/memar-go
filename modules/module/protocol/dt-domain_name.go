/* For license and copyright information please see the LEGAL file in the code repository */

package mp

import (
	container_p "memar/adt/container/protocol"
	function_p "memar/computer/function/protocol"
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

// DomainName
// flagSet.StringVar(&req.name, "n", "", "module domain name e.g. user-name")
type DomainName struct {
	value string
	datatype.DataType
}

//memar:impl computer/capsule/protocol.Accessor
func (f *DomainName) Get() string      { return f.value }
//memar:impl computer/capsule/protocol.Assignment
func (f *DomainName) Set(value string) { f.value = value }

func (f *DomainName) Access() function_p.Access { return function_p.Access_Unset }
func (f *DomainName) Optional() bool            { return false }
func (f *DomainName) Validate() error_p.Error {
	// TODO:::
	return nil
}
func (f *DomainName) SetDefault() {}

//memar:impl memar/datatype/protocol.DataType_ExpectedLen
func (f *DomainName) MinLength() container_p.NumberOfElement { return -1 }
func (f *DomainName) MaxLength() container_p.NumberOfElement { return 256 }

//memar:impl memar/protocol.Len
func (f *DomainName) Length() container_p.NumberOfElement {
	return container_p.NumberOfElement(len(f.value))
}

//memar:impl memar/codec/string/protocol.Stringer
func (f *DomainName) ToString() (s string, err error_p.Error) {
	s = f.value
	return
}
func (f *DomainName) FromString(str string) (err error_p.Error) {
	f.value = str
	return
}
