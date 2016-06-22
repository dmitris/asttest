package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to inspect the AST.

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "/Users/dsavints/go/src/github.com/dmitris/mytest/mytest.go", nil, 0)
	if err != nil {
		panic(err)
	}
	spec := f.Decls[0].(*ast.GenDecl).Specs[1].(*ast.ValueSpec)
	fmt.Printf("Name: %#v, Type: %#v\n", spec.Names[0], spec.Type)
}
