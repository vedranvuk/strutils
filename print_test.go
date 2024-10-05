package strutils

import (
	"fmt"
	"testing"
)

func TestByteString(t *testing.T) {
	var b = []byte("sinferopopokatepetl")
	if testing.Verbose() {
		fmt.Print(ByteString(b, 8))
	}
}

func BenchmarkByteString(b *testing.B) {
	for l := 1; l < 1e6; l *= 10 {
		b.Run(fmt.Sprintf("%d", l), func(b *testing.B) {
			var buf = []byte(RandomString(true, true, true, l))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ByteString(buf, 8)
			}
		})
	}
}
