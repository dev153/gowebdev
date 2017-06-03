package main

import "fmt"

func doPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover with: ", err)
		}
	}()
	panic("Just panicking for the sake of the demo")
	// fmt.Println("This will never be called")
}

func main() {
	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover")
}
