/* For license and copyright information please see LEGAL file in repository */

package protocol

// Cipher represents an implementation for a cipher
type Cipher interface {
	CipherSuite() CipherSuite
	PublicKey() Codec // DER, PEM, ...
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
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Decrypt(dst, src []byte)
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
	XORKeyStream(dst, src []byte)
}

// https://en.wikipedia.org/wiki/Cipher_suite
type CipherSuite interface {
	Stringer                   // e.g. TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

	ID() uint64                // hash of Stringer like GitiURN
	Protocol() string          // Defines the protocol that this cipher suite is for e.g. TLS
	KeyExchange() string       // indicates the key exchange algorithm being used e.g. ECDHE
	Authentication() string    // authentication mechanism during the handshake e.g. RSA
	SessionCipher() string     // Encryption and Decryption mechanism e.g. AES
	EncryptionKeySize() string // session encryption key size (bits) for cipher e.g. 128
	EncryptionType() string    // Type of encryption (cipher-block dependency and additional options) e.g. GCM
	Hash() string              // Signature mechanism. Indicates the message authentication algorithm which is used to authenticate a message. e.g. SHA(SHA2)

	// Insecure is true if the cipher suite has known security issues
	// due to its primitives, design, or implementation.
	Insecure() bool
}
