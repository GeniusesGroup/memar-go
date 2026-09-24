/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use desire services to autogenerate needed code before compile time
	and reduce runtime process to improve performance of the app and gain max performance from this protocol!
*/

// Marshal encodes the value of s to the buffer in runtime.
func Marshal(source any, buffer buffer_p.Buffer) (err error_p.Error) {
	return
}

// Unmarshal decode buffer and stores the result in the value pointed to by s in runtime.
func Unmarshal(buffer buffer_p.Buffer, destination any) (err error_p.Error) {
	return
}
