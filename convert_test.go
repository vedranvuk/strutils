package strutils

import (
	"reflect"
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
		Complex64     complex64
		Complex128    complex128
		Duration      time.Duration
		Time          time.Time
		Custom        Custom
		Unmarshalable Unmarshalable

		StringSlice     []string
		BoolSlice       []bool
		IntSlice        []int
		UIntSlice       []uint
		Int8Slice       []int8
		UInt8Slice      []uint8
		Int16Slice      []int16
		UInt16Slice     []uint16
		Int32Slice      []int32
		UInt32Slice     []uint32
		Int64Slice      []int64
		UInt64Slice     []uint64
		Float32Slice    []float32
		Float64Slice    []float64
		Complex64Slice  []complex64
		Complex128Slice []complex128
		DurationSlice   []time.Duration
		TimeSlice       []time.Time
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

	if err := c.StringToAny("1.0i", &Complex64); err != nil {
		t.Fatal(err)
	}
	if Complex64 != 1.0i {
		t.Fatal()
	}

	if err := c.StringToAny("1.0i", &Complex128); err != nil {
		t.Fatal(err)
	}
	if Complex128 != 1.0i {
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

	// Slice Tests
	if err := c.StringToAny("str1, str2", &StringSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(StringSlice, []string{"str1", "str2"}) {
		t.Fatalf("StringSlice conversion failed: got %v, want %v", StringSlice, []string{"str1", "str2"})
	}

	if err := c.StringToAny("true,false", &BoolSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(BoolSlice, []bool{true, false}) {
		t.Fatalf("BoolSlice conversion failed: got %v, want %v", BoolSlice, []bool{true, false})
	}

	if err := c.StringToAny("1,2", &IntSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(IntSlice, []int{1, 2}) {
		t.Fatalf("IntSlice conversion failed: got %v, want %v", IntSlice, []int{1, 2})
	}

	if err := c.StringToAny("1,2", &UIntSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(UIntSlice, []uint{1, 2}) {
		t.Fatalf("UIntSlice conversion failed: got %v, want %v", UIntSlice, []uint{1, 2})
	}

	if err := c.StringToAny("1,2", &Int8Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Int8Slice, []int8{1, 2}) {
		t.Fatalf("Int8Slice conversion failed: got %v, want %v", Int8Slice, []int8{1, 2})
	}

	if err := c.StringToAny("1,2", &UInt8Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(UInt8Slice, []uint8{1, 2}) {
		t.Fatalf("UInt8Slice conversion failed: got %v, want %v", UInt8Slice, []uint8{1, 2})
	}

	if err := c.StringToAny("1,2", &Int16Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Int16Slice, []int16{1, 2}) {
		t.Fatalf("Int16Slice conversion failed: got %v, want %v", Int16Slice, []int16{1, 2})
	}

	if err := c.StringToAny("1,2", &UInt16Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(UInt16Slice, []uint16{1, 2}) {
		t.Fatalf("UInt16Slice conversion failed: got %v, want %v", UInt16Slice, []uint16{1, 2})
	}

	if err := c.StringToAny("1,2", &Int32Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Int32Slice, []int32{1, 2}) {
		t.Fatalf("Int32Slice conversion failed: got %v, want %v", Int32Slice, []int32{1, 2})
	}

	if err := c.StringToAny("1,2", &UInt32Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(UInt32Slice, []uint32{1, 2}) {
		t.Fatalf("UInt32Slice conversion failed: got %v, want %v", UInt32Slice, []uint32{1, 2})
	}

	if err := c.StringToAny("1,2", &Int64Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Int64Slice, []int64{1, 2}) {
		t.Fatalf("Int64Slice conversion failed: got %v, want %v", Int64Slice, []int64{1, 2})
	}

	if err := c.StringToAny("1,2", &UInt64Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(UInt64Slice, []uint64{1, 2}) {
		t.Fatalf("UInt64Slice conversion failed: got %v, want %v", UInt64Slice, []uint64{1, 2})
	}

	if err := c.StringToAny("1.1,2.2", &Float32Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Float32Slice, []float32{1.1, 2.2}) {
		t.Fatalf("Float32Slice conversion failed: got %v, want %v", Float32Slice, []float32{1.1, 2.2})
	}

	if err := c.StringToAny("1.1,2.2", &Float64Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Float64Slice, []float64{1.1, 2.2}) {
		t.Fatalf("Float64Slice conversion failed: got %v, want %v", Float64Slice, []float64{1.1, 2.2})
	}

	if err := c.StringToAny("1.0i,2.0i", &Complex64Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Complex64Slice, []complex64{1.0i, 2.0i}) {
		t.Fatalf("Complex64Slice conversion failed: got %v, want %v", Complex64Slice, []complex64{1.0i, 2.0i})
	}

	if err := c.StringToAny("1.0i,2.0i", &Complex128Slice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(Complex128Slice, []complex128{1.0i, 2.0i}) {
		t.Fatalf("Complex128Slice conversion failed: got %v, want %v", Complex128Slice, []complex128{1.0i, 2.0i})
	}

	if err := c.StringToAny("1s,2s", &DurationSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(DurationSlice, []time.Duration{time.Second, 2 * time.Second}) {
		t.Fatalf("DurationSlice conversion failed: got %v, want %v", DurationSlice, []time.Duration{time.Second, 2 * time.Second})
	}

	ts1 := "2024-10-06T12:00:00.000000000+02:00"
	ts2 := "2024-10-07T12:00:00.000000000+02:00"
	timeStr := ts1 + "," + ts2
	tm1, _ := time.Parse(c.TimeFormat, ts1)
	tm2, _ := time.Parse(c.TimeFormat, ts2)
	expectedTimeSlice := []time.Time{tm1, tm2}

	if err := c.StringToAny(timeStr, &TimeSlice); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(TimeSlice, expectedTimeSlice) {
		t.Fatalf("TimeSlice conversion failed: got %v, want %v", TimeSlice, expectedTimeSlice)
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
	b.Run("complex64", func(b *testing.B) {
		var out complex64
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.0i", &out)
		}
	})
	b.Run("complex128", func(b *testing.B) {
		var out complex128
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.0i", &out)
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

	b.Run("stringSlice", func(b *testing.B) {
		var out []string
		for i := 0; i < b.N; i++ {
			c.StringToAny("str1,str2", &out)
		}
	})
	b.Run("boolSlice", func(b *testing.B) {
		var out []bool
		for i := 0; i < b.N; i++ {
			c.StringToAny("true,false", &out)
		}
	})
	b.Run("intSlice", func(b *testing.B) {
		var out []int
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("uintSlice", func(b *testing.B) {
		var out []uint
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("int16Slice", func(b *testing.B) {
		var out []int16
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("uint16Slice", func(b *testing.B) {
		var out []uint16
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("int32Slice", func(b *testing.B) {
		var out []int32
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("uint32Slice", func(b *testing.B) {
		var out []uint32
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("int64Slice", func(b *testing.B) {
		var out []int64
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("uint64Slice", func(b *testing.B) {
		var out []uint64
		for i := 0; i < b.N; i++ {
			c.StringToAny("1,2", &out)
		}
	})
	b.Run("float32Slice", func(b *testing.B) {
		var out []float32
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.1,2.2", &out)
		}
	})
	b.Run("float64Slice", func(b *testing.B) {
		var out []float64
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.1,2.2", &out)
		}
	})
	b.Run("complex64Slice", func(b *testing.B) {
		var out []complex64
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.0i,2.0i", &out)
		}
	})
	b.Run("complex128Slice", func(b *testing.B) {
		var out []complex128
		for i := 0; i < b.N; i++ {
			c.StringToAny("1.0i,2.0i", &out)
		}
	})
	b.Run("durationSlice", func(b *testing.B) {
		var out []time.Duration
		for i := 0; i < b.N; i++ {
			c.StringToAny("1s,2s", &out)
		}
	})
	b.Run("timeSlice", func(b *testing.B) {
		const ts = "2024-10-06T12:00:00.000000000+02:00,2024-10-07T12:00:00.000000000+02:00"
		var out []time.Time
		for i := 0; i < b.N; i++ {
			c.StringToAny(ts, &out)
		}
	})
}
