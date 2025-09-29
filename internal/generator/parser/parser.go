package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Field struct {
	Name string
	Type string
}

type Entity struct {
	Name   string
	Fields []Field
	Rels   []Relationship
}

type Relationship struct {
	Type   string
	Target string
}

func ParseDomain(path string) ([]Entity, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var entities []Entity

	for _, decl := range node.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			entity := Entity{Name: ts.Name.Name}
			for _, field := range st.Fields.List {
				typeStr := exprToString(field.Type)

				var fieldName string
				if len(field.Names) > 0 {
					fieldName = field.Names[0].Name
				} else {
					// anonymous field (embedded type)
					fieldName = typeStr
				}

				f := Field{
					Name: fieldName,
					Type: typeStr,
				}
				entity.Fields = append(entity.Fields, f)
			}

			entities = append(entities, entity)
		}
	}

	return entities, nil
}

func exprToString(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name // e.g. "int", "string", "User"
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt) // e.g. []Post
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name // e.g. time.Time
	}
	return ""
}
