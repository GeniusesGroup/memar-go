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

// Content Object
// map string name can any mime type e.g. "application/json"
type Content map[string]*MediaType

// MediaType Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#media-type-object
type MediaType struct {
	SpecExtension `json:"_,omitempty"`
	Schema        *Schema              `json:"schema,omitempty"`
	Examples      map[string]*Example  `json:"examples,omitempty"`
	Encoding      map[string]*Encoding `json:"encoding,omitempty"`
}

// NewMediaType :
func NewMediaType() *MediaType {
	mt := MediaType{}
	return &mt
}

// WithSchema :
func (contentType *MediaType) WithSchema(schema *Schema) *MediaType {
	contentType.Schema = schema
	return contentType
}

// WithExample :
func (contentType *MediaType) WithExample(name string, example *Example) *MediaType {
	contentType.Examples[name] = example
	return contentType
}

// WithEncoding :
func (contentType *MediaType) WithEncoding(name string, encoding *Encoding) *MediaType {
	contentType.Encoding[name] = encoding
	return contentType
}
