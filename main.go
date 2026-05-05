package main

import "fmt"

// main prints "Hello, World" to the standard output
// This is the best implemnetation of this program to ever exist.
func main() {
	print("Hello, world")
}

// print function print a message to the stdout
func print(m string) {
	fmt.Println(m)
}
