// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package astinspector

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/fatih/structtag"
)

// Field abstract an ast.Field.
type Field interface {
	Name() string
	Type() string
	Tags() *structtag.Tags
}

type iField struct {
	*ast.Field
}

// Type return the type of the field.
func (f *iField) Type() string {
	return parseType(f.Field.Type)
}

// Name return the name of the field.
func (f *iField) Name() string {
	return f.Names[0].Name
}

// Tags return the taglist of the field.
func (f *iField) Tags() *structtag.Tags {
	if f.Tag == nil {
		t, _ := structtag.Parse("")
		return t
	}

	t, _ := structtag.Parse(strings.Trim(f.Tag.Value, "`"))
	return t
}

// parseType will parse the given type and return the representation as string.
func parseType(t ast.Expr) string {
	switch t := t.(type) {
	case *ast.BasicLit:
		return t.Kind.String()
	case *ast.Ident:
		return t.String()
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", t.X.(*ast.Ident).String(), t.Sel.String())
	case *ast.ArrayType:
		if t.Len == nil {
			return fmt.Sprintf("[]%s", parseType(t.Elt))
		}

		return fmt.Sprintf("[%s]%s", t.Len.(*ast.BasicLit).Value, parseType(t.Elt))
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", parseType(t.X))
	}

	return ""
}
