/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

// Path indicate Chapar switch route plan!
type Path struct {
	path [maxHopCount]byte
	len  byte
}

// Init sets path from the given frame
func (p *Path) Init(frame Frame) {
	var hopCount = frame.HopCount()
	copy(p.path[:], frame[frameFixedLength:frameFixedLength+hopCount])
	p.len = hopCount
}

func (p *Path) Set(path []byte) {
	copy(p.path[:], path)
	p.len = byte(len(path))
}

func (p *Path) Get() []byte {
	return p.path[:p.len]
}

func (p *Path) LenAsByte() byte {
	return p.len
}

// CopyReverseTo will copy p in reverse hops to given path.
func (p *Path) CopyReverseTo(reverse *Path) {
	var ln = p.len
	ln-- // Due to len & index mismatch
	var j byte
	for ln >= j {
		reverse.path[ln], reverse.path[j] = p.path[j], p.path[ln]
		ln--
		j++
	}

	reverse.len = ln
	return
}

//memar:impl memar/codec/string/protocol.Stringer
func (p *Path) ToString() (str string, err error_p.Error) {
	// TODO::: ??
	str = string(p.path[:p.len])
	return
}
func (p *Path) FromString(str string) (err error_p.Error) {
	// TODO:::
	return
}

//memar:impl memar/codec/protocol.Codec
// func (p *Path) MediaType() mediatype_p.MediaType    { return nil }
// func (p *Path) CompressType() compress_p.CompressType { return nil }

//memar:impl memar/codec/protocol.Field_Length
func (p *Path) SerializationLength() container_p.NumberOfElement {
	return container_p.NumberOfElement(p.len)
}

//memar:impl memar/protocol.Decoder
func (p *Path) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// TODO:::
	return
}

//memar:impl memar/protocol.Encoder
func (p *Path) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	// TODO:::
	return
}

func (p *Path) Unmarshal(path []byte) (err error_p.Error) {
	if len(path) == 0 {
		// err = &
		return
	}

	copy(p.path[:], path)
	p.len = byte(len(path))
	return
}
func (p *Path) Marshal() (path []byte) {
	return p.path[:p.len]
}
func (p *Path) MarshalTo(frame []byte) {
	copy(frame[frameFixedLength:], p.path[:p.len])
}
