/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	parser "../ChaparKhane-parser"
)

// MakeGoSDKReq is request structure of MakeGoSDK()
type MakeGoSDKReq struct {
	Repo *parser.Repository
}

// MakeGoSDKRes is response structure of MakeGoSDK()
type MakeGoSDKRes struct {
	Repo *parser.Repository
}

// MakeGoSDK will make sdk for logics with enough document.
func MakeGoSDK(req *MakeGoSDKReq) (res *MakeGoSDKRes, err error) {
	var goSDKRepo = parser.NewRepository()
	goSDKRepo.Name = "sdk-go"

	for _, function := range req.Repo.Functions {
		var (
			buffer = new(bytes.Buffer)
			obj    parser.File
		)

		if err = goSDKTemplate.Execute(buffer, function); err != nil {
			return nil, err
		}

		obj.Name = function.File.Name
		obj.Data, err = format.Source(buffer.Bytes())
		if err != nil {
			return nil, err
		}

		goSDKRepo.Files[obj.Name] = &obj
	}

	res = &MakeGoSDKRes{
		Repo: goSDKRepo,
	}

	return res, nil
}

var goSDKTemplate = template.Must(template.New("goSDKTemplate").Parse(`
package persiadb

import (
	chaparkhane "./ChaparKhane"
)

// {{.Parameter.Type}} : The request structure of "{{.Name}}()"
type {{.Parameter.Type}} struct {
	{{- range .Parameter.InnerType}}
		{{.Name}} {{.Type}}
	{{- end}}
}

// {{.Result.Type}} : The response structure of "{{.Name}}()"
type {{.Result.Type}} struct {
	{{- range .Result.InnerType}}
		{{.Name}} {{.Type}}
	{{- end}}
}

// {{.Name}} :
func {{.Name}}(req *{{.Parameter.Type}}) (res *{{.Result.Type}}, err error) {
	return nil, nil
}
`))
