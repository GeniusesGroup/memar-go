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

// Components :
type Components struct {
	SpecExtension `json:"_,omitempty"`
	Schemas         map[string]*Schema         `json:"schemas,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty"`
	Headers         map[string]*Header         `json:"headers,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
	Examples        map[string]*Example        `json:"examples,omitempty"`
	Links           map[string]*Link           `json:"links,omitempty"`
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty"`
}
