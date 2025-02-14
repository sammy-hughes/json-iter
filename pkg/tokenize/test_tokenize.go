package tokenize

import (
	"testing"
)

func TestScanLiteralNumber(t *testing.T) {
	scenarios := []string{
		"1",
		"1.1",
		"-1",
		"+1",
		"1e2.789",
	}

	for i := range scenarios {
		i, b, err := scanLiteralNumber(Token(scenarios[i]), false)
		switch {
		case i < 1:
			t.Errorf("expected to consume at least 1 byte; got %d bytes", i)
		case string(b) != "1":
			t.Errorf(`expected the number-literal %q; got %q`, scenarios[i], string(b))
		case err != nil:
			t.Errorf("expected a nil error; got %v", err)
		}
	}
}
