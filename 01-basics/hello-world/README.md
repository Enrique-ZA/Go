# GO basics:
## Go commands
### - Hello World example
```bash
go mod init example/hello
```
```bash
go run . # compiles code, creates and executable, and then runs it
```
```bash
go build hello.go # compiles code, creates and executable
```
```bash
go test # runs tests, including coverage and benchmarking
```
## Why are Go binaries so big?
>- Go binaries are **statically linked standalones**
>- Building a Go binary is the same, for debug or release!
## Declaring variables
### - Strings:
```go
// Example string, must be declared in scope and used
var Greeting = "Hell" // available outside of scope 
var greeting1 string
var greeting2 = "Hello, W"
greeting3 := "Hello, World!" greeting1 = "Hello" 
var greeting4 = "Hello, World!\nGoodbye, World!"
var raw = `Hello, World!
           Goodbye, World!`
```
## Basic data types in Go:
>- _numbers, strings, booleans_
>- **Composite types**: _arrays, structs, pointers, functions, interfaces, slices, maps,_ 
>  _and channels_
>- int8 (8 bits, -128 to 127)
>- int16 (16 bits, -32768 to 32767)
>- int32 (32 bits, -2147483648 to 2147483647)
>- int64 (64 bits, -9223372036854775808 to 9223372036854775807)
>- uint8 (8 bits, 0 to 255)
>- uint16 (16 bits, 0 to 65535)
>- uint32 (32 bits, 0 to 4294967295)
>- uint64 (64 bits, 0 to 18446744073709551615)
>- int (system dependent, system dependent)
>- uint (system dependent, system dependent)
>- **byte** is an alias for uint8 and **rune** is an alias for int32
>- 2 floating point types: _float32_ and _float64_
>  f := 123.45 will be allocated as a _float64_
>- **bool** is a boolean type - true or false 
>  uses &&, || and !
>  && and || follow short-circuiting rules
>- Unless explicity initialized, vars are automatically initialized with their zero value
>  Integer (0), Floating point (0.0), Boolean (false), String ("")
>- Data types are never nil and need not nil checked
```go
var (
    v0 int
    v1 float64
    v2 bool
    v3 string
)
// print example
fmt.Printf("[v0]: var type = %T; value = %v;\n", v0, v0)
```
>- **Pointer**: stores the memory address of a variable.
>  It provides a way to point to where the memory is located
>  and find the value stored at that address.
>  The * operator (dereferencing operator)
>  The & operator (address operator)
>  Zero value: nil
>  For example: 
```go
var a string = "Hello, World!"
var prt *string = &a
var b string = *prt
```
> Operations on nil pointers will cause a runtime error
