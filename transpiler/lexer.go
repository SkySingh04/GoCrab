package transpiler

import "fmt"

// Transpile attempts to transpile Rust code to Go code.
// It returns a RustCodeError if the Rust code itself has issues.

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

// Transpile attempts to transpile Rust code to Go code.
// It returns a RustCodeError if the Rust code itself has issues.
var hasError bool

func Transpile(code string) (string, error) {
	if code == "" {
		return "", &RustCodeError{Message: "Rust code is empty or invalid.", LineNum: 0}
	}

	scanner := NewScanner(code)
	tokens, errors := scanner.scanTokens()
	if len(errors) > 0 {
		hasError = true
		for _, err := range errors {
			fmt.Println(err)
		}
		return "", fmt.Errorf("transpilation failed with errors")
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	if hasError {
		return "", fmt.Errorf("why u give bad code my dude")
	}

	return "", nil
}

// scanTokensWithErrors goes through the source and generates tokens, collecting errors.
func (s *Scanner) scanTokens() ([]Token, []RustCodeError) {
	var errors []RustCodeError
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != (RustCodeError{}) {
			errors = append(errors, err)
		}
	}

	// Add final EOF token
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})

	return s.tokens, errors
}

// scanToken scans a single token based on the current character.
func (s *Scanner) scanToken() RustCodeError {
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
	default:
		if isDigit(char) {
			s.scanNumber()
		} else if isAlpha(char) {
			s.scanIdentifierOrKeyword()
		} else {
			// Unknown character error handling
			return RustCodeError{Message: fmt.Sprintf("Unexpected character: %c", char), LineNum: s.line}
		}
	}
	return RustCodeError{}
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

// scanNumber handles integer and floating-point literals.
func (s *Scanner) scanNumber() {
	for isDigit(s.peek()) {
		s.readChar()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.readChar() // consume the '.'

		for isDigit(s.peek()) {
			s.readChar()
		}
		s.addToken(FloatLiteral, s.source[s.start:s.current])
	} else {
		s.addToken(IntegerLiteral, s.source[s.start:s.current])
	}
}

// scanIdentifierOrKeyword handles identifiers and keywords.
func (s *Scanner) scanIdentifierOrKeyword() {
	// for isAlphaNumeric(s.peek()) {
	// 	s.readChar()
	// }

	// // Check if the identifier is a reserved keyword
	// text := s.source[s.start:s.current]
	// tokenType := lookupKeyword(text)
	// s.addToken(tokenType, nil)
}

// lookupKeyword determines if a lexeme is a Rust keyword.
// func lookupKeyword(text string) RustToken {
// 	switch text {
// 	case "let":
// 		return Let
// 	case "fn":
// 		return Fn
// 	case "struct":
// 		return Struct
// 	case "enum":
// 		return Enum
// 	case "impl":
// 		return Impl
// 	case "trait":
// 		return Trait
// 	// Add other keywords as needed
// 	default:
// 		return Identifier
// 	}
// }

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
