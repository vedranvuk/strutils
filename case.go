// Copyright 2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Original code in this file lifted from:
// https://github.com/segmentio/go-snakecase
// https://github.com/segmentio/go-camelcase

package strutils

import "errors"

// Camelcase the given string.
func CamelCase(s string) string {
	var b = camelcase(s)
	// the first byte must always be lowercase
	if len(b) != 0 {
		b[0] = ToLower(b[0])
	}
	return toString(b)
}

// Pascalcase the given string.
func PascalCase(s string) string {
	var b = camelcase(s)
	// the first byte must always be uppercase
	if len(b) != 0 {
		b[0] = ToUpper(b[0])
	}
	return toString(b)
}

// Snakecase the given string.
func SnakeCase(s string) string { return separatorCase(s, underscoreByte) }

// Kebabcase the given string.
func KebabCase(s string) string { return separatorCase(s, dashByte) }

func camelcase(s string) []byte {
	b := make([]byte, 0, 64)
	l := len(s)
	i := 0

	for i < l {

		// skip leading bytes that aren't letters or digits
		for i < l && !IsAlphanumeric(s[i]) {
			i++
		}

		// set the first byte to uppercase if it needs to
		if i < l {
			c := s[i]

			// simply append contiguous digits
			if IsDigit(c) {
				for i < l {
					if c = s[i]; !IsDigit(c) {
						break
					}
					b = append(b, c)
					i++
				}
				continue
			}

			// the sequence starts with and uppercase letter, we append
			// all following uppercase letters as equivalent lowercases
			if IsUpper(c) {
				b = append(b, c)
				i++

				for i < l {
					if c = s[i]; !IsUpper(c) {
						break
					}
					b = append(b, ToLower(c))
					i++
				}

			} else {
				b = append(b, ToUpper(c))
				i++
			}

			// append all trailing lowercase letters
			for i < l {
				if c = s[i]; !IsLower(c) {
					break
				}
				b = append(b, c)
				i++
			}
		}
	}
	return b
}

const (
	underscoreByte = '_'
	dashByte       = '-'
)

// separatorCase the given string.
func separatorCase(s string, separator byte) string {
	idx := 0
	hasLower := false
	hasSeparator := false
	lowercaseSinceSeparator := false

	// loop through all good characters:
	// - lowercase
	// - digit
	// - underscore (as long as the next character is lowercase or digit)
	for ; idx < len(s); idx++ {
		if IsLower(s[idx]) {
			hasLower = true
			if hasSeparator {
				lowercaseSinceSeparator = true
			}
			continue
		} else if IsDigit(s[idx]) {
			continue
		} else if s[idx] == separator && idx > 0 && idx < len(s)-1 && (IsLower(s[idx+1]) || IsDigit(s[idx+1])) {
			hasSeparator = true
			lowercaseSinceSeparator = false
			continue
		}
		break
	}

	if idx == len(s) {
		return s // no changes needed, can just borrow the string
	}

	// if we get here then we must need to manipulate the string
	b := make([]byte, 0, 64)
	b = append(b, s[:idx]...)

	if IsUpper(s[idx]) && (!hasLower || hasSeparator && !lowercaseSinceSeparator) {
		for idx < len(s) && (IsUpper(s[idx]) || IsDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (IsLower(s[idx]) || IsDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}

	for idx < len(s) {
		if !IsAlphanumeric(s[idx]) {
			idx++
			continue
		}

		if len(b) > 0 {
			b = append(b, separator)
		}

		for idx < len(s) && (IsUpper(s[idx]) || IsDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (IsLower(s[idx]) || IsDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}
	return toString(b) // return manipulated string
}

var asciiLowercaseArray = [256]byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	' ', '!', '"', '#', '$', '%', '&', '\'',
	'(', ')', '*', '+', ',', '-', '.', '/',
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', ':', ';', '<', '=', '>', '?',
	'@',

	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	'x', 'y', 'z',

	'[', '\\', ']', '^', '_',
	'`', 'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	'x', 'y', 'z', '{', '|', '}', '~', 0x7f,
	0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
	0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
	0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
	0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
	0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7,
	0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
	0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7,
	0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf,
	0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7,
	0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf,
	0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7,
	0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf,
	0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7,
	0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef,
	0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7,
	0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff,
}

// CaseMapping specifies one of case mapping supported by this package.
type CaseMapping int

const (
	// InvalidMapping is the invalid/undefined mapping.
	InvalidMapping CaseMapping = iota
	// NoMapping specifies no mapping.
	NoMapping
	// PascalMapping specifies PascalCase mapping.
	PascalMapping
	// SnakeMapping specifies snake_case mapping.
	SnakeMapping
	// CamelMapping specifies camelCase mapping.
	CamelMapping
	// KebabMapping specifies kebab-case mapping.
	KebabMapping
)

// String implements stringer on CaseMapping.
func (self CaseMapping) String() string {
	switch self {
	case NoMapping:
		return "NoMapping"
	case PascalMapping:
		return "PascalMapping"
	case SnakeMapping:
		return "SnakeMapping"
	case CamelMapping:
		return "CamelMapping"
	case KebabMapping:
		return "KebabMapping"
	default:
		return "InvalidMapping"
	}
}

// MarshalText implementes encoding.TextMarshaler on CaseMapping.
func (self CaseMapping) MarshalText() (text []byte, err error) {
	return []byte(self.String()), nil
}

// UnmarshalText implementes encoding.TextUnmarshaler on CaseMapping.
func (self *CaseMapping) UnmarshalText(text []byte) error {
	switch string(text) {
	case "NoMapping":
		*self = NoMapping
	case "PascalMapping":
		*self = PascalMapping
	case "SnakeMapping":
		*self = SnakeMapping
	case "CamelMapping":
		*self = CamelMapping
	case "KebabMapping":
		*self = KebabMapping
	default:
		*self = InvalidMapping
	}
	return errors.New("unknown mapping: " + string(text))
}

// Map case maps s depending on self value.
// If mapping value is unknown input string is returned unmodified.
func (self CaseMapping) Map(s string) string {
	switch self {
	case PascalMapping:
		return PascalCase(s)
	case SnakeMapping:
		return SnakeCase(s)
	case CamelMapping:
		return CamelCase(s)
	case KebabMapping:
		return KebabCase(s)
	}
	return s
}
