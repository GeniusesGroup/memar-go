/* For license and copyright information please see the LEGAL file in the code repository */

package tls

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"

	"libgo/log"
	"libgo/protocol"
)

// tcpNetwork store related data.
type AppMultiplexer struct {
	headerID    protocol.NetworkLink_NextHeaderID
	application protocol.Application
	appManifest protocol.ApplicationManifest
	connections connections

	port        uint16
	listener    *net.TCPListener
	tlsListener net.Listener
	certificate tls.Certificate
}

// MakeTCPTLSNetwork start a TCP/TLS listener and response request in given stream handler
func MakeTCPTLSNetwork(port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if Server.StreamProtocols.GetProtocolHandler(port) == nil {
		return &ErrProtocolHandler
	}

	var tcp = tcpNetwork{
		port: port,
	}

	tcp.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: Server.Networks.localIP[:], Port: int(port)})
	if err != nil {
		log.Warn(domainEnglish, "TCP/TLS - listen on port ", tcp.listener.Addr(), " failed due to: ", err)
		return
	}

	Server.Networks.RegisterTCPNetwork(&tcp)
	log.Info("TLS - Begin listen on ", tcp.listener.Addr())

	err = tcp.registerCertificates()

	// TODO::: Is it worth it to convert tls package usage and implement tls by this package!!??
	var config = new(tls.Config)
	config.NextProtos = []string{"http/1.1", "sRPC"}
	config.PreferServerCipherSuites = true
	// config.InsecureSkipVerify = true
	config.Certificates = append(config.Certificates, tcp.certificate)
	tcp.tlsListener = tls.NewListener(tcp.listener, config)

	go handleTCPTLSListener(&tcp, tcp.tlsListener)

	return
}

// handleTCPListener use to handle TCP/TLS networks connections with any application protocol.
func handleTCPTLSListener(tcp *tcpNetwork, ln net.Listener) {
	defer log.PanicHandler()
	for {
		var err error
		var tcpConn net.Conn
		tcpConn, err = ln.Accept()
		if err != nil {
			if protocol.AppMode_Dev {
				log.Debug(domainEnglish, "TCP/TLS - Accepting new connection occur error:", tcp.listener.Addr(), err)
			}
			continue
		}

		if protocol.AppMode_Dev {
			log.Debug(domainEnglish, "TCP/TLS - New connection:", tcpConn.RemoteAddr())
		}

		go handleTCPConn(tcp, tcpConn)
	}
}

func (tcp *tcpNetwork) registerCertificates() (err error) {
	var crtFile = Server.Assets.Secret.File(string(Server.Manifest.DomainName) + "-fullchain.crt")
	var keyFile = Server.Assets.Secret.File(string(Server.Manifest.DomainName) + ".key")
	if crtFile != nil && keyFile != nil {
		var goErr error
		tcp.certificate, goErr = tls.X509KeyPair(crtFile.Data, keyFile.Data)
		if goErr == nil {
			return
		}
		log.Warn(domainEnglish, "load exiting certificate face this err: ", goErr)
	}

	var goErr error
	tcp.certificate, goErr = makeEphemeralCertificate(string(Server.Manifest.DomainName))
	if goErr != nil {
		log.Warn(domainEnglish, "generate ephemeral certificate failed: ", goErr)
	}
	return
}

func makeEphemeralCertificate(domain string) (certificate tls.Certificate, err error) {
	if domain == "" {
		domain = "localhost"
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}

	serialLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialLimit)
	if err != nil {
		return
	}

	now := time.Now()
	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{CommonName: domain},
		NotBefore:             now.Add(-time.Hour),
		NotAfter:              now.AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{domain, "localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1"), net.IPv6loopback},
	}

	certificateDER, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return
	}

	privateKeyDER, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return
	}

	certificatePEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificateDER})
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyDER})
	return tls.X509KeyPair(certificatePEM, privateKeyPEM)
}
