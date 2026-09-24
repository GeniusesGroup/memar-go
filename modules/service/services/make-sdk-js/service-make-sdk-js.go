/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"memar/process/log"
	file_p "memar/storage/file/protocol"
)

// MakeJSSDK will make sdk for logics with enough document.
func MakeJSSDK(file file_p.File) (jsSDK, jsonSDK file_p.File, err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", file.Data, parser.ParseComments)
	if err != nil {
		return
	}

	var js = jsSDKMaker{
		File:            file,
		ExpireInFavorOf: "null", // if exist, replace by actual service detail.
	}
	var json = map[string]jsonLocalize{}

	// find some data from service.Service variable
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
						err = js.makeJSReqResType(gd)
						if err != nil {
							return
						}
						js.Request = js.Generated.String()
						// remove unneeded last new line
						js.Request = js.Request[:len(js.Request)-1]
						js.Generated.Reset()
					} else if strings.HasSuffix(gd.Name.Name, "Res") {
						// TODO::: add nested type support
						err = js.makeJSReqResType(gd)
						if err != nil {
							return
						}
						js.Response = js.Generated.String()
						// remove unneeded last new line
						js.Response = js.Response[:len(js.Response)-1]
						js.Generated.Reset()
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
											js.ID = elt.Value.(*ast.BasicLit).Value
											continue
										case "IssueDate":
											js.IssueDate = elt.Value.(*ast.BasicLit).Value
											continue
										case "ExpiryDate":
											js.ExpiryDate = elt.Value.(*ast.BasicLit).Value
											continue
										case "ExpireInFavorOf":
											js.ExpireInFavorOf = elt.Value.(*ast.BasicLit).Value
											continue
										case "ExpireInFavorOfID":
											js.ExpireInFavorOfID = elt.Value.(*ast.BasicLit).Value
											continue
										case "Status":
											js.Status = elt.Value.(*ast.SelectorExpr).Sel.Name
											continue

										case "Authorization":
											for _, innerElt := range elt.Value.(*ast.CompositeLit).Elts {
												var innerElement = innerElt.(*ast.KeyValueExpr)
												var key = innerElement.Key.(*ast.Ident).Name
												switch key {
												case "ActionType":
													js.ActionType = string(file.Data()[innerElement.Value.Pos()-1 : innerElement.Value.End()-1])
												case "UserType":
													js.UserType = string(file.Data[innerElement.Value.Pos()-1 : innerElement.Value.End()-1])
												}
											}
											continue

										case "Name":
											var unquote, _ = strconv.Unquote(elt.Value.(*ast.CompositeLit).Elts[0].(*ast.KeyValueExpr).Value.(*ast.BasicLit).Value)
											js.Name = strings.ReplaceAll(unquote, " ", "")

											for _, elts := range elt.Value.(*ast.CompositeLit).Elts {
												var lang = elts.(*ast.KeyValueExpr).Key.(*ast.SelectorExpr).Sel.Name
												var jsonLoc = jsonLocalize{}
												jsonLoc.Name, _ = strconv.Unquote(elts.(*ast.KeyValueExpr).Value.(*ast.BasicLit).Value)
												json[lang] = jsonLoc
											}
											continue
										case "Description":
											js.Description, _ = strconv.Unquote(elt.Value.(*ast.CompositeLit).Elts[0].(*ast.KeyValueExpr).Value.(*ast.BasicLit).Value)

											for _, elts := range elt.Value.(*ast.CompositeLit).Elts {
												var lang = elts.(*ast.KeyValueExpr).Key.(*ast.SelectorExpr).Sel.Name
												var jsonLoc = json[lang]
												jsonLoc.Description, _ = strconv.Unquote(elts.(*ast.KeyValueExpr).Value.(*ast.BasicLit).Value)
											}
											continue
										case "TAGS":
											js.TAGS = string(file.Data[elt.Value.(*ast.CompositeLit).Lbrace : elt.Value.(*ast.CompositeLit).Rbrace-1])
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

	var jsBuf = new(bytes.Buffer)
	err = jsSDKTemplate.Execute(jsBuf, js)
	if err != nil {
		return
	}

	var jsonBuf = new(bytes.Buffer)
	err = jsonSDKTemplate.Execute(jsonBuf, json)
	if err != nil {
		return
	}

	jsSDK = &file.File{
		FullName:  file.Name + ".js",
		Name:      file.Name,
		Extension: "js",
		Data:      jsBuf.Bytes(),
	}
	jsonSDK = &file.File{
		FullName:  file.Name + ".json",
		Name:      file.Name,
		Extension: "json",
		Data:      jsonBuf.Bytes(),
	}

	return
}

type jsSDKMaker struct {
	File file_p.File

	Name string

	URN               string
	ID                string
	IssueDate         string
	ExpiryDate        string
	ExpireInFavorOf   string
	ExpireInFavorOfID string
	Status            string

	ActionType string
	UserType   string

	Description string
	TAGS        string

	Request  string
	Response string

	Generated bytes.Buffer
	// Field Options
	FieldName string
	Dash      bool
	OmitEmpty bool
	String    bool
	Tuple     bool
}

func (js *jsSDKMaker) makeJSReqResType(t *ast.TypeSpec) (err error) {
	switch structType := t.Type.(type) {
	default:
		// Just occur if bad file pass to generator!!
		// return
	case *ast.BasicLit:
		// TODO::: very simple type
	case *ast.StructType:
		for _, structField := range structType.Fields.List {
			js.FieldName = ""
			js.OmitEmpty = false
			js.String = false
			js.Tuple = false

			if structField.Tag != nil && js.checkFieldTag(structField.Tag.Value) {
				continue
			}

			for _, field := range structField.Names {
				if js.FieldName == "" {
					js.FieldName = field.Name
				}

				// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Data_structures
				switch fieldType := structField.Type.(type) {
				case *ast.FuncType, *ast.InterfaceType, *ast.ChanType:
					log.Warn(domainEnglish, "Request or Response include bad fields! Field name:"+js.FieldName)
				case *ast.ArrayType:
					if js.String {
						js.Generated.WriteString(` * @property {string} ` + js.FieldName + " " + string(js.File.Data[fieldType.Pos()-1:fieldType.End()-1]) + " that encode||decode as base64\n")
					} else {
						js.Generated.WriteString(` * @property {array} ` + js.FieldName + " " + string(js.File.Data[fieldType.Pos()-1:fieldType.End()-1]) + "\n")
					}
				case *ast.StructType:
					js.Generated.WriteString(` * @property {Map} ` + js.FieldName + " " + string(js.File.Data[fieldType.Pos()-1:fieldType.End()-1]) + "\n")
				case *ast.MapType:
					js.Generated.WriteString(` * @property {Map} ` + js.FieldName + " " + string(js.File.Data[fieldType.Pos()-1:fieldType.End()-1]) + "\n")
				case *ast.Ident:
					switch fieldType.Name {
					case "bool":
						js.Generated.WriteString(` * @property {boolean} ` + js.FieldName + "\n")
					case "byte", "uint8", "int8", "uint16", "int16", "uint32", "int32":
						js.Generated.WriteString(` * @property {number} ` + js.FieldName + " " + fieldType.Name + "\n")
					case "uint64", "int64", "int", "uint":
						js.Generated.WriteString(` * @property {bigint} ` + js.FieldName + " " + fieldType.Name + "\n")
					case "string":
						js.Generated.WriteString(` * @property {string} ` + js.FieldName + "\n")
					default:
						js.Generated.WriteString(` * @property {Map} ` + js.FieldName + " " + fieldType.Name + "\n")
					}
				default:
					js.Generated.WriteString(` * @property {Map} ` + js.FieldName + " " + string(js.File.Data[fieldType.Pos()-1:fieldType.End()-1]) + "\n")
				}
			}
		}
	}
	return
}

func (js *jsSDKMaker) checkFieldTag(tagValue string) (notInclude bool) {
	var structFieldTag = reflect.StructTag(tagValue[1 : len(tagValue)-1])
	var structFieldTagJSON = structFieldTag.Get("json")
	if structFieldTagJSON != "" {
		if structFieldTagJSON == "-" {
			return true
		}

		var structFieldTagJSONParts = strings.Split(structFieldTagJSON, ",")
		if structFieldTagJSONParts[0] != "" {
			js.FieldName = structFieldTagJSONParts[0]
		}

		if len(structFieldTagJSONParts) > 1 {
			structFieldTagJSONParts = structFieldTagJSONParts[1:]

			for i := 0; i < len(structFieldTagJSONParts); i++ {
				switch structFieldTagJSONParts[i] {
				case "omitempty":
					js.OmitEmpty = true
				case "string":
					js.String = true
				case "tuple":
					js.Tuple = true
				}
			}
		}
	}
	return
}

var jsSDKTemplate = template.Must(template.New("jsSDKTemplate").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */
// Auto-generated, edits will be overwritten

const {{.Name}}Service = {
    URN: "{{.URN}},
    ID: "{{.ID}}",
    Status: {{.Status}}, // https://en.wikipedia.org/wiki/Software_release_life_cycle
    IssueDate: {{.IssueDate}},
    ExpireDate: {{.ExpiryDate}},
    ExpireInFavorOf: {{.ExpireInFavorOf}},

	Authorization: {
		ActionType: {{.ActionType}},
		UserType:   {{.UserType}},
	},

	Detail: {
		Name:        "LocaleText[0]",
		Description: "LocaleText[1]",
		TAGS:        ["LocaleText[2]"],
	},
}
Application.RegisterService({{.Name}}Service)

/**
 * @typedef {Object} {{.Name}}Req is the request structure of {{.Name}}()
{{.Request}}
 */

/**
 * @typedef {Object} {{.Name}}Res is the response structure of {{.Name}}()
{{.Response}}
 */

 /**
 * {{.Description}}
 * @param {{"{"}}{{.Name}}Req} req
 * @returns {{"{"}}{{.Name}}Res}
 * @throws {GitiError} related error
 */
{{.Name}}Service.Do = async function (req) {
    // TODO::: check network availability and other requirements like WASM (e.g. for encoding and decoding)
	// and decide to do request by best one e.g. http, grpc, srpc, ...
}

/**
 * {{.Description}}
 * @param {{"{"}}{{.Name}}Req} req
 * @returns {{"{"}}{{.Name}}Res}
 * @throws {GitiError} related error
 */
 {{.Name}}Service.DoHTTP = async function (req) {
    // Check cache if service responses cache-able before send to apis server!
    if (cacheDB && useCacheDB) {
        const transaction = cacheDB.transaction({{.Name}}DBName, "readonly")
        const objectStore = transaction.objectStore({{.Name}}DBName)
        const IDBRequest = objectStore.get(req.ID)
        const result = await getIndexDBResultAsync(IDBRequest)
        if (result && result.CacheTime + {{.Name}}DBCacheTTL > Date.now()) {
            return result
        }
    }

    // TODO::: Validate req before send to apis server!

    const httpReq = new Request('/m?{{.ID}}', {
        method: Method_POST,
        // compress: true,
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(req)
    })
	try {
        var httpRes = await fetch(httpReq)
    } catch (err) {
        throw GitiError.ConvertToGitiError(err)
    }

    /**
     * @type {{"{"}}{{.Name}}Res}
     */
    let res

	switch (httpRes.status) {
		case 200:
			const contentType = httpRes.headers.get('content-type')
			switch (contentType) {
				case 'application/json':
					try {
						res = await httpRes.json()
                        // res.ID = req.ID
					} catch (err) {
						throw GitiError.ConvertToGitiError(err)
					}
					break
				// case 'text/plain':
				//     try {
				//         return await httpRes.text()
				//     } catch (err) {
				//         throw GitiError.ConvertToGitiError(err)
				//     }
				default:
					throw Application.GetErrorByURN("domain/http.wg.ietf.org; type=error; name=unsupported-media-type")
			}
			break
		case 201:
			return null
		case 400:
		case 500:
		default:
            throw httpRes.GetError()
	}

    if (cacheDB) {
        res.CacheTime = Date.now()
        const transaction = cacheDB.transaction({{.Name}}DBName, "readwrite")
        const objectStore = transaction.objectStore({{.Name}}DBName)
        objectStore.put(res)
    }
    return res
}
`))

type jsonLocalize struct {
	Name        string
	Description string
}

var jsonSDKTemplate = template.Must(template.New("jsonSDKTemplate").Parse(`{
	{{ range $key, $value := . }}
	"{{$key}}": [
        "{{$value.Name}}",
        "{{$value.Description}}",
		""
    ],
	{{ end }}
}`))
