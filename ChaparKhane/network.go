/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Network :
type Network struct {
	UIPRange [14]byte
	MTU      uint16 // Maximum Transmission Unit. UIP packet size
}

// RegisterUIP use to get new UIP & MTU from OS router!
func (n *Network) RegisterUIP() (err error) {
	// send PublicKey to router and get IP if user granted. otherwise log error.
	n.UIPRange = [14]byte{}

	// Get MTU from router
	n.MTU = 1200

	// Because ChaparKhane is server based application must have IP access.
	// otherwise close server app and return err

	return nil
}