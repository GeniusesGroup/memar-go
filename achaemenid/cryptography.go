/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"../crypto"
)

// cryptography : Public-key for related domain.
type cryptography struct {
	publicKey         [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	privateKey        [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	ChecksumGenerator crypto.Hash256
}

// init make and register cryptography data for given server
func (c *cryptography) init() {
	// make public & private key for desire node e.g. node12.sabz.city and store them
	c.publicKey = [32]byte{}
	c.privateKey = [32]byte{}
	// ecdsa.GenerateKey()

	return
}

func (c *cryptography) shutdown() {
	// TODO:::

	// Send signal to DNS & Certificate server to revoke app data.
}
