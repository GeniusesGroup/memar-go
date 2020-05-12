/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"

	"../assets"
)

/*
Before pass file to safe||unsafe function, dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!

func ({{DesireName}} *{{DesireType}}) syllabDecoder(buf []byte) (err error) {
	return nil
}

func ({{DesireName}} *{{DesireType}}) syllabEncoder(buf []byte) (err error) {
	return nil
}
*/

// CompleteEncoderMethodSafe use to update given go files and complete syllab encoder to any struct type in it!
// It will overwrite given file slice! If you need it clone it before pass it here!
func CompleteEncoderMethodSafe(file *assets.File) (err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", file.DataString, parser.ParseComments)
	if err != nil {
		return
	}

	var fileTypes = map[string]*ast.TypeSpec{}
	var cpyFile = []*copyToFileReq{}

	// find syllabDecoder || syllabEncoder method
	for _, decl := range fileParsed.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range d.Specs {
				switch gd := gDecl.(type) {
				case *ast.TypeSpec:
					fileTypes[gd.Name.Name] = gd
				}
			}
		case *ast.FuncDecl:
			// Just needed methods!
			if d.Name.Name == "syllabDecoder" {
				var data = syllabMaker{
					RN:    d.Recv.List[0].Names[0].Name,
					FRN:   d.Recv.List[0].Names[0].Name + ".",
					RTN:   d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name,
					Types: fileTypes,
				}
				err = data.makeSyllabDecoderSafe()
				cpyFile = append(cpyFile, &copyToFileReq{data.GData, int(d.Body.Lbrace), int(d.Body.Rbrace)})
			} else if d.Name.Name == "syllabEncoder" {
				var data = syllabMaker{
					RN:    d.Recv.List[0].Names[0].Name,
					FRN:   d.Recv.List[0].Names[0].Name + ".",
					RTN:   d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name,
					Types: fileTypes,
				}
				err = data.makeSyllabEncoderSafe()
				if data.HeapCreated {
					data.GData = "	var lhi int = " + strconv.Itoa(data.LSI) + " // Heap start index\n" + data.GData
				}
				cpyFile = append(cpyFile, &copyToFileReq{data.GData, int(d.Body.Lbrace), int(d.Body.Rbrace)})
			}
		}
	}

	copyToFile(file, cpyFile)
	file.Status = assets.StateChanged
	file.Data = []byte(file.DataString)
	return
}

type syllabMaker struct {
	RN          string                   // Receiver Name
	FRN         string                   // Flat Receiver Name e.g. req.Time.
	RTN         string                   // Receiver Type Name
	Types       map[string]*ast.TypeSpec // All types
	GData       string                   // Generated Data
	LSI         int                      // Last Stack Index
	HeapCreated bool
}

func (sm *syllabMaker) makeSyllabDecoderSafe() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrNeededTypeNotExist
	}

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {

			} else {
				// Get array len
				var len, err = strconv.Atoi(t.Len.(*ast.BasicLit).Value)
				if err != nil {
					return ErrArrayLenNotSupported
				}

				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					return ErrTypeIncludeIllegalChild
				case "bool":
				case "byte":
					sm.GData += "	copy(" + sm.FRN + in + "[:], buf[" + strconv.Itoa(sm.LSI) + ":])\n"
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
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}

				sm.LSI += len
			}
		case *ast.StructType:
			var tmp = sm.FRN
			sm.FRN += in + "."
			sm.RTN = t.Fields.List[0].Names[0].Name
			err = sm.makeSyllabDecoderSafe()
			sm.FRN = tmp
		case *ast.FuncType:
			return ErrTypeIncludeIllegalChild
		case *ast.InterfaceType:
			return ErrTypeIncludeIllegalChild
		case *ast.MapType:

		case *ast.ChanType:
			return ErrTypeIncludeIllegalChild
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return ErrTypeIncludeIllegalChild
			case "bool":
				sm.GData += "	" + sm.FRN + in + " = buf[" + strconv.Itoa(sm.LSI) + "] == 1 \n"
				sm.LSI++
			case "byte":
				sm.GData += "	" + sm.FRN + in + " = buf[" + strconv.Itoa(sm.LSI) + "]\n"
				sm.LSI++
			case "uint8":
				sm.GData += "	" + sm.FRN + in + " = uint8(buf[" + strconv.Itoa(sm.LSI) + "])\n"
				sm.LSI++
			case "int8":
				sm.GData += "	" + sm.FRN + in + " = int8(buf[" + strconv.Itoa(sm.LSI) + "])\n"
				sm.LSI++
			case "uint16":
				sm.GData += "	" + sm.FRN + in + " = uint16(buf[" + strconv.Itoa(sm.LSI) + "]) | uint16(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8\n"
				sm.LSI += 2
			case "int16":
				sm.GData += "	" + sm.FRN + in + " = int16(buf[" + strconv.Itoa(sm.LSI) + "]) | int16(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8\n"
				sm.LSI += 2
			case "uint32":
				sm.GData += "	" + sm.FRN + in + " = uint32(buf[" + strconv.Itoa(sm.LSI) + "]) | uint32(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8 | uint32(buf[" + strconv.Itoa(sm.LSI+2) + "])<<16 | uint32(buf[" + strconv.Itoa(sm.LSI+3) + "])<<24\n"
				sm.LSI += 4
			case "int32":
				sm.GData += "	" + sm.FRN + in + " = int32(buf[" + strconv.Itoa(sm.LSI) + "]) | int32(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8 | int32(buf[" + strconv.Itoa(sm.LSI+2) + "])<<16 | int32(buf[" + strconv.Itoa(sm.LSI+3) + "])<<24\n"
				sm.LSI += 4
			case "uint64":
				sm.GData += "	" + sm.FRN + in + " = uint64(buf[" + strconv.Itoa(sm.LSI) + "]) | uint64(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8 | uint64(buf[" + strconv.Itoa(sm.LSI+2) + "])<<16 | uint64(buf[" + strconv.Itoa(sm.LSI+3) + "])<<24 | uint64(buf[" + strconv.Itoa(sm.LSI+4) + "])<<32 | uint64(buf[" + strconv.Itoa(sm.LSI+5) + "])<<40 | uint64(buf[" + strconv.Itoa(sm.LSI+6) + "])<<48 | uint64(buf[" + strconv.Itoa(sm.LSI+7) + "])<<56\n"
				sm.LSI += 8
			case "int64":
				sm.GData += "	" + sm.FRN + in + " = int64(buf[" + strconv.Itoa(sm.LSI) + "]) | int64(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8 | int64(buf[" + strconv.Itoa(sm.LSI+2) + "])<<16 | int64(buf[" + strconv.Itoa(sm.LSI+3) + "])<<24 | int64(buf[" + strconv.Itoa(sm.LSI+4) + "])<<32 | int64(buf[" + strconv.Itoa(sm.LSI+5) + "])<<40 | int64(buf[" + strconv.Itoa(sm.LSI+6) + "])<<48 | int64(buf[" + strconv.Itoa(sm.LSI+7) + "])<<56\n"
				sm.LSI += 8
			case "string":
				sm.GData += "	var " + in + "Add = uint32(buf[" + strconv.Itoa(sm.LSI) + "]) | uint32(buf[" + strconv.Itoa(sm.LSI+1) + "])<<8 | uint32(buf[" + strconv.Itoa(sm.LSI+2) + "])<<16 | uint32(buf[" + strconv.Itoa(sm.LSI+3) + "])<<24\n"
				sm.GData += "	var " + in + "Len = uint32(buf[" + strconv.Itoa(sm.LSI+4) + "]) | uint32(buf[" + strconv.Itoa(sm.LSI+5) + "])<<8 | uint32(buf[" + strconv.Itoa(sm.LSI+6) + "])<<16 | uint32(buf[" + strconv.Itoa(sm.LSI+7) + "])<<24\n"
				sm.GData += "	" + sm.FRN + in + " = string(buf[" + in + "Add:" + in + "Len])\n"
				sm.LSI += 8
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	sm.GData += "\n	return"

	return
}

func (sm *syllabMaker) makeSyllabEncoderSafe() (err error) {
	// Check needed type exist!!
	t, found := sm.Types[sm.RTN]
	if !found {
		return ErrNeededTypeNotExist
	}

	var in string
	for _, c := range t.Type.(*ast.StructType).Fields.List {
		in = c.Names[0].Name
		switch t := c.Type.(type) {
		case *ast.ArrayType:
			// Check array is slice?
			if t.Len == nil {

			} else {
				// Get array len
				var len, err = strconv.Atoi(t.Len.(*ast.BasicLit).Value)
				if err != nil {
					return ErrArrayLenNotSupported
				}

				if t.Len.(*ast.BasicLit).Kind == token.STRING {
					// Its common to use const to indicate number of array like in IP type as [16]byte!
					// TODO::: get related const value by its name as t.Len.(*ast.BasicLit).Value
				}

				switch t.Elt.(*ast.Ident).Name {
				case "int", "uint":
					return ErrTypeIncludeIllegalChild
				case "bool":
				case "byte":
					sm.GData += "	copy(buf[" + strconv.Itoa(sm.LSI) + ":], " + sm.FRN + in + "[:])\n"
					// TODO::: Performance check assignment vs copy??
					// for i:= 0; i<len; i++ {
					// 	sm.GData += "	buf["+strconv.Itoa(sm.LSI+i)+"] = "+sm.FRN+in+"["+strconv.Itoa(i)+"];"
					// }
					// sm.GData += "\n"
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
					// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
				}

				sm.LSI += len
			}
		case *ast.StructType:
			var tmp = sm.FRN
			sm.FRN += in + "."
			sm.RTN = t.Fields.List[0].Names[0].Name
			err = sm.makeSyllabEncoderSafe()
			sm.FRN = tmp
		case *ast.FuncType:
			return ErrTypeIncludeIllegalChild
		case *ast.InterfaceType:
			return ErrTypeIncludeIllegalChild
		case *ast.MapType:

		case *ast.ChanType:
			return ErrTypeIncludeIllegalChild
		case *ast.Ident:
			switch t.Name {
			case "int", "uint":
				return ErrTypeIncludeIllegalChild
			case "bool":
				sm.GData += "	if " + sm.FRN + in + " {\n	buf[" + strconv.Itoa(sm.LSI) + "] = 1\n	}\n"
				sm.LSI++
			case "byte":
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = " + sm.FRN + in + "\n"
				sm.LSI++
			case "uint8", "int8":
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = byte(" + sm.FRN + in + ")\n"
				sm.LSI++
			case "uint16", "int16":
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = byte(" + sm.FRN + in + ")\n	buf[" + strconv.Itoa(sm.LSI+1) + "] = byte(" + sm.FRN + in + " >> 8)\n"
				sm.LSI += 2
			case "uint32", "int32":
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = byte(" + sm.FRN + in + ")\n	buf[" + strconv.Itoa(sm.LSI+1) + "] = byte(" + sm.FRN + in + " >> 8)\n	buf[" + strconv.Itoa(sm.LSI+2) + "] = byte(" + sm.FRN + in + " >> 16)\n	buf[" + strconv.Itoa(sm.LSI+3) + "] = byte(" + sm.FRN + in + " >> 24)\n"
				sm.LSI += 4
			case "uint64", "int64":
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = byte(" + sm.FRN + in + ")\n	buf[" + strconv.Itoa(sm.LSI+1) + "] = byte(" + sm.FRN + in + " >> 8)\n	buf[" + strconv.Itoa(sm.LSI+2) + "] = byte(" + sm.FRN + in + " >> 16)\n	buf[" + strconv.Itoa(sm.LSI+3) + "] = byte(" + sm.FRN + in + " >> 24)\n	buf[" + strconv.Itoa(sm.LSI+4) + "] = byte(" + sm.FRN + in + " >> 32)\n	buf[" + strconv.Itoa(sm.LSI+5) + "] = byte(" + sm.FRN + in + " >> 40)\n	buf[" + strconv.Itoa(sm.LSI+6) + "] = byte(" + sm.FRN + in + " >> 48)\n	buf[" + strconv.Itoa(sm.LSI+7) + "] = byte(" + sm.FRN + in + " >> 56)\n"
				sm.LSI += 8
			case "string":
				if !sm.HeapCreated {
					sm.GData += "	var ln int // len of string, slices, maps, ...\n"
					sm.HeapCreated = true
				}
				sm.GData += "	ln = len(" + sm.FRN + in + ")\n"
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI) + "] = byte(lhi)\n	buf[" + strconv.Itoa(sm.LSI+1) + "] = byte(lhi >> 8)\n	buf[" + strconv.Itoa(sm.LSI+2) + "] = byte(lhi >> 16)\n	buf[" + strconv.Itoa(sm.LSI+3) + "] = byte(lhi >> 24)\n"
				sm.GData += "	buf[" + strconv.Itoa(sm.LSI+4) + "] = byte(ln)\n	buf[" + strconv.Itoa(sm.LSI+5) + "] = byte(ln >> 8)\n	buf[" + strconv.Itoa(sm.LSI+6) + "] = byte(ln >> 16)\n	buf[" + strconv.Itoa(sm.LSI+7) + "] = byte(ln >> 24)\n"
				sm.GData += "	copy(buf[lhi:], " + sm.FRN + in + "[:])\n	lhi += ln\n"
				sm.LSI += 8
			default:
				// TODO::: get related type by its name as t.Elt.(*ast.Ident).Name
			}
		case *ast.BasicLit:
			// fmt.Fprintf(os.Stderr, "BasicLit : %v\n", t.Kind)
		}
	}

	sm.GData += "\n	return"

	return
}

type copyToFileReq struct {
	Data  string
	Start int
	End   int
}

func copyToFile(file *assets.File, gData []*copyToFileReq) {
	var addedSize int
	for _, cpy := range gData {
		cpy.Start += addedSize
		cpy.End += addedSize

		addedSize += len(cpy.Data) - (cpy.End - cpy.Start) + 3 // 3 is for bracket, new line, and len!!

		file.DataString = file.DataString[:cpy.Start+1] + cpy.Data + file.DataString[cpy.End-2:]
	}
}
