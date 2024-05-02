package say

// this file and its related file say_test.go show how to perform
// a basic test in go

import (
	"fmt"
)

func Say(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
