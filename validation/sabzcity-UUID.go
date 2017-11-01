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
	"strings"

	"github.com/asaskevich/govalidator"

	"github.com/SabzCity/go-library/convert"
)

// IsUserID : Validate a (real/organization) UUID.
func IsUserID(in interface{}) bool {
	if str, ok := in.(string); ok {
		if UUID, err := convert.StringToUInt64(str); err == nil {
			if UUID >= 100000000000000 && UUID <= 999999999999999 {
				return true
			}
		}
	}
	return false
}

// IsRealUserID : Validate a real user UUID.
func IsRealUserID(in interface{}) bool {
	if str, ok := in.(string); ok {
		if UUID, err := convert.StringToUInt64(str); err == nil {
			if UUID >= 100000000000000 && UUID <= 799999999999999 {
				return true
			}
		}
	}
	return false
}

// IsOrganizationUserID : Validate a real user UUID.
func IsOrganizationUserID(in interface{}) bool {
	if str, ok := in.(string); ok {
		if UUID, err := convert.StringToUInt64(str); err == nil {
			if UUID >= 800000000000000 && UUID <= 999999999999999 {
				return true
			}
		}
	}
	return false
}

// IsGroupID : Validate a group ID.
func IsGroupID(in interface{}) bool {
	if str, ok := in.(string); ok {
		if govalidator.IsNumeric(str) && strings.Count(str, "")-1 == 35 {
			return true
		}
	}
	return false
}

// IsCoordinateID : Validate coordinate id.
func IsCoordinateID(in interface{}) bool {
	// TODO : Inquiring from Coordinate microservice and verify this id.
	if str, ok := in.(string); ok {
		if govalidator.IsNumeric(str) && strings.Count(str, "")-1 == 35 {
			return true
		}
	}
	return false
}

// IsPhoneID : Validate Phone ID.
func IsPhoneID(in interface{}) bool {
	if str, ok := in.(string); ok {
		if strings.Count(str, "")-1 < 20 {
			return true
		}
	}
	return false
}

// IsWarehouseID : Validate Warehouse ID.
func IsWarehouseID(in interface{}) bool {
	// TODO : Inquiring from Warehouse microservice and verify this id.
	if str, ok := in.(string); ok {
		if govalidator.IsNumeric(str) && strings.Count(str, "")-1 == 35 {
			return true
		}
	}
	return false
}

// IsProductID : Validate Product ID.
func IsProductID(in interface{}) bool {
	// TODO : Inquiring from Product microservice and verify this id.
	if str, ok := in.(string); ok {
		if govalidator.IsNumeric(str) && strings.Count(str, "")-1 == 35 {
			return true
		}
	}
	return false
}

// IsWikiRecordID : Validate Wiki record ID.
func IsWikiRecordID(in interface{}) bool {
	// TODO : Inquiring from wiki microservice and verify this id.
	if str, ok := in.(string); ok {
		if govalidator.IsNumeric(str) && strings.Count(str, "")-1 == 35 {
			return true
		}
	}
	return false
}
