package strutils

import (
	"encoding"
	"errors"
	"strconv"
	"time"
)

// Converter converts strings and string slices into basic go types.
type Converter struct {
	TimeFormat string
}

// NewConverter returns a new Converter with default values.
func NewConverter() Converter {
	return Converter{
		TimeFormat: time.RFC3339Nano,
	}
}

// Value is an interface which can convert a string to self.
type Value interface {
	Set(string) error
}

// StringToAny sets v which must be a pointer to a supported type from raw
// or returns an error if conversion error occured.
func (self Converter) StringToAny(in, out any) (err error) {
	switch tin := in.(type) {
	case string:
		return self.setString(tin, out)
	case []string:
		switch tout := out.(type) {
		case *[]string:
			*tout = tin
			return nil
		}
	}
	return errors.New("unsupported input type")
}

func (self Converter) setString(in string, out any) (err error) {
	switch p := out.(type) {
	case *string:
		*p = in
	case *bool:
		var b bool
		if b, err = strconv.ParseBool(in); err == nil {
			*p = b
		}
	case *int:
		var v int64
		if v, err = strconv.ParseInt(in, 10, 0); err == nil {
			*p = int(v)
		}
	case *uint:
		var v uint64
		if v, err = strconv.ParseUint(in, 10, 0); err == nil {
			*p = uint(v)
		}
	case *int8:
		var v int64
		if v, err = strconv.ParseInt(in, 10, 8); err == nil {
			*p = int8(v)
		}
	case *uint8:
		var v uint64
		if v, err = strconv.ParseUint(in, 10, 8); err == nil {
			*p = uint8(v)
		}
	case *int16:
		var v int64
		if v, err = strconv.ParseInt(in, 10, 16); err == nil {
			*p = int16(v)
		}
	case *uint16:
		var v uint64
		if v, err = strconv.ParseUint(in, 10, 16); err == nil {
			*p = uint16(v)
		}
	case *int32:
		var v int64
		if v, err = strconv.ParseInt(in, 10, 32); err == nil {
			*p = int32(v)
		}
	case *uint32:
		var v uint64
		if v, err = strconv.ParseUint(in, 10, 32); err == nil {
			*p = uint32(v)
		}
	case *int64:
		*p, err = strconv.ParseInt(in, 10, 64)
	case *uint64:
		*p, err = strconv.ParseUint(in, 10, 64)
	case *float32:
		var v float64
		if v, err = strconv.ParseFloat(in, 64); err == nil {
			*p = float32(v)
		}
	case *float64:
		*p, err = strconv.ParseFloat(in, 64)
	case *time.Duration:
		*p, err = time.ParseDuration(in)
	case *time.Time:
		*p, err = time.Parse(self.TimeFormat, in)
	default:
		if v, ok := p.(encoding.TextUnmarshaler); ok {
			return v.UnmarshalText([]byte(in))
		}
		if v, ok := p.(Value); ok {
			err = v.Set(in)
		} else {
			return errors.New("incompatible target var")
		}
	}
	return
}
