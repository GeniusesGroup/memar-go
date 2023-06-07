/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"libgo/crypto"
	"libgo/protocol"
)

// cryptography : Public-key for related domain.
type cryptography struct {
	publicKey         [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	privateKey        [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	ChecksumGenerator crypto.Hash256
}

// init make and register cryptography data for given server
func (c *cryptography) init() (err protocol.Error) {
	// make public & private key for desire node e.g. node12.geniuses.group and store them
	c.publicKey = [32]byte{}
	c.privateKey = [32]byte{}
	// ecdsa.GenerateKey()

	return
}

func (c *cryptography) Deinit() (err protocol.Error) {
	// TODO:::

	// Send signal to DNS & Certificate server to revoke app data.
	return
}
