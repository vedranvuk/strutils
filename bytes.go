// Copyright 2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package strutils

// IsUpper returns true if c is an uppercase alpha.
func IsUpper(c byte) bool { return c >= 'A' && c <= 'Z' }

// IsLower returns true if c is a lowercase alpha.
func IsLower(c byte) bool { return c >= 'a' && c <= 'z' }

// IsLetter returns true if c is lowercase or uppercase alpha.
func IsLetter(c byte) bool { return IsLower(c) || IsUpper(c) }

// IsDigit returns true if c is a digit.
func IsDigit(c byte) bool { return c >= '0' && c <= '9' }

// IsAlphanumeric returns true if c is a digit, lower or upper alpha.
func IsAlphanumeric(c byte) bool { return IsLower(c) || IsUpper(c) || IsDigit(c) }

// ToLower returns lowercased c.
func ToLower(c byte) byte {
	if IsUpper(c) {
		return c + ('a' - 'A')
	}
	return c
}

// ToLower returns uppercased c.
func ToUpper(c byte) byte {
	if IsLower(c) {
		return c - ('a' - 'A')
	}
	return c
}
