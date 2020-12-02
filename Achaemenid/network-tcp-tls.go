/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/tls"
	"net"

	"../log"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support TCP/TLS over IP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

// MakeTCPTLSNetwork start a TCP/TLS listener and response request in given stream handler
func MakeTCPTLSNetwork(s *Server, port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if s.StreamProtocols.GetProtocolHandler(port) == nil {
		return ErrProtocolHandler
	}

	var tcp = tcpNetwork{
		s:    s,
		port: port,
	}

	tcp.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: s.Networks.localIP[:], Port: int(port)})
	if err != nil {
		log.Warn("TCP/TLS - listen on port ", tcp.listener.Addr(), " failed due to: ", err)
		return
	}

	s.Networks.RegisterTCPNetwork(&tcp)
	log.Info("TCP/TLS - Begin listen on ", tcp.listener.Addr())

	err = tcp.registerCertificates(s)

	// TODO::: Is it worth it to convert tls package usage and implement tls by this package!!??
	var config = new(tls.Config)
	config.NextProtos = []string{"http/1.1", "sRPC"}
	config.PreferServerCipherSuites = true
	// config.InsecureSkipVerify = true
	config.Certificates = append(config.Certificates, tcp.certificate)
	tcp.tlsListener = tls.NewListener(tcp.listener, config)

	go handleTCPTLSListener(s, &tcp, tcp.tlsListener)

	return
}

// handleTCPListener use to handle TCP/TLS networks connections with any application protocol.
func handleTCPTLSListener(s *Server, tcp *tcpNetwork, ln net.Listener) {
	defer s.PanicHandler()
	for {
		var err error
		var tcpConn net.Conn
		tcpConn, err = ln.Accept()
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP/TLS - Accepting new connection occur error:", tcp.listener.Addr(), err)
			}
			continue
		}

		if log.DebugMode {
			log.Debug("TCP/TLS - New connection:", tcpConn.RemoteAddr())
		}

		go handleTCPConn(s, tcp, tcpConn)
	}
}

func (tcp *tcpNetwork) registerCertificates(s *Server) (err error) {
	// Try to load certificate from secret folder
	var crtFile = s.Assets.Secret.GetFile(s.Manifest.DomainName + "-fullchain.crt")
	var keyFile = s.Assets.Secret.GetFile(s.Manifest.DomainName + ".key")
	if crtFile != nil && keyFile != nil {
		tcp.certificate, err = tls.X509KeyPair(crtFile.Data, keyFile.Data)
		if err == nil {
			return
		}
		log.Warn("load exiting certificate face this err: ", err)
	}

	// If letsencrypt failed, can't let certificate be empty!
	log.Warn("Can't load certificates from secret folder, Just use self-sign certificate")
	tcp.certificate, err = tls.X509KeyPair(certificate, privateKey)
	if err != nil {
		log.Warn("load self-sign certificate face this err: ", err)
	}

	return
}

/*
	Default self-sign certificate and Private key when no certificate exist in secret folder!
*/

var privateKey = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGaOKZrTqnBNuebC3WVxTkFSVRxPbsZRrhlbBlUZqeogoAoGCCqGSM49
AwEHoUQDQgAE4swE/yaIMVN5FTUrOJ6jlQFZjLFyUjvF2RR6DbEf4v9XaiPHguAf
VY4DipKxLYzDiGmw5Jd2dKAA1ugySWglsA==
-----END EC PRIVATE KEY-----`)

var certificate = []byte(`-----BEGIN CERTIFICATE-----
MIICPjCCAeOgAwIBAgIHXP21q9ofmTAKBggqhkjOPQQDAjBmMQswCQYDVQQGEwIt
LTERMA8GA1UECAwIVW5pdmVyc2UxDjAMBgNVBAcMBUVhcnRoMRIwEAYDVQQKDAlT
YWJ6LkNpdHkxCzAJBgNVBAsMAklUMRMwEQYDVQQDDApBY2hhZW1lbmlkMB4XDTIw
MDYwODA3NDMwMloXDTMwMDYwNjA3NDMwMlowZjELMAkGA1UEBhMCLS0xETAPBgNV
BAgMCFVuaXZlcnNlMQ4wDAYDVQQHDAVFYXJ0aDESMBAGA1UECgwJU2Fiei5DaXR5
MQswCQYDVQQLDAJJVDETMBEGA1UEAwwKQWNoYWVtZW5pZDBZMBMGByqGSM49AgEG
CCqGSM49AwEHA0IABOLMBP8miDFTeRU1Kzieo5UBWYyxclI7xdkUeg2xH+L/V2oj
x4LgH1WOA4qSsS2Mw4hpsOSXdnSgANboMkloJbCjfDB6MB0GA1UdDgQWBBSSUIKv
1ZobYwPIwkx8y4ZStC6J2jAfBgNVHSMEGDAWgBSSUIKv1ZobYwPIwkx8y4ZStC6J
2jAMBgNVHRMEBTADAQH/MAsGA1UdDwQEAwIDqDAdBgNVHSUEFjAUBggrBgEFBQcD
AgYIKwYBBQUHAwEwCgYIKoZIzj0EAwIDSQAwRgIhAOCvpLIh1u187Kc4M3dKJbJ9
hSJrBqmtA4OmlE2o1ZeLAiEA9tPfVMmV7rst/3CV9fARISVA1ABdqjlpOi6dqbzR
vhM=
-----END CERTIFICATE-----`)

var csr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIIBITCByAIBADBmMQswCQYDVQQGEwItLTERMA8GA1UECAwIVW5pdmVyc2UxDjAM
BgNVBAcMBUVhcnRoMRIwEAYDVQQKDAlTYWJ6LkNpdHkxCzAJBgNVBAsMAklUMRMw
EQYDVQQDDApBY2hhZW1lbmlkMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4swE
/yaIMVN5FTUrOJ6jlQFZjLFyUjvF2RR6DbEf4v9XaiPHguAfVY4DipKxLYzDiGmw
5Jd2dKAA1ugySWglsKAAMAoGCCqGSM49BAMCA0gAMEUCIFlRr4ZdvVDc3pQtjaHf
gV5zOcSSmgYQtvz4aM74TX29AiEAwXxSkPYuFWne/gcQYtDsuzmUAa6zQxxv8uhK
Pixsl74=
-----END CERTIFICATE REQUEST-----`)
