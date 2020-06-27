/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unsafe"

	"../assets"
)

// MakeJSSDK will make sdk for logics with enough document.
func MakeJSSDK(file *assets.File) (jsSDK *assets.File, err error) {
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
						// TODO::: add nested type support
						service.Request, err = makeJSReqResType(gd)
						if err != nil {
							return
						}
					} else if strings.HasSuffix(gd.Name.Name, "Res") {
						// TODO::: add nested type support
						service.Response, err = makeJSReqResType(gd)
						if err != nil {
							return
						}
					}
				case *ast.ValueSpec:
					// name of const||var: gd.Names[0].Name
					// type of const||var (if any package name): gd.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name)
					// type of const||var: gd.Type.(*ast.SelectorExpr).Sel.Name)
					if len(gd.Values) == 0 {
						continue
					}
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
											service.TAGS = file.DataString[elt.Value.(*ast.CompositeLit).Lbrace : elt.Value.(*ast.CompositeLit).Rbrace-1]
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

	var sf = new(bytes.Buffer)
	err = jsSDKTemplate.Execute(sf, service)
	if err != nil {
		return
	}

	jsSDK = &assets.File{
		FullName:  file.Name + ".js",
		Name:      file.Name,
		Extension: "js",
		State:     assets.StateChanged,
		Data:      sf.Bytes(),
	}

	return
}

func makeJSReqResType(t *ast.TypeSpec) (string, error) {
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
				data = append(data, `	"`+in+`": [], // [`+number+"]"+aType+"\n"...)
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
			case "bool":
				data = append(data, `	"`+in+`": true, // boolean`+"\n"...)
			case "byte", "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64":
				data = append(data, `	"`+in+`": 0, // `+t.Name+"\n"...)
			case "string":
				data = append(data, `	"`+in+`": "", // string`+"\n"...)
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	// remove unneeded last new line
	if len(data) != 0 {
		data = data[:len(data)-1]
	}
	return *(*string)(unsafe.Pointer(&data)), nil
}

var jsSDKTemplate = template.Must(template.New("jsSDKTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */
// Auto-generated, edits will be overwritten

/*
Service Details :
	- Status : {{.Status}}  >> https://en.wikipedia.org/wiki/Software_release_life_cycle
	- IssueDate : {{.IssueDate}}
	- ExpireDate : {{.ExpiryDate}}
	- ExpireInFavorOf :  {{.ExpireInFavorOf}}
	- ExpireInFavorOfID : {{.ExpireInFavorOfID}}
	- TAGS : {{.TAGS}}

Usage :
	// {{.Name}}Req is the request structure of {{.Name}}()
	const {{.Name}}Req = {
		{{.Request}}
	}
	// {{.Name}}Res is the response structure of {{.Name}}()
	const {{.Name}}Res = {
    	{{.Response}}
	}
	{{.Name}}({{.Name}}Req)
		.then(res => {
			// Handle response
			console.log(res)
		})
		.catch(err => {
			// Handle error situation here
			console.log(err)
		})

Also you can use "async function (){ try{await func()}catch (err){} }" instead "func(req).then(res).catch(err)"
*/

// {{.Name}} {{.Description}}
async function {{.Name}}(req) {
    // TODO::: First validate req before send to apis server!

    // TODO::: Check Quic protocol availability!

    const request = new Request('/apis?{{.ID}}', {
        method: "POST",
        // compress: true,
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(req)
    })

    try {
        let res = await fetch(request)

        switch (res.status) {
            case 200:
                const contentType = res.headers.get('content-type')
                switch (contentType) {
                    case 'application/json':
                        try {
                            return await res.json()
                        } catch (err) {
                            throw err
                        }
                    // case 'text/plain':
                    //     try {
                    //         return await res.text()
                    //     } catch (err) {
                    //         throw err
                    //     }
                    default:
                        throw new TypeError("Oops, we haven't got valid data type in response!")
                }
            case 201:
                return null
            case 400:
            case 500:
            default:
                // Almost not reachable code!
                throw res.text()
        }
    } catch (err) {
		// TODO::: new more check here for error type!
        // TODO::: Toast to GUI about no network connectivity!
        throw err
    }
}
`))
