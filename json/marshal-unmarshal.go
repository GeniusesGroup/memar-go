/* For license and copyright information please see LEGAL file in repository */

package json

import "encoding/json"

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use CompleteEncoderMethodSafe() to autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain max performance from this protocol!
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
