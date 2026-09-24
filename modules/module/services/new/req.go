/* For license and copyright information please see the LEGAL file in the code repository */

package new

import (
	function_p "memar/computer/function/protocol"
	error_p "memar/process/error/protocol"
	cmd "memar/process/command"
	"memar/protocol"

	mp "memar/modules/module/protocol"
)

// Request :
type Request struct {
	name mp.DomainName
	path mp.Path
}

//memar:impl memar/protocol.New_Request
func (req *Request) Name() mp.DomainName { return req.name }
func (req *Request) Path() mp.Path       { return req.path }

func (req *Request) SetPath(p string) { req.path = mp.Path(p) }

//memar:impl memar/process/command/protocol.Object
func (req *Request) Fields() []protocol.DataType {
	return []protocol.DataType{(*mp.DomainName)(&req.name), (*mp.Path)(&req.path)}
}
func (req *Request) Methods() []function_p.Method { return nil }

//memar:impl memar/process/command/protocol.Decoder
func (req *Request) FromCLA(arguments []string) (remaining []string, err error_p.Error) {
	remaining, err = cmd.FromCLA(req, arguments)
	return
}
//memar:impl memar/process/command/protocol.Encoder
func (req *Request) ToCLA() (arguments []string, err error_p.Error) {
	arguments, err = cmd.ToCLA(req)
	return
}

//memar:impl memar/json/protocol.JSON
//memar:json
func (req *Request) FromJSON(payload []byte) (remaining []byte, err error_p.Error) { return }
func (req *Request) ToJSON(payload []byte) (added []byte, err error_p.Error)       { return payload, nil }
func (req *Request) JSON_Length() int                                             { return 0 }
