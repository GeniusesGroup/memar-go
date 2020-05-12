/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// PublicKeyCryptography : Public-key for related domain.
type PublicKeyCryptography struct {
	PublicKey  [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	PrivateKey [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
}

// RegisterPublicKey use to register public key in new domain name systems like apis.sabz.city
func (p *PublicKeyCryptography) RegisterPublicKey() (err error) {
	// make public & private key and store them
	p.PublicKey = [32]byte{}
	p.PrivateKey = [32]byte{}

	return nil
}
