/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"go/ast"
	"go/parser"
	"go/token"
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
	HelperFuncs bool // true means use lib function that can't inline instead of pure code!
	UnSafe      bool // true means allow use unsafe package in generated codes!
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
			sm.RN = d.Recv.List[0].Names[0].Name
			sm.FRN = d.Recv.List[0].Names[0].Name + "."
			sm.RTN = d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name

			// Just needed methods!
			if d.Name.Name == "syllabDecoder" || d.Name.Name == "SyllabDecoder" {
				if sm.LSI == 0 {
					err = sm.make()
					if err != nil {
						return
					}
				}
				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data:  sm.Decoder.String(),
					Start: int(d.Body.Lbrace),
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			} else if d.Name.Name == "syllabEncoder" || d.Name.Name == "SyllabEncoder" {
				if sm.LSI == 0 {
					err = sm.make()
					if err != nil {
						return
					}
				}
				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data:  sm.Encoder.String(),
					Start: int(d.Body.Lbrace),
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			} else if d.Name.Name == "syllabStackLen" || d.Name.Name == "SyllabStackLen" {
				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data: "\n	return " + sm.getSLIAsString(0) + " // fixed size data + variables data add&&len\n",
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
	StackSize strings.Builder          // Stack len data to make slice size
	HeapSize  strings.Builder          // Heap len data to make slice size
	Encoder   strings.Builder          // Generated Data
	Decoder   strings.Builder          // Generated Data
}

func (sm *syllabMaker) make() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrSyllabNestedStruct
	}

	// Add some common data if ...
	if sm.LSI == 0 {
		sm.Decoder.WriteString(
			"\n	var add, ln uint32\n	var tempSlice []byte\n\n" +
				"	if uint32(len(buf)) < " + sm.RN + ".syllabStackLen() {\n" +
				"	err = syllab.ErrSyllabDecodeSmallSlice\n" +
				"	return\n" +
				"	}\n\n")

		sm.Encoder.WriteString(
			"\n	// buf = make([]byte, " + sm.RN + ".syllabLen()+offset)\n" +
				"	var hsi uint32 = " + sm.RN + ".syllabStackLen() // Heap start index || Stack size!\n" +
				"	var ln uint32 // len of strings, slices, maps, ...\n\n")

		sm.HeapSize.WriteString("\n")
	}

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {
				// Slice generator
				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					log.Warn(ErrSyllabFieldType, in)
				case "bool", "uint8", "int8":
					sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + "))\n")
				case "byte":
					if sm.Options.HelperFuncs {
						if sm.Options.UnSafe {
							sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.UnsafeGetByteArray(buf, " + sm.getSLIAsString(0) + ")\n")
						} else {
							sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetByteArray(buf, " + sm.getSLIAsString(0) + ")\n")
						}
					} else {
						if sm.Options.UnSafe {
							sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
							sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
							sm.Decoder.WriteString("	" + sm.FRN + in + " = buf[add : add+ln]\n")
						} else {
							sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
							sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
							sm.Decoder.WriteString("	" + sm.FRN + in + " = make([]byte, ln)\n")
							sm.Decoder.WriteString("	copy(" + sm.FRN + in + ", buf[add : add+ln])\n")
						}
					}

					sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + "))\n")
				case "uint16", "int16":
					sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + ") * 2)\n")
				case "uint32", "int32":
					sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + ") * 4)\n")
				case "uint64", "int64":
					sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + ") * 8)\n")
				case "string":
					sm.HeapSize.WriteString("	for i:=0; i<len(" + sm.FRN + in + "); i++ {\n")
					sm.HeapSize.WriteString("		ln += uint32(len(" + sm.FRN + in + "[i]))\n")
					sm.HeapSize.WriteString("	}\n")
				default:
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}

				// In any case we need 8 byte for address and len of array!
				sm.LSI += 8
			} else {
				// Get array len
				var ln, err = strconv.ParseUint(t.Len.(*ast.BasicLit).Value, 10, 64)
				if err != nil {
					return ErrSyllabArrayLen
				}

				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					log.Warn(ErrSyllabFieldType, in)
				case "bool":
					sm.LSI += ln
				case "byte":
					sm.Encoder.WriteString("	copy(buf[" + sm.getSLIAsString(0) + ":], " + sm.FRN + in + "[:])\n")
					sm.Decoder.WriteString("	copy(" + sm.FRN + in + "[:], buf[" + sm.getSLIAsString(0) + ":])\n")
					sm.LSI += ln
				case "uint8":
					sm.LSI += ln
				case "int8":
					sm.LSI += ln
				case "uint16":
					// for i:= 0; i<ln; i++ {
					// 	sm.Encoder.WriteString("	buf["+strconv.FormatUint(sm.LSI+i)+"] = "+sm.FRN+in+"["+strconv.FormatUint(i)+"];")
					// }
					// sm.Encoder.WriteString("\n")
					sm.LSI += ln * 2
				case "int16":
					sm.LSI += ln * 2
				case "uint32":
					sm.LSI += ln * 4
				case "int32":
					sm.LSI += ln * 4
				case "uint64":
					sm.LSI += ln * 8
				case "int64":
					sm.LSI += ln * 8
				case "string":
					sm.HeapSize.WriteString("	for i:=0; i<" + t.Len.(*ast.BasicLit).Value + "; i++ {\n")
					sm.HeapSize.WriteString("		ln += len(" + sm.FRN + in + "[i])\n")
					sm.HeapSize.WriteString("	}\n")
				default:
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}
			}
		case *ast.StructType:
			var tmp = sm.FRN
			sm.FRN += in + "."
			sm.RTN = t.Fields.List[0].Names[0].Name
			err = sm.make()
			sm.FRN = tmp
		case *ast.FuncType:
			log.Warn(ErrSyllabFieldType, in)
		case *ast.InterfaceType:
			log.Warn(ErrSyllabFieldType, in)
		case *ast.MapType:

		case *ast.ChanType:
			log.Warn(ErrSyllabFieldType, in)
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				log.Warn(ErrSyllabFieldType, in)
			case "bool":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetBool(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetBool(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI++
			case "byte":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetByte(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetByte(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI++
			case "int8":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetInt8(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetInt8(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI++
			case "uint8":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetUInt8(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetUInt8(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI++
			case "int16":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetInt16(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetInt16(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 2
			case "uint16":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetUInt16(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetUInt16(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 2
			case "int32":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetInt32(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetInt32(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 4
			case "uint32":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 4
			case "int64":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetInt64(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetInt64(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 8
			case "uint64":
				// Inlined by go compiler! So don't respect dev wants not use HelperFuncs
				sm.Encoder.WriteString("	syllab.SetUInt64(buf, " + sm.getSLIAsString(0) + ", " + sm.FRN + in + ")\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetUInt64(buf, " + sm.getSLIAsString(0) + ")\n")
				sm.LSI += 8
			case "string":
				sm.Encoder.WriteString("	ln = uint32(len(" + sm.FRN + in + "))\n")
				if sm.Options.HelperFuncs {
					sm.Encoder.WriteString("	syllab.SetString(buf, " + sm.FRN + in + ", " + sm.getSLIAsString(0) + ", hsi)\n")
					if sm.Options.UnSafe {
						sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.UnsafeGetString(buf, " + sm.getSLIAsString(0) + ")\n")
					} else {
						sm.Decoder.WriteString("	" + sm.FRN + in + " = syllab.GetString(buf, " + sm.getSLIAsString(0) + ")\n")
					}
				} else {
					sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
					sm.Encoder.WriteString("	syllab.SetUInt32(buf, " + sm.getSLIAsString(4) + ", ln)\n")
					sm.Encoder.WriteString("	copy(buf[hsi:], " + sm.FRN + in + ")\n")
					if sm.Options.UnSafe {
						sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
						sm.Decoder.WriteString("	tempSlice = buf[add : add+ln]\n")
						sm.Decoder.WriteString("	" + sm.FRN + in + " = *(*string)(unsafe.Pointer(&tempSlice))\n")
					} else {
						sm.Decoder.WriteString("	add = syllab.GetUInt32(buf, " + sm.getSLIAsString(0) + ")\n")
						sm.Decoder.WriteString("	ln = syllab.GetUInt32(buf, " + sm.getSLIAsString(4) + ")\n")
						sm.Decoder.WriteString("	" + sm.FRN + in + " = string(buf[add : add+ln])\n")
					}
				}
				sm.Encoder.WriteString("	hsi += ln\n")

				sm.HeapSize.WriteString("	ln += uint32(len(" + sm.FRN + in + "))\n")
				sm.LSI += 8
			default:
				// TODO::: below code not work for very simple type e.g. type test uint8
				sm.Encoder.WriteString("	hsi = " + sm.FRN + in + ".syllabEncoder(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + ".syllabDecoder(buf, " + sm.getSLIAsString(0) + ")\n")

				sm.StackSize.WriteString(" +" + sm.FRN + in + ".syllabStackLen()")
				sm.HeapSize.WriteString("	ln += " + sm.FRN + in + ".syllabHeapLen()\n")
			}
		case *ast.SelectorExpr:
			sm.Encoder.WriteString("	hsi = " + sm.FRN + in + ".SyllabEncoder(buf, " + sm.getSLIAsString(0) + ", hsi)\n")
			sm.Decoder.WriteString("	" + sm.FRN + in + ".SyllabDecoder(buf," + sm.getSLIAsString(0) + ")\n")

			sm.StackSize.WriteString(" + " + sm.FRN + in + ".SyllabStackLen()")
			sm.HeapSize.WriteString("	ln += " + sm.FRN + in + ".SyllabHeapLen()\n")
		case *ast.BasicLit:
			// log.Info("BasicLit :", t.Kind)
		}
	}

	sm.Encoder.WriteString("\n	// return buf\n")
	sm.Decoder.WriteString("\n	return\n")

	sm.HeapSize.WriteString("	return\n")

	return
}

func (sm *syllabMaker) getSLIAsString(plus uint64) (s string) {
	// TODO::: Improve below line!
	s += strconv.FormatUint(sm.LSI+plus, 10)
	s += sm.StackSize.String()
	return
}
