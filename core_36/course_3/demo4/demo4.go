package main

import (
	"flag"
	lib3 "go-test/core_36/course_3/demo3/lib"
	lib "go-test/core_36/course_3/demo4/lib"
	//in "go-test/core_36/course_3/demo4/lib/internal"
	//"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	//lib4.Hello(name)
	//in.Hello(os.Stdout, name)
	test01()
	test02()
}

func test01() {
	lib.Hello(name)
}

func test02() {
	lib3.Hello(name)
}
