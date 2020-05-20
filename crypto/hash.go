/* For license and copyright information please see LEGAL file in repository */

package crypto

// Hash32 represents an implementation for a 32 bit hash!
type Hash32 interface {
	Generate(buf []byte) (hash uint32)
}

// Hash64 represents an implementation for a 64 bit hash!
type Hash64 interface {
	Generate(buf []byte) (hash uint64)
}

// Hash128 represents an implementation for a 128 bit hash!
type Hash128 interface {
	Generate(buf []byte) (hash [16]byte)
}

// Hash256 represents an implementation for a 256 bit hash!
type Hash256 interface {
	Generate(buf []byte) (hash [32]byte)
}
