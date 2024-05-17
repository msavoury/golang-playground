# golang-playground
Playground for Learning golang

"Clear is better than clever"

# Run Code
use `go run <file>`
```
go run helloworld.go
```

# Run Tests
use `go test <file>`
```
// with say/ dir:
go test say_test.go
```

# Errors
If you see `cannot find main module` try running `go mod init <module-name>`


## Short declaration operator
Can only be used in functions
```
func main() {
  c :=2
  // print the type of something using %T
  fmt.Printf("a: %T %v", c, c)
}
```

## Slices
Slices have variable length.
```
var a [] int
```
Slices have a descriptor, they are passed by reference; no copyting, updating OK

Slices are indexed like [8:11]
(read as the staring element and one past the ending element, so this way we have 11-8 = 3 elements in our slice)

## Slices vs Arrays

Slices vs Arrays
- variable length / length fixed
- passed by reference / passed by value (copied)
- not comparable / comparable (==)
- cannot be used as map key / can be used as map key
- has copy & append helpers 
- useful as function params / usefuls as pseudo constants

# Maps
var m map[string]int // nil, no storage
p := make(map[string]int) // non-nil but empty

the type used for the key must have == and != defined.

creating a map literal:
```
var m = map[string]int {
  "and": 1,
  "the": 1,
  "or": 2,

}
```

Getting 2 values from a map
```
b, ok := p["and"] // ok will tell you if the key was missing
```

# Loops
// i is index, v is value
```
for i, v := range myArray {
    fmt.Println(i,v)
}
```

```
// note that there is no order in keys
for k, v := range myMap {
  fmt.Println(k,v)
}
```

// infinite loop
for {
    
 // this will require a `break` statement
}

# Switches
```
switch a := f.Get() {
  case 0, 1, 2:
    fmt.println("Underflow possible")
  case 3,4,5: // noop

  default: 
    fmt.Println("warning: overload"
}
```

Alternatives may be empty and do NOT fall through (break is not required)

# Packages
Every standalone program has a main package

Nothing is global it's either in your package or another
It's either at package scope or function scope

You can declare anything at package scope, but you can't use the short declarator

Packages control visibility

Every name that's capitalized is exported

A package can have more than one file. Within a package, everything is visible, even across files.

Each _source_ file in your package must import what it needs
Generally, files of the same package live in the same directory
golang doesn't allow cyclic dependencies.  (in part to keep compiles fast)

A package shoudl embed deep functionality behind a simple API
