package main

import (
	"flag"
	"fmt"
)

func main() {
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")

	//way one
	//var name = flag.String("name", "everyone", "The greeting object.")

	//way two
	name := flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	//fmt.Printf("Hello, %v!\n", name)

	//way one and two
	fmt.Printf("Hello, %v!\n", *name)
}
