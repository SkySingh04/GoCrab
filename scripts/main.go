// package temporarily must be set to main to run this script
// package scripts

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	outputDir := "./expr"
	defineAst(outputDir, "Expr", []string{
		"Binary   : Left Expr, Operator lexer.Token, Right Expr",
		"Grouping : Expression Expr",
		"Literal  : Value interface{}",
		"Unary    : Operator lexer.Token, Right Expr",
	})
}

func defineAst(outputDir string, baseName string, types []string) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Could not create directory: %v\n", err)
		os.Exit(1)
	}
	path := fmt.Sprintf("%s/%s.go", outputDir, strings.ToLower(baseName))
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write the package declaration.
	file.WriteString("package expr\n\n")

	// Write imports.
	file.WriteString("import (\n\t\"GoCrab/lexer\"\n)\n\n")

	// Write the base interface.
	file.WriteString(fmt.Sprintf("// %s is the base interface for all expression types.\n", baseName))
	file.WriteString(fmt.Sprintf("type %s interface {\n", baseName))
	file.WriteString(fmt.Sprintf("\tAccept(visitor %sVisitor) interface{}\n", baseName))
	file.WriteString("}\n\n")

	// Write each subclass.
	for _, typeDef := range types {
		parts := strings.Split(typeDef, ":")
		className := strings.TrimSpace(parts[0])
		fields := strings.TrimSpace(parts[1])
		defineType(file, baseName, className, fields)
	}

	// Write the visitor interface.
	defineVisitor(file, baseName, types)
}

func defineType(file *os.File, baseName, className, fieldList string) {
	// Write the struct definition.
	file.WriteString(fmt.Sprintf("// %s represents the %s expression type.\n", className, className))
	file.WriteString(fmt.Sprintf("type %s struct {\n", className))

	// Write the fields.
	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		parts := strings.SplitN(field, " ", 2)
		fieldType := strings.TrimSpace(parts[1])
		fieldName := strings.TrimSpace(parts[0])
		file.WriteString(fmt.Sprintf("\t%s %s\n", fieldName, fieldType))
	}
	file.WriteString("}\n\n")

	// Write the Accept method.
	file.WriteString(fmt.Sprintf("// Accept allows %s to implement the %s interface.\n", className, baseName))
	file.WriteString(fmt.Sprintf("func (e *%s) Accept(visitor %sVisitor) interface{} {\n", className, baseName))
	file.WriteString(fmt.Sprintf("\treturn visitor.Visit%s%s(e)\n", className, baseName))
	file.WriteString("}\n\n")
}

func defineVisitor(file *os.File, baseName string, types []string) {
	// Write the visitor interface.
	file.WriteString(fmt.Sprintf("// %sVisitor is the visitor interface for %s types.\n", baseName, baseName))
	file.WriteString(fmt.Sprintf("type %sVisitor interface {\n", baseName))

	for _, typeDef := range types {
		className := strings.TrimSpace(strings.Split(typeDef, ":")[0])
		file.WriteString(fmt.Sprintf("\tVisit%s%s(*%s) interface{}\n", className, baseName, className))
	}

	file.WriteString("}\n")
}
