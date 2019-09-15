// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package astinspector

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// File abstract an ast.File.
type File interface {
	StructByName(name string) Struct
	PackageName() string
	Structs(...string) StructList
}

// iFile
type iFile struct {
	*ast.File
	fset *token.FileSet
}

// StructByName get informations of an structure by his name.
func (f *iFile) StructByName(name string) Struct {
	return StructByName(f.File, name)
}

// PackageName return the name of the package.
func (f *iFile) PackageName() string {
	return f.Name.String()
}

// Structs return list of structure.
func (f *iFile) Structs(names ...string) StructList {
	return Structs(f.File, names...)
}

// ParseFile parse the given filename.
func ParseFile(filename string) File {
	f := &iFile{
		fset: token.NewFileSet(),
	}

	f.File, _ = parser.ParseFile(f.fset, filename, nil, parser.ParseComments)
	return f
}
