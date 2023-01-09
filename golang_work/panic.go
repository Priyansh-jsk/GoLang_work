package main

import "fmt"

func main() {
	fmt.Println("one")
	defer fmt.Println("three")
	panic("a panic message show")
	fmt.Println("four")
}
