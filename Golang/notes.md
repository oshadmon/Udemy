# Week 1 
## basics 
* Workspace is a hierarchy of directories
  * src - source files 
  * pkg - libraries 
    * group of related source code files
  * bin - executables 

**Commands**: 
* `go doc` - `man` for _Golang_
* `go fmt` - format source-code files 
* `go get` - install packages 
* `go list` - view installed packages
* `go run` - build and execute 
  * `go build` - converts go into executable 
* `go test` - runs files with "_test.go" in naming

[hello_world](hello_world.go)

## pointers & variables
* **pointers**: an address of data in memory
  * & returns the address of the variable / function
  * \* returns the data at an address
  * `new()` - returns pointer to variable
* **blocks**: A sequence of declarations and statements within  curly brackets `{}`
* **lexical scope**: layers of blocks 
* variable is accessible from block 
  * functions can access variables that are within their block || an outter block 

[pointers](pointer.go)
[variable scope](variable_scope.go)

## memory allocation & garbage collection
* we need to deallocate a variable when it's no longer needed  
* **stack** - dedicated to function calls 
  * automatically deallocated automatically 
* **heap** - persistent memory 
  * needs to be **explicitly** deallocated 
* hard to know when to deallocate a variable
* **garbage collection** - automatic deallocate variable(s), used by interpreter (ex. Java and Python)
* _Golang_ is a compiled langauge with garbage collection built-in 


## Data Types 
* users can specify generic data type (`int`) or specific (`int32`)
  * once a variable is declared as a (specific) type it _cannot_ be changed 
```go
var x int32 = 1 
var y int16 = 2 

// convert y to int32 
var z int32 = int32(y) 
```

* **ASCII** - (English) character 8-bit code per character 
  * A - 0x41 
* **Unicode** - 32-bit code 
* **UTF-8** - variable length
  * first set (128 characters) of codes match _ASCII_ code
  * default for Golang
* **IOTA** - generates distinct but related constants 
  * example - days of the week 
  * has several distinct values 
```go
type Grade int
const (
    A Grade = iota
    B
    C
    D
    F
)
  
func main(){
    fmt.Printf("A - %d | B - %d | C - %d | D - %d | F - %d\n", A, B, C, D, F)
}
```
* **Array** - fixed series of elements of a given type 
* **Slice** - window on an underlying array 
  * has changeable size upto the size of the associated array 
  * _Point_: start of the slice
  * _Length_: number of elements in the slize 
  * _Capacity_: number of values in the slice

* **hash table**: key/value pair table 
* 