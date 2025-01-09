// package temporarily must be set to main to run this script

package printer

// package main

import (
	"GoCrab/expr"
	"GoCrab/lexer"
	"fmt"
	"strings"
)

func main() {
	// Construct the AST manually.
	expression := &expr.Binary{
		Left: &expr.Unary{
			Operator: lexer.Token{
				Type:    lexer.Minus,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: &expr.Literal{Value: 123},
		},
		Operator: lexer.Token{
			Type:    lexer.Star,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &expr.Grouping{
			Expression: &expr.Literal{Value: 45.67},
		},
	}

	// Print the AST.
	printer := AstPrinter{}
	fmt.Println(printer.Print(expression))
}

// AstPrinter implements the ExprVisitor interface to print expressions.
type AstPrinter struct{}

// Print takes an Expr and returns its string representation.
func (p *AstPrinter) Print(expr expr.Expr) string {
	return expr.Accept(p).(string)
}

// VisitBinaryExpr handles Binary expressions.
func (p *AstPrinter) VisitBinaryExpr(expr *expr.Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// VisitGroupingExpr handles Grouping expressions.
func (p *AstPrinter) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return p.parenthesize("group", expr.Expression)
}

// VisitLiteralExpr handles Literal expressions.
func (p *AstPrinter) VisitLiteralExpr(expr *expr.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

// VisitUnaryExpr handles Unary expressions.
func (p *AstPrinter) VisitUnaryExpr(expr *expr.Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

// parenthesize formats an expression in a Lisp-style string.
func (p *AstPrinter) parenthesize(name string, exprs ...expr.Expr) string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(name)
	for _, subExpr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(subExpr.Accept(p).(string))
	}
	builder.WriteString(")")

	return builder.String()
}
