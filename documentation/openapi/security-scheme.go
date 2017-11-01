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

// SecurityScheme Object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#security-scheme-object
type SecurityScheme struct {
	SpecExtension `json:"_,omitempty"`
	Reference
	Type         string      `json:"type,omitempty"`
	Description  string      `json:"description,omitempty"`
	Name         string      `json:"name,omitempty"`
	In           string      `json:"in,omitempty"`
	Scheme       string      `json:"scheme,omitempty"`
	BearerFormat string      `json:"bearerFormat,omitempty"`
	Flow         *OAuthFlows `json:"flow,omitempty"`
}

// NewSecurityScheme :
func NewSecurityScheme() *SecurityScheme {
	return &SecurityScheme{}
}

// NewCSRFSecurityScheme :
func NewCSRFSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type: "apiKey",
		In:   "header",
		Name: "X-XSRF-TOKEN",
	}
}

// NewJWTSecurityScheme :
func NewJWTSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
	}
}

// WithType :
func (ss *SecurityScheme) WithType(value string) *SecurityScheme {
	ss.Type = value
	return ss
}

// WithDescription :
func (ss *SecurityScheme) WithDescription(value string) *SecurityScheme {
	ss.Description = value
	return ss
}

// WithName :
func (ss *SecurityScheme) WithName(value string) *SecurityScheme {
	ss.Name = value
	return ss
}

// WithIn :
func (ss *SecurityScheme) WithIn(value string) *SecurityScheme {
	ss.In = value
	return ss
}

// WithScheme :
func (ss *SecurityScheme) WithScheme(value string) *SecurityScheme {
	ss.Scheme = value
	return ss
}

// WithBearerFormat :
func (ss *SecurityScheme) WithBearerFormat(value string) *SecurityScheme {
	ss.BearerFormat = value
	return ss
}

// OAuthFlows :
type OAuthFlows struct {
	SpecExtension     `json:"_,omitempty"`
	Implicit          *OAuthFlow `json:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}

// OAuthFlow :
type OAuthFlow struct {
	SpecExtension    `json:"_,omitempty"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes"`
}
