/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	datatype_p "memar/computer/datatype/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/process/service"
	"memar/protocol"
	user_p "memar/user/protocol"
)

var GenerateGoSyllabMethodsService = generateGoSyllabMethodsService{
	Service: service.New(protocol.SyllabServiceGenerateGoSyllabMethods, "", datatype_p.LifeCycle_PreAlpha, 1605027929).
		SetDetail(lang_p.LanguageEnglish, domainEnglish, "Generate Go Syllab Methods",
			"",
			``,
			[]string{"Syllab"}).
		SetAuthorization(operation_p.ActionType_Read, user_p.Type_All).Expired(0, ""),
}

type generateGoSyllabMethodsService struct {
	service.Service
}

// GenerateGoSyllabMethodsReq is the request of GenerateGoSyllabMethods()
type GenerateGoSyllabMethodsReq struct {
	GoType       string
	goTypeSpec   *ast.TypeSpec `json:"-" syllab:"-"`
	ReceiverName string
	Safe         bool // true means copy data from given buffer and don't use just point to it for decoding fields! buffer can GC after struct decode!
}

// GenerateGoSyllabMethodsReq is the response of GenerateGoSyllabMethods()
type GenerateGoSyllabMethodsRes struct {
	CheckSyllab      string
	FromSyllab       string
	ToSyllab         string
	Syllab_StackLength string
	Syllab_HeapLength  string
	Syllab_Length      string
	FieldsAccesser   []string
}

// GenerateGoSyllabMethods use to update given go files and complete Syllab encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func GenerateGoSyllabMethods(req *GenerateGoSyllabMethodsReq) (res *GenerateGoSyllabMethodsRes, err error_p.Error) {
	const packageTemp = "package temp\n"
	var goType = make([]byte, len(req.GoType)+len(packageTemp))
	copy(goType, packageTemp)
	copy(goType[len(packageTemp):], req.GoType)

	var fileParsed, goErr = parser.ParseFile(token.NewFileSet(), "", goType, parser.ParseComments)
	if goErr != nil {
		// log.Warn(domainEnglish, "Syllab generator error:", goErr)
		err = &ErrNeededTypeNotExist
		return
	}

	// find FromSyllab || ToSyllab method
	for _, decl := range fileParsed.Decls {
		switch declType := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range declType.Specs {
				switch gd := gDecl.(type) {
				case *ast.TypeSpec:
					req.goTypeSpec = gd
				default:
					err = &ErrNeededTypeNotExist
					return
				}
			}
		default: // *ast.FuncDecl
			err = &ErrNeededTypeNotExist
			return
		}
	}

	res, err = generateGoSyllabMethods(req)
	return
}

func generateGoSyllabMethods(req *GenerateGoSyllabMethodsReq) (res *GenerateGoSyllabMethodsRes, err error_p.Error) {
	var sm = syllabMaker{
		Request: req,
		FRN:     req.ReceiverName + ".",
		RTN:     req.goTypeSpec.Name.Name,
	}

	sm.Encoder.WriteString("func (" + req.ReceiverName + " *" + sm.RTN + ") ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {\n" +
		"	// pass buf by: buf = make([]byte, " + sm.Request.ReceiverName + ".Syllab_Length())\n" +
		"	// var i, ln uint32 // len of strings, slices, maps, ...\n\n")
	sm.Decoder.WriteString("func (" + req.ReceiverName + " *" + sm.RTN + ") FromSyllab(buf []byte, stackIndex uint32) {\n" +
		"	// var add, ln uint32\n	// var tempSlice []byte\n	// var heapIndex uint32 \n\n")
	sm.HeapSize.WriteString("func (" + req.ReceiverName + " *" + sm.RTN + ") Syllab_HeapLength() (ln uint32) {\n")

	err = sm.make()
	if err != nil {
		return
	}

	sm.Encoder.WriteString("	return\n}")
	sm.Decoder.WriteString("	return\n}")
	sm.HeapSize.WriteString("	return\n}")

	res = &GenerateGoSyllabMethodsRes{
		CheckSyllab: "func (" + req.ReceiverName + " *" + sm.RTN + ") CheckSyllab(buf []byte) (err error_p.Error) {" +
			"	if len(buf) < int(" + req.ReceiverName + ".Syllab_StackLength()) {\n" +
			"		err = syllab_errs.ShortArrayDecode\n" +
			"	}\n" +
			"	return\n}",
		FromSyllab:       sm.Decoder.String(),
		ToSyllab:         sm.Encoder.String(),
		Syllab_StackLength: "func (" + req.ReceiverName + " *" + sm.RTN + ") Syllab_HeapLength() (ln uint32) { return " + sm.getSLIAsString(0) + " }",
		Syllab_HeapLength:  sm.HeapSize.String(),
		Syllab_Length:      "func (" + req.ReceiverName + " *" + sm.RTN + ") Syllab_HeapLength() (ln uint32) { return uint64(" + req.ReceiverName + ".Syllab_StackLength() + " + req.ReceiverName + ".Syllab_HeapLength()) }",
		FieldsAccesser:   sm.FieldsAccesser,
	}

	// else if strings.HasPrefix(declType.Name.Name, "get") {
	// TODO:::
	// fileReplaces = append(fileReplaces, file.ReplaceReq{
	// 	Data: "\n	return " + req.ReceiverName + ".Syllab_StackLength() + " + req.ReceiverName + ".Syllab_HeapLength()n",
	// 	Start: int(declType.Body.Lbrace),
	// 	End:   int(declType.Body.Rbrace) - 1}) // -1 to not remove end brace
	// }
	return
}

type syllabMaker struct {
	Request        *GenerateGoSyllabMethodsReq
	FRN            string          // Flat Receiver Name e.g. req.Time.
	RTN            string          // Receiver Type Name
	LSI            uint64          // Last Stack Index
	StackSize      strings.Builder // Stack len data to make slice size
	HeapSize       strings.Builder // Heap len data to make slice size
	Encoder        strings.Builder // Generated Data
	Decoder        strings.Builder // Generated Data
	FieldsAccesser []string
}

func (sm *syllabMaker) reset() {
	sm.LSI = 0
	sm.StackSize.Reset()
	sm.HeapSize.Reset()
	sm.Encoder.Reset()
	sm.Decoder.Reset()
}

func (sm *syllabMaker) make() (err error_p.Error) {
	var fieldName string
	switch structType := sm.Request.goTypeSpec.Type.(type) {
	default:
		// Just occur if bad file pass to generator!!
		// err =
		return
	case *ast.BasicLit:
		// TODO::: very simple type
	case *ast.StructType:
		for _, structField := range structType.Fields.List {
			if structField.Tag != nil && sm.checkFieldTag(structField.Tag.Value) {
				continue
			}

			for _, field := range structField.Names {
				fieldName = field.Name

				switch fieldType := structField.Type.(type) {
				case *ast.FuncType, *ast.InterfaceType, *ast.ChanType:
					// log.Warn(domainEnglish, ErrFieldType, fieldName)
					// err = TODO:::
					return
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
								// log.Warn(domainEnglish, ErrFieldType, fieldName)
								// err = TODO:::
								return
							case "bool":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetBoolArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetBoolArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetBoolArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "byte", "uint8":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetByteArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetByteArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetByteArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "int8":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt8ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt8Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetInt8Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "uint16":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt16ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetUInt16Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 2)\n")
							case "int16":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt16ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetInt16Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 2)\n")
							case "uint32":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt32ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetUInt32Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 4)\n")
							case "int32":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt32ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetInt32Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 4)\n")
							case "uint64":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt64ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetUInt64Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
							case "int64":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt64ArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetInt64Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 8)\n")
							case "string":
								if sm.Request.Safe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetStringArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	for i:=0; i<len(" + sm.FRN + fieldName + "); i++ {\n")
								sm.HeapSize.WriteString("		ln += uint32(len(" + sm.FRN + fieldName + "[i]))\n")
								sm.HeapSize.WriteString("	}\n")
							default:
								// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
								// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", heapIndex)\n")
								// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(4) + ", ln)\n")
								// sm.Encoder.WriteString("	copy(buf[heapIndex:], " + sm.FRN + fieldName + ")\n")
								// if sm.Request.Safe {
								// 	sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
								// 	sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
								// 	sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = buf[add : add+ln]\n")
								// } else {
								// 	sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
								// 	sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
								// 	sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = make([]byte, ln)\n")
								// 	sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + ", buf[add : add+ln])\n")
								// }
							}
						}
						// In any case we need 8 byte for address and len of array!
						sm.LSI += 8
					} else {
						// Get array len
						var ln, err = strconv.ParseUint(fieldType.Len.(*ast.BasicLit).Value, 10, 64)
						if err != nil {
							return &ErrArrayLen
						}

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
								// log.Warn(domainEnglish, ErrFieldType, fieldName)
								return &ErrFieldType
							case "bool":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeBoolSliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToBoolSlice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln
							case "byte", "uint8":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], " + sm.FRN + fieldName + "[:])\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], buf[" + sm.getSLIAsString(0) + ":])\n")
								sm.LSI += ln
							case "int8":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeInt8SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToInt8Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln
							case "uint16":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeUInt16SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToUInt16Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 2
							case "int16":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeInt16SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToInt16Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 2
							case "uint32":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeUInt32SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToUInt32Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 4
							case "int32":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeInt32SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToInt32Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 4
							case "uint64":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeUInt64SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToUInt64Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 8
							case "int64":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeInt64SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToInt64Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 8
							case "float32":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeFloat32SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToFloat32Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 4
							case "float64":
								sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], convert.UnsafeFloat64SliceToByteSlice(" + sm.FRN + fieldName + "[:]))\n")
								sm.Decoder.WriteString("	copy(" + sm.FRN + fieldName + "[:], convert.UnsafeByteSliceToFloat64Slice(buf[" + sm.getSLIAsString(0) + ":]))\n")
								sm.LSI += ln * 4
							case "string":
								if sm.Request.Safe {
									// sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArraySafely(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									// sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	heapIndex = syllab.SetStringArray(buf, " + sm.FRN + fieldName + "[:], " + sm.getSLIAsString(0) + ", heapIndex)\n")
								sm.HeapSize.WriteString("	for i:=0; i<" + fieldType.Len.(*ast.BasicLit).Value + "; i++ {\n")
								sm.HeapSize.WriteString("		ln += len(" + sm.FRN + fieldName + "[i])\n")
								sm.HeapSize.WriteString("	}\n")
							default:
								// TODO::: get related type by its name as fieldType.Elt.(*ast.Ident).Name
							}
						}
					}
				case *ast.StructType:
					var tmp = sm.FRN
					sm.FRN += fieldName + "."
					sm.RTN = fieldType.Fields.List[0].Names[0].Name
					// TODO::: add struct itself to sm.Types
					err = sm.make()
					sm.FRN = tmp
				case *ast.MapType:

				case *ast.Ident:
					switch fieldType.Name {
					case "int", "uint":
						// log.Warn(domainEnglish, ErrFieldType, fieldName)
						// err = TODO:::
						return
					case "bool":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetBool(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetBool(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI++
					case "byte":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetByte(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetByte(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI++
					case "int8":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetInt8(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt8(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI++
					case "uint8":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetUInt8(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt8(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI++
					case "int16":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetInt16(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt16(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 2
					case "uint16":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetUInt16(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt16(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 2
					case "int32":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetInt32(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt32(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 4
					case "uint32":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 4
					case "int64":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetInt64(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt64(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 8
					case "uint64":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetUInt64(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt64(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 8
					case "float32":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetFloat32(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetFloat32(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 4
					case "float64":
						// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
						sm.Encoder.WriteString("	syllab.SetFloat64(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + fieldName + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetFloat64(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.LSI += 8
					case "string":
						if sm.Request.Safe {
							sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringSafely(buf, " + sm.getSLIAsString(0) + ")\n")
						} else {
							sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetString(buf, " + sm.getSLIAsString(0) + ")\n")
						}
						sm.Encoder.WriteString("	heapIndex = syllab.SetString(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", heapIndex)\n")
						sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
						sm.LSI += 8
					default:
						// TODO::: below code not work for very simple type e.g. type test uint8
						sm.Encoder.WriteString("	heapIndex = " + sm.FRN + fieldName + ".ToSyllab(buf, " + sm.getSLIAsString(0) + ", heapIndex)\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + ".FromSyllab(buf, " + sm.getSLIAsString(0) + ")\n")

						sm.StackSize.WriteString(" +" + sm.FRN + fieldName + ".Syllab_StackLength()")
						sm.HeapSize.WriteString("	ln += " + sm.FRN + fieldName + ".Syllab_HeapLength()\n")
					}
				case *ast.SelectorExpr:
					sm.Encoder.WriteString("	heapIndex = " + sm.FRN + fieldName + ".ToSyllab(buf, " + sm.getSLIAsString(0) + ", heapIndex)\n")
					sm.Decoder.WriteString("	" + sm.FRN + fieldName + ".FromSyllab(buf," + sm.getSLIAsString(0) + ")\n")

					sm.StackSize.WriteString(" + " + sm.FRN + fieldName + ".Syllab_StackLength()")
					sm.HeapSize.WriteString("	ln += " + sm.FRN + fieldName + ".Syllab_HeapLength()\n")
				case *ast.BasicLit:
					// log.Info(domainEnglish, "BasicLit :", t.Kind)
					// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", heapIndex)\n")
					// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(4) + ", ln)\n")
					// sm.Encoder.WriteString("	copy(buf[heapIndex:], " + sm.FRN + fieldName + ")\n")
					// if sm.Request.Safe {
					// 	sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
					// 	sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
					// 	sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = convert.UnsafeByteSliceToString(buf[add : add+ln])\n")
					// } else {
					// 	sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
					// 	sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
					// 	sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = string(buf[add : add+ln])\n")
					// }
				}
			}
		}
	}
	return
}

func (sm *syllabMaker) getSLIAsString(plus uint64) (s string) {
	// TODO::: Improve below line!
	s += strconv.FormatUint(sm.LSI+plus, 10)
	s += sm.StackSize.String()
	return
}

type fieldTag struct {
	id         int  // orgID  [32]byte `syllab:"id={{ID}}"`
	deprecated bool // orgID  [32]byte `deprecated:"{{Time}}"`
	notInclude bool // orgID  [32]byte `syllab:"-"`
}

func (sm *syllabMaker) checkFieldTag(tagValue string) (notInclude bool) {
	var structFieldTag = reflect.StructTag(tagValue[1 : len(tagValue)-1]) // remove ``
	var structFieldTagSyllab = structFieldTag.Get("syllab")
	if structFieldTagSyllab == "-" {
		return true
	}
	return
}
