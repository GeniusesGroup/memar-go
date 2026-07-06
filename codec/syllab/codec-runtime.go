/* For license and copyright information please see LEGAL file in repository */

package syllab

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use CompleteEncoderMethodSafe() to autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain max performance from this protocol!
*/

// Marshal encodes the value of s to the payload buffer in runtime.
// offset add free space by given number at begging of return slice that almost just use in sRPC protocol! It can be 0!!
func Marshal(s interface{}, offset int) (p []byte, err error) {
	return
}

// UnMarshal decode payload and stores the result in the value pointed to by s in runtime.
func UnMarshal(p []byte, s interface{}) (err error) {
	return
}
