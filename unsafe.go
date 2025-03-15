package strutils

import "unsafe"

// UnsafeStringBytes returns bytes of a string without allocating.
//
// Result is a weak, non-garbage collected byte slice that is valid for the 
// lifetime of s.
func UnsafeStringBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}