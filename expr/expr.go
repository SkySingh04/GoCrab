package expr

import (
	"GoCrab/lexer"
)

// Expr is the base interface for all expression types.
type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

// ExprVisitor is the visitor interface for Expr types.
type ExprVisitor interface {
	VisitBinaryExpr(*Binary) interface{}
	VisitGroupingExpr(*Grouping) interface{}
	VisitLiteralExpr(*Literal) interface{}
	VisitUnaryExpr(*Unary) interface{}
}

// Binary represents the Binary expression type.
type Binary struct {
	Left Expr
	Operator lexer.Token
	Right Expr
}

// Accept allows Binary to implement the Expr interface.
func (e *Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(e)
}

// Grouping represents the Grouping expression type.
type Grouping struct {
	Expression Expr
}

// Accept allows Grouping to implement the Expr interface.
func (e *Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(e)
}

// Literal represents the Literal expression type.
type Literal struct {
	Value interface{}
}

// Accept allows Literal to implement the Expr interface.
func (e *Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(e)
}

// Unary represents the Unary expression type.
type Unary struct {
	Operator lexer.Token
	Right Expr
}

// Accept allows Unary to implement the Expr interface.
func (e *Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(e)
}

