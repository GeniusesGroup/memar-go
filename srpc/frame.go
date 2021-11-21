/* For license and copyright information please see LEGAL file in repository */

package srpc

/*
type frame struct {
	Type     byte
	Payload []byte
}
*/
type frame []byte

func (f frame) Type() byte      { return byte(f[0]) }
func (f frame) Payload() []byte { return f[1:] }

// Frame type ID is like service ID but fixed ID with 8bit length. Just some few services get one byte length service ID
// Common services must register by 64bit unsigned integer.
const (
	frameTypePadding byte = iota
	frameTypePing
	frameTypeCallService
	frameTypeOpenStream
	frameTypeCloseStream
	frameTypeData
	frameTypeSignature
)
