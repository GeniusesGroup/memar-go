/* For license and copyright information please see LEGAL file in repository */

package crypto

// Cipher represents an implementation for a block||stream cipher!
type Cipher interface {
	// Encrypt encrypts the buf.
	// If original buf needed for any other proccess, must clone it before pass it!!
	Encrypt(buf []byte) (err error)

	// Decrypt decrypts the buf.
	// If original buf needed for any other proccess, must clone it before pass it!!
	Decrypt(buf []byte) (err error)
}

// A BlockCipher64 represents an implementation of 64 bit block cipher size!
type BlockCipher64 interface {
	// Encrypt encrypts the block.
	// If original block needed for any other proccess, must clone it before pass it!!
	Encrypt(block *[8]byte)

	// Decrypt decrypts the block.
	// If original block needed for any other proccess, must clone it before pass it!!
	Decrypt(block *[8]byte)
}

// A BlockCipher128 represents an implementation of 128 bit block cipher size!
type BlockCipher128 interface {
	// Encrypt encrypts the block.
	// If original block needed for any other proccess, must clone it before pass it!!
	Encrypt(block *[16]byte)

	// Decrypt decrypts the block.
	// If original block needed for any other proccess, must clone it before pass it!!
	Decrypt(block *[16]byte)
}
