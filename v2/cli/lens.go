// Copyright (c) 2023 - 2025 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"go/ast"
	"text/template"

	C "github.com/urfave/cli/v3"
)

const (
	keyLensDir         = "dir"
	keyVerbose         = "verbose"
	keyIncludeTestFile = "include-test-files"
	lensAnnotation     = "fp-go:Lens"
)

var (
	flagLensDir = &C.StringFlag{
		Name:  keyLensDir,
		Value: ".",
		Usage: "Directory to scan for Go files",
	}

	flagVerbose = &C.BoolFlag{
		Name:    keyVerbose,
		Aliases: []string{"v"},
		Value:   false,
		Usage:   "Enable verbose output",
	}

	flagIncludeTestFiles = &C.BoolFlag{
		Name:    keyIncludeTestFile,
		Aliases: []string{"t"},
		Value:   false,
		Usage:   "Include test files (*_test.go) when scanning for annotated types",
	}
)

// structInfo holds information about a struct that needs lens generation
type structInfo struct {
	Name           string
	TypeParams     string // e.g., "[T any]" or "[K comparable, V any]" - for type declarations
	TypeParamNames string // e.g., "[T]" or "[K, V]" - for type usage in function signatures
	Fields         []fieldInfo
	Imports        map[string]string // package path -> alias
}

// fieldInfo holds information about a struct field
type fieldInfo struct {
	Name         string
	TypeName     string
	BaseType     string // TypeName without leading * for pointer types
	IsOptional   bool   // true if field is a pointer or has json omitempty tag
	IsComparable bool   // true if the type is comparable (can use ==)
	IsEmbedded   bool   // true if this field comes from an embedded struct
}

// templateData holds data for template rendering
type templateData struct {
	PackageName string
	Structs     []structInfo
}

const lensStructTemplate = `
// {{.Name}}Lenses provides [lenses] for accessing fields of [{{.Name}}]
//
// [lenses]: __lens.Lens
type {{.Name}}Lenses{{.TypeParams}} struct {
	// mandatory fields
{{- range .Fields}}
	{{.Name}} __lens.Lens[{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
	// optional fields
{{- range .Fields}}
{{- if .IsComparable}}
	{{.Name}}O __lens_option.LensO[{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
{{- end}}
}

// {{.Name}}RefLenses provides [lenses] for accessing fields of [{{.Name}}] via a reference to [{{.Name}}]
//
//
// [lenses]: __lens.Lens
type {{.Name}}RefLenses{{.TypeParams}} struct {
	// mandatory fields
{{- range .Fields}}
	{{.Name}} __lens.Lens[*{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
	// optional fields
{{- range .Fields}}
{{- if .IsComparable}}
	{{.Name}}O __lens_option.LensO[*{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
{{- end}}
}

// {{.Name}}Prisms provides [prisms] for accessing fields of [{{.Name}}]
//
// [prisms]: __prism.Prism
type {{.Name}}Prisms{{.TypeParams}} struct {
{{- range .Fields}}
	{{.Name}} __prism.Prism[{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
}

// {{.Name}}RefPrisms provides [prisms] for accessing fields of [{{.Name}}] via a reference to [{{.Name}}]
//
// [prisms]: __prism.Prism
type {{.Name}}RefPrisms{{.TypeParams}} struct {
{{- range .Fields}}
	{{.Name}} __prism.Prism[*{{$.Name}}{{$.TypeParamNames}}, {{.TypeName}}]
{{- end}}
}
`

const lensConstructorTemplate = `
// Make{{.Name}}Lenses creates a new [{{.Name}}Lenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func Make{{.Name}}Lenses{{.TypeParams}}() {{.Name}}Lenses{{.TypeParamNames}} {
	// mandatory lenses
{{- range .Fields}}
	lens{{.Name}} := __lens.MakeLensWithName(
		func(s {{$.Name}}{{$.TypeParamNames}}) {{.TypeName}} { return s.{{.Name}} },
		func(s {{$.Name}}{{$.TypeParamNames}}, v {{.TypeName}}) {{$.Name}}{{$.TypeParamNames}} { s.{{.Name}} = v; return s },
		"{{$.Name}}{{$.TypeParamNames}}.{{.Name}}",
	)
{{- end}}
	// optional lenses
{{- range .Fields}}
{{- if .IsComparable}}
	lens{{.Name}}O := __lens_option.FromIso[{{$.Name}}{{$.TypeParamNames}}](__iso_option.FromZero[{{.TypeName}}]())(lens{{.Name}})
{{- end}}
{{- end}}
	return {{.Name}}Lenses{{.TypeParamNames}}{
		// mandatory lenses
{{- range .Fields}}
		{{.Name}}: lens{{.Name}},
{{- end}}
		// optional lenses
{{- range .Fields}}
{{- if .IsComparable}}
		{{.Name}}O: lens{{.Name}}O,
{{- end}}
{{- end}}
	}
}

// Make{{.Name}}RefLenses creates a new [{{.Name}}RefLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func Make{{.Name}}RefLenses{{.TypeParams}}() {{.Name}}RefLenses{{.TypeParamNames}} {
	// mandatory lenses
{{- range .Fields}}
{{- if .IsComparable}}
	lens{{.Name}} := __lens.MakeLensStrictWithName(
		func(s *{{$.Name}}{{$.TypeParamNames}}) {{.TypeName}} { return s.{{.Name}} },
		func(s *{{$.Name}}{{$.TypeParamNames}}, v {{.TypeName}}) *{{$.Name}}{{$.TypeParamNames}} { s.{{.Name}} = v; return s },
		"(*{{$.Name}}{{$.TypeParamNames}}).{{.Name}}",
	)
{{- else}}
	lens{{.Name}} := __lens.MakeLensRefWithName(
		func(s *{{$.Name}}{{$.TypeParamNames}}) {{.TypeName}} { return s.{{.Name}} },
		func(s *{{$.Name}}{{$.TypeParamNames}}, v {{.TypeName}}) *{{$.Name}}{{$.TypeParamNames}} { s.{{.Name}} = v; return s },
		"(*{{$.Name}}{{$.TypeParamNames}}).{{.Name}}",
	)
{{- end}}
{{- end}}
	// optional lenses
{{- range .Fields}}
{{- if .IsComparable}}
	lens{{.Name}}O := __lens_option.FromIso[*{{$.Name}}{{$.TypeParamNames}}](__iso_option.FromZero[{{.TypeName}}]())(lens{{.Name}})
{{- end}}
{{- end}}
	return {{.Name}}RefLenses{{.TypeParamNames}}{
		// mandatory lenses
{{- range .Fields}}
		{{.Name}}: lens{{.Name}},
{{- end}}
		// optional lenses
{{- range .Fields}}
{{- if .IsComparable}}
		{{.Name}}O: lens{{.Name}}O,
{{- end}}
{{- end}}
	}
}

// Make{{.Name}}Prisms creates a new [{{.Name}}Prisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func Make{{.Name}}Prisms{{.TypeParams}}() {{.Name}}Prisms{{.TypeParamNames}} {
{{- range .Fields}}
{{- if .IsComparable}}
	_fromNonZero{{.Name}} := __option.FromNonZero[{{.TypeName}}]()
	_prism{{.Name}} := __prism.MakePrismWithName(
		func(s {{$.Name}}{{$.TypeParamNames}}) __option.Option[{{.TypeName}}] { return _fromNonZero{{.Name}}(s.{{.Name}}) },
		func(v {{.TypeName}}) {{$.Name}}{{$.TypeParamNames}} {
			{{- if .IsEmbedded}}
			var result {{$.Name}}{{$.TypeParamNames}}
			result.{{.Name}} = v
			return result
			{{- else}}
			return {{$.Name}}{{$.TypeParamNames}}{ {{.Name}}: v }
			{{- end}}
		},
		"{{$.Name}}{{$.TypeParamNames}}.{{.Name}}",
	)
{{- else}}
	_prism{{.Name}} := __prism.MakePrismWithName(
		func(s {{$.Name}}{{$.TypeParamNames}}) __option.Option[{{.TypeName}}] { return __option.Some(s.{{.Name}}) },
		func(v {{.TypeName}}) {{$.Name}}{{$.TypeParamNames}} {
			{{- if .IsEmbedded}}
			var result {{$.Name}}{{$.TypeParamNames}}
			result.{{.Name}} = v
			return result
			{{- else}}
			return {{$.Name}}{{$.TypeParamNames}}{ {{.Name}}: v }
			{{- end}}
		},
		"{{$.Name}}{{$.TypeParamNames}}.{{.Name}}",
	)
{{- end}}
{{- end}}
	return {{.Name}}Prisms{{.TypeParamNames}} {
{{- range .Fields}}
		{{.Name}}: _prism{{.Name}},
{{- end}}
	}
}

// Make{{.Name}}RefPrisms creates a new [{{.Name}}RefPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func Make{{.Name}}RefPrisms{{.TypeParams}}() {{.Name}}RefPrisms{{.TypeParamNames}} {
{{- range .Fields}}
{{- if .IsComparable}}
	_fromNonZero{{.Name}} := __option.FromNonZero[{{.TypeName}}]()
	_prism{{.Name}} := __prism.MakePrismWithName(
		func(s *{{$.Name}}{{$.TypeParamNames}}) __option.Option[{{.TypeName}}] { return _fromNonZero{{.Name}}(s.{{.Name}}) },
		func(v {{.TypeName}}) *{{$.Name}}{{$.TypeParamNames}} {
			{{- if .IsEmbedded}}
			var result {{$.Name}}{{$.TypeParamNames}}
			result.{{.Name}} = v
			return &result
			{{- else}}
			return &{{$.Name}}{{$.TypeParamNames}}{ {{.Name}}: v }
			{{- end}}
		},
		"{{$.Name}}{{$.TypeParamNames}}.{{.Name}}",
	)
{{- else}}
	_prism{{.Name}} := __prism.MakePrismWithName(
		func(s *{{$.Name}}{{$.TypeParamNames}}) __option.Option[{{.TypeName}}] { return __option.Some(s.{{.Name}}) },
		func(v {{.TypeName}}) *{{$.Name}}{{$.TypeParamNames}} {
			{{- if .IsEmbedded}}
			var result {{$.Name}}{{$.TypeParamNames}}
			result.{{.Name}} = v
			return &result
			{{- else}}
			return &{{$.Name}}{{$.TypeParamNames}}{ {{.Name}}: v }
			{{- end}}
		},
		"{{$.Name}}{{$.TypeParamNames}}.{{.Name}}",
	)
{{- end}}
{{- end}}
	return {{.Name}}RefPrisms{{.TypeParamNames}} {
{{- range .Fields}}
		{{.Name}}: _prism{{.Name}},
{{- end}}
	}
}
`

var (
	structTmpl      *template.Template
	constructorTmpl *template.Template
)

func init() {
	var err error
	structTmpl, err = template.New("struct").Parse(lensStructTemplate)
	if err != nil {
		panic(err)
	}
	constructorTmpl, err = template.New("constructor").Parse(lensConstructorTemplate)
	if err != nil {
		panic(err)
	}
}

// hasLensAnnotation checks if a comment group contains the lens annotation
func hasLensAnnotation(doc *ast.CommentGroup) bool { _ = "STUB: not implemented"; return false }

// getTypeName extracts the type name from a field type expression
func getTypeName(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

// Generic type with single type parameter (Go 1.18+)
// e.g., Option[string]

// Generic type with multiple type parameters (Go 1.18+)
// e.g., Map[string, int]

// extractImports extracts package imports from a type expression
// Returns a map of package path -> package name
func extractImports(expr ast.Expr, imports map[string]string) { _ = "STUB: not implemented"; return }

// This is a qualified identifier like "option.Option"

// ident.Name is the package name (e.g., "option")
// We need to track this for import resolution

// Generic type with single type parameter

// Generic type with multiple type parameters

// hasOmitEmpty checks if a struct tag contains json omitempty
func hasOmitEmpty(tag *ast.BasicLit) bool { _ = "STUB: not implemented"; return false }

// Parse the struct tag

// Check if omitempty is present

// isPointerType checks if a type expression is a pointer
func isPointerType(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// isComparableType checks if a type expression represents a comparable type.
// Comparable types in Go include:
// - Basic types (bool, numeric types, string)
// - Pointer types
// - Channel types
// - Interface types
// - Structs where all fields are comparable
// - Arrays where the element type is comparable
//
// Non-comparable types include:
// - Slices
// - Maps
// - Functions
//
// typeParams is a map of type parameter names to their constraints (e.g., "T" -> "any", "K" -> "comparable")
func isComparableType(expr ast.Expr, typeParams map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if this is a type parameter

// Type parameter - check its constraint

// Basic types and named types
// We assume named types are comparable unless they're known non-comparable types

// Known non-comparable built-in types

// error is an interface, which is comparable

// Most basic types and named types are comparable
// We can't determine if a custom type is comparable without type checking,
// so we assume it is (conservative approach)

// Pointer types are always comparable

// Arrays are comparable if their element type is comparable

// This is a slice (no length), slices are not comparable

// Fixed-size array, check element type

// Maps are not comparable

// Functions are not comparable

// Interface types are comparable

// Structs are comparable if all fields are comparable
// We can't easily determine this without full type information,
// so we conservatively return false for struct literals

// Qualified identifier (e.g., pkg.Type)
// We can't determine comparability without type information
// Check for known non-comparable types from standard library

// Check for known non-comparable types

// context.Context is an interface, which is comparable

// For other qualified types, we assume they're comparable
// This is a conservative approach

// Generic types - we can't determine comparability without type information
// For common generic types, we can make educated guesses

// Check for known non-comparable generic types

// Option types are not comparable (they contain a slice internally)

// Either types are not comparable

// For other generic types, conservatively assume not comparable

// Channel types are comparable

// Unknown type, conservatively assume not comparable

// embeddedFieldResult holds both the field info and its AST type for import extraction
type embeddedFieldResult struct {
	fieldInfo fieldInfo
	fieldType ast.Expr
}

// extractEmbeddedFields extracts fields from an embedded struct type
// It returns a slice of embeddedFieldResult for all exported fields in the embedded struct
// typeParamsMap contains the type parameters of the parent struct (for checking comparability)
func extractEmbeddedFields(embedType ast.Expr, fileImports map[string]string, file *ast.File, typeParamsMap map[string]string) []embeddedFieldResult {
	_ = "STUB: not implemented"
	return nil
}

// Get the type name of the embedded field

// Direct embedded type: type MyStruct struct { EmbeddedType }

// Pointer embedded type: type MyStruct struct { *EmbeddedType }

// Qualified embedded type: type MyStruct struct { pkg.EmbeddedType }
// We can't easily resolve this without full type information
// For now, skip these

// Find the struct definition in the same file

// Struct not found in this file, might be from another package

// Extract fields from the embedded struct

// Skip embedded fields within embedded structs (for now, to avoid infinite recursion)

// Generate lenses for both exported and unexported fields

// Keep the block structure for minimal changes

// Check if field is optional

// Check if the type is comparable

// extractTypeParams extracts type parameters from a type spec
// Returns two strings: full params like "[T any]" and names only like "[T]"
func extractTypeParams(typeSpec *ast.TypeSpec) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// buildTypeParamsMap creates a map of type parameter names to their constraints
// e.g., for "type Box[T any, K comparable]", returns {"T": "any", "K": "comparable"}
func buildTypeParamsMap(typeSpec *ast.TypeSpec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// parseFile parses a Go file and extracts structs with lens annotations
func parseFile(filename string) ([]structInfo, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Build import map: package name -> import path

// Extract package name from path (last component)

// First pass: collect all GenDecls with their doc comments

// Second pass: process type specs

// Look for type declarations

// Check if it's a struct type

// Get the doc comment from our map

// Extract field information and collect imports

// Build type parameters map for this struct

// Embedded field - promote its fields

// Extract imports from embedded field's type

// Resolve package names to full import paths

// Generate lenses for both exported and unexported fields

// Keep the block structure for minimal changes

// Check if field is optional:
// 1. Pointer types are always optional
// 2. Non-pointer types with json omitempty tag are optional

// Strip leading * for base type

// Non-pointer type with omitempty is also optional

// Check if the type is comparable (for non-optional fields)
// For optional fields, we don't need to check since they use LensO

// log.Printf("field %s, type: %v, isComparable: %b\n", name, field.Type, isComparable)

// Extract imports from this field's type

// Resolve package names to full import paths

// generateLensHelpers scans a directory for Go files and generates lens code
func generateLensHelpers(dir, filename string, verbose, includeTestFiles bool) error {
	_ = "STUB: not implemented"
	// Get absolute path
	return nil
}

// Find all Go files in the directory

// Parse all files and collect structs, separating test and non-test files

// Skip generated lens files (both regular and test)

// Skip test files unless includeTestFiles is true

// Separate structs based on source file type

// Generate regular lens file if there are regular structs

// Generate test lens file if there are test structs

// generateLensFile generates a lens file for the given structs
func generateLensFile(absDir, filename, packageName string, structs []structInfo, verbose bool) error {
	_ = "STUB: not implemented"
	// Collect all unique imports from all structs
	return nil
}

// import path -> alias

// Create output file

// Write header

// Write imports

// Standard fp-go imports always needed

// Add additional imports collected from field types

// Generate lens code for each struct using templates

// Generate struct type

// Generate constructor

// Write to file

// LensCommand creates the CLI command for lens generation
func LensCommand() *C.Command { _ = "STUB: not implemented"; return nil }
