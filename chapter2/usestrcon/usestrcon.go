package main

import (
	"fmt"

	strcon "github.com/dev153/gowebdev/strcon"
)

func printSliceStat(sl []int) {
	fmt.Println(sl)
	fmt.Println("Length is", len(sl))
	fmt.Println("Capacity is", cap(sl))
}

func main() {
	str := "AbCdEf"
	stralt := strcon.SwapCase(str)
	fmt.Println(stralt)
	x := []int{4: 1}
	fmt.Println(x)
	y := make([]int, 2, 5)
	y[0] = 10
	y[1] = 20
	printSliceStat(y)
	y = append(y, 30, 40, 50)
	printSliceStat(y)
	y = append(y, 60)
	printSliceStat(y)
}
