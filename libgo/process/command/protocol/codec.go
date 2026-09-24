/* For license and copyright information please see the LEGAL file in the code repository */

package command_p

import (
	error_p "memar/process/error/protocol"
)

// Codec is implement by operation request and response.
type Codec interface {
	Encoder
	Decoder
}

type Decoder interface {
	FromCLA(args Arguments) (remaining Arguments, err error_p.Error)
}

type Encoder interface {
	ToCLA() (args Arguments, err error_p.Error)
}
