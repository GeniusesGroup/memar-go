/* For license and copyright information please see LEGAL file in repository */

package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
)

// Repository :
type Repository struct {
	Name         string
	ID           uint16                 // Each microservice has ID e.g. "service-name--000"
	DependencyID [16]byte               // DependencyID in SabzCity version control
	FSPath       string                 // Folder location in FileSystems
	Files        map[string]*File       // Name
	Imports      map[string]*Import     // UsageName
	Functions    map[string]*Function   // Name
	Types        map[string]*Type       // Name
	Dependencies map[string]*Repository // Name
}

// AddFile use to add file to Repository
func (r *Repository) AddFile(f *File) {
	r.Files[f.Name] = f
}

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
	DependencyID [16]byte // DependencyID in SabzCity version control
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

// Parse use to add new file & parsed FileData and add to repo object!
func (repo *Repository) Parse(FileName string, FileData []byte, FolderID uint16) (err error) {
	// Parsed FileData
	var (
		file    = File{Name: FileName, Data: FileData}
		fileSet = token.NewFileSet()
	)
	repo.Files[FileName] = &file
	repo.ID = FolderID

	// Just parsed needed file!
	if strings.HasSuffix(FileName, ".go") ||
		!strings.HasSuffix(FileName, "_test.go") {

		file.Parsed, err = parser.ParseFile(fileSet, "", FileData, parser.ParseComments)
		if err != nil {
			return err
		}

		// Set package name
		repo.Name = file.Parsed.Name.Name

		err = repo.parseFile(&file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) parseFile(file *File) (err error) {
	for _, imp := range file.Parsed.Imports {
		var impor = Import{
			FSPath:     imp.Path.Value[1 : len(imp.Path.Value)-1],
			ImportSpec: imp,
		}
		if imp.Name != nil {
			impor.UsageName = imp.Name.Name
			impor.PackageName = path.Base(imp.Path.Value[1 : len(imp.Path.Value)-1])
		} else {
			impor.UsageName = path.Base(imp.Path.Value[1 : len(imp.Path.Value)-1])
			impor.PackageName = path.Base(imp.Path.Value[1 : len(imp.Path.Value)-1])
		}

		repo.Imports[impor.UsageName] = &impor
	}

	for _, decl := range file.Parsed.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, gDecl := range d.Specs {
				switch gd := gDecl.(type) {
				case *ast.ImportSpec:
					// Check this before in file.Parsed.Imports!!!! WHY go WHY!!!? duplicate data!!!???
				case *ast.ValueSpec:
					// Don't need this so far!
				case *ast.TypeSpec:
					t := repo.parseType(gd.Type)
					t.Name = gd.Name.Name
					t.Exported = gd.Name.IsExported()
					t.File = file

					repo.Types[t.Name] = &t
				}
			}

		case *ast.FuncDecl:
			// Just exported function use in chaparkhane generation!
			if !d.Name.IsExported() {
				continue
			}

			function := Function{
				Name: d.Name.Name,
				File: file,
				Decl: d,
			}

			if d.Type != nil && d.Type.Params != nil {
				if len(d.Type.Params.List) != 2 {
					continue
				}
				fp := repo.parseType(d.Type.Params.List[0].Type)
				fp.Name = d.Type.Params.List[0].Names[0].Name
				fp.Exported = d.Type.Params.List[0].Names[0].IsExported()
				fp.File = file
				function.Parameter = &fp
			}

			if d.Type != nil && d.Type.Results != nil {
				if len(d.Type.Params.List) > 2 {
					continue
				}

				for _, rField := range d.Type.Results.List {
					for _, name := range rField.Names {
						fr := repo.parseType(rField.Type)
						fr.Name = name.Name
						fr.Exported = name.IsExported()
						fr.File = file
						if fr.Type == "error" {
							function.Err = &fr
						} else {
							function.Result = &fr
						}
					}
				}
			}

			repo.Functions[function.Name] = &function
		}
	}

	return nil
}

func (repo *Repository) parseType(expr ast.Expr) (ft Type) {
	switch v := expr.(type) {
	case *ast.StarExpr:
		ft = repo.parseType(v.X)
		ft.Pointer = true
	case *ast.SelectorExpr:
		if imp, ok := repo.Imports[v.X.(*ast.Ident).Name]; ok {
			ft.Package = imp
		}
		ft.Type = v.Sel.Name
	case *ast.ArrayType:
		//ft.Type = "[" + repo.parseType(v.Len).Type + "]" + repo.parseType(v.Elt).Type
	case *ast.MapType:
		//"map[" + parseType(v.Key) + "]" + parseType(v.Value)
	case *ast.InterfaceType:
		//"interface{}"
		// interface type forbidden
	case *ast.Ident:
		if v.Obj != nil {
			switch obj := v.Obj.Decl.(type) {
			case *ast.TypeSpec:
				innerFT := repo.parseType(obj.Type)
				for _, each := range innerFT.InnerType {
					ft.InnerType = append(ft.InnerType, each)
				}
			}
		}
		ft.Type = v.Name
	case *ast.StructType:
		for _, field := range v.Fields.List {
			if field.Names == nil {
				// embedded struct forbidden
			}
			for _, innerField := range field.Names {
				innerFT := repo.parseType(field.Type)
				innerFT.Name = innerField.Name
				innerFT.ID = len(ft.InnerType)
				ft.InnerType = append(ft.InnerType, &innerFT)
			}
		}
	case *ast.BasicLit:
		// embedded basic type forbidden
		// type test struct {
		// 	 string
		// }
	}

	return ft
}
