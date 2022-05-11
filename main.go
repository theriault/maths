package main

import (
	"fmt"

	"github.com/theriault/maths/number"
)

func main() {
	v := number.Sqrt(25)
	fmt.Printf("%T %v\n", v, v)
}
