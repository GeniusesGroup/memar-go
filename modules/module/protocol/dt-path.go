/* For license and copyright information please see the LEGAL file in the code repository */

package mp

import (
	container_p "memar/adt/container/protocol"
	function_p "memar/computer/function/protocol"
	error_p "memar/process/error/protocol"
)

// Path
type Path string

//memar:impl computer/capsule/protocol.Accessor
func (dt *Path) Get() Path { return *dt }
//memar:impl computer/capsule/protocol.Assignment
func (dt *Path) Set(value Path) (err error_p.Error) {
	*dt = value
	err = dt.Validate()
	return
}

func (dt *Path) Access() function_p.Access { return function_p.Access_Unset }
func (dt *Path) Optional() bool            { return true }
func (dt *Path) Validate() error_p.Error {
	if !dt.Optional() && *dt == "" {
		// err = &
		// return
	}
	// TODO::: isDirectory(), isValidPath(), ...
	return nil
}
func (dt *Path) SetDefault() {
	// TODO::: {project root}/modules/
}

//memar:impl memar/protocol.Quiddity
func (dt *Path) Name() string         { return "path" }
func (dt *Path) Abbreviation() string { return "p" }
func (dt *Path) Aliases() []string    { return []string{} }

//memar:impl memar/datatype/protocol.DataType_ExpectedLen
func (dt *Path) MinLength() container_p.NumberOfElement { return -1 }
func (dt *Path) MaxLength() container_p.NumberOfElement { return 256 }

//memar:impl memar/protocol.Len
func (dt *Path) Length() container_p.NumberOfElement { return container_p.NumberOfElement(len(*dt)) }

//memar:impl memar/codec/string/protocol.Stringer
func (dt *Path) ToString() (str string, err error_p.Error) {
	str = string(*dt)
	return
}
func (dt *Path) FromString(str string) (err error_p.Error) {
	*dt = Path(str)
	return
}
