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
	"../log"
)

/*
Before pass file to safe||unsafe function, dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!
for just decoder method, jsonlen() can omit!

func ({{DesireName}} *{{DesireType}}) jsonDecoder(buf []byte) (err error) {
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
	Minifed           bool // without any space or new line, false for stylish version
	Strict            bool // Full check of key names && decode in any order not force by its struct(just last pair is important to check)
	UnSafe            bool
	AllowNoDefinedKey bool // allow not defined key in json encoded string to decode it! true return related error.
	NilMapAsNil       bool // false(empty map) >> "map":{},	true >> "map":nil,
	ForceUpdate       bool // true means delete exiting codes and update encoders && decoders codes anyway!
}

// CompleteMethods use to update given go files and complete json encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func CompleteMethods(file *assets.File, gos *GenerationOptions) (err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", file.Data, parser.ParseComments)
	if err != nil {
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
						return
					}

					jm.Encoder.WriteString("	return encoder.Buf\n")
					jm.Decoder.WriteString("		}\n\n" +
						"		err = decoder.IterationCheck()\n		if err != nil {\n		return\n		}" +
						"	}\n" +
						"	return\n")
					jm.TemplateSize-- // due to one unneeded leading comma!
					jm.Len.WriteString("	ln += " + strconv.FormatUint(uint64(jm.TemplateSize), 10) + "\n	return\n")
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

func (jm *jsonMaker) make() (err error) {
	// Check needed type exist!!
	typ, found := jm.Types[jm.RTN]
	if !found {
		return ErrJSONNeededTypeNotExist
	}

	// Add some common data if ...
	if jm.TemplateSize == 0 {
		jm.TemplateSize += 2 // due to have len("{}") || len("[]")

		jm.Decoder.WriteString(
			"\n	var decoder = json.DecoderUnsafeMinifed{\n		Buf: buf,\n	}\n" +
				"	for len(decoder.Buf) > 2 {\n")
		if jm.Options.Strict {

		} else {
			jm.Decoder.WriteString("decoder.Offset(2)\n")
			jm.Decoder.WriteString("switch decoder.Buf[0] {\n")

		}

		jm.Encoder.WriteString(
			"\n	var encoder = json.Encoder{\n		Buf: make([]byte, 0, " + jm.RN + ".jsonLen()),\n	}\n" +
				"\n	encoder.EncodeString(`{\"")

		jm.Len.WriteString("\n")
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

				if jm.Options.Strict {

				} else {
					jm.Decoder.WriteString("case '" + string(jm.FieldName[0]) + "':\n")
					jm.Decoder.WriteString("decoder.SetFounded()\n")
				}

				switch fieldType := structField.Type.(type) {
				case *ast.FuncType, *ast.InterfaceType, *ast.ChanType:
					log.Warn(ErrJSONFieldType, field.Name)
				case *ast.ArrayType:
					jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+3), 10) + ")\n") // +3 due to have ":[ after name
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
							case "bool", "uint8", "int8":
							case "byte":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsBase64(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(\",\"")
									}

									jm.Len.WriteString("	ln += base64.StdEncoding.EncodedLen(len(" + jm.FRN + field.Name + "))\n")
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsNumber(" + jm.FRN + field.Name + ")\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(],\"")
									}

									jm.Len.WriteString("	ln += len(" + jm.FRN + field.Name + ") * 4\n")
								}

								jm.TemplateSize += 2 // len(`[]`) || len(`""`)
							case "uint16", "int16":
							case "uint32", "int32":
							case "uint64", "int64":
							case "string":
							default:
								// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
							}
						}
					} else {
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
							case "byte":
								if jm.String {
									jm.Encoder.WriteString("\"`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsBase64(" + jm.FRN + field.Name + "[:])\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\"}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"")
									}

									jm.Decoder.WriteString("		err = decoder.DecodeArrayAsBase64(" + jm.FRN + field.Name + "[:])\n			if err != nil {\n				return\n				}\n")

									jm.Len.WriteString("	ln += base64.StdEncoding.EncodedLen(" + fieldType.Len.(*ast.BasicLit).Value + ")\n")
								} else {
									jm.Encoder.WriteString("[`)\n")
									jm.Encoder.WriteString("	encoder.EncodeByteSliceAsNumber(" + jm.FRN + field.Name + "[:])\n")
									if jm.LastField {
										jm.Encoder.WriteString("\n	encoder.EncodeString(`]}`)\n")
									} else {
										jm.Encoder.WriteString("\n	encoder.EncodeString(],\"")
									}

									jm.Decoder.WriteString("		for i := 0; i <" + fieldType.Len.(*ast.BasicLit).Value +
										"; i++ {\n		var value uint8\n		value, err = decoder.DecodeUInt8()\n		if err != nil {\n	return\n	}\n		" +
										jm.FRN + field.Name + "[i] = value\n	}\n")
									if !jm.Options.Strict {
										jm.Decoder.WriteString("	decoder.Offset(1)\n")
									}

									jm.Len.WriteString("	ln += " + fieldType.Len.(*ast.BasicLit).Value + " * 4\n")
								}

								jm.TemplateSize += 2 // len(`[]`) || len(`""`)
							case "uint8":
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
						}
					}
				case *ast.MapType:
					jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+3), 10) + ")\n") // +3 due to have ":{ after name

				case *ast.Ident:
					switch fieldType.Name {
					case "bool":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeBoolean(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(',\"")
						}

						jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+2), 10) + ")\n") // +2 due to have ": after name
						jm.Decoder.WriteString("	" + jm.FRN + field.Name + ", err = decoder.DecodeBool()\n	if err != nil {\n	return\n	}\n")

						jm.Len.WriteString("	ln += 5\n")
					case "uint", "byte", "uint8", "uint16", "uint32", "uint64":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeUInt64(uint64(" + jm.FRN + field.Name + "))\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(',\"")
						}

						jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+2), 10) + ")\n") // +2 due to have ": after name
						jm.Decoder.WriteString("		var num uint64\n	num, err = decoder.DecodeUInt64()\n		if err != nil {\n" +
							"		return\n		}\n		" + jm.FRN + field.Name + " = " + fieldType.Name + "(num)\n")

						jm.Len.WriteString("	ln += 20\n")
					case "int", "int8", "int16", "int32", "int64":
						jm.Encoder.WriteString("`)\n")
						jm.Encoder.WriteString("	encoder.EncodeInt64(uint64(" + jm.FRN + field.Name + "))\n")
						if jm.LastField {
							jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(',\"")
						}

						jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+2), 10) + ")\n") // +2 due to have ": after name
						jm.Decoder.WriteString("		var num int64\n	num, err = decoder.DecodeInt64()\n		if err != nil {\n" +
							"		return\n		}\n		" + jm.FRN + field.Name + " = " + fieldType.Name + "(num)\n")

						jm.Len.WriteString("	ln += 20\n")
					case "string":
						jm.Encoder.WriteString("\"`)\n")
						jm.Encoder.WriteString("	encoder.EncodeString(" + jm.FRN + field.Name + ")\n")
						if jm.LastField {
							jm.Encoder.WriteString("	encoder.EncodeString(`\"}`)\n")
						} else {
							jm.Encoder.WriteString("\n	encoder.EncodeString(`\",\"`")
						}

						jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+3), 10) + ")\n") // +3 due to have ":" after name
						jm.Decoder.WriteString("			" + jm.FRN + field.Name + " = decoder.DecodeString()\n")

						jm.TemplateSize += 2 // len(`""`)
						jm.Len.WriteString("	ln += len(" + jm.FRN + field.Name + ")\n")
					default:
					}
				case *ast.SelectorExpr:
					// TODO::: below code is not what it must be!
					jm.Encoder.WriteString("`)\n")
					jm.Encoder.WriteString("	encoder.EncodeUInt8(uint8(" + jm.FRN + field.Name + "))\n")
					if jm.LastField {
						jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
					} else {
						jm.Encoder.WriteString("\n	encoder.EncodeString(',\"")
					}

					jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+2), 10) + ")\n") // +2 due to have ": after name
					jm.Decoder.WriteString("		var num uint8\n	num, err = decoder.DecodeUInt8()\n		if err != nil {\n" +
						"		return\n		}\n		" + jm.FRN + field.Name + " = " + fieldType.X.(*ast.Ident).Name + "." + fieldType.Sel.Name + "(num)\n")

					jm.Len.WriteString("	ln += 3\n")
				case *ast.BasicLit:
					// TODO::: below code is not what it must be!
					jm.Encoder.WriteString("`)\n")
					jm.Encoder.WriteString("	encoder.EncodeUInt8(uint8(" + jm.FRN + field.Name + "))\n")
					if jm.LastField {
						jm.Encoder.WriteString("\n	encoder.EncodeByte('}')\n")
					} else {
						jm.Encoder.WriteString("\n	encoder.EncodeString(',\"")
					}

					jm.Decoder.WriteString("	decoder.Offset(" + strconv.FormatUint(uint64(len(jm.FieldName)+2), 10) + ")\n") // +2 due to have ": after name
					jm.Decoder.WriteString("		var num uint8\n	num, err = decoder.DecodeUInt8()\n		if err != nil {\n" +
						"		return\n		}\n		" + jm.FRN + field.Name + " = " + fieldType.Value + "(num)\n")

					jm.Len.WriteString("	ln += 3\n")
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

		if len(structFieldTagJSONParts) > 1 {
			structFieldTagJSONParts = structFieldTagJSONParts[1:]

			for i := 0; i < len(structFieldTagJSONParts); i++ {
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
	}
	return
}
