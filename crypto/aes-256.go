/* For license and copyright information please see the LEGAL file in the code repository */

package crypto

// AESBlockSize is block size in bytes.
const AESBlockSize = 16

// A aes256 is an instance of AES-256 encryption using a particular key.
type aes256 struct {
	enc [60]uint32
	dec [60]uint32
}

// NewAES256 use to create the aes256 that implement BlockCipher interface!
func NewAES256(key [32]byte) *aes256 {
	var c = aes256{}
	return &c
}

// Encrypt encrypts the buf.
// If original buf needed for any other process, must clone it before pass it!!
func (aes *aes256) Encrypt(block []byte) {}

// Decrypt decrypts the buf.
// If original buf needed for any other process, must clone it before pass it!!
func (aes *aes256) Decrypt(block []byte) {}
