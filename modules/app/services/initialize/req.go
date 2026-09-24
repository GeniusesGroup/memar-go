/* For license and copyright information please see the LEGAL file in the code repository */

package initialize

import (
	function_p "memar/computer/function/protocol"
	error_p "memar/process/error/protocol"
	ap "memar/modules/app/protocol"
	cmd "memar/process/command"
	"memar/protocol"
)

// Request :
type Request struct {
	domain      ap.Domain
	templateURI string
	blank       bool
}

//memar:impl /memar/modules/app/protocol.New_Request
func (req *Request) Domain() ap.Domain   { return req.domain }
func (req *Request) TemplateURI() string { return req.templateURI }
func (req *Request) Blank() bool         { return req.blank }

func (req *Request) TemplateURIField() string { return req.templateURI }
func (req *Request) BlankField() bool         { return req.blank }

//memar:impl memar/process/command/protocol.Object
func (req *Request) Fields() []protocol.DataType {
	return []protocol.DataType{(*ap.Domain)(&req.domain), req.templateURI, req.blank}
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
func (req *Request) FromJSON(buf []byte) (remaining []byte, err error_p.Error) { return }
func (req *Request) ToJSON(buf []byte) (added []byte, err error_p.Error)       { return buf, nil }
func (req *Request) JSON_Length() int                                             { return 0 }

func (req *Request) CheckAndSetDefaults() {
	if req.domain == "" && !req.DomainField().Optional() {
		req.DomainField().SetDefault()
	}
}
