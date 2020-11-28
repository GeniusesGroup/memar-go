/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"crypto/tls"

	etime "../../earth-time"
)

const (
	maxOpenConnection = 100
	tlsConnTimeOut    = 60 * etime.Second
)

var tlsConns = [maxOpenConnection]tlsConn{}

type tlsConn struct {
	conn     *tls.Conn
	expireIn etime.Time
	use      bool
}

func (tlsConn *tlsConn) free() {
	tlsConn.use = false
}

func (tlsConn *tlsConn) stablish() (err error) {
	var tlsConf = tls.Config{
		ServerName: posDomain,
	}
	tlsConn.conn, err = tls.Dial("tcp", posDomainPort, &tlsConf)
	return
}

func getTLSConn() (tlsConn *tlsConn, err error) {
	var timeNow = etime.Now()
	for i := 0; i < maxOpenConnection; i++ {
		if tlsConns[i].use == true {
			continue
		}
		if tlsConns[i].expireIn.Pass(timeNow) {
			err = tlsConns[i].stablish()
			if err != nil {
				continue
			}
		}
		tlsConns[i].expireIn = timeNow.AddDuration(tlsConnTimeOut)
		tlsConns[i].use = true
		return &tlsConns[i], nil
	}
	return
}
