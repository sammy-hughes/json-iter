package tokenize

import (
	"testing"
)

func TestConsumeLiteralNumber(t *testing.T) {
	scenarios := []string{
		"1",
		"1.1",
		"-1",
		"+1",
		"1e2.789",
		"NaN",
		"Inf",
		"-Inf",
	}

	for i := range scenarios {
		n, b, err := consumeLiteralNumber(Token(scenarios[i]), true)
		switch {
		case n < 1:
			t.Errorf("expected to consume at least 1 byte; got %d bytes (%v, %v, %v)", n, scenarios[i], string(b), err)
		case string(b) != string(scenarios[i]):
			t.Errorf(`expected the number-literal %q; got %q`, scenarios[i], string(b))
		case err != nil:
			t.Errorf("expected a nil error; got %v", err)
		}
	}
}
