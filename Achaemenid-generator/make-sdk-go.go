/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"text/template"

	"../assets"
)

// MakeGoSDK will make sdk for logics with enough document.
func MakeGoSDK(folder *assets.Folder) (err error) {
	var goSDKRepo = assets.NewFolder("sdk-go")
	goSDKRepo.Status = assets.StateChanged
	folder.SetDependency(goSDKRepo)

	// for _, function := range req.Repo.Functions {
	// 	var (
	// 		buffer = new(bytes.Buffer)
	// 		obj    parser.File
	// 	)

	// 	if err = goSDKTemplate.Execute(buffer, function); err != nil {
	// 		return err
	// 	}

	// 	obj.Name = function.File.Name
	// 	obj.Data, err = format.Source(buffer.Bytes())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	goSDKRepo.Files[obj.Name] = &obj
	// }

	return nil
}

var goSDKTemplate = template.Must(template.New("goSDKTemplate").Parse(`
package sdk

import (
	"../libgo/achaemenid"
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
