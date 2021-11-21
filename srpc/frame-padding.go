/* For license and copyright information please see LEGAL file in repository */

package srpc

import "../syllab"

/*
type paddingFrame struct {
	Length  [2]byte // including the header fields
	Padding []byte
}
*/
type paddingFrame []byte

func (f paddingFrame) Length() uint16    { return syllab.GetUInt16(f, 0) }
func (f paddingFrame) Payload() []byte   { return f[2:f.Length()] }
func (f paddingFrame) NextFrame() []byte { return f[f.Length():] }
