// For license and copyright information please see LEGAL file in repository

package binary

func Bool(b []byte) bool   { return b[0] != 0 }
func Uint8(b []byte) uint8 { return uint8(b[0]) }

func PutBool(b []byte, v bool) {
	if v {
		b[0] = 1
	} else {
		b[0] = 0
	}
}
func PutUint8(b []byte, v uint8) {
	b[0] = v
}
