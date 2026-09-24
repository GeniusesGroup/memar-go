/* For license and copyright information please see the LEGAL file in the code repository */

package crypto_p

import (
	datatype_p "memar/computer/datatype/protocol"
	uuid_p "memar/identifier/protocol"
	string_p "memar/codec/string/protocol"
)

// https://en.wikipedia.org/wiki/Cipher_suite
type CipherSuite interface {
	CipherSuite() string_p.String // e.g. TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

	datatype_p.DataType
}

type CipherSuite_Parsed interface {
	string_p.Stringer // e.g. TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

	Protocol() string_p.String          // Defines the protocol that this cipher suite is for e.g. TLS
	KeyExchange() string_p.String       // indicates the key exchange algorithm being used e.g. ECDHE
	Authentication() string_p.String    // authentication mechanism during the handshake e.g. RSA
	SessionCipher() string_p.String     // Encryption and Decryption mechanism e.g. AES
	EncryptionKeySize() string_p.String // session encryption key size (bits) for cipher e.g. 128
	EncryptionType() string_p.String    // Type of encryption (cipher-block dependency and additional options) e.g. GCM
	Hash() string_p.String              // Signature mechanism. Indicates the message authentication algorithm which is used to authenticate a message. e.g. SHA(SHA2)

	// Insecure is true if the cipher suite has known security issues
	// due to its primitives, design, or implementation.
	Insecure() bool

	uuid_p.UUID_Hash // hash of Stringer like MediaType
}
