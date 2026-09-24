/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"
	"time"

	errs "memar/modules/service/errors"
)

// MakeServiceGO use to make new service template file!
func MakeServiceGO(req MakeServiceGOReq) (res MakeServiceGORes, err error) {
	if req.ServiceName == "" {
		err = &errs.ServiceNameIsNotValid
		return
	}

	var serviceName = strings.Title(req.ServiceName)
	var tempName = struct {
		DomainName         string
		ServiceName        string
		ServiceEnglishName string
		ServiceUpperName   string
		ServiceLowerName   string
		IssueDate          time.Time
	}{
		DomainName:         req.DomainName,
		ServiceName:        req.ServiceName,
		ServiceEnglishName: strings.Title(strings.ReplaceAll(serviceName, "-", " ")),
		ServiceUpperName:   strings.ReplaceAll(serviceName, "-", ""),
		IssueDate:          time.Now(),
	}
	tempName.ServiceLowerName = strings.ToLower(tempName.ServiceUpperName[0:1]) + tempName.ServiceUpperName[1:]

	var sf bytes.Buffer
	err = goServiceTemplate.Execute(&sf, tempName)
	if err != nil {
		return
	}
	res.Service, err = format.Source(sf.Bytes())

	sf.Reset()
	err = goServiceSDKTemplate.Execute(&sf, tempName)
	if err != nil {
		return
	}
	res.SDK, err = format.Source(sf.Bytes())
	return
}

var goServiceTemplate = template.Must(template.New("goServiceTemplate").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"memar/data"
	"memar/libgo/net/http"
	"memar/libgo/codec/json"
	"memar/protocol"
)

func (ser *{{.ServiceLowerName}}Service) ServeSRPC(sk net_p.Socket, rawReq codec_p.Codec) (rawRes codec_p.Codec, err error_p.Error) {
	var req = &{{.ServiceLowerName}}Req{}
	err = req.FromSyllab(srpc.GetPayload(st.IncomePayload))
	if err != nil {
		return
	}

	var res *{{.ServiceUpperName}}Response
	res, err = {{.ServiceLowerName}}(st, req)
	// Check if any error occur in business logic
	if err != nil {
		return
	}

	st.OutcomePayload = make([]byte, srpc.MinLength+res.Syllab_Length())
	res.ToSyllab(srpc.GetPayload(st.OutcomePayload))
}

func (ser *{{.ServiceLowerName}}Service) ServeHTTP(sk net_p.Socket, httpReq http_p.Request, httpRes http_p.Response) (err error_p.Error) {
	var req = &{{.ServiceLowerName}}Req{}
	err = req.FromJSON(httpReq.Body().Marshal())
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *{{.ServiceUpperName}}Response
	res, err = {{.ServiceLowerName}}(st, req)
	// Check if any error occur in bussiness logic
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.SetBody(json.NewCodec(res))
}

func {{.ServiceLowerName}}(sk net_p.Socket, req {{.ServiceLowerName}}Request) (res *{{.ServiceUpperName}}Response, err error_p.Error) {
	// Validate data here due to service use internally by other services!
	err = req.validator()
	if err != nil {
		return
	}

	// TODO::: Service logic: Authenticate & Authorizing request first by service policy.

	// TODO::: Remove me and write some code e.g. save a record and related indexes if you need!

	res = &{{.ServiceUpperName}}Res{}
	return
}

func (req *{{.ServiceLowerName}}Req) validator() (err error_p.Error) {
	return
}
`))

var goServiceSDKTemplate = template.Must(template.New("goServiceSDKTemplate").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"memar/protocol"
	"memar/process/service"
	"memar/srpc"
	"memar/codec/syllab"
)

var {{.ServiceUpperName}}Service = {{.ServiceLowerName}}{
	Service: service.New("domain/{{.DomainName}}.service; name={{.ServiceName}}", "", datatype_p.LifeCycle_PreAlpha, {{.IssueDate}}).
		SetDetail(lang_p.LanguageEnglish, "{{.ServiceEnglishName}}",
		"",
		[]string{}).
		SetAuthorization(operation_p.ActionType_None, user_p.Type_All).Expired(0, ""),
}

type {{.ServiceLowerName}}Service struct {
	service.Service
}

func (ser *{{.ServiceLowerName}}Service) DoSRPC(req {{.ServiceLowerName}}Request) (res {{.ServiceLowerName}}Response, err error_p.Error) {
	return
}

/*
	Service request and response shape
*/

type {{.ServiceLowerName}}Request interface {
	validator() (err error_p.Error)
}

type {{.ServiceLowerName}}Response interface {
}

/*
	Service Request
*/

type {{.ServiceUpperName}}Request struct {

}

func (req *{{.ServiceUpperName}}Request) validator() (err error_p.Error) {
	return
}

// methods to implements {{.ServiceLowerName}}Request interface
// func (req *{{.ServiceUpperName}}Request) ObjectID() [32]byte { return req.ObjectID }

// methods to implements protocol.Syllab interface
func (req *{{.ServiceUpperName}}Request) CheckSyllab(buf []byte) (err error_p.Error) {
	if len(buf) < int(req.Syllab_StackLength()) {
		err = syllab_errs.ShortArrayDecode
	}
	return
}
func (req *{{.ServiceUpperName}}Request) FromSyllab(buf []byte, stackIndex uint32) {
	// TODO::: Use syllab generator to generate needed codes
}
func (req *{{.ServiceUpperName}}Request) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32)
	// TODO::: Use syllab generator to generate needed codes
	return heapIndex
}
func (req *{{.ServiceUpperName}}Request) Syllab_Length() uint64          { return uint64(req.Syllab_StackLength() + req.Syllab_HeapLength()) }
func (req *{{.ServiceUpperName}}Request) Syllab_StackLength() uint32     { return 0 }
func (req *{{.ServiceUpperName}}Request) Syllab_HeapLength() (ln uint32) {
	// TODO::: Use syllab generator to generate needed codes
	return
}

// methods to implements memar/json/protocol.JSON interface
func (req *{{.ServiceUpperName}}Request) FromJSON(buf []byte) (remaining []byte, err error_p.Error) {
	// TODO::: Use json generator to have better performance!
	err = json.Unmarshal(buf, req)
	return
}
func (req *{{.ServiceUpperName}}Request) ToJSON(buf []byte) (added []byte, err error_p.Error) {
	var encoder = json.Encoder{Buf: buf}
	// TODO::: generate by generator!
	return encoder.Buf
}
func (req *{{.ServiceUpperName}}Request) JSON_Length() (ln int) {
	return
}

type {{.ServiceLowerName}}RequestSyllab []byte

// methods to implements {{.ServiceLowerName}}Request interface
// func (req {{.ServiceLowerName}}RequestSyllab) ObjectID() (objectID [32]byte) { copy(objectID[:], req[0:]); return }

// methods to implements protocol.Syllab interface
func (req {{.ServiceLowerName}}RequestSyllab) CheckSyllab(buf []byte) (err error_p.Error) {
	if len(req) < int(req.Syllab_StackLength()) {
		err = syllab_errs.ShortArrayDecode
	}
	return
}
func (req {{.ServiceLowerName}}RequestSyllab) FromSyllab(buf []byte, stackIndex uint32) {
	*req = buf[stackIndex : stackIndex+req.Syllab_StackLength()]
}
func (req {{.ServiceLowerName}}RequestSyllab) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(buf[stackIndex:], req)
	return heapIndex
}
func (req {{.ServiceLowerName}}RequestSyllab) Syllab_Length() uint64          { return uint64(req.Syllab_StackLength() + req.Syllab_HeapLength()) }
func (req {{.ServiceLowerName}}RequestSyllab) Syllab_StackLength() uint32     { return 0 }
func (req {{.ServiceLowerName}}RequestSyllab) Syllab_HeapLength() (ln uint32) { return }

/*
	Service Response
*/

type {{.ServiceUpperName}}Response struct {}

// methods to implements {{.ServiceLowerName}}Response interface
// func (res *{{.ServiceUpperName}}Response) ObjectID() [32]byte { return res.ObjectID }

// methods to implements protocol.Syllab interface
func (res *{{.ServiceUpperName}}Response) CheckSyllab(buf []byte) (err error_p.Error) {
	if len(buf) < int(res.Syllab_StackLength()) {
		err = syllab_errs.ShortArrayDecode
	}
	return
}
func (res *{{.ServiceUpperName}}Response) FromSyllab(buf []byte, stackIndex uint32) {
	// TODO::: Use syllab generator to generate needed codes
}
func (res *{{.ServiceUpperName}}Response) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32)
	// TODO::: Use syllab generator to generate needed codes
	return heapIndex
}
func (res *{{.ServiceUpperName}}Response) Syllab_Length() uint64          { return uint64(res.Syllab_StackLength() + res.Syllab_HeapLength()) }
func (res *{{.ServiceUpperName}}Response) Syllab_StackLength() uint32     { return 0 }
func (res *{{.ServiceUpperName}}Response) Syllab_HeapLength() (ln uint32) {
	// TODO::: Use syllab generator to generate needed codes
	return
}

// methods to implements memar/json/protocol.JSON interface
func (res *{{.ServiceUpperName}}Response) FromJSON(buf []byte) (remaining []byte, err error_p.Error) {
	// TODO::: Use json generator to have better performance!
	err = json.Unmarshal(buf, res)
	return
}
func (res *{{.ServiceUpperName}}Response) ToJSON(buf []byte) (added []byte, err error_p.Error) {
	var encoder = json.Encoder{Buf: buf}

	// TODO::: Use json generator to have better performance!
	buf, _ = json.Marshal(res)
	return
}
func (res *{{.ServiceUpperName}}Response) JSON_Length() (ln int) {
	return
}

type {{.ServiceLowerName}}ResponseSyllab []byte

// methods to implements {{.ServiceLowerName}}Response interface
// func (res {{.ServiceLowerName}}ResponseSyllab) ObjectID() (objectID [32]byte) { copy(objectID[:], res[0:]); return }

// methods to implements protocol.Syllab interface
func (res {{.ServiceLowerName}}ResponseSyllab) CheckSyllab(buf []byte) (err error_p.Error) {
	if len(res) < int(res.Syllab_StackLength()) {
		err = syllab_errs.ShortArrayDecode
	}
	return
}
func (res {{.ServiceLowerName}}ResponseSyllab) FromSyllab(buf []byte, stackIndex uint32) {
	*res = buf[stackIndex : stackIndex+res.Syllab_StackLength()]
}
func (res {{.ServiceLowerName}}ResponseSyllab) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(buf[stackIndex:], res)
	return heapIndex
}
func (res {{.ServiceLowerName}}ResponseSyllab) Syllab_Length() uint64          { return uint64(res.Syllab_StackLength() + res.Syllab_HeapLength()) }
func (res {{.ServiceLowerName}}ResponseSyllab) Syllab_StackLength() uint32     { return 0 }
func (res {{.ServiceLowerName}}ResponseSyllab) Syllab_HeapLength() (ln uint32) { return }
`))
