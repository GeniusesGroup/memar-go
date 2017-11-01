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

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncryptAES : Encrypt data from source with AES algorithm.
func EncryptAES(src, key []byte) ([]byte, error) {

	// Create the AES cipher.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Include the IV at the beginning.
	dst := make([]byte, aes.BlockSize+len(src))

	// Slice of first 16 bytes.
	iv := dst[:aes.BlockSize]

	// Write 16 rand bytes to fill iv.
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	// Create an encrypted stream.
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from source to destination.
	stream.XORKeyStream(dst[aes.BlockSize:], src)

	return dst, nil
}

// DecryptAES : Decrypt data from source with AES algorithm.
func DecryptAES(src, key []byte) ([]byte, error) {

	// Create the AES cipher.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Get the 16 bytes IV.
	iv := src[:aes.BlockSize]

	// Remove the IV from the source.
	esrc := src[aes.BlockSize:]

	// Create a decrypted stream.
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from source.
	stream.XORKeyStream(esrc, esrc)

	return esrc, nil
}
