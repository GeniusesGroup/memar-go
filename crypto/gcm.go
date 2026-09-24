/* For license and copyright information please see the LEGAL file in the code repository */

package crypto

import (
	crypto_p "memar/crypto/protocol"
	error_p "memar/process/error/protocol"
)

// GCM cipher mode
type gcm struct {
	b             crypto_p.BlockCipher
	peerPublicKey [32]byte
}

// NewGCM use to create the gcm that implement Cipher interface.
func NewGCM(b crypto_p.BlockCipher) *gcm {
	var g = gcm{
		b: b,
	}
	return &g
}

// Encrypt encrypts the buf.
// If original buf needed for any other process, must clone it before pass it.
func (g *gcm) Encrypt(buf []byte) (err error_p.Error) {
	return
}

// Decrypt decrypts the buf.
// If original buf needed for any other process, must clone it before pass it.
func (g *gcm) Decrypt(buf []byte) (err error_p.Error) {
	return
}

// PeerPublicKey return peer PublicKey.
func (g *gcm) PeerPublicKey() []byte {
	return g.peerPublicKey[:]
}
