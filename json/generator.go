/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"../assets"
	"../convert"
	er "../error"
	"../log"
)

/*
Before pass file to safe||unsafe function, dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!
for just decoder method, jsonlen() can omit!

func ({{DesireName}} *{{DesireType}}) jsonDecoder(buf []byte) (err *er.Error) {
	return
}

func ({{DesireName}} *{{DesireType}}) jsonEncoder() (buf []byte) {
	return
}

func ({{DesireName}} *{{DesireType}}) jsonLen() (ln int) {
	return
}

Standards by https://www.json.org/json-en.html
*/

// GenerationOptions indicate jsonMaker behavior!
type GenerationOptions struct {
	Minifed           bool // without any space or new line, false for stylish version, true for better performance!
	Strict            bool // If a key not exist return error!
	UnSafe            bool // All decoded data will unsafe pointer to given buffer!
	AllowNoDefinedKey bool // allow not defined key in json encoded string to decode it! true return related error.
	NilMapAsNil       bool // false(empty map) >> "map":{},	true >> "map":nil,
	ForceUpdate       bool // true means delete exiting codes and update encoders && decoders codes anyway!
}

// CompleteMethods use to update given go files and complete json encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func CompleteMethods(file *assets.File, gos *GenerationOptions) (err *er.Error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	var goErr error
	fileParsed, goErr = parser.ParseFile(fileSet, "", file.Data, parser.ParseComments)
	if goErr != nil {
		log.Warn("JSON generator error:", goErr)
		return
	}

	var fileReplaces = make([]assets.ReplaceReq, 0, 4)
	var jm = jsonMaker{
		Options: gos,
		Types:   map[string]*ast.TypeSpec{},
	}

	// find jsonDecoder || jsonEncoder method
	for _, decl := range fileParsed.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range d.Specs {
				switch gd := gDecl.(type) {
				case *ast.TypeSpec:
					jm.Types[gd.Name.Name] = gd
				}
			}
		case *ast.FuncDecl:
			if d.Recv != nil {
				if jm.RN != d.Recv.List[0].Names[0].Name {
					jm.reset()
					jm.RN = d.Recv.List[0].Names[0].Name
					jm.FRN = d.Recv.List[0].Names[0].Name + "."
					jm.RTN = d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name

					err = jm.make()
					if err != nil {
						continue
					}

					jm.Encoder.WriteString("	return encoder.Buf\n")

					if jm.Options.Strict {
						jm.Decoder.WriteString("		default:\n" +
							"			err = decoder.NotFoundKeyStrict()\n		}\n\n" +
							"		if len(decoder.Buf) < 3 {\n			return\n		}\n	}\n" +
							"	return\n")
					} else {
						jm.Decoder.WriteString("		default:\n" +
							"			err = decoder.NotFoundKey()\n		}\n\n" +
							"		if len(decoder.Buf) < 3 {\n			return\n		}\n	}\n" +
							"	return\n")
					}

					jm.TemplateSize-- // due to one unneeded leading comma!
					jm.Len.WriteString("\n ln += " + strconv.FormatUint(uint64(jm.TemplateSize), 10) + "\n	return\n")
				}

				// Just needed methods!
				if d.Name.Name == "jsonDecoder" || d.Name.Name == "JSONDecoder" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  jm.Decoder.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "jsonEncoder" || d.Name.Name == "JSONEncoder" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  jm.Encoder.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "jsonLen" || d.Name.Name == "JSONLen" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  jm.Len.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				}
			}
		}
	}

	file.Replace(fileReplaces)
	file.State = assets.StateChanged
	return
}

type jsonMaker struct {
	Options      *GenerationOptions
	Types        map[string]*ast.TypeSpec // All types
	RN           string                   // Receiver Name
	FRN          string                   // Flat Receiver Name e.g. req.Time.
	RTN          string                   // Receiver Type Name
	TemplateSize int                      // e.g. len(`{"CaptchaID":[],"Image":""}`)
	Encoder      bytes.Buffer             // Generated Data
	Decoder      bytes.Buffer             // Generated Data
	Len          bytes.Buffer             // Len data to make slice size

	LastField bool
	// Field Options
	FieldName string
	Dash      bool
	OmitEmpty bool
	String    bool
	Tuple     bool
}

func (jm *jsonMaker) reset() {
	jm.TemplateSize = 0
	jm.Encoder.Reset()
	jm.Decoder.Reset()
	jm.Len.Reset()
	jm.LastField = false
}

func (jm *jsonMaker) make() (err *er.Error) {
	// Check needed type exist!!
	typ, found := jm.Types[jm.RTN]
	if !found {
		return ErrJSONNeededTypeNotExist
	}

	// Add some common data if ...
	if jm.TemplateSize == 0 {
		jm.TemplateSize += 2 // due to have len("{}") || len("[]")

		if jm.Options.Minifed && jm.Options.UnSafe {
			jm.Decoder.WriteString("\n	var decoder = json.DecoderUnsafeMinifed{\n		Buf: buf,\n	}\n")
		} else if jm.Options.Minifed && !jm.Options.UnSafe {
			jm.Decoder.WriteString("\n	var decoder = json.DecoderMinifed{\n		Buf: buf,\n	}\n")
		} else if !jm.Options.Minifed && !jm.Options.UnSafe {
			jm.Decoder.WriteString("\n	var decoder = json.Decoder{\n		Buf: buf,\n	}\n")
		} else {
			jm.Decoder.WriteString("\n	var decoder = json.DecoderUnsafe{\n		Buf: buf,\n	}\n")
		}

		jm.Decoder.WriteString("	for err == nil {\n")
		jm.Decoder.WriteString("		var keyName = decoder.DecodeKey()\n")
		jm.Decoder.WriteString("		switch keyName {\n")
		if !jm.Options.Minifed {
			jm.Decoder.WriteString("		case \"\":\n		return\n")
		}

		jm.Encoder.WriteString("\n	var encoder = json.Encoder{\n		Buf: make([]byte, 0, " + jm.RN + ".jsonLen()),\n	}\n" +
			"\n	encoder.EncodeString(`{\"")

		jm.Len.WriteString("\n	ln = 0")
	}

	switch structType := typ.Type.(type) {
	default:
		// Just occur if bad file pass to generator!!
		return
	case *ast.BasicLit:
		// TODO::: very simple type
	case *ast.StructType:
		for structFieldLoc, structField := range structType.Fields.List {
			if structFieldLoc == len(structType.Fields.List)-1 {
				jm.LastField = true
			}
			jm.FieldName = ""
			jm.OmitEmpty = false
			jm.String = false
			jm.Tuple = false

			if structField.Tag != nil && jm.checkFieldTag(structField.Tag.Value) {
				continue
			}

			for _, field := range structField.Names {
				if jm.FieldName == "" {
					jm.FieldName = field.Name
				}
				jm.Encoder.WriteString(jm.FieldName + `":`)
				jm.TemplateSize += len(field.Name) + 4 // +4 due to have len(,"":)

				jm.Decoder.WriteString("		case \"" + jm.FieldName + "\":\n")

				switch fieldType := structField.Type.(type) {
				case *ast.FuncType, *ast.InterfaceType, *ast.ChanType:
					log.Warn(ErrJSONFieldType, field.Name)
				case *ast.ArrayType:
					// Check array is slice?
					if fieldType.Len == nil {
						// Slice generator
						switch sliceType := fieldType.Elt.(type) {
						case *ast.ArrayType:
							// Check array is slice?
							if fieldType.Len == nil {
							} else {
							}
						case *ast.Ident:
							switch sliceType.Name {
							case "int", "uint":
							case "bool":
							case "byte", "uint8":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsBase64(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
									}
									jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeByteSliceAsBase64()\n")
									jm.Len.WriteString(" + ((len(" + jm.FRN + field.Name + ")*8+5)/6)")
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsNumber(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`],\"")
									}
									jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeByteSliceAsNumber()\n")
									jm.Len.WriteString(" + (len(" + jm.FRN + field.Name + ") * 4)") // 4 = 3(uint8 digits) + 1(comma)
								}
							case "int8":
							case "uint16":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeUInt16SliceAsBase64(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
									}
									jm.Decoder.WriteString("		err = decoder.DecodeUInt16ArrayAsBase64(" + jm.FRN + field.Name + ")\n")
									jm.Len.WriteString(" + (len(" + jm.FRN + field.Name + ")*2*8+5)/6)")
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeUInt16SliceAsNumber(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`],\"")
									}
									jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt16SliceAsNumber()\n")
									jm.Len.WriteString(" + (len(" + jm.FRN + field.Name + ") * 6)") // 6 = 5(uint16 digits) + 1(comma)
								}
							case "int16":
							case "uint32":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeUInt32SliceAsBase64(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
									}

									jm.Decoder.WriteString("		err = decoder.DecodeUInt32ArrayAsBase64(" + jm.FRN + field.Name + ")\n")

									jm.Len.WriteString(" + (len(" + jm.FRN + field.Name + ")*4*8+5)/6)")
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeUInt32SliceAsNumber(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`],\"")
									}

									jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt32SliceAsNumber()\n")

									jm.Len.WriteString(" + (len(" + jm.FRN + field.Name + ") * 11)") // 11 = 10(uint32 digits) + 1(comma)
								}
							case "int32":
							case "uint64":
							case "int64":
							case "string":
							default:
								// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
							}
							jm.TemplateSize += 2 // len(`[]`) || len(`""`)
						}
					} else {
						var arrayLen, _ = convert.Base10StringToUint32(fieldType.Len.(*ast.BasicLit).Value)

						switch arrayType := fieldType.Elt.(type) {
						case *ast.BasicLit:
							if arrayType.Kind == token.STRING {
								// Its common to use const to indicate number of array like in IP type as [16]byte!
								// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
							}
						case *ast.ArrayType:
							// Check array is slice?
							if fieldType.Len == nil {
							} else {
							}
						case *ast.Ident:
							switch arrayType.Name {
							case "int", "uint":
							case "bool":
							case "byte", "uint8":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsBase64(" + jm.FRN + field.Name + "[:])\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
									}
									jm.Decoder.WriteString("		err = decoder.DecodeByteArrayAsBase64(" + jm.FRN + field.Name + "[:])\n")
									jm.TemplateSize += (int(arrayLen)*8 + 5) / 6 // base64 encode len
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsNumber(" + jm.FRN + field.Name + "[:])\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`],\"")
									}
									jm.Decoder.WriteString("		err = decoder.DecodeByteArrayAsNumber(" + jm.FRN + field.Name + "[:])\n")
									jm.TemplateSize += int(arrayLen) * 4 // 4 = 3(uint8 digits) + 1(comma)
								}
							case "int8":
							case "uint16":
							case "int16":
							case "uint32":
							case "int32":
							case "uint64":
							case "int64":
							case "string":
							default:
								// TODO::: get related type by its name as fieldType.Elt.(*ast.Ident).Name
							}
							jm.TemplateSize += 2 // len(`[]`) || len(`""`)
						}
					}
				case *ast.MapType:

				case *ast.Ident:
					switch fieldType.Name {
					case "bool":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeBoolean(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("	" + jm.FRN + field.Name + ", err = decoder.DecodeBool()\n")
						jm.TemplateSize += 5
					case "uint":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt64(uint64(" + jm.FRN + field.Name + "))\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		var num uint64\n	num, err = decoder.DecodeUInt64()\n" + jm.FRN + field.Name + " = " + fieldType.Name + "(num)\n")
						jm.TemplateSize += 20
					case "byte", "uint8":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt8(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt8()\n")
						jm.TemplateSize += 3
					case "uint16":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt16(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt16()\n")
						jm.TemplateSize += 5
					case "uint32":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt32(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt32()\n")
						jm.TemplateSize += 10
					case "uint64":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt64(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeUInt64()\n")
						jm.TemplateSize += 20
					case "int", "int8", "int16":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeInt64(int64(" + jm.FRN + field.Name + "))\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("			var num int64\n		num, err = decoder.DecodeInt64()\n" + jm.FRN + field.Name + " = " + fieldType.Name + "(num)\n")
						jm.TemplateSize += 5
					case "int32":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeInt32(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeInt32()\n")
						jm.TemplateSize += 10
					case "int64":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeInt64(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
						}
						jm.Decoder.WriteString("		" + jm.FRN + field.Name + ", err = decoder.DecodeInt64()\n")
						jm.TemplateSize += 20
					case "string":
						jm.Encoder.WriteString("\"`)\n")
						jm.Encoder.WriteString("	encoder.EncodeString(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
						}
						jm.Decoder.WriteString("			" + jm.FRN + field.Name + ", err = decoder.DecodeString()\n")
						jm.TemplateSize += 2 // len(`""`)
						jm.Len.WriteString(" + len(" + jm.FRN + field.Name + ")")
					default:
					}
				case *ast.SelectorExpr:
					// TODO::: below code is not what it must be!
					jm.Encoder.WriteString("`)\n")
					jm.Encoder.WriteString("	encoder.EncodeUInt8(uint8(" + jm.FRN + field.Name + "))\n")
					if jm.LastField {
						jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
					} else {
						jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
					}
					jm.Decoder.WriteString("		var num uint8\n	num, err = decoder.DecodeUInt8()\n" + jm.FRN + field.Name + " = " + fieldType.X.(*ast.Ident).Name + "." + fieldType.Sel.Name + "(num)\n")
					jm.TemplateSize += 3
				case *ast.BasicLit:
					// TODO::: below code is not what it must be!
					jm.Encoder.WriteString("`)\n")
					jm.Encoder.WriteString("	encoder.EncodeUInt8(uint8(" + jm.FRN + field.Name + "))\n")
					if jm.LastField {
						jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
					} else {
						jm.Encoder.WriteString("\n	encoder.EncodeString(`,\"")
					}
					jm.Decoder.WriteString("		var num uint8\n	num, err = decoder.DecodeUInt8()\n" + jm.FRN + field.Name + " = " + fieldType.Value + "(num)\n")
					jm.TemplateSize += 3
				}
			}
		}
	}
	return
}

func (jm *jsonMaker) checkFieldTag(tagValue string) (notInclude bool) {
	var structFieldTag = reflect.StructTag(tagValue[1 : len(tagValue)-1])
	var structFieldTagJSON = structFieldTag.Get("json")
	if structFieldTagJSON != "" {
		if structFieldTagJSON == "-" {
			return true
		}

		var structFieldTagJSONParts = strings.Split(structFieldTagJSON, ",")
		if structFieldTagJSONParts[0] != "" {
			jm.FieldName = structFieldTagJSONParts[0]
		}

		for i := 1; i < len(structFieldTagJSONParts); i++ {
			switch structFieldTagJSONParts[i] {
			case "omitempty":
				jm.OmitEmpty = true
			case "string":
				jm.String = true
			case "tuple":
				jm.Tuple = true
			}
		}
	}
	return
}

const (
	returnError = "			if err != nil {\n				return\n			}\n"
)
