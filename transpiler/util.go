package transpiler

import "fmt"

type RustCodeError struct {
	Message string
	LineNum int
}

func (e *RustCodeError) Error() string {
	return fmt.Sprintf("At line %d: %s", e.LineNum, e.Message)
}

