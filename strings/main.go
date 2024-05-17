package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "élite"
	fmt.Printf("%8T %[1]v\n", s)         // print the string
	fmt.Printf("%8T %[1]v\n", []rune(s)) //print the runes of the string - sequence of 32bit ints
	fmt.Printf("%8T %[1]v\n", []byte(s)) //print the bytes of the string -

	// the length of a string is the length of the number of bytes it takes

	// in go, the length of a string is stored in a string descriptor

	// strings as passed by reference, thus they aren't copied

	//sp := "a string"

	k := strings.ToUpper(s)

	fmt.Println(k)
}
