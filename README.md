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
