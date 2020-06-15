/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// cryptography : Public-key for related domain.
type cryptography struct {
	publicKey  [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	privateKey [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
}

// RegisterPublicKey use to register public key in new domain name systems like apis.sabz.city
func (c *cryptography) registerPublicKey() (err error) {
	// make public & private key and store them
	c.publicKey = [32]byte{}
	c.privateKey = [32]byte{}

	return nil
}

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
