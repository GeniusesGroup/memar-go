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

package code

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

// NumericCode : Generate a random number with 7 digits.
func NumericCode() string {

	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9999999-1000000) + 1000000

	return strconv.Itoa(code)
}

// StringCode : Generate a UUID (version 4).
func StringCode() string {

	return uuid.NewV4().String()
}
