/* For license and copyright information please see the LEGAL file in the code repository */

package crypto

import (
	crypto_p "memar/crypto/protocol"
	error_p "memar/process/error/protocol"
)

// CCM cipher mode
type ccm struct {
	b             crypto_p.BlockCipher
	peerPublicKey [32]byte
}

// NewCCM use to create the ccm that implement Cipher interface!
func NewCCM(b crypto_p.BlockCipher) *ccm {
	var c = ccm{
		b: b,
	}
	return &c
}

// Encrypt encrypts the buf.
// If original buf needed for any other process, must clone it before pass it!!
func (c *ccm) Encrypt(buf []byte) (err error_p.Error) {
	return
}

// Decrypt decrypts the buf.
// If original buf needed for any other process, must clone it before pass it!!
func (c *ccm) Decrypt(buf []byte) (err error_p.Error) {
	return
}

// PeerPublicKey return peer PublicKey!
func (c *ccm) PeerPublicKey() []byte {
	return c.peerPublicKey[:]
}
