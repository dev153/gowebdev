package main

import (
	"fmt"

	"github.com/dev153/gowebdev/chapter1/calc"
)

func main() {
	fmt.Println("Hello Go")
	x, y := 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
