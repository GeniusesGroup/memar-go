/* For license and copyright information please see the LEGAL file in the code repository */

package crypto_p

import (
	buffer_p "memar/buffer/protocol"
	codec_p "memar/codec/protocol"
	error_p "memar/error/protocol"
)

// Cipher represents an implementation for a cipher
type Cipher interface {
	CipherSuite() CipherSuite
	PublicKey() codec_p.Codec // DER, PEM, ...
	// SymmetricKey() []byte // length depend on CipherSuite Can't store it due to security impact
}

// A BlockCipher represents an implementation of block cipher
// using a given key. It provides the capability to encrypt
// or decrypt individual blocks. The mode implementations
// extend that capability to streams of blocks.
type BlockCipher interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Encrypt(dst, src buffer_p.Buffer) (err error_p.Error)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Decrypt(dst, src buffer_p.Buffer) (err error_p.Error)
}

// A Stream represents a stream cipher.
type StreamCipher interface {
	// XORKeyStream XORs each byte in the given slice with a byte from the
	// cipher's key stream. Dst and src must overlap entirely or not at all.
	//
	// If len(dst) < len(src), XORKeyStream should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to XORKeyStream behave as if the concatenation of
	// the src buffers was passed in a single run. That is, Stream
	// maintains state and does not reset at each XORKeyStream call.
	XORKeyStream(dst, src buffer_p.Buffer) (err error_p.Error)
}
