package lib3

import (
	in "go-test/core_36/course_3/demo3/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
