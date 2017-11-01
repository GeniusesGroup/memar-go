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

// Responses Object
type Responses map[string]*Response

// Response Object
type Response struct {
	SpecExtension `json:"_,omitempty"`
	Reference
	Description string             `json:"description,omitempty"`
	Headers     map[string]*Schema `json:"headers,omitempty"`
	Content     Content            `json:"content,omitempty"`
	Links       map[string]*Link   `json:"links,omitempty"`
}

// NewResponse :
func NewResponse(description string) *Response {
	r := Response{}
	r.Description = description
	return &r
}

// AddContent :
func (r *Response) AddContent(mimeType string, mediaType *MediaType) *Response {

	if r.Content == nil {
		r.Content = Content{}
	}

	r.Content[mimeType] = mediaType
	return r
}
