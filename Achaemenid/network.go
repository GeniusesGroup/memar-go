/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Network :
type Network struct {
	GPRange [14]byte
	MTU      uint16 // Maximum Transmission Unit. GP packet size
}

// RegisterGP use to get new GP & MTU from OS router!
func (n *Network) RegisterGP() (err error) {
	// send PublicKey to router and get IP if user granted. otherwise log error.
	n.GPRange = [14]byte{}

	// Get MTU from router
	n.MTU = 1200

	// Because ChaparKhane is server based application must have IP access.
	// otherwise close server app and return err

	return nil
}