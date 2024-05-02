package say

// example of how to perform a test in go
// run this file with `go test`

import (
	"testing"
)

// All tests take a `*testing.T` param
func TestSayHello(t *testing.T) {
	want := "Hello Marcos!"
	got := Say("Marcos")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

}
