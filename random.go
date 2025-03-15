// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"math/rand"
	"strings"
	"unsafe"
)

// Default special characters set used in passwords.
// < and > may cause issues on some systems.
var DefSpecialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

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

var randomFuncs []func() string

func init() {
	randomFuncs = []func() string{
		RandomLower,
		RandomUpper,
		RandomNum,
	}
}

// Returns a random string of "length".
// If "lo" includes lowercase letters.
// If "up" includes uppercase letters.
// If "num" includes numbers.
func RandomString(lo, up, nums bool, length int) string {
	if length < 1 {
		return ""
	}
	var f = randomFuncs[:0]
	if lo {
		f = append(f, RandomLower)
	}
	if up {
		f = append(f, RandomUpper)
	}
	if nums {
		f = append(f, RandomNum)
	}
	var sb strings.Builder
	sb.Grow(length)
	if len(f) > 0 {
		for i := 0; i < length; i++ {
			sb.WriteString(f[rand.Intn(len(f))]())
		}
	}
	return sb.String()
}

// Foo returns various random texts.
type Foo struct {
	// MinName is the minimum length for various name generation functions.
	MinName int
	// MaxName is the maximum length for various name generation functions.
	MaxName int
	// Domains is a set of domains used for generating fake urls or emails.
	Domains []string
}

// NewFoo returns a new Foo.
func NewFoo() Foo {
	return Foo{
		MinName: 2,
		MaxName: 10,

		Domains: []string{".com", ".net", ".org"},
	}
}

// word returns a random word of specified length l.
// It will be a combination of wovels and consonants.
func (self Foo) word(l int) string {
	var (
		out = make([]byte, 0, l)
		nw  = 0
		nc  = 0
	)
	for i := 0; i < l; i++ {
		if nc == 2 {
			goto Wovel
		}
		if nw == 1 {
			goto Consonant
		}
		if self.Bool() {
			goto Wovel
		}
		goto Consonant
	Wovel:
		out = append(out, []byte(Random(Wovels))...)
		nw++
		nc = 0
		continue
	Consonant:
		out = append(out, []byte(Random(Consonants))...)
		nc++
		nw = 0
		continue
	}
	return string(out)
}

// intRng returns a random int between lo and hi.
func (self Foo) intRng(lo, hi int) int { return rand.Intn(hi-lo) + lo }

// Bool returns a random bool.
func (self Foo) Bool() bool { return rand.Intn(2) > 0 }

// Name returns a random name.
func (self Foo) Name() string {
	return PascalCase(self.word(self.intRng(self.MinName, self.MaxName)))
}

// EMail returns a random email where name and domain name are randomly
// generated words and the domain will be one of configured domains.
func (self Foo) EMail() string {
	s := self.word(self.intRng(self.MinName, self.MaxName))
	s += "@"
	s += self.word(self.intRng(self.MinName, self.MaxName))
	s += self.Domains[rand.Intn(len(self.Domains))]
	return s
}
