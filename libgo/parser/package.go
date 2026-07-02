/* For license and copyright information please see the LEGAL file in the code repository */

package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
)

// Package :
type Package struct {
	Name         string
	DependencyID [16]byte             //
	FSPath       string               // Folder location in FileSystems
	Files        map[string]*File     // Name
	Imports      map[string]*Import   // UsageName
	Functions    map[string]*Function // Name
	Types        map[string]*Type     // Name
	Dependencies map[string]*Package  // Name
}

// Init initialize the object.
func (p *Package) Init() {
	p.Files = make(map[string]*File)
	p.Imports = make(map[string]*Import)
	p.Functions = make(map[string]*Function)
	p.Types = make(map[string]*Type)
	p.Dependencies = make(map[string]*Package)
}

// AddFile use to add file to Package
func (p *Package) AddFile(f *File) {
	p.Files[f.Name] = f
}

// Parse use to add new file & parsed FileData and add to repo object!
func (p *Package) Parse(FileName string, FileData []byte) (err error) {
	// Parsed FileData
	var (
		file    = File{Name: FileName, Data: FileData}
		fileSet = token.NewFileSet()
	)
	p.Files[FileName] = &file

	// Just parsed needed file!
	if strings.HasSuffix(FileName, ".go") ||
		!strings.HasSuffix(FileName, "_test.go") {

		file.Parsed, err = parser.ParseFile(fileSet, "", FileData, parser.ParseComments)
		if err != nil {
			return
		}

		// Set package name
		p.Name = file.Parsed.Name.Name

		err = p.parseFile(&file)
		if err != nil {
			return
		}
	}

	return nil
}

func (p *Package) parseFile(file *File) (err error) {
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

		p.Imports[impor.UsageName] = &impor
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
					t := p.parseType(gd.Type)
					t.Name = gd.Name.Name
					t.Exported = gd.Name.IsExported()
					t.File = file

					p.Types[t.Name] = &t
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
				fp := p.parseType(d.Type.Params.List[0].Type)
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
						fr := p.parseType(rField.Type)
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

			p.Functions[function.Name] = &function
		}
	}

	return nil
}

func (p *Package) parseType(expr ast.Expr) (ft Type) {
	switch v := expr.(type) {
	case *ast.StarExpr:
		ft = p.parseType(v.X)
		ft.Pointer = true
	case *ast.SelectorExpr:
		if imp, ok := p.Imports[v.X.(*ast.Ident).Name]; ok {
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
				innerFT := p.parseType(obj.Type)
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
				innerFT := p.parseType(field.Type)
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
