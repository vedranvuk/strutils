package strutils

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	var foo = NewFoo()
	for i := 0; i < 100; i++ {
		if testing.Verbose() {
			fmt.Println(foo.Name())
			fmt.Println(foo.EMail())
		}
	}
}

func BenchmarkWord(b *testing.B) {
	var foo =  NewFoo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo.word(10)
	}
}