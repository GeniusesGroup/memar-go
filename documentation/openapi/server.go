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

// Servers :
type Servers []*Server

// Server Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#server-object
type Server struct {
	SpecExtension `json:"_,omitempty"`
	URL           string                     `json:"url,omitempty"`
	Description   string                     `json:"description,omitempty"`
	Variables     map[string]*ServerVariable `json:"variables,omitempty"`
}

// ServerVariable Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#server-variable-object
type ServerVariable struct {
	SpecExtension `json:"_,omitempty"`
	Enum          []interface{} `json:"enum,omitempty"`
	Default       interface{}   `json:"default,omitempty"`
	Description   string        `json:"description,omitempty"`
}
