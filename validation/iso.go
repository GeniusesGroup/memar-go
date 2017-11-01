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

package validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/ttacon/libphonenumber"
)

// IsNationality : Validate a national string code.
// e.g. IRN.
func IsNationality(in interface{}) bool {

	if str, ok := in.(string); ok {
		if govalidator.IsISO3166Alpha3(str) {
			return true
		}
	}

	return false
}

// IsValidPhone : Validate phone number.
func IsValidPhone(in interface{}) bool {

	if str, ok := in.(string); ok {
		if IsPhone(str, "IR") {
			return true
		}
	}

	return false
}

// IsPhone : Validate phone number.
func IsPhone(str, region string) bool {

	number, err := libphonenumber.Parse(str, region)
	if err != nil {
		return false
	}

	return libphonenumber.IsValidNumber(number)
}

// FormatPhoneE164 : Convert any format of phone number to E164 format.
// XXX : Don't use it for not valid numbers.
func FormatPhoneE164(str, region string) string {

	number, _ := libphonenumber.Parse(str, region)

	return libphonenumber.Format(number, libphonenumber.E164)
}
