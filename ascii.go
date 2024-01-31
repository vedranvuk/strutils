// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"strings"
)

const (
	Nums       = "0123456789"
	AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	AlphaLower = "abcdefghijklmnopqrstuvxyz"
	Alpha      = AlphaUpper + AlphaLower
	AlphaNums  = Nums + Alpha
)

// Checks if "s" consists exclusively of numeric characters.
func IsNumsOnly(s string) bool {
	return strings.IndexAny(s, Nums) > -1
}

// Checks if "s" consists exclusively of lowercase alpha characters.
func IsAlphaLowerOnly(s string) bool {
	return strings.IndexAny(s, AlphaLower) > -1
}

// Checks if "s" consists exclusively of uppercase alpha characters.
func IsAlphaUpperOnly(s string) bool {
	return strings.IndexAny(s, AlphaUpper) > -1
}

// Checks if "s" consists exclusively of alpha characters.
func IsAlphaOnly(s string) bool {
	return strings.IndexAny(s, Alpha) > -1
}

// Checks if "s" consists exclusively of alphanumeric characters.
func IsAlphaNumsOnly(s string) bool {
	return strings.IndexAny(s, AlphaNums) > -1
}
