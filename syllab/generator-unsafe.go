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

// CompleteEncoderMethodUnsafe use to update unsafely given go files and complete syllab encoder to any struct type in it!
func CompleteEncoderMethodUnsafe(ass *assets.File) (err error) {
	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	fileParsed, err = parser.ParseFile(fileSet, "", ass.DataString, parser.ParseComments)
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
				err = data.makeSyllabDecoderUnsafe()
				cpyFile = append(cpyFile, &copyToFileReq{data.GData, int(d.Body.Lbrace), int(d.Body.Rbrace)})
			} else if d.Name.Name == "syllabEncoder" {
				var data = syllabMaker{
					RN:    d.Recv.List[0].Names[0].Name,
					FRN:   d.Recv.List[0].Names[0].Name + ".",
					RTN:   d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name,
					Types: fileTypes,
				}
				err = data.makeSyllabEncoderUnsafe()
				if data.HeapCreated {
					data.GData += "	var lhi int = " + strconv.Itoa(data.LSI) + " // Heap start index\n"
				}
				cpyFile = append(cpyFile, &copyToFileReq{data.GData, int(d.Body.Lbrace), int(d.Body.Rbrace)})
			}
		}
	}
	copyToFile(ass, cpyFile)
	return
}

func (sm *syllabMaker) makeSyllabDecoderUnsafe() (err error) {
	// TODO:::

	sm.GData += "\n	return"

	return
}

func (sm *syllabMaker) makeSyllabEncoderUnsafe() (err error) {
	// TODO:::

	sm.GData += "\n	return"

	return
}
