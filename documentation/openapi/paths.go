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

// Paths Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#paths-object
type Paths map[string]*PathItem

// PathItem Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#path-item-object
type PathItem struct {
	SpecExtension `json:"_,omitempty"`
	Reference
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	DELETE      *Operation `json:"delete,omitempty"`
	GET         *Operation `json:"get,omitempty"`
	HEAD        *Operation `json:"head,omitempty"`
	OPTIONS     *Operation `json:"options,omitempty"`
	PATCH       *Operation `json:"patch,omitempty"`
	POST        *Operation `json:"post,omitempty"`
	PUT         *Operation `json:"put,omitempty"`
	TRACE       *Operation `json:"trace,omitempty"`
	Servers     Servers    `json:"servers,omitempty"`
	Parameters  Parameters `json:"parameters,omitempty"`
}
