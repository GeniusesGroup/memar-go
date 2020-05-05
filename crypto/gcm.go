/* For license and copyright information please see LEGAL file in repository */

package crypto

// GCM cipher mode
type gcm struct {
	b BlockCipher128
}

// NewGCM use to create the gcm that implement Cipher interface!
func NewGCM(b BlockCipher128) Cipher {
	var g = gcm{
		b: b,
	}
	return &g
}

// Encrypt encrypts the buf.
// If original buf needed for any other proccess, must clone it before pass it!!
func (g *gcm) Encrypt(buf []byte) (err error) {
	return
}

// Decrypt decrypts the buf.
// If original buf needed for any other proccess, must clone it before pass it!!
func (g *gcm) Decrypt(buf []byte) (err error) {
	return
}