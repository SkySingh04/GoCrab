package parser

import "GoCrab/lexer"

// AST represents the abstract syntax tree of the Rust code.

type AST struct {
	Root *Node
}

// Node represents a node in the AST.
type Node struct {
	Children []*Node
}

func Parse(tokens []lexer.Token) (*AST, error) {
	return nil, nil
}
