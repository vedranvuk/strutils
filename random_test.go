package strutils

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	var foo = NewFoo()
	for i := 0; i < 100; i++ {
		fmt.Println(foo.Name())
		fmt.Println(foo.EMail())
	}
}