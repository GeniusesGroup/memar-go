/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"

	"../assets"
)

/*
Before pass file to CompleteMethods(), dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!
For just syllabDecoder() method you can omit syllabFixedSize() & syllabLen()

func ({{DesireName}} *{{DesireType}}) syllabDecoder(buf []byte) (err error) {
	return
}

// offset add free space by given number at begging of return slice that almost just use in sRPC protocol! It can be 0!!
func ({{DesireName}} *{{DesireType}}) syllabEncoder(offset int) (buf []byte) {
	return
}

func ({{DesireName}} *{{DesireType}}) syllabFixedSize() (ln int) {
	return
}

func ({{DesireName}} *{{DesireType}}) syllabLen() (ln int) {
	return
}
*/

// GenerationOptions indicate generator behavior!
type GenerationOptions struct {
	HelperFuncs bool // true means use lib function such as GetUInt64() instead of pure code!
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
		Types: map[string]*ast.TypeSpec{},
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
			if d.Name.Name == "syllabDecoder" {
				// Add some common data
				sm.Decoder.WriteString(
					"\n	var add, ln uint32\n\n" +
						"	if len(buf) < " + sm.RN + ".syllabFixedSize() {\n" +
						"	err = syllab.ErrSyllabDecodeSmallSlice\n" +
						"	return\n" +
						"	}\n\n")

				err = sm.makeDecoder()
				if err != nil {
					return
				}

				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data:  sm.Decoder.String(),
					Start: int(d.Body.Lbrace) + 1,  // +1 for bracket itself
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			} else if d.Name.Name == "syllabEncoder" {
				// Add some common data
				sm.Encoder.WriteString(
					"\n	buf = make([]byte, " + sm.RN + ".syllabLen()+offset)\n" +
						"	var hsi int = " + sm.RN + ".syllabFixedSize() // Heap start index || Stack size!\n" +
						"	var ln int // len of strings, slices, maps, ...\n" +
						"	var b = buf[offset:]\n\n")

				err = sm.makeEncoder()
				if err != nil {
					return
				}

				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data:  sm.Encoder.String(),
					Start: int(d.Body.Lbrace) + 1,  // +1 for bracket itself
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			} else if d.Name.Name == "syllabFixedSize" {
				sm.HeapSize.WriteString(" ln += " + sm.RN + ".syllabFixedSize()\n")
				err = sm.makeStackHeapSize()
				if err != nil {
					return
				}

				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data: "\n	return " + sm.getSLIAsString(0) + " // fixed size data + variables data add&&len\n",
					Start: int(d.Body.Lbrace) + 1,  // +1 for bracket itself
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			} else if d.Name.Name == "syllabLen" {
				fileReplaces = append(fileReplaces, assets.ReplaceReq{
					Data:  sm.HeapSize.String(),
					Start: int(d.Body.Lbrace) + 1,  // +1 for bracket itself
					End:   int(d.Body.Rbrace) - 1}) // -1 to not remove end brace
			}
		}
	}

	file.Replace(fileReplaces)
	file.State = assets.StateChanged
	return
}

type syllabMaker struct {
	Types    map[string]*ast.TypeSpec // All types
	RN       string                   // Receiver Name
	FRN      string                   // Flat Receiver Name e.g. req.Time.
	RTN      string                   // Receiver Type Name
	LSI      uint64                   // Last Stack Index
	HeapSize strings.Builder          // Heap len data to make slice size
	Encoder  strings.Builder          // Generated Data
	Decoder  strings.Builder          // Generated Data
}

func (sm *syllabMaker) makeDecoder() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrSyllabNestedStruct
	}

	// reset sm.LSI to zero!
	sm.LSI = 0

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {
				// TODO::: Slice generator
			} else {
				// Get array len
				var len, err = strconv.ParseUint(t.Len.(*ast.BasicLit).Value, 10, 64)
				if err != nil {
					return ErrSyllabArrayLen
				}

				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					return ErrSyllabFieldType
				case "bool":
					sm.LSI += len
				case "byte":
					sm.Decoder.WriteString("	copy(" + sm.FRN + in + "[:], buf[" + sm.getSLIAsString(0) + ":])\n")
					sm.LSI += len
				case "uint8":
					sm.LSI += len
				case "int8":
					sm.LSI += len
				case "uint16":
					sm.LSI += len * 2
				case "int16":
					sm.LSI += len * 2
				case "uint32":
					sm.LSI += len * 4
				case "int32":
					sm.LSI += len * 4
				case "uint64":
					sm.LSI += len * 8
				case "int64":
					sm.LSI += len * 8
				case "string":
				default:
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}
			}
		case *ast.StructType:
			var tmp = sm.FRN
			sm.FRN += in + "."
			sm.RTN = t.Fields.List[0].Names[0].Name
			err = sm.makeDecoder()
			sm.FRN = tmp
		case *ast.FuncType:
			return ErrSyllabFieldType
		case *ast.InterfaceType:
			return ErrSyllabFieldType
		case *ast.MapType:

		case *ast.ChanType:
			return ErrSyllabFieldType
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return ErrSyllabFieldType
			case "bool":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = buf[" + sm.getSLIAsString(0) + "] == 1 \n")
				sm.LSI++
			case "byte":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = buf[" + sm.getSLIAsString(0) + "]\n")
				sm.LSI++
			case "uint8":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = uint8(buf[" + sm.getSLIAsString(0) + "])\n")
				sm.LSI++
			case "int8":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = int8(buf[" + sm.getSLIAsString(0) + "])\n")
				sm.LSI++
			case "uint16":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = uint16(buf[" + sm.getSLIAsString(0) + "]) | uint16(buf[" + sm.getSLIAsString(1) + "])<<8\n")
				sm.LSI += 2
			case "int16":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = int16(buf[" + sm.getSLIAsString(0) + "]) | int16(buf[" + sm.getSLIAsString(1) + "])<<8\n")
				sm.LSI += 2
			case "uint32":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = uint32(buf[" + sm.getSLIAsString(0) + "]) | uint32(buf[" + sm.getSLIAsString(1) + "])<<8 | uint32(buf[" + sm.getSLIAsString(2) + "])<<16 | uint32(buf[" + sm.getSLIAsString(3) + "])<<24\n")
				sm.LSI += 4
			case "int32":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = int32(buf[" + sm.getSLIAsString(0) + "]) | int32(buf[" + sm.getSLIAsString(1) + "])<<8 | int32(buf[" + sm.getSLIAsString(2) + "])<<16 | int32(buf[" + sm.getSLIAsString(3) + "])<<24\n")
				sm.LSI += 4
			case "uint64":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = uint64(buf[" + sm.getSLIAsString(0) + "]) | uint64(buf[" + sm.getSLIAsString(1) + "])<<8 | uint64(buf[" + sm.getSLIAsString(2) + "])<<16 | uint64(buf[" + sm.getSLIAsString(3) + "])<<24 | uint64(buf[" + sm.getSLIAsString(4) + "])<<32 | uint64(buf[" + sm.getSLIAsString(5) + "])<<40 | uint64(buf[" + sm.getSLIAsString(6) + "])<<48 | uint64(buf[" + sm.getSLIAsString(7) + "])<<56\n")
				sm.LSI += 8
			case "int64":
				sm.Decoder.WriteString("	" + sm.FRN + in + " = int64(buf[" + sm.getSLIAsString(0) + "]) | int64(buf[" + sm.getSLIAsString(1) + "])<<8 | int64(buf[" + sm.getSLIAsString(2) + "])<<16 | int64(buf[" + sm.getSLIAsString(3) + "])<<24 | int64(buf[" + sm.getSLIAsString(4) + "])<<32 | int64(buf[" + sm.getSLIAsString(5) + "])<<40 | int64(buf[" + sm.getSLIAsString(6) + "])<<48 | int64(buf[" + sm.getSLIAsString(7) + "])<<56\n")
				sm.LSI += 8
			case "string":
				sm.Decoder.WriteString("	add = uint32(buf[" + sm.getSLIAsString(0) + "]) | uint32(buf[" + sm.getSLIAsString(1) + "])<<8 | uint32(buf[" + sm.getSLIAsString(2) + "])<<16 | uint32(buf[" + sm.getSLIAsString(3) + "])<<24\n")
				sm.Decoder.WriteString("	ln = uint32(buf[" + sm.getSLIAsString(4) + "]) | uint32(buf[" + sm.getSLIAsString(5) + "])<<8 | uint32(buf[" + sm.getSLIAsString(6) + "])<<16 | uint32(buf[" + sm.getSLIAsString(7) + "])<<24\n")
				sm.Decoder.WriteString("	" + sm.FRN + in + " = string(buf[add : add+ln])\n")
				sm.LSI += 8
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	sm.Decoder.WriteString("\n	return\n")

	return
}

func (sm *syllabMaker) makeEncoder() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrSyllabNestedStruct
	}

	// reset sm.LSI to zero!
	sm.LSI = 0

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {
				// TODO::: Slice generator
			} else {
				// Get array len
				var len, err = strconv.ParseUint(t.Len.(*ast.BasicLit).Value, 10, 64)
				if err != nil {
					return ErrSyllabArrayLen
				}

				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					return ErrSyllabFieldType
				case "bool":
					sm.LSI += len
				case "byte":
					sm.Encoder.WriteString("	copy(b[" + sm.getSLIAsString(0) + ":], " + sm.FRN + in + "[:])\n")
					sm.LSI += len
				case "uint8":
					sm.LSI += len
				case "int8":
					sm.LSI += len
				case "uint16":
					// for i:= 0; i<len; i++ {
					// 	sm.Encoder.WriteString("	b["+strconv.FormatUint(sm.LSI+i)+"] = "+sm.FRN+in+"["+strconv.FormatUint(i)+"];")
					// }
					// sm.Encoder.WriteString("\n")
					sm.LSI += len * 2
				case "int16":
					sm.LSI += len * 2
				case "uint32":
					sm.LSI += len * 4
				case "int32":
					sm.LSI += len * 4
				case "uint64":
					sm.LSI += len * 8
				case "int64":
					sm.LSI += len * 8
				case "string":
				default:
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}
			}
		case *ast.StructType:
			var tmp = sm.FRN
			sm.FRN += in + "."
			sm.RTN = t.Fields.List[0].Names[0].Name
			err = sm.makeEncoder()
			sm.FRN = tmp
		case *ast.FuncType:
			return ErrSyllabFieldType
		case *ast.InterfaceType:
			return ErrSyllabFieldType
		case *ast.MapType:

		case *ast.ChanType:
			return ErrSyllabFieldType
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return ErrSyllabFieldType
			case "bool":
				sm.Encoder.WriteString("	if " + sm.FRN + in + " {\n	b[" + sm.getSLIAsString(0) + "] = 1\n	}\n")
				sm.LSI++
			case "byte":
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = " + sm.FRN + in + "\n")
				sm.LSI++
			case "uint8", "int8":
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = byte(" + sm.FRN + in + ")\n")
				sm.LSI++
			case "uint16", "int16":
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = byte(" + sm.FRN + in + ")\n	b[" + sm.getSLIAsString(1) + "] = byte(" + sm.FRN + in + " >> 8)\n")
				sm.LSI += 2
			case "uint32", "int32":
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = byte(" + sm.FRN + in + ")\n	b[" + sm.getSLIAsString(1) + "] = byte(" + sm.FRN + in + " >> 8)\n	b[" + sm.getSLIAsString(2) + "] = byte(" + sm.FRN + in + " >> 16)\n	b[" + sm.getSLIAsString(3) + "] = byte(" + sm.FRN + in + " >> 24)\n")
				sm.LSI += 4
			case "uint64", "int64":
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = byte(" + sm.FRN + in + ")\n	b[" + sm.getSLIAsString(1) + "] = byte(" + sm.FRN + in + " >> 8)\n	b[" + sm.getSLIAsString(2) + "] = byte(" + sm.FRN + in + " >> 16)\n	b[" + sm.getSLIAsString(3) + "] = byte(" + sm.FRN + in + " >> 24)\n	b[" + sm.getSLIAsString(4) + "] = byte(" + sm.FRN + in + " >> 32)\n	b[" + sm.getSLIAsString(5) + "] = byte(" + sm.FRN + in + " >> 40)\n	b[" + sm.getSLIAsString(6) + "] = byte(" + sm.FRN + in + " >> 48)\n	b[" + sm.getSLIAsString(7) + "] = byte(" + sm.FRN + in + " >> 56)\n")
				sm.LSI += 8
			case "string":
				sm.Encoder.WriteString("	ln = len(" + sm.FRN + in + ")\n")
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(0) + "] = byte(hsi)\n	b[" + sm.getSLIAsString(1) + "] = byte(hsi >> 8)\n	b[" + sm.getSLIAsString(2) + "] = byte(hsi >> 16)\n	b[" + sm.getSLIAsString(3) + "] = byte(hsi >> 24)\n")
				sm.Encoder.WriteString("	b[" + sm.getSLIAsString(4) + "] = byte(ln)\n	b[" + sm.getSLIAsString(5) + "] = byte(ln >> 8)\n	b[" + sm.getSLIAsString(6) + "] = byte(ln >> 16)\n	b[" + sm.getSLIAsString(7) + "] = byte(ln >> 24)\n")
				sm.Encoder.WriteString("	copy(b[hsi:], " + sm.FRN + in + "[:])\n	hsi += ln\n")
				sm.LSI += 8
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	sm.Encoder.WriteString("\n	return\n")

	return
}

func (sm *syllabMaker) makeStackHeapSize() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrSyllabNestedStruct
	}

	// reset sm.LSI to zero!
	sm.LSI = 0

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {
				// In any case we need 8 byte for address and len of array!
				sm.LSI += 8

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					return ErrSyllabFieldType
				case "bool", "byte", "uint8", "int8":
					sm.HeapSize.WriteString("	ln += len(" + sm.FRN + in + ")\n")
				case "uint16", "int16":
					sm.HeapSize.WriteString("	ln += len(" + sm.FRN + in + ")*2\n")
				case "uint32", "int32":
					sm.HeapSize.WriteString("	ln += len(" + sm.FRN + in + ")*4\n")
				case "uint64", "int64":
					sm.HeapSize.WriteString("	ln += len(" + sm.FRN + in + ")*8\n")
				case "string":
					sm.HeapSize.WriteString("	for i:=0; i<len(" + sm.FRN + in + "); i++ {\n")
					sm.HeapSize.WriteString("		ln += len(" + sm.FRN + in + "[i])\n")
					sm.HeapSize.WriteString("	}\n")
				default:
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}
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
					return ErrSyllabFieldType
				case "bool", "byte", "uint8", "int8":
					sm.LSI += ln
				case "uint16", "int16":
					sm.LSI += ln * 2
				case "uint32", "int32":
					sm.LSI += ln * 4
				case "uint64", "int64":
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
			err = sm.makeEncoder()
			sm.FRN = tmp
		case *ast.FuncType:
			return ErrSyllabFieldType
		case *ast.InterfaceType:
			return ErrSyllabFieldType
		case *ast.MapType:

		case *ast.ChanType:
			return ErrSyllabFieldType
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return ErrSyllabFieldType
			case "bool", "byte", "uint8", "int8":
				sm.LSI++
			case "uint16", "int16":
				sm.LSI += 2
			case "uint32", "int32":
				sm.LSI += 4
			case "uint64", "int64":
				sm.LSI += 8
			case "string":
				sm.HeapSize.WriteString("	ln += len(" + sm.FRN + in + ")\n")
				sm.LSI += 8
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	sm.Encoder.WriteString("\n	return\n")

	return
}

func (sm *syllabMaker) getSLIAsString(plus uint64) (s string) {
	// TODO::: Improve below line!
	s = strconv.FormatUint(sm.LSI+plus, 10)
	return
}
