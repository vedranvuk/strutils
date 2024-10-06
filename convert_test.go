package strutils

import (
	"strconv"
	"testing"
	"time"
)

type Custom int64

func (self *Custom) Set(s string) (err error) {
	var i int64
	if i, err = strconv.ParseInt(s, 10, 64); err == nil {
		*self = Custom(i)
	}
	return
}

type Unmarshalable int64

func (self *Unmarshalable) UnmarshalText(text []byte) (err error) {
	var i int64
	if i, err = strconv.ParseInt(string(text), 10, 64); err == nil {
		*self = Unmarshalable(i)
	}
	return
}

func TestConverter(t *testing.T) {

	var (
		InvalidTarget struct{}

		String        string
		Bool          bool
		Int           int
		UInt          uint
		Int8          int8
		UInt8         uint8
		Int16         int16
		UInt16        uint16
		Int32         int32
		UInt32        uint32
		Int64         int64
		UInt64        uint64
		Float32       float32
		Float64       float64
		Duration      time.Duration
		Time          time.Time
		Custom        Custom
		Unmarshalable Unmarshalable
	)

	c := NewConverter()

	if err := c.StringToAny("String", &InvalidTarget); err == nil {
		t.Fatal("did not detect unsupported target")
	}

	if err := c.StringToAny("String", &String); err != nil {
		t.Fatal(err)
	}
	if String != "String" {
		t.Fatal()
	}

	if err := c.StringToAny("true", &Bool); err != nil {
		t.Fatal(err)
	}
	if Bool != true {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Int); err != nil {
		t.Fatal(err)
	}
	if Int != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &UInt); err != nil {
		t.Fatal(err)
	}
	if UInt != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Int8); err != nil {
		t.Fatal(err)
	}
	if Int8 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &UInt8); err != nil {
		t.Fatal(err)
	}
	if UInt8 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Int16); err != nil {
		t.Fatal(err)
	}
	if Int16 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &UInt16); err != nil {
		t.Fatal(err)
	}
	if UInt16 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Int32); err != nil {
		t.Fatal(err)
	}
	if Int32 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &UInt32); err != nil {
		t.Fatal(err)
	}
	if UInt32 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Int64); err != nil {
		t.Fatal(err)
	}
	if Int64 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &UInt64); err != nil {
		t.Fatal(err)
	}
	if UInt64 != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("3.14", &Float32); err != nil {
		t.Fatal(err)
	}
	if Float32 != 3.14 {
		t.Fatal()
	}

	if err := c.StringToAny("3.14", &Float64); err != nil {
		t.Fatal(err)
	}
	if Float64 != 3.14 {
		t.Fatal()
	}

	if err := c.StringToAny("5s", &Duration); err != nil {
		t.Fatal(err)
	}
	if Duration != (5 * time.Second) {
		t.Fatal()
	}
	const ts = "2024-10-06T12:00:00.000000000+02:00"
	if err := c.StringToAny(ts, &Time); err != nil {
		t.Fatal(err)
	}
	var tm, err = time.Parse(c.TimeFormat, ts)
	if err != nil {
		t.Fatal(err)
	}
	if !Time.Equal(tm) {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Custom); err != nil {
		t.Fatal(err)
	}
	if Custom != 69 {
		t.Fatal()
	}

	if err := c.StringToAny("69", &Unmarshalable); err != nil {
		t.Fatal(err)
	}
	if Unmarshalable != 69 {
		t.Fatal()
	}
}

func BenchmarkConverter(b *testing.B) {
	c := NewConverter()
	b.Run("string", func(b *testing.B) {
		var out string
		for i := 0; i < b.N; i++ {
			c.StringToAny("string", &out)
		}
	})
	b.Run("bool", func(b *testing.B) {
		var out bool
		for i := 0; i < b.N; i++ {
			c.StringToAny("true", &out)
		}
	})
	b.Run("int", func(b *testing.B) {
		var out int
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("uint", func(b *testing.B) {
		var out uint
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("int16", func(b *testing.B) {
		var out int16
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("uint16", func(b *testing.B) {
		var out uint16
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("int32", func(b *testing.B) {
		var out int32
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("uint32", func(b *testing.B) {
		var out uint32
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("int64", func(b *testing.B) {
		var out int64
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("uint64", func(b *testing.B) {
		var out uint64
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
	b.Run("float32", func(b *testing.B) {
		var out float32
		for i := 0; i < b.N; i++ {
			c.StringToAny("3.14", &out)
		}
	})
	b.Run("float64", func(b *testing.B) {
		var out float64
		for i := 0; i < b.N; i++ {
			c.StringToAny("3.14", &out)
		}
	})
	b.Run("duration", func(b *testing.B) {
		var out time.Duration
		for i := 0; i < b.N; i++ {
			c.StringToAny("5s", &out)
		}
	})
	b.Run("time", func(b *testing.B) {
		const ts = "2024-10-06T12:00:00.000000000+02:00"
		var out time.Time
		for i := 0; i < b.N; i++ {
			c.StringToAny(ts, &out)
		}
	})
	b.Run("custom", func(b *testing.B) {
		var out Custom
		for i := 0; i < b.N; i++ {
			c.StringToAny("69", &out)
		}
	})
}
