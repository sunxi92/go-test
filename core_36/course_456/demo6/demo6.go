package main

import . "fmt"

var block = "1"

func main() {
	block := "2"
	{
		block := "3"
		Printf("The block is %s.\n", block)
	}
	Printf("The block is %s.\n", block)

	var fmt = "hello"

}
