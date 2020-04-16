/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"text/template"
)

// UpdateServiceFileReq is request structure of UpdateServiceFile()
type UpdateServiceFileReq struct {
	ServiceFile []byte
}

// UpdateServiceFileRes is response structure of UpdateServiceFile()
type UpdateServiceFileRes struct {
	ServiceFile []byte
}

// UpdateServiceFile use to update service file and complete or edit some auto generate part.
func UpdateServiceFile(req *UpdateServiceFileReq) (res *UpdateServiceFileRes, err error) {
	// Check and update service detail essentially service ID!
	// Update handler function
	// Update encoders
	// Update validator

	const tagName = "valid"

	// Validate each data with tagName with related function from validators folder.
	// valid data can be required, optional, ...

	var sf = new(bytes.Buffer)
	err = sRPCTemplate.Execute(sf, "sg.ReqRepo.Functions")
	if err != nil {
		return nil, err
	}

	return res, nil
}

var sRPCTemplate = template.Must(template.New("sRPCTemplate").Funcs(template.FuncMap{
	"minus": func(a, b int) int {
		return a - b
	},
	"plus": func(a, b int) int {
		return a + b
	},
}).Parse(`
// HandleSRPC : sRPC is our experimental SabzCity remote procedure call Protocol!!
func HandleSRPC(st *chaparkhane.Stream) {

	switch st.ServiceID {
	{{- range . }}
	case {{.ID}}:
		var req = logic.{{.Parameter.Type}}{
			{{- range .Parameter.InnerType }}
				{{- if (eq .Type "string")}}
					{{.Name}}: string(st.Payload[st.Offsets[{{.ID}}] : st.Offsets[{{plus .ID 1}}]-1]),
				{{- else if (eq .Type "array")}}
				{{- else if (eq .Type "int8")}}
					{{.Name}}: int8(st.Payload[st.Offsets[{{.ID}}]]),
				{{- else if (eq .Type "int16")}}
					{{.Name}}: int16(st.Payload[st.Offsets[{{.ID}}]]) | int16(st.Payload[st.Offsets[{{.ID}}]+1])<<8,
				{{- else if (eq .Type "int32")}}
					{{.Name}}: int32(st.Payload[st.Offsets[{{.ID}}]]) | int32(st.Payload[st.Offsets[{{.ID}}]+1])<<8 | int32(st.Payload[st.Offsets[{{.ID}}]+2])<<16 | int32(st.Payload[st.Offsets[{{.ID}}]+3])<<24,
				{{- else if (eq .Type "int64")}}
					{{.Name}}: int64(st.Payload[st.Offsets[{{.ID}}]]) | int64(st.Payload[st.Offsets[{{.ID}}]+1])<<8 | int64(st.Payload[st.Offsets[{{.ID}}]+2])<<16 | int64(st.Payload[st.Offsets[{{.ID}}]+3])<<24 | int64(st.Payload[st.Offsets[{{.ID}}]+4])<<32 | int64(st.Payload[st.Offsets[{{.ID}}]+5])<<40 | int64(st.Payload[st.Offsets[{{.ID}}]+6])<<48 | int64(st.Payload[st.Offsets[{{.ID}}]+7])<<56,
				{{- else if (eq .Type "uint8")}}
					{{.Name}}: uint8(st.Payload[st.Offsets[{{.ID}}]]),
				{{- else if (eq .Type "uint16")}}
					{{.Name}}: uint16(st.Payload[st.Offsets[{{.ID}}]]) | uint16(st.Payload[st.Offsets[{{.ID}}]+1])<<8,
				{{- else if (eq .Type "uint32")}}
					{{.Name}}: uint32(st.Payload[st.Offsets[{{.ID}}]]) | uint32(st.Payload[st.Offsets[{{.ID}}]+1])<<8 | uint32(st.Payload[st.Offsets[{{.ID}}]+2])<<16 | uint32(st.Payload[st.Offsets[{{.ID}}]+3])<<24,
				{{- else if (eq .Type "uint64")}}
					{{.Name}}: uint64(st.Payload[st.Offsets[{{.ID}}]]) | uint64(st.Payload[st.Offsets[{{.ID}}]+1])<<8 | uint64(st.Payload[st.Offsets[{{.ID}}]+2])<<16 | uint64(st.Payload[st.Offsets[{{.ID}}]+3])<<24 | uint64(st.Payload[st.Offsets[{{.ID}}]+4])<<32 | uint64(st.Payload[st.Offsets[{{.ID}}]+5])<<40 | uint64(st.Payload[st.Offsets[{{.ID}}]+6])<<48 | uint64(st.Payload[st.Offsets[{{.ID}}]+7])<<56,
				{{- else if (eq .Type "string")}}
					{{.Name}}: *(*string)(unsafe.Pointer(&st.Payload[st.Offsets[{{.ID}}]:]))
				{{- end}}
				
			{{- end }}
		}
		{{if .Result}} var res *logic.{{.Result.Type}} {{- end}}
		{{if .Result}} res {{- end}} {{- if .Err}}, st.Err {{- end}} = logic.{{.Name}}(&req, sd)
		// handle res & st.Err if exist
		if res != nil {
			
		}
	{{- end }}
	default:
		st.Err = chaparkhane.ServiceNotFound
		// handle st.Err
	}
}
`))
