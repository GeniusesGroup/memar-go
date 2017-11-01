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
	"reflect"
	"strings"

	"github.com/SabzCity/go-library/errors"
)

const tagName = "valid"

// Validator : This can manage special field of object.
type Validator func(interface{}) bool

var validators = map[string]Validator{}

// ValidateObject : Validate any objects with validation tags.
func ValidateObject(object interface{}) error {
	var validatedFailed []string

	// Empty objects are valid by default.
	if object == nil {
		return nil
	}

	value := reflect.ValueOf(object)
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	// The type of object most be struct.
	if value.Kind() != reflect.Struct {
		return errors.ContentMustBeStruct
	}

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := value.Type().Field(i)

		// Check Private field
		if fieldType.PkgPath != "" {
			continue
		}

		if fieldValue.Kind() == reflect.Struct {
			err := ValidateObject(fieldValue.Interface())
			if err != nil {
				if errors.IsEqual(err, errors.RequestDataIsNotValid) {
					failedFields := err.(*errors.ExtendedError).Information.([]string)

					for i := 0; i < len(failedFields); i++ {
						validatedFailed = append(validatedFailed, fieldType.Name+"."+failedFields[i])
					}
				} else {
					return err
				}
			}
		} else {
			tagValue := fieldType.Tag.Get(tagName)

			// Check if the field should be ignored.
			if tagValue == "" {
				continue
			}

			validatorNames := strings.Split(tagValue, ",")

		L1:
			for _, validatorName := range validatorNames {
				switch validatorName {
				case "required":
					if isEmptyValue(fieldValue) {
						validatedFailed = append(validatedFailed, fieldType.Name)
						break L1
					}
				case "optional":
					if isEmptyValue(fieldValue) {
						break L1
					}
				default:
					validator, exist := validators[validatorName]
					// Check validator exist or not.
					if !exist {
						return errors.AddInformation(errors.ValidatorNotFound, validatorName)
					}
					if !validator(fieldValue.Interface()) {
						validatedFailed = append(validatedFailed, fieldType.Name)
						break L1
					}
				}
			}
		}
	}

	if len(validatedFailed) != 0 {
		return errors.AddInformation(errors.RequestDataIsNotValid, validatedFailed)
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
