# Structures and Interfaces
## Struct basics
>- Allows for the creation of custom types and the grouping of a collection of fields
>- Provides a way to integrate state and behaviour
>- To create a struct we use the type and struct keywords together
```go
type person struct
```
>- Can also be exported outside of package scope by capitalizing the name
```go
type Person struct
```
>- Structs contain named fields
```go
type person struct {
    name string
    age uint8
}
func main() {
    p := person{
        name: "John",
        age:  30,
    }
    fmt.Println(p.name, p.age)
}
```
>- Structs are mutable - The dot (.) operator allows both reading and writing to fields
>- Fields can be optional upon initialization and default to their type zero value
```go
type person struct {
    name string
    age uint8
}
func main() {
    p := person{
        name: "John",
    }
    fmt.Println(p.name, p.age)
}
```
>- It's idiomatic to use functions that encapsulates struct creation
```go
type person struct {
    name string
    age uint8
}

// The below function is idiomatic in Go
func newPerson(name string, age uint8) person {
    return person{
        name: name,
        age: age,
    }
}
```
## Methods
>- Methods are types with special receiver args
>- The receiver name and type appear in brackets between the func keyword and name
```go
func (p person) print() {
    fmt.Println(p.name, p.age)
}
```
>- The above example is equivalent to the below function
```go
func print(p person) {
    fmt.Println(p.name, p.age)
}
```
>- In Go method and field names cannot clash
```go
// The below code will NOT compile since name is both a field and method name for the
// struct person

type person struct {
    name string
}

func (p person) name() {
    return p.name
}
```
>- Method naming should be in MixedCaps. It's not considered idiomatic to explicitly
>- prefix Getters and Setters using Get* and Set*
>- We use pointers for changes to persist outside of the function scope
```go
func (c *city) tempF() {
    c.tempC = (c.tempC * 9 / 5) + 32
}
```
## Modules
>- Modules are how Go manages dependencies
>- A new module is create using **go mod init import_path**
```bash
go mod init example/hello
```
>- The go mod command should be run inside the root directory of the project
>- This will create a go.mod file
>- Dependency source code in Go is downloaded, this can be browsed to see implemented
>- details
>- We can only define one package per directory
>- A method or field cannot be exported if its struct is not exported, however
>- a struct can be exported with only some of its methods and fields
## Interfaces
>- Go does not implement inheritance, but interfaces allows for polymorphism
>- Interfaces are collections of method signatures
>- The compiler will check and enforce interface types
>- We define interfaces using the **typt** and **interface** keywords
```go
type CityTemp interface
```
>- A struct can implement multiple interfaces
>- The empty interface, interface{}, doesn't specify any methods
>- The zero value of interfaces is nil
>- Interfaces are built-in pointer types - unlike structs, they don't need to be 
>- explicitly passed by pointer to functions
>- A common use is to export the interface and leave the underlying struct as package
>- scope
```go
type city struct {
    name string
    tempC float64
}

type CityTemp interface {
    Name() string
    TempC() float64
    TempF() float64
}

func NewCity(n string, t float64) CityTemp {
   return &city{
       name: n,
       tempC: t,
   } 
}

func (c city) Name() string {
    return c.name
}

func (c city) TempC() float64 {
    return c.tempC
}

func (c city) TempF() float64 {
    return (c.tempC * 9 / 5) + 32
}
```
