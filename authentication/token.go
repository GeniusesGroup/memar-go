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

package authentication

import (
	"encoding/json"
	"time"

	"github.com/SabzCity/go-library/cryptography/aes"
	"github.com/SabzCity/go-library/encoding/base"
	"github.com/SabzCity/go-library/encoding/ejson"
	"github.com/SabzCity/go-library/microservice"
)

// Token : The standard struct of authentication Token
type Token struct {
	Delegate  bool   //SabzCity Platform or Delegate Requests
	TokenID   string //
	ExpiredAt int64
	Audience  string
	Plan      bool
}

// EncryptToken : Encrypt "data" with AES algorithm.
func EncryptToken(decryptToken *Token) (string, error) {

	jDoc, _ := ejson.Marshal(decryptToken)

	encrypted, err := aes.EncryptAES([]byte(jDoc), microservice.Key(time.Now()))
	if err != nil {
		return "", err
	}

	return base.StringToBase64(string(encrypted)), nil
}

// DecryptToken : Decrypt to a "interface".
func DecryptToken(token string, decryptToken *Token) error {

	source := []byte(base.Base64ToString(token))

	jDoc, err := aes.DecryptAES(source, microservice.Key(time.Now()))
	if err != nil {
		return err
	}

	err = json.Unmarshal(jDoc, decryptToken)
	if err != nil {
		return err
	}

	return nil
}
