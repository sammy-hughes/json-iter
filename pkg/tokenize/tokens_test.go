package tokenize_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/sammy-hughes/json-iter/pkg/tokenize"
)

func TestTokens(t *testing.T) {
	testInput := strings.NewReader(`{"id": 1, "name": "apple"}`)
	testOutput := slices.Collect(tokenize.Tokens(testInput))

	switch {
	case string(testOutput[0]) != "{":
		t.Fail()
	case string(testOutput[1]) != `"id"`:
		t.Fail()
	case string(testOutput[2]) != `:`:
		t.Fail()
	case string(testOutput[3]) != `1`:
		t.Fail()
	case string(testOutput[4]) != `,`:
		t.Fail()
	case string(testOutput[5]) != `"name"`:
		t.Fail()
	case string(testOutput[6]) != `:`:
		t.Fail()
	case string(testOutput[7]) != `"apple"`:
		t.Fail()
	case string(testOutput[8]) != `}`:
		t.Fail()
	}
}

//	{"id": 1, "name": "orange", "sectionAngles": [0, 45, 90, 135, 180, 225, 270, 315]},
//	[true,false,null,1,1.1e0,1.1e+1,1.1e-1,0,0.0," "]
//]
//`
//}
