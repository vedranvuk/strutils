// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"fmt"
	"strings"
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

	if segment, next := Segment("sinferopopokatepetl", "notfound", 0); segment != "sinferopopokatepetl" || next != -1 {
		t.Fatal("Segment failed")
	}
	if segment, next := Segment("sinferopopokatepetl", "notfound", 7); segment != "popokatepetl" || next != -1 {
		t.Fatal("Segment failed")
	}
	if segment, next := Segment("sinferopopokatepetl", "popo", 0); segment != "sinfero" || next != 11 {
		t.Fatal("Segment failed")
	}
	if segment, next := Segment("sinferopopokatepetl", "te", 11); segment != "ka" || next != 15 {
		t.Fatal("Segment failed")
	}
	if segment, next := Segment("sinferopopokatepetl", "pe", 15); segment != "tl" || next != 19 {
		t.Fatal("Segment failed")
	}

	if segment, next := SegmentFold("sinferopopokatepetl", "pOpO", 0); segment != "sinfero" || next != 11 {
		t.Fatal("Segment failed")
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

const wrapDisText = "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?"

func TestWrapText(t *testing.T) {
	for i := 10; i < 11; i += 10 {
		for _, l := range WrapText(wrapDisText, i, false) {
			fmt.Printf("%s\n", strings.Repeat("-", i))
			fmt.Printf("%s\n", l)
		}
	}
}

func BenchmarkWrapText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WrapText(wrapDisText, 80, false)
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
