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

package openapi

import (
	"reflect"
	"strings"
)

// NewSimpleSchema make new simple schema
// type/format can be : integer/int32 || integer/int64 || number/float || number/double || string/""
// || string/byte || string/binary || boolean/"" || string/date || string/date-time || string/password
func NewSimpleSchema(Type, Format string) *Schema {
	s := Schema{Type: Type, Format: Format}
	return &s
}

// NewObjectSchema make new schema by given object
func NewObjectSchema(object interface{}) *Schema {
	// Empty objects.
	if object == nil {
		return nil
	}

	st := reflect.ValueOf(object)
	// Turn pointers into simpler types so further checks are correct.
	if st.Kind() == reflect.Interface || st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	// The type of object most be struct.
	if st.Kind() != reflect.Struct {
		return nil
	}

	objectSchema := Schema{
		Required:   []string{},
		Properties: map[string]*Schema{},
	}

	for i := 0; i < st.NumField(); i++ {
		fieldValue := st.Field(i)
		fieldType := st.Type().Field(i)
		var innerSchema *Schema
		if fieldValue.Kind() == reflect.Struct {
			innerSchema = NewObjectSchema(fieldValue.Interface())
			objectSchema.Properties[nameOfField(fieldType)] = innerSchema
			if isPropertyRequired(fieldType) {
				objectSchema.Required = append(objectSchema.Required, nameOfField(fieldType))
			}
		} else if fieldValue.Kind() == reflect.Slice || fieldValue.Kind() == reflect.Array {
			innerSchema = &Schema{Items: &Schema{Type: jsonSchemaType(fieldType.Type.Elem().Name()), Format: jsonSchemaFormat(fieldType.Type.Elem().Name())}}
			objectSchema.Properties[nameOfField(fieldType)] = innerSchema
			if isPropertyRequired(fieldType) {
				objectSchema.Required = append(objectSchema.Required, nameOfField(fieldType))
			}
		} else {
			innerSchema = &Schema{Type: jsonSchemaType(fieldType.Type.Name()), Format: jsonSchemaFormat(fieldType.Type.Name())}
			objectSchema.Properties[nameOfField(fieldType)] = innerSchema
			if isPropertyRequired(fieldType) {
				objectSchema.Required = append(objectSchema.Required, nameOfField(fieldType))
			}
		}
	}

	return &objectSchema
}

// nameOfField returns the name of the field as it should appear in JSON format
// An empty string indicates that this field is not part of the JSON representation
func nameOfField(field reflect.StructField) string {
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		s := strings.Split(jsonTag, ",")
		if s[0] == "-" {
			// empty name signals skip property
			return ""
		} else if s[0] != "" {
			return s[0]
		}
	}
	return field.Name
}

func isPropertyRequired(field reflect.StructField) bool {
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		s := strings.Split(jsonTag, ",")
		if len(s) > 1 && s[1] == "omitempty" {
			return false
		}
	}
	return true
}

// see also http://json-schema.org/latest/json-schema-core.html#anchor8
func jsonSchemaType(modelName string) string {
	mapped, ok := JSONSchemaTypes[modelName]
	if !ok {
		return modelName // use as is (custom or struct)
	}
	return mapped
}

func jsonSchemaFormat(modelName string) string {
	mapped, ok := JSONSchemaFormats[modelName]
	if !ok {
		return "" // no format
	}
	return mapped
}

var (
	// JSONSchemaFormats is json standard formats
	JSONSchemaFormats map[string]string
	// JSONSchemaTypes is json standard types
	JSONSchemaTypes map[string]string
)

func init() {
	JSONSchemaFormats = map[string]string{
		"int":        "int32",
		"int32":      "int32",
		"int64":      "int64",
		"byte":       "byte",
		"uint":       "integer",
		"uint8":      "byte",
		"float64":    "double",
		"float32":    "float",
		"time.Time":  "date-time",
		"*time.Time": "date-time",
	}

	JSONSchemaTypes = map[string]string{
		"uint":   "integer",
		"uint8":  "integer",
		"uint16": "integer",
		"uint32": "integer",
		"uint64": "integer",

		"int":   "integer",
		"int8":  "integer",
		"int16": "integer",
		"int32": "integer",
		"int64": "integer",

		"byte":      "integer",
		"float64":   "number",
		"float32":   "number",
		"bool":      "boolean",
		"time.Time": "string",
	}
}
