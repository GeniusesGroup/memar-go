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

// Schema Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#schema-object
type Schema struct {
	SpecExtension `json:"_,omitempty"`
	Reference

	OneOf        []*Schema     `json:"oneOf,omitempty"`
	AnyOf        []*Schema     `json:"anyOf,omitempty"`
	AllOf        []*Schema     `json:"allOf,omitempty"`
	Not          *Schema       `json:"not,omitempty"`
	Type         string        `json:"type,omitempty"`   // https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#data-types
	Format       string        `json:"format,omitempty"` // https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#data-types
	Description  string        `json:"description,omitempty"`
	Enum         []interface{} `json:"enum,omitempty"`
	Default      interface{}   `json:"default,omitempty"`
	Example      interface{}   `json:"example,omitempty"`
	Examples     []interface{} `json:"examples,omitempty"`
	ExternalDocs interface{}   `json:"externalDocs,omitempty"`

	// Properties
	Nullable  bool        `json:"nullable,omitempty"`
	ReadOnly  bool        `json:"readOnly,omitempty"`
	WriteOnly bool        `json:"writeOnly,omitempty"`
	XML       interface{} `json:"xml,omitempty"`

	// Number
	ExclusiveMinimum float64 `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum float64 `json:"exclusiveMaximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty"`
	Maximum          float64 `json:"maximum,omitempty"`
	MultipleOf       int64   `json:"multipleOf,omitempty"` // Numbers can be restricted to a multiple of a given number. It may be set to any positive number.

	// String
	MinLength uint64 `json:"minLength,omitempty"`
	MaxLength uint64 `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"` // This string SHOULD be a valid regular expression, according to the ECMA 262 regular expression dialect

	// Array
	MinItems uint64  `json:"minItems,omitempty"`
	MaxItems uint64  `json:"maxItems,omitempty"`
	Items    *Schema `json:"items,omitempty"`

	// Object
	Required                    []string           `json:"required,omitempty"`
	Properties                  map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties        *Schema            `json:"-" multijson:"additionalProperties,omitempty"`
	AdditionalPropertiesAllowed bool               `json:"-" multijson:"additionalProperties,omitempty"`
	Discriminator               string             `json:"discriminator,omitempty"`
	PatternProperties           string             `json:"patternProperties,omitempty"`
}

// WithEnum :
func (s *Schema) WithEnum(values ...interface{}) *Schema {
	s.Enum = values
	return s
}

// Number Methods

// WithMin add minimum value for number item schema
func (s *Schema) WithMin(min float64) *Schema {
	s.Minimum = min
	return s
}

// WithMax add maximum value for number item schema
func (s *Schema) WithMax(max float64) *Schema {
	s.Maximum = max
	return s
}

// WithExclusiveMin add exclusive minimum value for number item schema
func (s *Schema) WithExclusiveMin(eMin float64) *Schema {
	s.ExclusiveMinimum = eMin
	return s
}

// WithExclusiveMax add exclusive maximum value for number item schema
func (s *Schema) WithExclusiveMax(eMax float64) *Schema {
	s.ExclusiveMaximum = eMax
	return s
}

// String Methods

// WithMinLength add minimum length for string item schema
func (s *Schema) WithMinLength(n uint64) *Schema {
	s.MinLength = n
	return s
}

// WithMaxLength add maximum length for string item schema
func (s *Schema) WithMaxLength(n uint64) *Schema {
	s.MaxLength = n
	return s
}

// WithPattern add regex pattern for string item schema
func (s *Schema) WithPattern(pattern string) *Schema {
	s.Pattern = pattern
	return s
}

// Array Methods

// WithItems add schema object for array type schema
func (s *Schema) WithItems(schema *Schema) *Schema {
	s.Items = schema
	return s
}

// WithMinItems add minimum number of array length
func (s *Schema) WithMinItems(n uint64) *Schema {
	s.MinItems = n
	return s
}

// WithMaxItems add maximum number of array length
func (s *Schema) WithMaxItems(n uint64) *Schema {
	s.MaxItems = n
	return s
}

// Object Methods

// WithProperty :
func (s *Schema) WithProperty(name string, propertySchema *Schema) *Schema {
	return s.WithPropertyRef(name, propertySchema)
}

// WithPropertyRef :
func (s *Schema) WithPropertyRef(name string, ref *Schema) *Schema {
	properties := s.Properties
	if properties == nil {
		properties = make(map[string]*Schema)
		s.Properties = properties
	}
	properties[name] = ref
	return s
}

// WithProperties :
func (s *Schema) WithProperties(properties map[string]*Schema) *Schema {
	result := make(map[string]*Schema, len(properties))
	for k, v := range properties {
		result[k] = v
	}
	s.Properties = result
	return s
}

// WithAnyAdditionalProperties :
func (s *Schema) WithAnyAdditionalProperties() *Schema {
	s.AdditionalProperties = nil
	s.AdditionalPropertiesAllowed = true
	return s
}

// WithAdditionalProperties :
func (s *Schema) WithAdditionalProperties(v *Schema) *Schema {
	if v == nil {
		s.AdditionalProperties = nil
	} else {
		s.AdditionalProperties = v
	}
	return s
}
