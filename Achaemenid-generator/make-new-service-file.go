/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"hash/crc32"
	"strings"
	"text/template"
	"time"

	"../assets"
)

// MakeNewServiceFile use to make new service template file!
// Pass desire service name in ```kebab-case``` like ```register-new-person``` in file.Name
func MakeNewServiceFile(file *assets.File) (err error) {
	file.FullName = file.Name + ".go"
	file.Extension = "go"

	var sn = strings.Title(file.Name)
	sn = strings.ReplaceAll(sn, "-", "")

	var tempName = struct {
		ServiceID        uint32
		ServiceUpperName string
		ServiceLowerName string
		IssueDate        int64
	}{
		ServiceID:        crc32.ChecksumIEEE([]byte(sn)), // hash sn for its ID
		ServiceUpperName: sn,
		ServiceLowerName: strings.ToLower(sn[0:1]) + sn[1:],
		IssueDate:        time.Now().Unix(),
	}

	var sf = new(bytes.Buffer)
	err = serviceFileTemplate.Execute(sf, tempName)
	if err != nil {
		return
	}

	file.Data, err = format.Source(sf.Bytes())
	// Indicate file had been changed
	file.State = assets.StateChanged

	return
}

var serviceFileTemplate = template.Must(template.New("serviceFileTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package services

import (
	"../../libgo/achaemenid"
	"../../libgo/http"
	"../../libgo/json"
)

var {{.ServiceLowerName}}Service = achaemenid.Service{
	ID:                {{.ServiceID}},
	URI:               "", // API services can set like "/apis?{{.ServiceID}}" but it is not efficient, find services by ID.
	Name:              "{{.ServiceUpperName}}",
	IssueDate:         {{.IssueDate}},
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,
	Description:       []string{
		"",
	},
	TAGS:        []string{""},
	SRPCHandler: {{.ServiceUpperName}}SRPC,
	HTTPHandler: {{.ServiceUpperName}}HTTP,
}

// {{.ServiceUpperName}}SRPC is sRPC handler of {{.ServiceUpperName}} service.
func {{.ServiceUpperName}}SRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	var req = &{{.ServiceLowerName}}Req{}
	st.ReqRes.Err = req.syllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		st.Connection.FailedPacketsReceived++
		// Attack??
		return
	}

	var res = &{{.ServiceLowerName}}Res{}
	res, st.ReqRes.Err = {{.ServiceLowerName}}(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		st.Connection.FailedServiceCall++
		// Attack??
		return
	}

	st.ReqRes.Payload = res.syllabEncoder(4)
}

// {{.ServiceUpperName}}HTTP is HTTP handler of {{.ServiceUpperName}} service.
func {{.ServiceUpperName}}HTTP(s *achaemenid.Server, st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &{{.ServiceLowerName}}Req{}
	st.ReqRes.Err = req.jsonDecoder(httpReq.Body)
	if st.ReqRes.Err != nil {
		st.Connection.FailedPacketsReceived++
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		httpRes.Body = []byte(st.ReqRes.Err.Error())
		// Attack??
		return
	}

	var res = &{{.ServiceLowerName}}Res{}
	res, st.ReqRes.Err = {{.ServiceLowerName}}(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		st.Connection.FailedServiceCall++
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		httpRes.Body = []byte(st.ReqRes.Err.Error())
		// Attack??
		return
	}

	httpRes.Body, st.ReqRes.Err = res.jsonEncoder()
	// st.ReqRes.Err make occur on just memory full!

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.SetValue(http.HeaderKeyContentType, "application/json")
}

type {{.ServiceLowerName}}Req struct {}

type {{.ServiceLowerName}}Res struct {}

func {{.ServiceLowerName}}(st *achaemenid.Stream, req *{{.ServiceLowerName}}Req) (res *{{.ServiceLowerName}}Res, err error) {
	// Validate data here due to service use internally!
	err = req.validator()
	if err != nil {
		return
	}

	return
}

func (req *{{.ServiceLowerName}}Req) validator() (err error) {
	return
}

func (req *{{.ServiceLowerName}}Req) syllabDecoder(buf []byte) (err error) {
	return
}

func (req *{{.ServiceLowerName}}Req) jsonDecoder(buf []byte) (err error) {
	// TODO::: Help to complete json generator package to have better performance!
	err = json.Unmarshal(buf, req)
	return
}

// offset add free space by given number at begging of return slice that almost just use in sRPC protocol! It can be 0!!
func (res *{{.ServiceLowerName}}Res) syllabEncoder(offset int) (buf []byte) {
	return
}

func (res *{{.ServiceLowerName}}Res) jsonEncoder() (buf []byte, err error) {
	// TODO::: Help to complete json generator package to have better performance!
	buf, err = json.Marshal(res)
	return
}
`))
