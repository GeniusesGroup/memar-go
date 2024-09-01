/* For license and copyright information please see the LEGAL file in the code repository */

package crypto_p

import (
	uuid_p "memar/identifier/protocol"
	"memar/protocol"
	string_p "memar/string/protocol"
)

// https://en.wikipedia.org/wiki/Cipher_suite
type CipherSuite interface {
	CipherSuite() string_p.String // e.g. TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

	protocol.DataType
}

type CipherSuite_Parsed[STR string_p.String] interface {
	string_p.Stringer[STR] // e.g. TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

	Protocol() string       // Defines the protocol that this cipher suite is for e.g. TLS
	KeyExchange()           // indicates the key exchange algorithm being used e.g. ECDHE
	Authentication() STR    // authentication mechanism during the handshake e.g. RSA
	SessionCipher() STR     // Encryption and Decryption mechanism e.g. AES
	EncryptionKeySize() STR // session encryption key size (bits) for cipher e.g. 128
	EncryptionType() STR    // Type of encryption (cipher-block dependency and additional options) e.g. GCM
	Hash() STR              // Signature mechanism. Indicates the message authentication algorithm which is used to authenticate a message. e.g. SHA(SHA2)

	// Insecure is true if the cipher suite has known security issues
	// due to its primitives, design, or implementation.
	Insecure() bool

	uuid_p.UUID_Hash // hash of Stringer like MediaType
}
