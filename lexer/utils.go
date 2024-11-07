package lexer

import "GoCrab/errors"

func (s *Scanner) scanString() errors.RustCodeError {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.readChar()
	}

	if s.isAtEnd() {
		// Unterminated string error handling
		return errors.RustCodeError{Message: "Unterminated string", LineNum: s.line}

	}

	// Consume the closing '"'
	s.readChar()

	// Trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addToken(StringLiteral, value)
	return errors.RustCodeError{}
}

// match consumes the next character if it matches the expected character.
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

// scanNumber handles integer and floating-point literals and throws errors for invalid numbers.
func (s *Scanner) scanNumber() errors.RustCodeError {
	for isDigit(s.peek()) {
		s.readChar()
	}

	if s.peek() == '.' {
		s.readChar() // consume the '.'

		if !isDigit(s.peek()) {
			return errors.RustCodeError{Message: "Invalid number format", LineNum: s.line}
		}

		for isDigit(s.peek()) {
			s.readChar()
		}
		s.addToken(FloatLiteral, s.source[s.start:s.current])
	} else {
		s.addToken(IntegerLiteral, s.source[s.start:s.current])
	}

	return errors.RustCodeError{}
}

// scanIdentifierOrKeyword handles identifiers and keywords.
func (s *Scanner) scanIdentifierOrKeyword() errors.RustCodeError {
	for isAlphaNumeric(s.peek()) {
		s.readChar()
	}

	// Check if the identifier is a reserved keyword
	text := s.source[s.start:s.current]
	tokenType := lookupKeyword(text)

	// If it matches a keyword, add it as a keyword token
	if tokenType != Keyword(Identifier) {
		s.addToken(RustToken(tokenType), nil)
	} else {
		// Otherwise, add it as an identifier
		s.addToken(Identifier, nil)
	}
	return errors.RustCodeError{}
}

// lookupKeyword determines if a lexeme is a Rust keyword.
func lookupKeyword(text string) Keyword {
	switch text {
	case "break":
		return Break
	case "const":
		return Const
	case "continue":
		return Continue
	case "else":
		return Else
	case "enum":
		return Enum
	case "extern":
		return Extern
	case "false":
		return False
	case "for":
		return For
	case "fn":
		return Fn
	case "if":
		return If
	case "impl":
		return Impl
	case "in":
		return In
	case "let":
		return Let
	case "loop":
		return Loop
	case "match":
		return Match
	case "mod":
		return Mod
	case "move":
		return Move
	case "mut":
		return Mut
	case "pub":
		return Pub
	case "ref":
		return Ref
	case "return":
		return Return
	case "self":
		return Self
	case "Self":
		return SelfType
	case "static":
		return Static
	case "struct":
		return Struct
	case "super":
		return Super
	case "trait":
		return Trait
	case "true":
		return True
	case "type":
		return Type
	case "union":
		return Union
	case "unsafe":
		return Unsafe
	case "use":
		return Use
	case "where":
		return Where
	case "while":
		return While
	case "as":
		return As
	case "async":
		return Async
	case "await":
		return Await
	case "dyn":
		return Dyn
	case "abstract":
		return Abstract
	case "become":
		return Become
	case "box":
		return Box
	case "do":
		return Do
	case "final":
		return Final
	case "macro":
		return Macro
	case "override":
		return Override
	case "priv":
		return Priv
	case "typeof":
		return Typeof
	case "unsized":
		return Unsized
	case "virtual":
		return Virtual
	case "yield":
		return Yield
	case "crate":
		return Crate
	case "derive":
		return Derive
	case "try":
		return Try
	default:
		return Keyword(Identifier)
	}
}

// peek returns the current character without consuming it.
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current]
}

// peekNext returns the character after the current one without consuming it.
func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\x00'
	}
	return s.source[s.current+1]
}

// Utility functions for character classification
func isDigit(c byte) bool        { return c >= '0' && c <= '9' }
func isAlpha(c byte) bool        { return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' }
func isAlphaNumeric(c byte) bool { return isAlpha(c) || isDigit(c) }

// readChar consumes and returns the current character.
func (s *Scanner) readChar() byte {
	s.current++
	return s.source[s.current-1]
}

// isAtEnd checks if the scanner has reached the end of the source.
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// addToken adds a new token to the tokens slice.
func (s *Scanner) addToken(tokenType RustToken, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType, text, literal, s.line})
}
