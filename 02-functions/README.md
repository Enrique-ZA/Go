# Functions in Go
## Function definition
>- Defined using **func**
>- takes 0 or more values and returns 0 or more values
>- functions are exported by capitalizing the name
```go
func add(a int, b int) int {
   return a + b c 
}

func main() {
    sub := func(a int, b int) int {
        return a - b
    }
    fmt.Println(add(10, 5))
    fmt.Println(sub(10, 5))
}
```
>- Go does not support method overloading
>- Go does however support variadic arguments
```go
func add(a ...int) {
    fmt.Println(a)
}
```
### Pointers and Functions
```go
func negate(x *int) {
    neg := -*x
    *x = neg
}

func main() {
    x := 10
    fmt.Println(x) // 10
    negate(&x)
    fmt.Println(x) // -10
}
```
## Multiple return values
>- Multiple return values replace tuple in Go.
```go
func rectangle(x int, y int) (int, int) {
   area := x * y
   circumference := 2 * (x + y)
   return area, circumference
}
```
>- Named return values
```go
func rectangle(x int, y int) (area int, circumference int) {
   area = x * y
   circumference = 2 * (x + y)
   return // the Go compiler can infer returns ... also called naked return
}
```
>- The _ operator a.k.a blank identifier
>- Use it to ignore return values during assignment
```go
func rectangle(x int, y int) (area int, circumference int) {
   area = x * y
   circumference = 2 * (x + y)
   return // the Go compiler can infer returns ... also called naked return
}

func main() {
    area, _ := rectangle(10, 5)
    fmt.Println(area)
}
```
## If/Else
```go
func printParity(x int) {
    if x % 2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```
>- The else can be removed in the above example
```go
func printParity(x int) {
    if x % 2 == 0 {
        fmt.Println("Even")
    }
    fmt.Println("Odd")
}
```
>- This concise minimal form is preferred in Go
>- If statement support shorthand variable definition
>- The variable definition precedes the conditional, separated by a semicolon
>- The scope is only in the if-else block.
```go
func printParity(x int) {
    if x := x % 2; x == 0 {
        fmt.Println("Even")
        return
    }
    // fmt.Println(x) // error: x is not defined
    fmt.Println("Odd")
}
```
## Error handling
>- Go has no exception type. Error handling/checking is part of the regular program flow
>- Go provides the builtin **error** type - its zero value is nil
>- Errors are initialized in 2 ways:
>- using errors.New("msg") or using
>- fmt.Errorf("msg:", x)
>- Error handling is explicit by design
>- Go uses multiple return values, pointers and if/else to handle errors
```go
func area(x int, y int) (*int, error) {
    if x == 0 || y == 0 {
        return nil, fmt.Errorf(zero area: [%v, %v]", x, y)
    }
    area := x * y
    return &area, nil
}

func main() {
    area, err := area(10, 5)
    if err != nil {
        fmt.Println("error", err)
        return // we print the error and stop code execution at this point
    }
    fmt.Printf("Area: %v\n", *area)
}
```
>- Another way to handle errors is by forcefully exciting using the **panic** keyword
>- also see recover ... Use them both sparingly
```go
func main() {
    area, err := area(10, 5)
    if err != nil {
        panic("error", err)
    }
    fmt.Printf("Area: %v\n", *area)
}
```
## Deferred functions
>- defer keyword is applied before a function call - defer myfunc()
>- Delays function execution until the surrounding function finishes
>- In case of multiple deferred functions, the execution order will be LIFO (last in,
>- first out.
>- Guaranteed run - deferred functions are often used fir cleanup operations
>- A deferred function's parameters are evaluated first, then used later during the
>- deferred function invocation
```go
package main

import "fmt"

func printWord(id string, word string) {
    fmt.Printf("[%v]: %v\n", id, word) 
}

func main() {
    word := "word"
    defer printWord("Defer 0", word)
    defer printWord("Defer 1", word)
    
    word = "other"
    fmt.Println("Main has exited")
}
// Output
// Main has exited.
// Defer 1 word
// Defer 0 word
```
