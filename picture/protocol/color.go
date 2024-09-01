/* For license and copyright information please see the LEGAL file in the code repository */

package picture_p

// True color (24-bit) 2^24 gives 16,777,216 color variations.
// The human eye can discriminate up to ten million colors, and since the gamut of a display is smaller than the range of human vision,
// this means this should cover that range with more detail than can be perceived.
// https://en.wikipedia.org/wiki/Color_depth
type RGB struct {
	Red   uint32
	Green uint32
	Blue  uint32
}

// type ColorHex [6]byte
