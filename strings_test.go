// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"testing"
)

func TestStringFunctions(t *testing.T) {
	in := "teststring"

	if FetchLeft(in, "str") != "test" {
		t.Error("FetchLeft() failed.")
	}

	if FetchRight(in, "est") != "string" {
		t.Error("FetchRight() failed.")
	}

	if FetchLeftFold(in, "STR") != "test" {
		t.Error("FetchLeftFold() failed.")
	}

	if FetchRightFold(in, "EST") != "string" {
		t.Error("FetchRightFold() failed.")
	}

	if !HasPrefixFold(in, "TEST") {
		t.Error("HasPrefixFold() failed.")
	}

	if !HasSuffixFold(in, "STRING") {
		t.Error("HasSuffixFold() failed.")
	}

	i := Indexes("a.b.c.d.e", ".")
	if len(i) != 4 {
		t.Error("Indexes() failed.")
	}
	if i[0] != 1 && i[1] != 3 && i[2] != 5 && i[3] != 7 {
		t.Error("Indexes() failed.")
	}

	j := IndexesFold("1a2a3a4a5", "A")
	if len(j) != 4 {
		t.Error("Indexes() failed.")
	}
	if j[0] != 1 && j[1] != 3 && j[2] != 5 && j[3] != 7 {
		t.Error("Indexes() failed.")
	}

	if !MatchesWildcard("Dickson", "?ic*n") {
		t.Error("MatchesWildcard() failed")
	}
	if !MatchesWildcard("Sensible", "se?s*le") {
		t.Error("MatchesWildcard() failed")
	}
	if MatchesWildcard("ThisIsWrong", "?hi*IswrongO") {
		t.Error("MatchesWildcard() failed")
	}
}

func TestAsciiFunctions(t *testing.T) {
	if !IsNumsOnly("12345") {
		t.Error("IsNumsOnly() failed.")
	}
	if !IsAlphaLowerOnly("abcde") {
		t.Error("IsAlphaLowerOnly() failed.")
	}
	if !IsAlphaUpperOnly("ABCDE") {
		t.Error("IsAlphaUpperOnly() failed.")
	}
	if !IsAlphaOnly("abcdeABCDE") {
		t.Error("IsAlphaOnly() failed.")
	}
	if !IsAlphaNumsOnly("12345abcdeABCDE") {
		t.Error("IsAlphaNumsOnly() failed.")
	}
}

func TestUnpack(t *testing.T) {
	if res, ok := Unwrap("leftwordright", "left", "right"); !ok {
		t.Fatal("unpack failed to find prefix and/or suffix")
	} else if res != "word" {
		t.Fatal("unpack failed")
	}

	if res, ok := Unwrap("word", "", ""); !ok {
		t.Fatal("unpack failed to find prefix and/or suffix")
	} else if res != "word" {
		t.Fatal("unpack failed")
	}
}

func BenchmarkMatchesWildCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MatchesWildcard("Sinferopopokatepetl", "Si?fero*ka?epe?l")
	}
}

func BenchmarkRandomUppers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomUppers(10)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(true, true, true, 10)
	}
}

