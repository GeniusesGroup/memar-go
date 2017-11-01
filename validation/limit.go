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
	"github.com/SabzCity/go-library/convert"
)

// IsLimit : Check Limit is not more than Standard limit range. now it is 30!
func IsLimit(in interface{}) bool {

	if str, ok := in.(string); ok {
		if limit, err := convert.StringToUInt64(str); err == nil {
			// Standard limit range for platform can set it here!
			if limit != 0 && limit < 30 {
				return true
			}
		}
	}

	return false
}

// IsOffset : Check Offset is valid uint!
func IsOffset(in interface{}) bool {

	if str, ok := in.(string); ok {
		if _, err := convert.StringToUInt64(str); err == nil {
			return true
		}
	}

	return false
}
