/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"go/ast"
	"go/parser"
	"go/token"

	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	"memar/file"
	lang_p "memar/language/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/process/service"
	"memar/protocol"
	directory_p "memar/storage/directory/protocol"
	user_p "memar/user/protocol"
)

var CompleteGoSyllabMethodsService = completeGoSyllabMethodsService{
	Service: service.New(protocol.SyllabServiceCompleteGoSyllabMethods, "", datatype_p.LifeCycle_PreAlpha, 1605027929).
		SetDetail(lang_p.LanguageEnglish, domainEnglish, "Complete Go Syllab Methods",
			"",
			`Before request from or pass file to CompleteGoSyllabMethods(), It is better that developer respect protocol.Syllab interface and implements its empty methods!
Otherwise all method add to end of desire file!!`,
			[]string{"Syllab"}).
		SetAuthorization(operation_p.ActionType_Update, user_p.Type_App).Expired(0, ""),
}

type completeGoSyllabMethodsService struct {
	service.Service
}

// CompleteGoSyllabMethodsReq indicate generator behavior!
type CompleteGoSyllabMethodsReq struct {
	DIR         directory_p.Directory
	FilePath    string
	TypeNames   []string
	Safe        bool // true means copy data from given payload||buffer and don't use just point to it for decoding fields! buffer can GC after struct decode!
	ForceUpdate bool // true means delete exiting codes and update encoders && decoders codes anyway!
}

// CompleteGoSyllabMethods use to update given go files and complete Syllab encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func CompleteGoSyllabMethods(req CompleteGoSyllabMethodsReq) (err error_p.Error) {
	var goFile file_p.File
	goFile, err = req.DIR.FileByPath(req.FilePath)
	if err != nil || goFile == nil {
		return
	}

	var fileSet *token.FileSet = token.NewFileSet()
	var fileParsed *ast.File
	var goErr error
	fileParsed, goErr = parser.ParseFile(fileSet, "", goFile.Data().Marshal(), parser.ParseComments)
	if goErr != nil {
		return
	}

	var fileReplaces = make([]file.ReplaceReq, 0, 4)

	// find FromSyllab || ToSyllab method
	for _, decl := range fileParsed.Decls {
		switch declType := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range declType.Specs {
				switch gd := gDecl.(type) {
				case *ast.TypeSpec:
					for _, desireType := range req.TypeNames {
						if desireType == gd.Name.Name {
							// var sm = syllabMaker{
							// Options: req,
							// Types:   map[string]*ast.TypeSpec{},
							// }

						}
					}
				}
			}
		case *ast.FuncDecl:
			if declType.Recv != nil {
				switch declType.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					if declType.Name.Name == "CheckSyllab" || declType.Name.Name == "FromSyllab" ||
						declType.Name.Name == "ToSyllab" || declType.Name.Name == "Syllab_StackLength" ||
						declType.Name.Name == "Syllab_HeapLength" || declType.Name.Name == "Syllab_Length" {
						// TODO::: check by req.ForceUpdate and remove them from decl.Pos() to decl.End() if allow to update them
					}
				}
			}
		}
	}

	file.ReplaceLocation(goFile, fileReplaces)
	return
}
