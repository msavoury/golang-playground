package main

// in this example we show how to read arguments from the command line

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s\n", os.Args[1])
	} else {
		fmt.Printf("Hello, %s\n", "Marcos")
	}
}
