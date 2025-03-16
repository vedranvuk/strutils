package strutils

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestUnsafeStringBytes(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		s := ""
		expected := []byte(nil)
		actual := UnsafeStringBytes(s)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("UnsafeStringBytes(\"\") = %v, want %v", actual, expected)
		}
	})

	t.Run("non-empty string", func(t *testing.T) {
		s := "hello"
		expected := []byte("hello")
		actual := UnsafeStringBytes(s)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("UnsafeStringBytes(%q) = %v, want %v", s, actual, expected)
		}

		// Verify that modifying the slice does not modify the original string.
		// This test can only verify that directly editing the returned byte slice affects the original string,
		// which is not intended.
		// Therefore, it is commented out to prevent false positive test failures.
		/*
		original := s
		actual[0] = 'H'
		if s == original {
			t.Errorf("UnsafeStringBytes(%q) modification did not affect original string. This should not happen", s)
		}
		*/
	})

	t.Run("string with unicode", func(t *testing.T) {
		s := "你好世界"
		expected := []byte(s)
		actual := UnsafeStringBytes(s)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("UnsafeStringBytes(%q) = %v, want %v", s, actual, expected)
		}
	})

	t.Run("long string", func(t *testing.T) {
		s := ""
		for i := 0; i < 1000; i++ {
			s += "a"
		}
		expected := []byte(s)
		actual := UnsafeStringBytes(s)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("UnsafeStringBytes(long string) = %v, want %v", actual, expected)
		}
	})

	t.Run("verify data ptr", func(t *testing.T) {
		s := "hello"
		actual := UnsafeStringBytes(s)

		stringDataPtr := unsafe.StringData(s)
		sliceDataPtr := unsafe.SliceData(actual)

		if stringDataPtr != sliceDataPtr {
			t.Errorf("String data pointer and slice data pointer are different. String Data Ptr: %p, Slice Data Ptr: %p", stringDataPtr, sliceDataPtr)
		}

		if len(s) != len(actual) {
			t.Errorf("String length and slice length are different. String Length: %d, Slice Length: %d", len(s), len(actual))
		}

	})
}

func BenchmarkUnsafeStringBytes_ShortString(b *testing.B) {
	s := "hello"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UnsafeStringBytes(s)
	}
}

func BenchmarkUnsafeStringBytes_LongString(b *testing.B) {
	s := ""
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UnsafeStringBytes(s)
	}
}

func BenchmarkStringBytes_ShortString(b *testing.B) {
	s := "hello"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStringBytes_LongString(b *testing.B) {
	s := ""
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}
