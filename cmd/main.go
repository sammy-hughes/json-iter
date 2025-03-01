package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/sammy-hughes/json-iter/pkg/tokenize"
)

var TestData = `
[
	{"id": 1, "name": "apple"},
	{"id": 1, "name": "orange", "sectionAngles": [0, 45, 90, 135, 180, 225, 270, 315]}
]
`
var TestLiterals = `[true,false,null,1,1.1e0,1.1e+1,1.1e-1,0,0.0," "]`

func TestSplitTokens_TestLiterals() {
	reader := bytes.NewReader(tokenize.Token(TestLiterals))
	scanner := bufio.NewScanner(reader)
	scanner.Split(tokenize.Tokens)

	fmt.Println("starting up")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func TestSplitTokens_TestData() {
	reader := bytes.NewReader(tokenize.Token(TestData))
	scanner := bufio.NewScanner(reader)
	scanner.Split(tokenize.Tokens)

	fmt.Println("starting up")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	TestSplitTokens_TestLiterals()
	TestSplitTokens_TestData()
}
