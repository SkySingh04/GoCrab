package lexer

import (
	"GoCrab/errors"
	"fmt"
)

//The core of the scanner is a loop. Starting at the first character of the source code, the scanner figures out what lexeme the character belongs to, and consumes it and any following characters that are part of that lexeme.
//When it reaches the end of that lexeme, it emits a token.

// Scanner holds the state of the current scanning process.
type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

// NewScanner initializes a new Scanner with the provided source code.
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

// scanTokensWithErrors goes through the source and generates tokens, collecting errors.
func (s *Scanner) ScanTokens() ([]Token, []errors.RustCodeError) {
	var errorsList []errors.RustCodeError
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != (errors.RustCodeError{}) {
			errorsList = append(errorsList, err)
		}
	}

	// Add final EOF token
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})

	return s.tokens, errorsList
}

// scanToken scans a single token based on the current character.
func (s *Scanner) scanToken() errors.RustCodeError {
	char := s.readChar()
	switch char {
	// Single-character tokens
	case '+':
		s.addToken(Plus, nil)
	case '-':
		if s.match('>') {
			s.addToken(RArrow, nil)
		} else {
			s.addToken(Minus, nil)
		}
	case '*':
		s.addToken(Star, nil)
	case '/':
		if s.match('/') {
			// Single-line comment
			for s.peek() != '\n' && !s.isAtEnd() {
				s.readChar()
			}
		} else if s.match('*') {
			// Multi-line comment
			for !(s.peek() == '*' && s.peekNext() == '/') && !s.isAtEnd() {
				if s.peek() == '\n' {
					s.line++
				}
				s.readChar()
			}
			// Consume the closing '*/'
			if !s.isAtEnd() {
				s.readChar()
				s.readChar()
			}
		} else {
			s.addToken(Slash, nil)
		}
	case '%':
		s.addToken(Percent, nil)
	case '^':
		s.addToken(Caret, nil)
	case '!':
		if s.match('=') {
			s.addToken(Ne, nil)
		} else {
			s.addToken(Not, nil)
		}
	case '&':
		if s.match('&') {
			s.addToken(AndAnd, nil)
		} else {
			s.addToken(And, nil)
		}
	case '|':
		if s.match('|') {
			s.addToken(OrOr, nil)
		} else {
			s.addToken(Or, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EqEq, nil)
		} else {
			s.addToken(Eq, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(Le, nil)
		} else if s.match('<') {
			s.addToken(Shl, nil)
		} else {
			s.addToken(Lt, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(Ge, nil)
		} else if s.match('>') {
			s.addToken(Shr, nil)
		} else {
			s.addToken(Gt, nil)
		}
	case '@':
		s.addToken(At, nil)
	case '_':
		s.addToken(Underscore, nil)
	case '.':
		if s.match('.') {
			if s.match('.') {
				s.addToken(DotDotDot, nil)
			} else {
				s.addToken(DotDot, nil)
			}
		} else {
			s.addToken(Dot, nil)
		}
	case ',':
		s.addToken(Comma, nil)
	case ';':
		s.addToken(Semi, nil)
	case ':':
		s.addToken(Colon, nil)
	case '$':
		s.addToken(Dollar, nil)
	case '?':
		s.addToken(Question, nil)
	case '#':
		s.addToken(Pound, nil)
	case '{':
		s.addToken(CurlyOpen, nil)
	case '}':
		s.addToken(CurlyClose, nil)
	case '[':
		s.addToken(SquareOpen, nil)
	case ']':
		s.addToken(SquareClose, nil)
	case '(':
		s.addToken(ParenOpen, nil)
	case ')':
		s.addToken(ParenClose, nil)
	// Whitespace and newlines
	case ' ', '\r', '\t':
		// Ignore whitespace
		break
	case '\n':
		s.line++
	case '"':
		err := s.scanString()
		if err != (errors.RustCodeError{}) {
			return err
		}
	default:
		if isDigit(char) {
			err := s.scanNumber()
			if err != (errors.RustCodeError{}) {
				return err
			}
		} else if isAlpha(char) {
			err := s.scanIdentifierOrKeyword()
			if err != (errors.RustCodeError{}) {
				return err
			}
		} else {
			// Unknown character error handling
			return errors.RustCodeError{Message: fmt.Sprintf("Unexpected character: %c", char), LineNum: s.line}
		}
	}
	return errors.RustCodeError{}
}
