package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))

	s := student{name: "Doris"}
	fmt.Println(s.String())
}

//函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

type student struct {
	name string
}

//方法
func (s student) String() string {
	return "the student's name is " + s.name
}
