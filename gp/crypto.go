/* For license and copyright information please see LEGAL file in repository */

package gp

import "../protocol"

// Encrypt use in encrypted connection from Apps to Apps!
func Encrypt(frames []byte, cipher protocol.Cipher) (payload []byte, err protocol.Error) {
	err = cipher.Encrypt(frames)
	return
}

// Decrypt use in encrypted connection from Apps to Apps!
func Decrypt(payload []byte, cipher protocol.Cipher) (frames []byte, err protocol.Error) {
	// Decrypt packet by encryptionKey & Checksum data in this protocol :
	// We check packet errors with encryption proccess together
	// and needed checksum data will be add to encrypted data. 32 bit checksum in end of Packet
	err = cipher.Decrypt(payload)
	return
}

// EncryptRouting usually use in encrypted connection from OS to GP Router!
func EncryptRouting(packet []byte, cipher protocol.Cipher) (err protocol.Error) {
	err = cipher.Encrypt(packet[:32])
	return
}

// DecryptRouting usually use in encrypted connection from OS to GP Router!
func DecryptRouting(packet []byte, cipher protocol.Cipher) (err protocol.Error) {
	err = cipher.Decrypt(packet[:32])
	return
}

func checkSignature() {}
