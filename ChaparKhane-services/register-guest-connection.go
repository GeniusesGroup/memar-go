/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var registerGuestConnectionService = achaemenid.Service{
	Name:       "RegisterGuestConnection",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    RegisterGuestConnection,
	Description: []string{
		`use to make new connection by ECC algorithm`,
	},
	TAGS: []string{},
}

// RegisterGuestConnection use to make new connection by ECC algorithm.
func RegisterGuestConnection(s *achaemenid.Server, st *achaemenid.Stream) {
	// Make Connection.EncryptionKey by ECC algorithm with :
	// req.ConnectionPublicKey
	// RunningServerData.PublicKeyCryptography.PrivateKey

	// Check server allow to register guest connection and max guest connection availability

	// If everything go well register connection
	s.Connections.RegisterConnection(st.Connection)
}

type registerGuestConnectionReq struct {
	ConnectionPublicKey  [32]byte `valid:"PublicKey"`
	SuggestedCipherSuite uint16
	UserAgent            string
}

type registerGuestConnectionRes struct {
}

func (req *registerGuestConnectionReq) validate() error {
	return nil
}

func (req *registerGuestConnectionReq) sRPCDecoder(buf []byte) error {
	return nil
}

func (res *registerGuestConnectionRes) sRPCEncoder(buf []byte) error {
	return nil
}
