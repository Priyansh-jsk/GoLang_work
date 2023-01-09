package main

import "fmt"

/*func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}
func main() {
	add(20, 7)
}/*

// function Example2 //

/*func mymessage() {
	fmt.Println("Hello, priyansh")
}
func main() {
	mymessage()
}*/

// Recursion //
func test(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return test(x + 1)
}
func main() {
	test(1)
}
