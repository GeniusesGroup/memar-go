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

// The OpenAPI 3.0 specification can be found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.md

package openapi

// OpenAPI is the root document object of the OpenAPI document.
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#openapi-object
type OpenAPI struct {
	SpecExtension `json:"_,omitempty"`
	OpenAPI       string               `json:"openapi"`
	Info          Info                 `json:"info"`
	Servers       Servers              `json:"servers,omitempty"`
	Paths         Paths                `json:"paths,omitempty"`
	Components    *Components          `json:"components,omitempty"`
	Security      SecurityRequirements `json:"security,omitempty"`
	Tags          Tags                 `json:"tags,omitempty"`
	ExternalDocs  *ExternalDocs        `json:"externalDocs,omitempty"`
}

// NewOpenAPI make new OpenAPI document.
func NewOpenAPI(info Info) *OpenAPI {
	oapi := OpenAPI{
		OpenAPI: "3.0",
		Info:    info,
		Paths:   Paths{}}

	return &oapi
}
