/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// cryptography : Public-key for related domain.
type cryptography struct {
	publicKey  [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	privateKey [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
}

// init use to register public key in new domain name systems like sabz.city
func (c *cryptography) init(s *Server) (err error) {
	// make public & private key and store them
	c.publicKey = [32]byte{}
	c.privateKey = [32]byte{}
	// ecdsa.GenerateKey()

	return
}

func (c *cryptography) shutdown() {
	// TODO:::

	// Send signal to DNS & Certificate server to revoke app data.
}