// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutils

import (
	"fmt"
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

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?

At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat.`

func TestWrapText(t *testing.T) {
	if testing.Verbose() {
		for _, line := range WrapText(loremIpsum, 100, false) {
			fmt.Println(line)
		}
		return
	}
	WrapText(loremIpsum, 100, false)
}

func BenchmarkWrapText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WrapText(loremIpsum, 80, false)
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
