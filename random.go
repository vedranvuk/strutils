// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"math/rand"
	"strings"
)

var (
	Wovels        = "aeiou"
	Consonants    = "bcdfghjklmnpqrstvxyz"
	SpecialChars  = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	Numerals      = "0123456789"
	AlphaUpper    = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	AlphaLower    = "abcdefghijklmnopqrstuvxyz"
	Alpha         = AlphaUpper + AlphaLower
	AlphaNumerals = Numerals + Alpha
)

var (
	WovelsBytes        = []byte(Wovels)
	ConsonantsBytes    = []byte(Consonants)
	SpecialCharsBytes  = []byte(SpecialChars)
	NumeralsBytes      = []byte(Numerals)
	AlphaUpperBytes    = []byte(AlphaUpper)
	AlphaLowerBytes    = []byte(AlphaLower)
	AlphaBytes         = []byte(Alpha)
	AlphaNumeralsBytes = []byte(AlphaNumerals)
)

// Random returns a string containing a single character from set.
func Random(set []byte) string {
	var i = rand.Intn(len(set))
	return UnsafeString(set[i : i+1])
}

// Random returns a string containing a single character from set.
func RandomByte(set []byte) []byte {
	var i = rand.Intn(len(set))
	return set[i : i+1]
}

// Randoms returns a string of length consisting of random characters from s.
func Randoms(set []byte, length int) string {
	if length < 1 {
		return ""
	}
	var r = make([]byte, length, length)
	for i := 0; i < length; i++ {
		r[i] = set[rand.Intn(len(set))]
	}
	return UnsafeString(r)
}

// Returns a string containing a random number.
func RandomNum() string {
	return Random(NumeralsBytes)
}

// Returns a string of random numbers of length.
func RandomNums(length int) string {
	return Randoms(NumeralsBytes, length)
}

// Returns a string containing a random uppercase letter.
func RandomUpper() string {
	return Random(AlphaUpperBytes)
}

// Returns a string of random uppercase letters of length.
func RandomUppers(length int) string {
	return Randoms(AlphaUpperBytes, length)
}

// Returns a string containing a random lowercase letter.
func RandomLower() string {
	return Random(AlphaLowerBytes)
}

// Returns a string of random lowercase letters of length.
func RandomLowers(length int) string {
	return Randoms(AlphaLowerBytes, length)
}

// Returns a string containing a random password special character.
func RandomSpecial() string {
	return Random(SpecialCharsBytes)
}

// Returns a string of random special characters of length.
func RandomSpecials(length int) string {
	return Randoms(SpecialCharsBytes, length)
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
		out = append(out, RandomByte(WovelsBytes)...)
		nw++
		nc = 0
		continue
	Consonant:
		out = append(out, RandomByte(ConsonantsBytes)...)
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
