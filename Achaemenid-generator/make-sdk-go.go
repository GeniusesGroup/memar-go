/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unsafe"

	"../assets"
)

// Errors use in SDK making proccess.
var (
	ErrInvalidServiceFile      = errors.New("Given service file doesn't match with standard service structure")
	ErrTypeIncludeIllegalChild = errors.New("Request||Response type may include function, interface, int, uint, ... type that can't encode||decode")
)

// MakeGoSDK will make sdk for logics with enough document.
func MakeGoSDK(file *assets.File) (goSDK *assets.File, err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", file.DataString, parser.ParseComments)
	if err != nil {
		return
	}

	var service = struct {
		ID                uint64
		Name              string
		IssueDate         string
		ExpiryDate        string
		ExpireInFavorOf   string
		ExpireInFavorOfID string
		Status            string
		Description       string
		TAGS              string
		Request           string
		Response          string
	}{}

	// find some data from achaemenid.Service variable
	for _, decl := range fileParsed.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range d.Specs {
				switch gd := gDecl.(type) {
				case *ast.ImportSpec:
					// Don't need import
				case *ast.TypeSpec:
					if strings.HasSuffix(gd.Name.Name, "Req") {
						service.Request, err = makeGOReqResType(gd)
						if err != nil {
							return
						}
					} else if strings.HasSuffix(gd.Name.Name, "Res") {
						service.Response, err = makeGOReqResType(gd)
						if err != nil {
							return
						}
					}
				case *ast.ValueSpec:
					// name of const||var: gd.Names[0].Name
					// type of const||var (if any package name): gd.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name)
					// type of const||var: gd.Type.(*ast.SelectorExpr).Sel.Name)
					switch val := gd.Values[0].(type) {
					case *ast.Ident:
						// Don't need ident here
					case *ast.BasicLit:
						// Don't need basic lit here
					case *ast.CompositeLit:
						if val.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name == "achaemenid" && val.Type.(*ast.SelectorExpr).Sel.Name == "Service" {
							for _, element := range val.Elts {
								switch elt := element.(type) {
								case *ast.KeyValueExpr:
									switch key := elt.Key.(type) {
									case *ast.Ident:
										// fmt.Printf(">> %v\n", key.Name)
										switch key.Name {
										case "ID":
											// fmt.Printf("ID %v\n", elt.Value.(*ast.BasicLit).Value)
											service.ID, err = strconv.ParseUint(elt.Value.(*ast.BasicLit).Value, 10, 32)
											if err != nil {
												return nil, ErrInvalidServiceFile
											}
											continue
										case "Name":
											// fmt.Printf("Name %v\n", elt.Value.(*ast.BasicLit).Value)
											service.Name, _ = strconv.Unquote(elt.Value.(*ast.BasicLit).Value)
											continue
										case "IssueDate":
											var ut, _ = strconv.ParseInt(elt.Value.(*ast.BasicLit).Value, 10, 64)
											var t = time.Unix(ut, 0)
											service.IssueDate = t.Format("02/01/2006 15:04:05 MST")
											continue
										case "ExpiryDate":
											var ut, _ = strconv.ParseInt(elt.Value.(*ast.BasicLit).Value, 10, 64)
											var t = time.Unix(ut, 0)
											service.ExpiryDate = t.Format("02/01/2006 15:04:05 MST")
											continue
										case "ExpireInFavorOf":
											service.ExpireInFavorOf = elt.Value.(*ast.BasicLit).Value
											continue
										case "ExpireInFavorOfID":
											service.ExpireInFavorOfID = elt.Value.(*ast.BasicLit).Value
											continue
										case "Status":
											service.Status = elt.Value.(*ast.SelectorExpr).Sel.Name
											continue
										case "Description":
											// fmt.Printf("Description %v\n", elt.Value.(*ast.CompositeLit).Elts[0].(*ast.BasicLit).Value)
											service.Description, _ = strconv.Unquote(elt.Value.(*ast.CompositeLit).Elts[0].(*ast.BasicLit).Value)
											continue
										case "TAGS":
											service.TAGS = file.DataString[elt.Value.(*ast.CompositeLit).Lbrace:elt.Value.(*ast.CompositeLit).Rbrace-1]
										}
									}
								}
							}
						}
					}
				}
			}
		case *ast.FuncDecl:
			// TODO::: It is better to find service function and indicate Req&&Res type!! but lazy for now!!
		}
	}

	var buff = new(bytes.Buffer)
	err = goSDKTemplate.Execute(buff, service)
	if err != nil {
		return
	}

	goSDK = &assets.File{
		FullName:  file.Name + ".go",
		Name:      file.Name,
		Extension: "go",
		State:     assets.StateChanged,
	}
	goSDK.Data, err = format.Source(buff.Bytes())

	return
}

// TODO::: add nested type support
func makeGOReqResType(t *ast.TypeSpec) (string, error) {
	var data = make([]byte, 0, 1024)
	var in string
	for _, f := range t.Type.(*ast.StructType).Fields.List {
		in = f.Names[0].Name
		switch t := f.Type.(type) {
		case *ast.ArrayType:
			var number string = ""
			// Check array is slice?
			if t.Len == nil {
				// type is slice nothing to do!
			} else {
				number = t.Len.(*ast.BasicLit).Value
				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}
			}
			var aType = t.Elt.(*ast.Ident).Name
			switch aType {
			case "int", "uint":
				return "", ErrTypeIncludeIllegalChild
			case "bool", "byte", "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "string":
				data = append(data, "	"+in+" ["+number+"]"+aType+"\n"...)
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.StructType:
			// TODO:::
		case *ast.FuncType:
			return "", ErrTypeIncludeIllegalChild
		case *ast.InterfaceType:
			return "", ErrTypeIncludeIllegalChild
		case *ast.MapType:
			// TODO:::
		case *ast.ChanType:
			return "", ErrTypeIncludeIllegalChild
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return "", ErrTypeIncludeIllegalChild
			case "bool", "byte", "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "string":
				data = append(data, "	"+in+" "+t.Name+"\n"...)
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	// remove unneeded last new line
	data = data[:len(data)-1]
	return *(*string)(unsafe.Pointer(&data)), nil
}

var goSDKTemplate = template.Must(template.New("goSDKTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */
// Auto-generated, edits will be overwritten

package sdk

import (
	"../libgo/achaemenid"
)

/*
Service Details:
	- Status : {{.Status}}  >> https://en.wikipedia.org/wiki/Software_release_life_cycle
	- IssueDate : {{.IssueDate}}
	- ExpireDate : {{.ExpiryDate}}
	- ExpireInFavorOf :  {{.ExpireInFavorOf}}
	- ExpireInFavorOfID : {{.ExpireInFavorOfID}}
	- TAGS : {{.TAGS}}

Usage :

*/

// {{.Name}}Req is the request structure of {{.Name}}()
type {{.Name}}Req struct {
{{.Request}}
}

// {{.Name}}Res is the response structure of {{.Name}}()
type {{.Name}}Res struct {
{{.Response}}
}

// {{.Name}} {{.Description}}
func {{.Name}}(s *achaemenid.Server, req *{{.Name}}Req) (res *{{.Name}}Res, err error) {
	var conn *achaemenid.Connection
	// TODO::: find connection or make it
	
	// Check if no connection exist to use!
	if conn == nil {
		return nil, err
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)

	// Set ServiceID
	reqStream.ServiceID = {{.ID}}

	reqStream.Payload = req.syllabEncoder()
	err = achaemenid.SrpcOutcomeRequestHandler(s, reqStream)
	if err == nil {
		return nil, err
	}

	res = &{{.Name}}Res{}
	err = res.syllabDecoder(resStream.Payload[4:])

	return
}

func (req *{{.Name}}Req) validator() (err error) {
	return
}

func (req *{{.Name}}Req) syllabEncoder() (buf []byte) {
	return
}

func (res *{{.Name}}Res) syllabDecoder(buf []byte) (err error) {
	return
}
`))
