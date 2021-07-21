package main

import (
	"flag"
	lib2 "go-test/core_36/course_3/demo2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib2.Hello(name)
}
