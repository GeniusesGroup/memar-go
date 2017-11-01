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

// Parameters Object
type Parameters []*Parameter

// Parameter Object
type Parameter struct {
	SpecExtension `json:"_,omitempty"`
	Reference
	Name            string        `json:"name,omitempty"`
	In              string        `json:"in,omitempty"`
	Description     string        `json:"description,omitempty"`
	Deprecated      bool          `json:"deprecated,omitempty"`
	Required        bool          `json:"required,omitempty"`
	Style           string        `json:"style,omitempty"`
	Explode         bool          `json:"explode,omitempty"`
	AllowEmptyValue bool          `json:"allowEmptyValue,omitempty"`
	AllowReserved   bool          `json:"allowReserved,omitempty"`
	Schema          *Schema       `json:"schema,omitempty"`
	Example         interface{}   `json:"example,omitempty"`
	Examples        []interface{} `json:"examples,omitempty"`
}

// PathParameter creates a path parameter, this is always required.
func PathParameter(name string) *Parameter {
	p := Parameter{Name: name, In: "path", Required: true}
	return &p
}

// QueryParameter creates a query parameter, this is always required.
func QueryParameter(name string) *Parameter {
	return &Parameter{Name: name, In: "query", Required: true}
}

// HeaderParameter creates a header parameter, this is always required.
func HeaderParameter(name string) *Parameter {
	return &Parameter{Name: name, In: "header", Required: true}
}

// CookieParameter creates a cookie parameter, this is always required.
func CookieParameter(name string) *Parameter {
	return &Parameter{Name: name, In: "cookie", Required: true}
}

// WithDescription a fluent builder method for the description of the parameter
func (p *Parameter) WithDescription(value string) *Parameter {
	p.Description = value
	return p
}

// WithRequired flags this parameter as required||optional
func (p *Parameter) WithRequired(value bool) *Parameter {
	p.Required = value
	return p
}

// WithSchema :
func (p *Parameter) WithSchema(schema *Schema) *Parameter {
	p.Schema = schema
	return p
}
