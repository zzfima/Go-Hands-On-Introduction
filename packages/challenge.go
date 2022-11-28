package main

import (
	"fmt"

	proverbs "github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getRandomProverb())
}

func getRandomProverb() string {
	return proverbs.Random().Saying
}
