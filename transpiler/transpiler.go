package transpiler

import (
	"GoCrab/errors"
	"GoCrab/lexer"
	"GoCrab/parser"
	"fmt"

)

var hasError bool

// Transpile attempts to transpile Rust code to Go code.
// It returns a RustCodeError if the Rust code itself has issues.
func Transpile(code string) (string, error) {
	if code == "" {
		return "", &errors.RustCodeError{Message: "Rust code is empty or invalid.", LineNum: 0}
	}

	scanner := lexer.NewScanner(code)
	tokens, errors := scanner.ScanTokens()
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

	// Call the parser here
	ast, err := parser.Parse(tokens)
	if err != nil {
		return "", err
	}

	fmt.Println(ast)

	if hasError {
		return "", fmt.Errorf("why u give bad code my dude")
	}

	return "", nil
}
