package main

import (
	"fmt"
	"os"
	"time"

	proverbs "github.com/jboursiquot/go-proverbs"
)

func main() {
	v1 := getRandomProverb()
	fmt.Println(v1)
	v2 := getRandomProverb()
	fmt.Fprint(os.Stdout, v2)
}

func getRandomProverb() string {
	p := proverbs.Random()
	time.Sleep(10)
	return p.Saying
}
