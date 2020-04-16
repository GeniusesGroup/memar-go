/* For license and copyright information please see LEGAL file in repository */

package uip

// DecryptRoutingPart usually use in encrypted connection from OS to UIP Router!
// Default frame size is 128bit due cipher block size but peer can change by settings.
func DecryptRoutingPart(p []byte, frameSize uint16, encryptionKey [32]byte, cipherSuite uint16) error {
	return nil
}

// DecryptDataPart use in encrypted connection from Apps to Apps!
// Default frame size is 128bit due cipher block size but peer can change by settings.
func DecryptDataPart(p []byte, frameSize uint16, encryptionKey [32]byte, cipherSuite uint16) error {
	// Decrypt packet by encryptionKey & Checksum data in this protocol :
	// We check packet errors with encryption proccess together
	// and needed checksum data will be add to encrypted data.
	
	// IPv6 Specification: In order to increase performance, and since current link layer technology
	// and transport or application layer protocols are assumed to provide sufficient error detection.

	// 16 bit checksum in end of Packet

	return nil
}