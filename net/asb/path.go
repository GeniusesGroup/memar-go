/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

// Path indicate Chapar switch route plan!
type Path struct {
	path [maxHopCount]byte
	len  byte
}

// Init sets path from the given frame
func (p *Path) Init(frame Frame) {
	var hopCount = frame.HopCount()
	copy(p.path[:], frame[fixedHeaderLength:fixedHeaderLength+hopCount])
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

//libgo:impl libgo/protocol.Stringer
func (p *Path) ToString() string {
	// TODO::: ??
	return string(p.path[:p.len])
}
func (p *Path) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}

//libgo:impl libgo/protocol.Codec
func (p *Path) MediaType() protocol.MediaType       { return nil }
func (p *Path) CompressType() protocol.CompressType { return nil }
func (p *Path) Len() int                            { return int(p.len) }
func (p *Path) Decode(source protocol.Codec) (n int, err protocol.Error) {
	// TODO:::
	return
}
func (p *Path) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	// TODO:::
	return
}
func (p *Path) Unmarshal(path []byte) (err protocol.Error) {
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
	copy(frame[fixedHeaderLength:], p.path[:p.len])
}
