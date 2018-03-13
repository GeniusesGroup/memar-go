//Copyright 2017 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package certificates

import (
	"crypto/tls"
	"time"

	"github.com/SabzCity/go-library/cache"
	"github.com/SabzCity/go-library/errors"
)

var certificateCache = cache.NewObjectCache(time.Hour * 24 * 7)

// GetCertificate : Return special certificate for client.
func GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	certificate, err := certificateCache.Get(clientHello.ServerName)
	if err != nil {
		newCertificate := getWithoutCache(clientHello.ServerName)
		if newCertificate == nil {
			return nil, errors.TLSCertificateNotExist
		}

		// And now we add "newCertificate" to local cache.
		certificateCache.Set(clientHello.ServerName, newCertificate, time.Hour*24)

		return newCertificate, nil
	}

	return certificate.(*tls.Certificate), nil
}

func getWithoutCache(DNS string) *tls.Certificate {

	// If certificate exist in Doamins MS convert it to "*tls.Certificate" type.
	certificate, err := tls.X509KeyPair([]byte(testCertificate), []byte(testPrivateKey))
	if err != nil {
		return nil
	}

	return &certificate
}
