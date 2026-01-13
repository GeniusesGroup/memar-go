/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype

import (
	string_p "memar/codec/string/protocol"
	error_p "memar/process/error/protocol"
)

// Parsed implement mediatype_p.Parsed interface
type Parsed[STR string_p.String] struct {
	mediaType STR

	mainType   STR
	tree       STR
	subType    STR
	suffix     STR
	parameters []STR
}

func (mt *Parsed[STR]) Init(mediatype STR) (err error_p.Error) {
	mt.mediaType = mediatype
	err = mt.parse()
	return
}

//memar:impl memar/protocol.MediaType
func (mt *Parsed[STR]) MediaType() STR    { return mt.mediaType }
func (mt *Parsed[STR]) MainType() STR     { return mt.mainType }
func (mt *Parsed[STR]) Tree() STR         { return mt.tree }
func (mt *Parsed[STR]) SubType() STR      { return mt.subType }
func (mt *Parsed[STR]) Suffix() STR       { return mt.suffix }
func (mt *Parsed[STR]) Parameters() []STR { return mt.parameters }

// func (mt *Parsed[STR]) FileExtension() STR{ return nil }

//memar:impl memar/codec/string/protocol.Stringer
func (mt *Parsed[STR]) ToString() (str string_p.String, err error_p.Error) { return mt.mediaType, nil }
func (mt *Parsed[STR]) FromString(str string_p.String) (err error_p.Error) {
	// TODO::: Can't change via this method.
	// err =
	return
}

// TODO::: complete extraction
func (mt *Parsed[STR]) parse() (err error_p.Error) {
	var mediatype = mt.mediaType

	err = mediatype.CopyTo(mt.mainType)
	if !err.IsNil() {
		return
	}
	// mt.subType, err = mediatype.SplitByElement('/')
	if !err.IsNil() {
		return
	}
	mt.mainType.Pop() // remove last '/' character
	// TODO:::
	return
}
