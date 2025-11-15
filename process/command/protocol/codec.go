/* For license and copyright information please see the LEGAL file in the code repository */

package command_p

import (
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	request_p "memar/process/request/protocol"
	response_p "memar/process/response/protocol"
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
