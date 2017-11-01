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

import (
	"fmt"
	"regexp"
)

// SchemaStringFormats :
var SchemaStringFormats = make(map[string]*regexp.Regexp, 8)

// DefineStringFormat :
func DefineStringFormat(name string, pattern string) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		err := fmt.Errorf("Format '%v' has invalid pattern '%v': %v", name, pattern, err)
		panic(err)
	}
	SchemaStringFormats[name] = re
}

func init() {
	// This pattern catches only some suspiciously wrong-looking email addresses.
	// Use DefineStringFormat(...) if you need something stricter.
	DefineStringFormat("email", `^[^@]+@[^@<>",\w]+$`)

	// Base64
	// The pattern supports base64 and b./ase64url. Padding ('=') is supported.
	DefineStringFormat("byte", `(^$|^[a-zA-Z0-9+/\-_]*=*$)`)

	// date
	DefineStringFormat("date", `^[0-9]{4}-(0[0-9]|11|12)-([0-2][0-9]|30|31)$`)

	// date-time
	DefineStringFormat("date-time", `^[0-9]{4}-(0[0-9]|11|12)-([0-2][0-9]|30|31)T[0-9]{2}:[0-9]{2}:[0-9]{2}(Z|(\+|-)[0-9]{2}:[0-9]{2})?$`)
}
