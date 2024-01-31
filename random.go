// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"math/rand"
	"unsafe"
)

// Default special characters set used in passwords.
// < and > may cause issues on some systems.
const DefSpecialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// toString performs unholy acts to avoid allocations
// https://github.com/kubernetes/kubernetes/blob/e4b74dd12fa8cb63c174091d5536a10b8ec19d34/staging/src/k8s.io/apiserver/pkg/authentication/token/cache/cached_token_authenticator.go#L288-L297
func toString(b []byte) string {
	// unsafe.SliceData relies on cap whereas we want to rely on len
	if len(b) == 0 {
		return ""
	}
	// Copied from go 1.20.1 strings.Builder.String
	// https://github.com/golang/go/blob/202a1a57064127c3f19d96df57b9f9586145e21c/src/strings/builder.go#L48
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// Random returns a string containing a single character from set.
func Random(set string) string {
	return toString([]byte{set[rand.Intn(len(set))]})
}

// Randoms returns a string of length consisting of random characters from s.
func Randoms(set string, length int) string {
	if length < 1 {
		return ""
	}
	var r = make([]byte, length, length)
	for i := 0; i < length; i++ {
		r[i] = set[rand.Intn(len(set))]
	}
	return toString(r)
}

// Returns a string containing a random number.
func RandomNum() string {
	return Random(Nums)
}

// Returns a string of random numbers of length.
func RandomNums(length int) string {
	return Randoms(Nums, length)
}

// Returns a string containing a random uppercase letter.
func RandomUpper() string {
	return Random(AlphaUpper)
}

// Returns a string of random uppercase letters of length.
func RandomUppers(length int) string {
	return Randoms(AlphaUpper, length)
}

// Returns a string containing a random lowercase letter.
func RandomLower() string {
	return Random(AlphaLower)
}

// Returns a string of random lowercase letters of length.
func RandomLowers(length int) string {
	return Randoms(AlphaLower, length)
}

// Returns a string containing a random password special character.
func RandomSpecial() string {
	return Random(DefSpecialChars)
}

// Returns a string of random special characters of length.
func RandomSpecials(length int) string {
	return Randoms(DefSpecialChars, length)
}

// Returns a random string of "length".
// If "lo" includes lowercase letters.
// If "up" includes uppercase letters.
// If "num" includes numbers.
func RandomString(lo, up, nums bool, length int) string {
	if length < 1 {
		return ""
	}
	f := []func() string{}
	if lo {
		f = append(f, RandomLower)
	}
	if up {
		f = append(f, RandomUpper)
	}
	if nums {
		f = append(f, RandomNum)
	}
	r := make([]byte, length, length)
	if lo || up || nums {
		for i := 0; i < length; i++ {
			r[i] = byte(f[rand.Intn(len(f))]()[0])
		}
	}
	return toString(r)
}
