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

package jwt

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

// GenerateJWT : Generate a new JWT with special parameters.
func GenerateJWT(payload map[string]interface{}) []byte {

	claims := jws.Claims{}

	for key, value := range payload {
		claims.Set(key, value)
	}

	token := jws.NewJWT(claims, crypto.SigningMethodHS256)

	data, _ := token.Serialize(signatureKey)

	return data
}

// ValidateToken : Validate the JWT.
func ValidateToken(encoded string) (bool, error) {

	token, err := jws.ParseJWT([]byte(encoded))
	if err != nil {
		return false, err
	}

	err = token.Validate(signatureKey, crypto.SigningMethodHS256)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetTokenData : Validate the token and Return the data in JWT.
func GetTokenData(encoded string) (map[string]interface{}, error) {

	token, err := jws.ParseJWT([]byte(encoded))
	if err != nil {
		return nil, err
	}

	err = token.Validate(signatureKey, crypto.SigningMethodHS256)
	if err != nil {
		return nil, err
	}

	return token.Claims(), nil
}
