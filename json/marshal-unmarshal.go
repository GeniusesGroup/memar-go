/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"encoding/json"
)

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use CompleteMethods() to autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain max performance from this protocol!
*/

/*
Field int `json:"{Name},{Option}"`
Options:
- Name		Other name than field name. Must be first option.
- dash(-)	Don't encode||decode field!
- omitempty Encode||Decode only field not nil!
- string	base64 string! Use when array||slice numbers is not represent meaningful data. Also encode|decode numbers as string with "".
- tuple		must assign to all fields to encode|decode as tuple array instead key-value object!

Examples of struct field tags and their meanings:
Field int `json:"myName"`           // Field appears in JSON as key "myName".
Field int `json:"myName,omitempty"` // Field appears in JSON as key "myName" and the field is omitted from the object if its value is empty.
Field int `json:",omitempty"`       // Field appears in JSON as key "Field" (the default), but the field is skipped if empty.
Field int `json:"-"`                // Field is ignored by this package.
Field int `json:"-,"`               // Field appears in JSON as key "-".
*/

// Marshal encodes the value of s to the payload buffer in runtime.
func Marshal(s interface{}) (p []byte, err error) {
	// TODO::: make better algorithm instead of below
	p, err = json.Marshal(s)
	return
}

// UnMarshal decode payload and stores the result in the value pointed to by s in runtime.
func UnMarshal(p []byte, s interface{}) (err error) {
	// TODO::: make better algorithm instead of below
	err = json.Unmarshal(p, s)
	return
}
