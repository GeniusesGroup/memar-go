/* For license and copyright information please see LEGAL file in repository */

package crypto

// CCM cipher mode
type ccm struct {
	b BlockCipher128
}

// NewCCM use to create the ccm that implement Cipher interface!
func NewCCM(b BlockCipher128) Cipher {
	var c = ccm{
		b: b,
	}
	return &c
}

// Encrypt encrypts the buf.
// If original buf needed for any other proccess, must clone it before pass it!!
func (c *ccm) Encrypt(buf []byte) (err error) {
	return
}

// Decrypt decrypts the buf.
// If original buf needed for any other proccess, must clone it before pass it!!
func (c *ccm) Decrypt(buf []byte) (err error) {
	return
}
