package strutils

import (
	"fmt"
	"strings"
)

func ByteString(buf []byte, width int) string {
	var sb = strings.Builder{}
	sb.Grow(len(buf) * 4)
	var c, s = 0, 0
	for i, b := range buf {
		fmt.Fprintf(&sb, "%X ", b)
		if c == width-1 {
			sb.WriteString(" ")
			sb.WriteString(str(buf[s : i+1]))
			sb.WriteRune('\n')
			s = i + 1
			c = 0
			continue
		}
		c++
	}
	if c > 0 {
		for i := 0; i < width-c; i++ {
			sb.WriteString("00 ")
		}
		sb.WriteString(" ")
		sb.WriteString(str(buf[len(buf)-c:]))
		sb.WriteRune('\n')
	}

	return sb.String()
}

func str(in []byte) string {
	var out = make([]byte, 0, len(in))
	for _, b := range in {

		if b < 32 {
			out = append(out, 32)
		}
		out = append(out, b)
	}
	return string(out)
}
