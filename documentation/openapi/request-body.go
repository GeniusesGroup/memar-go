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

// RequestBody Object
type RequestBody struct {
	SpecExtension `json:"_,omitempty"`
	Reference
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Content     Content `json:"content,omitempty"`
}

// NewRequestBody :
func NewRequestBody(description string) *RequestBody {
	rb := RequestBody{}
	rb.Description = description
	return &rb
}

// WithRequired :
func (rb *RequestBody) WithRequired(value bool) *RequestBody {
	rb.Required = value
	return rb
}

// AddContent :
func (rb *RequestBody) AddContent(mimeType string, mediaType *MediaType) *RequestBody {
	if rb.Content == nil {
		rb.Content = Content{}
	}

	rb.Content[mimeType] = mediaType
	return rb
}
