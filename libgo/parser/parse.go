/* For license and copyright information please see the LEGAL file in the code repository */

package parser

import (
	"go/ast"
)

// https://github.com/golang/mock/blob/main/mockgen/parse.go
// https://github.com/go-gad/sal/blob/master/looker/looker.go

// File :
type File struct {
	Name   string
	Data   []byte
	Parsed *ast.File
}

// Import :
type Import struct {
	UsageName    string
	PackageName  string
	DependencyID [16]byte //
	FSPath       string   // Folder location in FileSystems
	File         *File
	ImportSpec   *ast.ImportSpec
}

// Function store parsed data about logic Function!
type Function struct {
	Name      string
	Comment   string
	Parameter *Type // ChaparKhane just support one variable in Function input!
	Result    *Type // ChaparKhane just support one variable in Function output!
	Err       *Type // ChaparKhane just support one error in Function output!
	File      *File
	Decl      *ast.FuncDecl
}

// Type :
type Type struct {
	Name      string
	ID        int
	Package   *Import // If nil means local package not imported!
	Type      string  // struct: embedded struct in this struct.
	Len       uint64  // Use in Array, Slice, Map, ...
	Exported  bool
	Pointer   bool
	InnerType []*Type
	Tags      map[string]string
	Comment   string
	File      *File
}
