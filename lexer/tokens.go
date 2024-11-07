package lexer

import "fmt"

type RustToken int

// Token represents a lexical token with a type, lexeme, literal, and line number.
type Token struct {
	Type    RustToken
	Lexeme  string
	Literal interface{}
	Line    int
}

// NewToken creates and returns a new Token instance.
func NewToken(tokenType RustToken, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

// String returns a string representation of the Token.
func (t *Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}

const (
	// Keywords
	RustKeyword RustToken = iota

	// Identifiers
	Identifier
	Lifetime

	// Literals
	CharLiteral
	StringLiteral
	RawStringLiteral
	ByteLiteral
	ByteStringLiteral
	RawByteStringLiteral
	IntegerLiteral
	FloatLiteral
	CStringLiteral
	RawCStringLiteral

	// Punctuation Operators
	Plus
	Minus
	Star
	Slash
	Percent
	Caret
	Not
	And
	Or
	AndAnd
	OrOr
	Shl
	Shr
	PlusEq
	MinusEq
	StarEq
	SlashEq
	PercentEq
	CaretEq
	AndEq
	OrEq
	ShlEq
	ShrEq
	Eq
	EqEq
	Ne
	Gt
	Lt
	Ge
	Le
	At
	Underscore
	Dot
	DotDot
	DotDotDot
	DotDotEq
	Comma
	Semi
	Colon
	PathSep
	RArrow
	FatArrow
	LArrow
	Pound
	Dollar
	Question

	// Delimiters
	CurlyOpen
	CurlyClose
	SquareOpen
	SquareClose
	ParenOpen
	ParenClose

	// Reserved forms
	ReservedDoubleQuote
	ReservedSingleQuote
	ReservedPound

	EOF
)

// Rust Keywords Enum
type Keyword int

const (
	// Control flow keywords
	Break Keyword = iota
	Const
	Continue
	Else
	Enum
	Extern
	False
	For
	Fn
	If
	Impl
	In
	Let
	Loop
	Match
	Mod
	Move
	Mut
	Pub
	Ref
	Return
	Self
	SelfType
	Static
	Struct
	Super
	Trait
	True
	Type
	Union
	Unsafe
	Use
	Where
	While
	As
	Async
	Await
	Dyn
	Abstract
	Become
	Box
	Do
	Final
	Macro
	Override
	Priv
	Typeof
	Unsized
	Virtual
	Yield
	Crate
	Derive
	Try
)

// Integer and Float Literal Suffixes
type IntegerLiteralSuffix struct {
	Signed bool
	Bits   int
	Type   string
}

type FloatLiteralSuffix struct {
	Bits int
	Type string
}
