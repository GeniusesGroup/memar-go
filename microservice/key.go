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

package microservice

import (
	"time"
)

// Key : Global key for encryption or decryption
// iat means Issued at! in sabzcity platform everymoment can have diffrent key!
func Key(iat time.Time) []byte {

	// default sabzcity key string is:   -1m3Bf)y~*<kD]>\
	return []byte{45, 49, 109, 51, 66, 102, 41, 121, 126, 42, 60, 107, 68, 93, 62, 92}
}
