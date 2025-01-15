package lexer

import (
	"testing"
)

func TestNewScanner(t *testing.T) {
	src := "fn main() {}"
	scanner := NewScanner(src)
	if scanner.source != src {
		t.Errorf("Expected source %q, got %q", src, scanner.source)
	}
	if scanner.line != 1 {
		t.Errorf("Expected line to be 1, got %d", scanner.line)
	}
}

func TestScanTokens_EmptySource(t *testing.T) {
	scanner := NewScanner("")
	tokens, errs := scanner.ScanTokens()
	if len(tokens) != 1 {
		t.Errorf("Expected 1 token (EOF), got %d", len(tokens))
	}
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d", len(errs))
	}
	if tokens[0].Type != EOF {
		t.Errorf("Expected EOF token, got %v", tokens[0].Type)
	}
}

func TestScanTokens_ValidSource(t *testing.T) {
	scanner := NewScanner("fn main() {}")
	tokens, errs := scanner.ScanTokens()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d", len(errs))
	}
	if len(tokens) == 0 {
		t.Error("Expected some tokens, got none")
	}
}
