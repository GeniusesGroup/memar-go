/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"encoding/json"

	"../protocol"
)

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these codec instead use Codec and autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain better performance from this protocol!
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
func Marshal(s interface{}) (p []byte, err protocol.Error) {
	// TODO::: make better algorithm instead of below
	var goErr error
	p, goErr = json.Marshal(s)
	if goErr != nil {
		return nil, ErrEncodedCorrupted
	}
	return
}

// Unmarshal decode payload and stores the result in the value pointed to by s in runtime.
func Unmarshal(p []byte, s interface{}) (err protocol.Error) {
	// TODO::: make better algorithm instead of below
	var goErr error = json.Unmarshal(p, s)
	if goErr != nil {
		return ErrEncodedCorrupted
	}
	return
}

// RunTimeCodec is a wrapper to use anywhere need protocol.Codec interface instead of protocol.JSON interface
type RunTimeCodec struct {
	t       interface{}
	decoder interface{}
	encoder interface{}
	len     int
}

func NewRunTimeCodec(t interface{}) (codec *RunTimeCodec) {
	codec = &RunTimeCodec{
		t: t,
		// len: json.LenAsJSON(),
	}
	// codec.encoder.buf = make([]byte, 0, codec.len)
	return
}
