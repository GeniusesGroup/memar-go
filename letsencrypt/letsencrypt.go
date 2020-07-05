/* For license and copyright information please see LEGAL file in repository */

package letsencrypt

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"time"

	as "../assets"
	"../errors"
	"../log"
)

const (
	letsEncryptDomain     = "acme-v02.api.letsencrypt.org"
	letsEncryptACMEURL    = "/directory"
	letsEncryptAccountURL = "/acme/new-acct"
	letsEncryptNonceURL   = "/acme/new-nonce"
	letsEncryptOrderURL   = "/acme/new-order"
	letsEncryptRevokeURL  = "/acme/revoke-cert"
)

// Package Errors
var (
	ErrNoSecretFolderAvailable    = errors.New("NoSecretFolderAvailable", "No secret folder give in LetsEncrypt type")
	ErrNoCertificateExistInSecret = errors.New("NoCertificateExistInSecret", "Can't find private-key and-or certificate in secret folder")
	ErrCertificateExpireSoon      = errors.New("ErrCertificateExpireSoon", "Certificate expire in below 10 days")
)

// LetsEncrypt store all data to make, renew or revoke certificate.
type LetsEncrypt struct {
	SecretFolder *as.Folder

	CommonName       string
	EmailAddresses   []string
	Organization     []string
	OrganizationUnit []string
	Domains          []string

	ECKey    *ecdsa.PrivateKey
	ECKeyDER []byte
	ECKeyPEM []byte

	CertificateRequest    *x509.CertificateRequest
	CertificateRequestDER []byte
	CertificateRequestPEM []byte

	Certificate    *x509.Certificate
	CertificateDER []byte
	CertificatePEM []byte

	CertificateChain    *x509.Certificate
	CertificateChainDER []byte
	CertificateChainPEM []byte

	Nonce string
}

// init read any exiting file and generate new one for non exist file.
func (le *LetsEncrypt) init() (err error) {
	if le.SecretFolder == nil {
		// Almost this situation never occur!
		log.Warn("No secret folder set in given LetsEncrypt")
		return ErrNoSecretFolderAvailable
	}

	// check private key and generate if not exist
	var ecKeyFile = le.SecretFolder.GetFile(le.CommonName + ".key")
	if ecKeyFile != nil {
		le.ECKeyPEM = ecKeyFile.Data
		var ecKeyBlock, _ = pem.Decode(ecKeyFile.Data)
		le.ECKeyDER = ecKeyBlock.Bytes
		le.ECKey, err = x509.ParseECPrivateKey(le.ECKeyDER)
		if err != nil {
			log.Warn("Private Key in secret folder is not valid, It must replace with new one ...")
		}
	}
	if le.ECKey == nil {
		log.Info("No valid EC Key found in storage, generate new one and save to secret folder...")
		err = le.generateECKey()
		if err != nil {
			return
		}
	}

	// check certificate request file and generate if not exist
	var csrFile = le.SecretFolder.GetFile(le.CommonName + ".csr")
	if csrFile != nil {
		le.CertificateRequestPEM = csrFile.Data
		var csrBlock, _ = pem.Decode(csrFile.Data)
		le.CertificateRequestDER = csrBlock.Bytes
		le.CertificateRequest, err = x509.ParseCertificateRequest(le.CertificateRequestDER)
		if err != nil {
			log.Warn("Certificate request in secret folder is not valid, It must replace with new one ...")
		}
	}
	if le.CertificateRequest == nil {
		log.Info("No valid Certificate-Request found in storage, generate new one and save to secret folder...")
		err = le.generateCertificateRequest()
		if err != nil {
			return
		}
	}

	// check certificate and generate if not exist
	var certificateFile = le.SecretFolder.GetFile(le.CommonName + ".crt")
	if certificateFile == nil {
		log.Info("No valid Certificate found in storage")
	} else if certificateFile != nil && ecKeyFile != nil {
		le.CertificatePEM = certificateFile.Data
		var csrBlock, _ = pem.Decode(certificateFile.Data)
		le.CertificateDER = csrBlock.Bytes
		le.Certificate, err = x509.ParseCertificate(le.CertificateDER)
		if err != nil {
			log.Warn("Certificate in secret folder is not valid, It must replace with new one ...")
			return
		}
		// check certificate expire date. Don't accpet expire below 10 days!
		if time.Until(le.Certificate.NotAfter) < 240*time.Hour {
			log.Warn("Exist certificate in storage expired or expire in less than 10 days, It must replace with new one ...")
			le.Certificate = nil
		}
	}

	return
}

func (le *LetsEncrypt) generateECKey() (err error) {
	var pubkeyCurve = elliptic.P256()
	le.ECKey, err = ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		return
	}
	// Set x509 serialization DER version
	le.ECKeyDER, err = x509.MarshalECPrivateKey(le.ECKey)
	if err != nil {
		return
	}
	// Set PEM version
	var block = pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: le.ECKeyDER,
	}
	var privateBuf bytes.Buffer
	err = pem.Encode(&privateBuf, &block)
	if err != nil {
		return
	}
	le.ECKeyPEM = privateBuf.Bytes()
	return
}

// Generate Certificate Request - CSR
func (le *LetsEncrypt) generateCertificateRequest() (err error) {
	le.CertificateRequest = &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: le.CommonName,
			Country:    []string{"Giti"},
			// Province:           []string{""},
			// Locality:           []string{""},
			Organization:       le.Organization,
			OrganizationalUnit: le.OrganizationUnit,
		},
		DNSNames:       le.Domains,
		EmailAddresses: le.EmailAddresses,
	}
	le.CertificateRequestDER, err = x509.CreateCertificateRequest(rand.Reader, le.CertificateRequest, le.ECKey)
	if err != nil {
		return
	}
	// Set PEM version
	var csrBlock = pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: le.CertificateRequestDER,
	}
	var csrBuf bytes.Buffer
	err = pem.Encode(&csrBuf, &csrBlock)
	if err != nil {
		return
	}
	le.CertificateRequestPEM = csrBuf.Bytes()
	return
}

// SaveToSecretAssets make needed file and save to le.SecretFolder
func (le *LetsEncrypt) saveToSecretAssets() (err error) {
	// EC KEY
	var ecKeyFile = as.File{
		FullName:  le.CommonName + ".key",
		Name:      le.CommonName,
		Extension: "key",
		Data:      le.ECKeyPEM,
		State:     as.StateChanged,
	}
	le.SecretFolder.SetFile(&ecKeyFile)

	// CERTIFICATE REQUEST
	var csrFile = as.File{
		FullName:  le.CommonName + ".csr",
		Name:      le.CommonName,
		Extension: "csr",
		Data:      le.CertificateRequestPEM,
		State:     as.StateChanged,
	}
	le.SecretFolder.SetFile(&csrFile)

	// CERTIFICATE
	var crtFile = as.File{
		FullName:  le.CommonName + ".crt",
		Name:      le.CommonName,
		Extension: "crt", // CER is an X.509 certificate in binary form, DER encoded. CRT is a binary X.509 certificate, encapsulated in text (base-64) encoding.
		Data:      le.CertificatePEM,
		State:     as.StateChanged,
	}
	le.SecretFolder.SetFile(&crtFile)

	// CERTIFICATE CHAIN
	var crtChainFile = as.File{
		FullName:  le.CommonName + "-chain.crt",
		Name:      le.CommonName,
		Extension: "crt", // CER is an X.509 certificate in binary form, DER encoded. CRT is a binary X.509 certificate, encapsulated in text (base-64) encoding.
		Data:      le.CertificateChainPEM,
		State:     as.StateChanged,
	}
	le.SecretFolder.SetFile(&crtChainFile)

	// CERTIFICATE FULLCHAIN
	var fullChainPEM = make([]byte, len(le.CertificatePEM)+len(le.CertificateChainPEM))
	fullChainPEM = append(fullChainPEM, le.CertificatePEM...)
	fullChainPEM = append(fullChainPEM, le.CertificateChainPEM...)
	var crtFullChainFile = as.File{
		FullName:  le.CommonName + "-fullchain.crt",
		Name:      le.CommonName,
		Extension: "crt", // CER is an X.509 certificate in (binary form-DER encoded). CRT is a binary X.509 certificate, encapsulated in text (base-64) encoding.
		Data:      fullChainPEM,
		State:     as.StateChanged,
	}
	le.SecretFolder.SetFile(&crtFullChainFile)

	return
}

// https://github.com/golang/crypto/tree/master/acme/autocert
// https://github.com/hlandau/acmeapi
// https://github.com/caddyserver/caddy
// https://github.com/go-acme/lego
