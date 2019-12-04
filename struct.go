// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package astinspector

import (
	"go/ast"
	"go/token"
)

// StructList is a list of struct.
type StructList []Struct

// Len says the size of the slice.
func (sl StructList) Len() int {
	return len(sl)
}

// FindByName return the struct matching with the given name.
func (sl StructList) FindByName(name string) Struct {
	for _, s := range sl {
		if v := StructByName(s.AstTypeSpec(), name); v != nil {
			return v
		}
	}
	return nil
}

// First return the first element of the slice.
func (sl StructList) First() Struct {
	if sl.Len() > 0 {
		return sl[0]
	}
	return nil
}

// Names return the names of the structs.
func (sl StructList) Names() (out []string) {
	for _, s := range sl {
		out = append(out, s.Name())
	}
	return
}

// Struct abstract an ast.TypeSpec
type Struct interface {
	Name() string
	Fields() []Field
	AddField(name, typ string) Field
	AstTypeSpec() *ast.TypeSpec
	IsValid() bool
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
	return Structs(node, name).First()
}

// Structs return a list of structs defined in the given node.
func Structs(node ast.Node, names ...string) (out StructList) {
	ast.Inspect(node, func(n ast.Node) bool {
		switch d := n.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				break
			}

			for i := range d.Specs {
				ts := d.Specs[i].(*ast.TypeSpec)
				if len(names) > 0 {
					for _, name := range names {
						if ts.Name.Name == name {
							out = append(out, &iStruct{ts})
							break
						}
					}
				} else {
					out = append(out, &iStruct{ts})
				}
			}
			return false
		}
		return true
	})

	return
}
