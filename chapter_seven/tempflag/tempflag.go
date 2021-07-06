package main

import (
	"flag"
	"fmt"
	"go-test/chapter_seven/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
