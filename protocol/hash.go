/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Hash is the common interface implemented by all hash functions.
//
// Hash implementations in the standard library (e.g. hash/crc32 and
// crypto/sha256) implement the encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation
// allows its internal state to be saved and used for additional processing
// later, without having to re-write the data previously written to the hash.
// The hash state may contain portions of the input in its original form,
// which users are expected to handle for any possible security implications.
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

// Hash16 is the common interface implemented by all 16-bit hash functions.
type Hash16 interface {
	// Hash
	Sum16() uint16
}

// Hash32 is the common interface implemented by all 32-bit hash functions.
type Hash32 interface {
	// Hash
	Sum32() uint32
}

// Hash64 is the common interface implemented by all 64-bit hash functions.
type Hash64 interface {
	// Hash
	Sum64() uint64
}

// Hash128 is the common interface implemented by all 128-bit hash functions.
type Hash128 interface {
	// Hash
	Sum128() uint64
}

// Hash160 is the common interface implemented by all 160-bit hash functions.
type Hash160 interface {
	// Hash
	Sum160() [20]byte
}

// Hash256 is the common interface implemented by all 256-bit hash functions.
type Hash256 interface {
	// Hash
	Sum256() [32]byte
}

// Hash512 is the common interface implemented by all 512-bit hash functions.
type Hash512 interface {
	// Hash
	Sum512() [64]byte
}
