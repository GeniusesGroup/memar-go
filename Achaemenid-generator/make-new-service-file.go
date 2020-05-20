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
	file.Status = assets.StateChanged

	return
}

var serviceFileTemplate = template.Must(template.New("serviceFileTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package services

import "../../libgo/achaemenid"

var {{.ServiceLowerName}}Service = achaemenid.Service{
	ID:                {{.ServiceID}},
	Name:              "{{.ServiceUpperName}}",
	IssueDate:         {{.IssueDate}},
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,
	Handler:           {{.ServiceUpperName}},
	Description:       []string{
		"",
	},
	TAGS: []string{""},
}

// {{.ServiceUpperName}} will
func {{.ServiceUpperName}}(s *achaemenid.Server, st *achaemenid.Stream) {
	// Delete comments and write code here if you want directly work on input stream!
	// Don't delete these three comment line, If you want Achaemenid-generator make||update code for you!
	// StreamProtocol::: sRPC
}

type {{.ServiceLowerName}}Req struct {}

type {{.ServiceLowerName}}Res struct {}

func {{.ServiceLowerName}}(st *achaemenid.Stream, req *{{.ServiceLowerName}}Req) (res *{{.ServiceLowerName}}Res, err error) {
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

func (res *{{.ServiceLowerName}}Res) syllabEncoder(buf []byte) (err error) {
	return
}
`))
