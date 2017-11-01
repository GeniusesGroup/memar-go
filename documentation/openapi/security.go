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

// SecurityRequirements :
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#security-requirement-object
type SecurityRequirements []SecurityRequirement

// SecurityRequirement :
type SecurityRequirement map[string][]string

// AddSecurityRequirement :
func (srs *SecurityRequirements) AddSecurityRequirement(provider string, scopes ...string) *SecurityRequirements {

	if srs == nil {
		srs = &SecurityRequirements{}
	}

	var sr SecurityRequirement
	sr[provider] = scopes

	*srs = append(*srs, sr)

	return srs
}
