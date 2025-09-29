package parser

import (
	"go/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasrseDomain(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected []Entity
	}{
		{
			name: "single struct with primitive fields",
			source: `
				package test

				type User struct {
					ID   int
					Name string
				}
			`,
			expected: []Entity{
				{
					Name: "User",
					Fields: []Field{
						{Name: "ID", Type: "int"},
						{Name: "Name", Type: "string"},
					},
				},
			},
		},
		{
			name: "struct with slice and selector type",
			source: `
				package test

				import "time"

				type Post struct {
					Tags []string
					CreatedAt time.Time
				}
			`,
			expected: []Entity{
				{
					Name: "Post",
					Fields: []Field{
						{Name: "Tags", Type: "[]string"},
						{Name: "CreatedAt", Type: "time.Time"},
					},
				},
			},
		},
		{
			name: "multiple structs in one file",
			source: `
				package test

				type A struct {
					Value int
				}

				type B struct {
					Name string
				}
			`,
			expected: []Entity{
				{
					Name: "A",
					Fields: []Field{
						{Name: "Value", Type: "int"},
					},
				},
				{
					Name: "B",
					Fields: []Field{
						{Name: "Name", Type: "string"},
					},
				},
			},
		},
		{
			name: "struct with no fields",
			source: `
				package test

				type Empty struct {
				}
			`,
			expected: []Entity{
				{Name: "Empty"},
			},
		},
		{
			name: "nested slice of selector type",
			source: `
				package test

				import "time"

				type Schedule struct {
					Dates [][]time.Time
				}
			`,
			expected: []Entity{
				{
					Name: "Schedule",
					Fields: []Field{
						{Name: "Dates", Type: "[][]time.Time"},
					},
				},
			},
		},
		{
			name: "pointer and map fields (unsupported types return empty string)",
			source: `
				package test

				type Config struct {
					Parent *Config
					Meta   map[string]int
				}
			`,
			expected: []Entity{
				{
					Name: "Config",
					Fields: []Field{
						{Name: "Parent", Type: ""},
						{Name: "Meta", Type: ""},
					},
				},
			},
		},
		{
			name: "anonymous field (embedding)",
			source: `
				package test

				type Base struct {
					ID int
				}

				type Derived struct {
					Base
					Name string
				}
			`,
			expected: []Entity{
				{
					Name: "Base",
					Fields: []Field{
						{Name: "ID", Type: "int"},
					},
				},
				{
					Name: "Derived",
					Fields: []Field{
						// Anonymous fields appear in AST with no Names
						{Name: "Base", Type: "Base"},
						{Name: "Name", Type: "string"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			filePath := filepath.Join(tmpDir, "domain.go")
			err := os.WriteFile(filePath, []byte(tt.source), 0644)
			assert.NoError(t, err)

			entities, err := ParseDomain(filePath)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, entities)
		})
	}
}

func TestExprToString(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected string
	}{
		// Identifiers
		{"int", "int", "int"},
		{"string", "string", "string"},
		{"bool", "bool", "bool"},
		{"float64", "float64", "float64"},

		// Arrays
		{"slice of ident", "[]Post", "[]Post"},
		{"slice of slice", "[][]string", "[][]string"},
		{"slice of selector", "[]time.Time", "[]time.Time"},

		// Selectors
		{"selector expr", "time.Time", "time.Time"},
		{"nested selector", "pkg.sub.Type", "pkg.sub.Type"},

		// Unsupported (should return "")
		{"map type", "map[string]int", ""},
		{"pointer type", "*User", ""},
		{"struct type literal", "struct { ID int }", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.ParseExpr(tt.expr)
			assert.NoError(t, err)

			got := exprToString(expr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
