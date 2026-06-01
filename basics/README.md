# Basics

### Data Types
- int - integers
- float32 - floating point numbers
- string - stream of characters
- bool - boolean (true and false)

### Declaring Variables
``` go
var varname type = value; // using var keyword and specifying type
var varname = value; // no type declared, type is inferred from value
varname := value; // no var keyword and type, type is inferred
```

### Constants
Constants are declared using **const** keyword

```go
const varname type = value;
```


### outputs
Using functions - Print(), Println(), Printf()
```go
fmt.Print(i, "\n") // print with default format
fmt.Println(i,j) // print with a whitespace is added between the arguments, and a newline is added at the end
fmt.Printf("i has value: %v and type: %T\n", i, i) //%v is used to print the value of the arguments, %T is used to print the type of the arguments

```