// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package astinspector

import (
	"go/ast"
	"go/token"
	"strings"
)

// Struct abstract an ast.TypeSpec
type Struct interface {
	Name() string
	Fields() []Field
	AddField(name, typ string) Field
	AstTypeSpec() *ast.TypeSpec
}

type iStruct struct {
	*ast.TypeSpec
}

// AstTypeSpec return *ast.TypeSpec.
func (s *iStruct) AstTypeSpec() *ast.TypeSpec {
	return s.TypeSpec
}

// IsValid says if the struct is valid.
func (s *iStruct) IsValid() bool {
	return s.TypeSpec != nil
}

// Name return the name of the struct
func (s *iStruct) Name() string {
	return s.TypeSpec.Name.String()
}

// Fields return the list of fields.
func (s *iStruct) Fields() (fl []Field) {
	st := s.Type.(*ast.StructType)
	for k := range st.Fields.List {
		fl = append(fl, &iField{st.Fields.List[k]})
	}
	return
}

// AddField add a field to the structure.
func (s *iStruct) AddField(name string, t string) Field {
	f := ast.Field{
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type:  ast.NewIdent(t),
	}

	// Adding to the fields
	st := s.Type.(*ast.StructType)
	st.Fields.List = append(st.Fields.List, &f)

	return &iField{&f}
}

// StructByName return a struct composed of *ast.TypeSpec.
func StructByName(node ast.Node, name string) Struct {
	var s *iStruct

	ast.Inspect(node, func(n ast.Node) bool {
		switch d := n.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				break
			}

			for i := range d.Specs {
				ts := d.Specs[i].(*ast.TypeSpec)

				if strings.EqualFold(ts.Name.Name, name) {
					s = &iStruct{ts}

					return false
				}
			}
		}

		return true
	})
	return s
}
