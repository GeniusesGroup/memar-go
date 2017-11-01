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

// Info Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#info-object
type Info struct {
	SpecExtension  `json:"_,omitempty"`
	Title          string   `json:"title,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version,omitempty"`
}

// Contact Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#contact-object
type Contact struct {
	SpecExtension `json:"_,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
	Email         string `json:"email,omitempty"`
}

// License Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#license-object
type License struct {
	SpecExtension `json:"_,omitempty"`
	Name          string `json:"name,omitempty"`
	URL           string `json:"url,omitempty"`
}
