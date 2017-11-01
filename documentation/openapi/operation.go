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

import "strconv"

// Operation represents "operation" specified by" OpenAPI 3.0 standard.
type Operation struct {
	SpecExtension `json:"_,omitempty"`
	Tags          []string              `json:"tags,omitempty"`         // Optional tags for documentation.
	Summary       string                `json:"summary,omitempty"`      // Optional short summary.
	Description   string                `json:"description,omitempty"`  // Optional description. Should use CommonMark syntax.
	ExternalDocs  *ExternalDocs         `json:"externalDocs,omitempty"` //
	OperationID   string                `json:"operationId,omitempty"`  // Optional operation ID.
	Parameters    Parameters            `json:"parameters,omitempty"`   // Optional parameters.
	RequestBody   *RequestBody          `json:"requestBody,omitempty"`  // Optional body parameter.
	Responses     Responses             `json:"responses,omitempty"`    // Optional responses.
	Callbacks     map[string]*Callback  `json:"callbacks,omitempty"`    // Optional callbacks
	Deprecated    bool                  `json:"deprecated,omitempty"`   //
	Security      *SecurityRequirements `json:"security,omitempty"`     // Optional security requirements that overrides top-level security.
	Servers       *Servers              `json:"servers,omitempty"`      // Optional servers that overrides top-level servers.
}

// NewOperation creates a new operation instance.
// It expects an ID as parameter but not passing an ID is also valid.
func NewOperation(id string) *Operation {
	op := Operation{}
	op.OperationID = id
	return &op
}

// WithDescription sets the description on this operation, allows for chaining
func (o *Operation) WithDescription(description string) *Operation {
	o.Description = description
	return o
}

// WithSummary sets the summary on this operation, allows for chaining
func (o *Operation) WithSummary(summary string) *Operation {
	o.Summary = summary
	return o
}

// WithExternalDocs sets/removes the external docs for/from this operation.
func (o *Operation) WithExternalDocs(description, url string) *Operation {
	ed := ExternalDocs{description, url}
	o.ExternalDocs = &ed
	return o
}

// Deprecate marks the operation as deprecated
func (o *Operation) Deprecate() *Operation {
	o.Deprecated = true
	return o
}

// Undeprecate marks the operation as not deprected
func (o *Operation) Undeprecate() *Operation {
	o.Deprecated = false
	return o
}

// WithTags adds tags for this operation
func (o *Operation) WithTags(tags ...string) *Operation {
	o.Tags = append(o.Tags, tags...)
	return o
}

// AddParam adds a parameter to this operation.
func (o *Operation) AddParam(param *Parameter) *Operation {

	if param == nil {
		return o
	}

	//// Check parameter for that location and replaced if that name already exists.

	o.Parameters = append(o.Parameters, param)

	return o
}

// AddRequestBody adds a request body to the operation.
func (o *Operation) AddRequestBody(requestBody *RequestBody) *Operation {
	o.RequestBody = requestBody
	return o
}

// AddResponse adds a status code response to the operation.
func (o *Operation) AddResponse(statusCode int, response *Response) *Operation {

	if o.Responses == nil {
		o.Responses = Responses{}
	}

	// When the statusCode is 0 the value of the response will be used as default response value.
	if statusCode == 0 {
		o.Responses["default"] = response
	} else {
		o.Responses[strconv.FormatInt(int64(statusCode), 10)] = response
	}

	return o
}
