/* For license and copyright information please see LEGAL file in repository */

package syllab

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
Before pass file to CompleteMethods(), dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!
For just syllabDecoder() method you can omit syllabStackLen() & syllabLen()

func ({{DesireName}} *{{DesireType}}) syllabDecoder(buf []byte) (err error) {
	return
}

func ({{DesireName}} *{{DesireType}}) SyllabDecoder(buf []byte, stackIndex int) {
	return
}

func ({{DesireName}} *{{DesireType}}) syllabEncoder(buf []byte) {}

func ({{DesireName}} *{{DesireType}}) syllabEncoder() (buf []byte) {
	buf = make([]byte, {{DesireName}}.syllabLen())
	return
}

func ({{DesireName}} *{{DesireType}}) SyllabEncoder(buf []byte, stackIndex, heapIndex uint32) (hi uint32) {
	return heapIndex
}

func ({{DesireName}} *{{DesireType}}) syllabStackLen() (ln uint32) {
	return
}

func ({{DesireName}} *{{DesireType}}) syllabHeapLen() (ln uint32) {
	return
}

func ({{DesireName}} *{{DesireType}}) syllabLen() (ln int) {
	return int({{DesireName}}.SyllabStackLen() + {{DesireName}}.syllabHeapLen())
}
*/

// GenerationOptions indicate generator behavior!
type GenerationOptions struct {
	UnSafe      bool // true means don't copy data from given payload||buffer and just point to it for decoding fields! buffer can't GC until decoded struct free!
	ForceUpdate bool // true means delete exiting codes and update encoders && decoders codes anyway!
}

// CompleteMethods use to update given go files and complete Syllab encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func CompleteMethods(file *assets.File, gos *GenerationOptions) (err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", file.Data, parser.ParseComments)
	if err != nil {
		return
	}

	var fileReplaces = make([]assets.ReplaceReq, 0, 4)
	var sm = syllabMaker{
		Options: gos,
		Types:   map[string]*ast.TypeSpec{},
	}

	// find syllabDecoder || syllabEncoder method
	for _, decl := range fileParsed.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range d.Specs {
				switch gd := gDecl.(type) {
				case *ast.TypeSpec:
					sm.Types[gd.Name.Name] = gd
				}
			}
		case *ast.FuncDecl:
			if d.Recv != nil {
				if sm.RN != d.Recv.List[0].Names[0].Name {
					sm.reset()
					sm.RN = d.Recv.List[0].Names[0].Name
					sm.FRN = d.Recv.List[0].Names[0].Name + "."
					sm.RTN = d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name

					err = sm.make()
					if err != nil {
						return
					}

					sm.Encoder.WriteString("	return\n")
					sm.Decoder.WriteString("	return\n")
					sm.HeapSize.WriteString("	return\n")
				}

				// Just needed methods!
				if d.Name.Name == "syllabDecoder" || d.Name.Name == "SyllabDecoder" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  sm.Decoder.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "syllabEncoder" || d.Name.Name == "SyllabEncoder" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  sm.Encoder.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "syllabStackLen" || d.Name.Name == "SyllabStackLen" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data: "\n	return " + sm.getSLIAsString(0) + "\n",
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "syllabHeapLen" || d.Name.Name == "SyllabHeapLen" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data:  sm.HeapSize.String(),
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if d.Name.Name == "syllabLen" || d.Name.Name == "SyllabLen" {
					fileReplaces = append(fileReplaces, assets.ReplaceReq{
						Data: "\n	return int(" + sm.RN + ".syllabStackLen() + " + sm.RN + ".syllabHeapLen())\n",
						Start: int(d.Body.Lbrace),
						End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				} else if strings.HasPrefix(d.Name.Name, "get") {
					// TODO:::
					// fileReplaces = append(fileReplaces, assets.ReplaceReq{
					// 	Data: "\n	return " + sm.RN + ".syllabStackLen() + " + sm.RN + ".syllabHeapLen()\n",
					// 	Start: int(d.Body.Lbrace),
					// 	End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
				}
			}
		}
	}

	file.Replace(fileReplaces)
	file.State = assets.StateChanged
	return
}

type syllabMaker struct {
	Options   *GenerationOptions
	Types     map[string]*ast.TypeSpec // All types
	RN        string                   // Receiver Name
	FRN       string                   // Flat Receiver Name e.g. req.Time.
	RTN       string                   // Receiver Type Name
	LSI       uint64                   // Last Stack Index
	StackSize bytes.Buffer             // Stack len data to make slice size
	HeapSize  bytes.Buffer             // Heap len data to make slice size
	Encoder   bytes.Buffer             // Generated Data
	Decoder   bytes.Buffer             // Generated Data
}

func (sm *syllabMaker) reset() {
	sm.LSI = 0
	sm.StackSize.Reset()
	sm.HeapSize.Reset()
	sm.Encoder.Reset()
	sm.Decoder.Reset()
}

func (sm *syllabMaker) make() (err error) {
	// Check needed type exist!!
	typ, found := sm.Types[sm.RTN]
	if !found {
		return ErrSyllabNeededTypeNotExist
	}

	// Add some common data if ...
	if sm.LSI == 0 {
		sm.Decoder.WriteString(
			"\n	var add, ln uint32\n	// var tempSlice []byte\n\n" +
				"	if uint32(len(buf)) < " + sm.RN + ".syllabStackLen() {\n" +
				"		err = syllab.ErrSyllabDecodeSmallSlice\n" +
				"		return\n" +
				"	}\n\n")

		sm.Encoder.WriteString(
			"\n	// buf = make([]byte, " + sm.RN + ".syllabLen()+offset)\n" +
				"	var hsi uint32 = " + sm.RN + ".syllabStackLen() // Heap start index || Stack size!\n" +
				"	// var i, ln uint32 // len of strings, slices, maps, ...\n\n")

		sm.HeapSize.WriteString("\n")
	}

	var fieldName string
	switch structType := typ.Type.(type) {
	default:
		// Just occur if bad file pass to generator!!
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
					log.Warn(ErrSyllabFieldType, fieldName)
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
								log.Warn(ErrSyllabFieldType, fieldName)
							case "bool":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetBoolArray(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetBoolArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetBoolArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "byte", "uint8":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetByteArray(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetByteArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetByteArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "int8":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetInt8Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt8Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetInt8Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
							case "uint16":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetUInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetUInt16Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 2)\n")
							case "int16":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt16Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetInt16Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 2)\n")
							case "uint32":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetUInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetUInt32Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 4)\n")
							case "int32":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt32Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetInt32Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 4)\n")
							case "uint64":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetUInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetUInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetUInt64Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
							case "int64":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetInt64Array(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetInt64Array(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + ") * 8)\n")
							case "string":
								if sm.Options.UnSafe {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetStringArray(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
								sm.HeapSize.WriteString("	for i:=0; i<len(" + sm.FRN + fieldName + "); i++ {\n")
								sm.HeapSize.WriteString("		ln += uint32(len(" + sm.FRN + fieldName + "[i]))\n")
								sm.HeapSize.WriteString("	}\n")
							default:
								// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
								// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
								// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(4) + ", ln)\n")
								// sm.Encoder.WriteString("	copy(buf[hsi:], " + sm.FRN + fieldName + ")\n")
								// if sm.Options.UnSafe {
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
							return ErrSyllabArrayLen
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
								log.Warn(ErrSyllabFieldType, fieldName)
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
								if sm.Options.UnSafe {
									// sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								} else {
									// sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetStringArray(buf, " + sm.getSLIAsString(0) + ")\n")
								}
								sm.Encoder.WriteString("	hsi = syllab.SetStringArray(buf, " + sm.FRN + fieldName + "[:], " + sm.getSLIAsString(0) + ", hsi)\n")
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
						log.Warn(ErrSyllabFieldType, fieldName)
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
						if sm.Options.UnSafe {
							sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.UnsafeGetString(buf, " + sm.getSLIAsString(0) + ")\n")
						} else {
							sm.Decoder.WriteString("	" + sm.FRN + fieldName + " = syllab.GetString(buf, " + sm.getSLIAsString(0) + ")\n")
						}
						sm.Encoder.WriteString("	hsi = syllab.SetString(buf, " + sm.FRN + fieldName + ", " + sm.getSLIAsString(0) + ", hsi)\n")
						sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + fieldName + "))\n")
						sm.LSI += 8
					default:
						// TODO::: below code not work for very simple type e.g. type test uint8
						sm.Encoder.WriteString("	hsi = " + sm.FRN + fieldName + ".syllabEncoder(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
						sm.Decoder.WriteString("	" + sm.FRN + fieldName + ".syllabDecoder(buf, " + sm.getSLIAsString(0) + ")\n")

						sm.StackSize.WriteString(" +" + sm.FRN + fieldName + ".syllabStackLen()")
						sm.HeapSize.WriteString("	ln += " + sm.FRN + fieldName + ".syllabHeapLen()\n")
					}
				case *ast.SelectorExpr:
					sm.Encoder.WriteString("	hsi = " + sm.FRN + fieldName + ".SyllabEncoder(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
					sm.Decoder.WriteString("	" + sm.FRN + fieldName + ".SyllabDecoder(buf," + sm.getSLIAsString(0) + ")\n")

					sm.StackSize.WriteString(" + " + sm.FRN + fieldName + ".SyllabStackLen()")
					sm.HeapSize.WriteString("	ln += " + sm.FRN + fieldName + ".SyllabHeapLen()\n")
				case *ast.BasicLit:
					// log.Info("BasicLit :", t.Kind)
					// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
					// sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(4) + ", ln)\n")
					// sm.Encoder.WriteString("	copy(buf[hsi:], " + sm.FRN + fieldName + ")\n")
					// if sm.Options.UnSafe {
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

func (sm *syllabMaker) checkFieldTag(tagValue string) (notInclude bool) {
	var structFieldTag = reflect.StructTag(tagValue[1 : len(tagValue)-1])
	var structFieldTagSyllab = structFieldTag.Get("syllab")
	if structFieldTagSyllab == "-" {
		return true
	}
	return
}
