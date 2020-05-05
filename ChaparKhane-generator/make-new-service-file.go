/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"
	"time"
)

// MakeNewServiceFileReq is request structure of MakeNewServiceFile()
type MakeNewServiceFileReq struct {
	ServiceName string
}

// MakeNewServiceFileRes is response structure of MakeNewServiceFile()
type MakeNewServiceFileRes struct {
	ServiceFileName string
	ServiceFile     []byte
}

// MakeNewServiceFile use to make new service template file!
func MakeNewServiceFile(req *MakeNewServiceFileReq) (res *MakeNewServiceFileRes, err error) {
	res = &MakeNewServiceFileRes{
		ServiceFileName: req.ServiceName + ".go",
	}

	req.ServiceName = strings.Title(req.ServiceName)
	req.ServiceName = strings.ReplaceAll(req.ServiceName, "-", "")

	var tempName = struct {
		ServiceID        uint32
		ServiceUpperName string
		ServiceLowerName string
		IssueDate        int64
	}{
		ServiceID:        0, // hash req.ServiceName for its ID
		ServiceUpperName: req.ServiceName,
		ServiceLowerName: strings.ToLower(req.ServiceName[0:1]) + req.ServiceName[1:],
		IssueDate:        time.Now().Unix(),
	}

	var sf = new(bytes.Buffer)
	err = serviceFileTemplate.Execute(sf, tempName)
	if err != nil {
		return nil, err
	}

	res.ServiceFile, err = format.Source(sf.Bytes())
	if err != nil {
		return nil, err
	}

	return res, nil
}

var serviceFileTemplate = template.Must(template.New("serviceFileTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../../libgo/ChaparKhane"

var {{.ServiceLowerName}}Service = chaparkhane.Service{
	ID:              {{.ServiceID}},
	Name:            "{{.ServiceUpperName}}",
	IssueDate:       {{.IssueDate}},
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         {{.ServiceUpperName}},
	Description:     []string{
		"",
	},
	TAGS: []string{""},
}

// {{.ServiceUpperName}} will
func {{.ServiceUpperName}}(s *chaparkhane.Server, st *chaparkhane.Stream) {
	// Delete comments and write code here if you want directly work on input stream!
	// Don't delete these three comment line, If you want ChaparKhane-generator make||update code for you!
	// StreamProtocol::: sRPC
}

type {{.ServiceLowerName}}Req struct {}

type {{.ServiceLowerName}}Res struct {}

func {{.ServiceLowerName}}(st *chaparkhane.Stream, req *{{.ServiceLowerName}}Req) (res *{{.ServiceLowerName}}Res, err error) {
	return res, nil
}

func (req *{{.ServiceLowerName}}Req) validator() error {
	return nil
}

func (req *{{.ServiceLowerName}}Req) syllabDecoder(buf []byte) error {
	return nil
}

func (res *{{.ServiceLowerName}}Res) syllabEncoder(buf []byte) error {
	return nil
}
`))
