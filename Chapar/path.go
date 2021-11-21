/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"io"

	"../protocol"
)

// Path indicate Chapar switch route plan!
type Path struct {
	path [MaxHopCount]byte
	len  byte
}

// Init sets path from the given frame
func (p *Path) Init(frame []byte) {
	if len(frame) == 0 {
		p.len = MaxHopCount // broadcast frame
	} else {
		var hopCount = GetHopCount(frame)
		copy(p.path[:], frame[FixedHeaderLength:FixedHeaderLength+hopCount])
		p.len = hopCount
	}
}

func (p *Path) Set(path []byte) {
	copy(p.path[:], path)
	p.len = byte(len(path))
}

func (p *Path) Get() []byte {
	return p.path[:p.len]
}

func (p *Path) GetAsString() string {
	// TODO::: ??
	return string(p.path[:p.len])
}

func (p *Path) LenAsByte() byte {
	return p.len
}

// GetReverse return path in all hops just in reverse.
func (p *Path) GetReverse() (reverse Path) {
	reverse = Path{
		path: ReversePath(p.path[:]),
		len:  p.len,
	}
	return
}

/*
********** protocol.Codec interface **********
 */

func (p *Path) MediaType() protocol.MediaType       { return nil }
func (p *Path) CompressType() protocol.CompressType { return nil }
func (p *Path) Len() int                            { return int(p.len) }

// Marshal return the path in the given frame.
func (p *Path) Decode(reader io.Reader) (err protocol.Error) {
	return
}

// Marshal return the path in the given frame.
func (p *Path) Encode(writer io.Writer) (err error) {
	return
}

// Unmarshal sets path from the given path
func (p *Path) Unmarshal(path []byte) (err protocol.Error) {
	if len(path) == 0 {
		p.len = MaxHopCount // broadcast frame
	} else {
		copy(p.path[:], path)
		p.len = byte(len(path))
	}
}

// Marshal return the path in the given frame.
func (p *Path) Marshal() (path []byte) {
	return p.path[:p.len]
}

// MarshalTo sets the path in the given frame.
func (p *Path) MarshalTo(frame []byte) {
	copy(frame[FixedHeaderLength:], p.path[:p.len])
}
