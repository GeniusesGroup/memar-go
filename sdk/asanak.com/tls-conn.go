/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"crypto/tls"
	"time"
)

const (
	maxOpenConnection = 100
	tlsConnTimeOut    = 5 * time.Second

	// Don't understand why such a fundamental service like asanak.com serve just by one server and one IP!!??
	domainPort = domain + ":443"
	ipPort     = "79.175.173.154:443"
)

var tlsConns = [maxOpenConnection]tlsConn{}

// Asanak store data to send and receive sms by asanak provider!
type tlsConn struct {
	conn     *tls.Conn
	maketime time.Time
	use      bool
}

func (tlsConn *tlsConn) free() {
	tlsConn.use = false
}

func (tlsConn *tlsConn) stablish() (err error) {
	var tlsConf = tls.Config{
		ServerName: domain,
	}
	// make TLS connection to asanak server by its IP:Port!
	tlsConn.conn, err = tls.Dial("tcp", ipPort, &tlsConf)
	if err != nil {
		// solve domain to IP and try again
		tlsConn.conn, err = tls.Dial("tcp", domainPort, &tlsConf)
		if err != nil {
			return
		}
	}

	return
}

func getTLSConn() (tlsConn *tlsConn, err error) {
	for i := 0; i < maxOpenConnection; i++ {
		if tlsConns[i].use == true {
			continue
		}
		if time.Since(tlsConns[i].maketime) > 4 {
			err = tlsConns[i].stablish()
			if err != nil {
				continue
			}
		}
		tlsConns[i].maketime = time.Now()
		tlsConns[i].use = true
		return &tlsConns[i], nil
	}
	return
}
