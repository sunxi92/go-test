package main

import (
	"flag"
	"go-test/core_36/course_3/demo3/lib"
	in "go-test/core_36/course_3/demo3/lib/internal"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib3.Hello(name)
	in.Hello(os.Stdout, name)
}
