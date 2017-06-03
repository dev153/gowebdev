package main

import "fmt"

func main() {
	strchan := make(chan string, 2)
	strchan <- "Golang"
	strchan <- "Gopher"
	fmt.Println(<-strchan)
	fmt.Println(<-strchan)
}
