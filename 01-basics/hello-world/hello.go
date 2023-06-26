package main

import "fmt" // stands for format, part of go standard lib

func main() {
    var Greeting = "Hell" // available outside of scope 
    var greeting1 string
    var greeting2 = "Hello, W"
    greeting3 := "Hello, World!"
    greeting1 = "Hello"
	fmt.Println(Greeting)
	fmt.Println(greeting1)
	fmt.Println(greeting2)
	fmt.Println(greeting3)
}
