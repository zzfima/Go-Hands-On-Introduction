package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

// run as "go run challenge.go data.txt"
// or add to run configuration: "args": ["data.txt"]
func mainFlowControl() {
	//defers
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panic!!!! :-( ", e)
		} else {
			fmt.Println("no panic :-)")
		}
	}()

	b, e := os.ReadFile(os.Args[1])
	if e != nil {
		fmt.Println(e)
		return
	}

	textMap := make(map[string]int)
	words := strings.Split(string(b), " ")
	textMap["words"] = len(words)

	for _, c := range b {
		r := int32(c)

		switch {
		case unicode.IsSpace(r):
			textMap["spaces"]++
		case unicode.IsLetter(r):
			textMap["letters"]++
		case unicode.IsDigit(r):
			textMap["numbers"]++
		default:
			textMap["symbols"]++
		}
	}

	spew.Dump(textMap)
}
