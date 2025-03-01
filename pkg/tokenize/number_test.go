package tokenize

import (
	"io"
	"testing"
)

func TestConsumeLiteralNumber(t *testing.T) {
	type input struct {
		Token []byte
		AtEOF bool
	}
	type output struct {
		N     int
		Token []byte
		Err   error
	}

	type testcase struct {
		Reason string
		In     input
		Out    output
	}
	type testgroup struct {
		Reason string
		Cases  []testcase
	}

	groups := []testgroup{
		{
			Reason: "Basic numeral tests",
			Cases: []testcase{
				{
					Reason: "atEOF, and completely consumed b",
					In:     input{Token: Token("1"), AtEOF: true},
					Out:    output{N: 1, Token: Token("1"), Err: nil},
				},
				{
					Reason: "completely consumed b, but !atEOF, so number may continue",
					In:     input{Token: Token("1"), AtEOF: false},
					Out:    output{N: 0, Token: nil, Err: nil},
				},
				{
					Reason: "!atOF, but encountered valid termination",
					In:     input{Token: Token("1,"), AtEOF: false},
					Out:    output{N: 1, Token: Token("1"), Err: nil},
				},
				{
					Reason: "!atOF, but encountered termination; validity is irrelevant",
					In:     input{Token: Token("1t"), AtEOF: false},
					Out:    output{N: 1, Token: Token("1"), Err: nil},
				},
			},
		},
		{
			Reason: "Decimal numeral tests",
			Cases: []testcase{
				{
					Reason: "atEOF and completely consumed b",
					In:     input{Token: Token("1.1"), AtEOF: true},
					Out:    output{N: 3, Token: Token("1.1"), Err: nil},
				},
				{
					Reason: "completely consumed b, but !atEOF, so number may continue",
					In:     input{Token: Token("1.1"), AtEOF: false},
					Out:    output{N: 0, Token: nil, Err: nil},
				},
				{
					Reason: "!atOF, but encountered valid termination",
					In:     input{Token: Token("1.1"), AtEOF: true},
					Out:    output{N: 3, Token: Token("1.1"), Err: nil},
				},
				{
					Reason: "!atOF, but encountered termination; validity is irrelevant",
					In:     input{Token: Token("1.1"), AtEOF: true},
					Out:    output{N: 3, Token: Token("1.1"), Err: nil},
				},
			},
		},
		{
			Reason: "Signed integer and decimal tests",
			Cases: []testcase{
				{
					Reason: "",
					In:     input{Token: Token("-1"), AtEOF: true},
					Out:    output{N: 2, Token: Token("-1"), Err: nil},
				},

				{
					Reason: "",
					In:     input{Token: Token("+1"), AtEOF: true},
					Out:    output{N: 2, Token: Token("+1"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("-1.107"), AtEOF: true},
					Out:    output{N: 6, Token: Token("-1.107"), Err: nil},
				},

				{
					Reason: "",
					In:     input{Token: Token("+1.107"), AtEOF: true},
					Out:    output{N: 6, Token: Token("+1.107"), Err: nil},
				},
			},
		},

		{
			Reason: "Scientific-notation number tests",
			Cases: []testcase{
				{
					Reason: "",
					In:     input{Token: Token("1e2.789"), AtEOF: true},
					Out:    output{N: 7, Token: Token("1e2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("+1e2.789"), AtEOF: true},
					Out:    output{N: 8, Token: Token("+1e2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("-1e2.789"), AtEOF: true},
					Out:    output{N: 8, Token: Token("-1e2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("1e-2.789"), AtEOF: true},
					Out:    output{N: 8, Token: Token("1e-2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("1e+2.789"), AtEOF: true},
					Out:    output{N: 8, Token: Token("1e+2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("-1e+2.789"), AtEOF: true},
					Out:    output{N: 9, Token: Token("-1e+2.789"), Err: nil},
				},
				{
					Reason: "",
					In:     input{Token: Token("+1e-2.789"), AtEOF: true},
					Out:    output{N: 9, Token: Token("+1e-2.789"), Err: nil},
				},
			},
		},
		{
			Reason: "NaN tests",
			Cases: []testcase{
				{
					Reason: "plain NaN token",
					In:     input{Token: Token("NaN"), AtEOF: true},
					Out:    output{N: 3, Token: Token("NaN"), Err: nil},
				},
			},
		},
		{
			Reason: "Infinity tests",
			Cases: []testcase{
				{
					Reason: "plain Inf token with atEOF",
					In:     input{Token: Token("Inf"), AtEOF: true},
					Out:    output{N: 3, Token: Token("Inf"), Err: nil},
				},
				{
					Reason: "plain Inf token, but with !atEOF",
					In:     input{Token: Token("Inf"), AtEOF: false},
					Out:    output{N: 3, Token: Token("Inf"), Err: nil},
				},
				{
					Reason: "signed Inf token, negative",
					In:     input{Token: Token("-Inf"), AtEOF: true},
					Out:    output{N: 4, Token: Token("-Inf"), Err: nil},
				},
				{
					Reason: "signed Inf token, positive (does that make it a fool?)",
					In:     input{Token: Token("+Inf"), AtEOF: true},
					Out:    output{N: 4, Token: Token("+Inf"), Err: nil},
				},
				{
					Reason: "signed Inf token, negative, affirmatively terminated",
					In:     input{Token: Token("-Inf,"), AtEOF: true},
					Out:    output{N: 4, Token: Token("-Inf"), Err: nil},
				},
				{
					Reason: "signed Inf token, positive, affirmatively terminated",
					In:     input{Token: Token("-Inf,"), AtEOF: true},
					Out:    output{N: 4, Token: Token("-Inf"), Err: nil},
				},
				{
					Reason: "plain numeral followed by Inf, should consume only the numeral",
					In:     input{Token: Token("1Inf"), AtEOF: true},
					Out:    output{N: 1, Token: Token("1"), Err: nil},
				},
			},
		},
		{
			Reason: "zero-length tokens",
			Cases: []testcase{
				{
					Reason: "zero-length, atEOF",
					In:     input{Token: nil, AtEOF: true},
					Out:    output{N: 0, Token: nil, Err: io.EOF},
				},
				{
					Reason: "",
					In:     input{Token: nil, AtEOF: false},
					Out:    output{N: 0, Token: nil, Err: nil},
				},
			},
		},
	}

	for i := range groups {
		t.Log(groups[i].Reason)
		scenarios := groups[i].Cases
		for j := range scenarios {
			t.Log(scenarios[j].Reason)

			n, b, err := consumeLiteralNumber(Token(scenarios[j].In.Token), scenarios[j].In.AtEOF)
			result := output{N: n, Token: b, Err: err}
			switch {
			case n != scenarios[j].Out.N:
				t.Errorf("expected to consume %v; got %v", scenarios[j].Out, result)
			case !Token.Equals(scenarios[j].Out.Token, result.Token):
				t.Errorf(`expected the number-literal %q; got %q`, scenarios[j].Out.Token, result.Token)
			case scenarios[j].Out.Err != result.Err:
				t.Errorf("expected err as %v; got %v", scenarios[j].Out.Err, result.Err)
			}
		}
	}
}
