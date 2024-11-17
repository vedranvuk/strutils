// Copyright 2013-2024 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings adds additional string utility functions.
package strutils

import (
	"sort"
	"strings"
)

// Compare returns 1 if a > b, -1 if a < b or 0 if a == b.
func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// CompareFold returns 1 if a > b, -1 if a < b or 0 if a == b.
// Comparison is not case-sensitive.
func CompareFold(a, b string) int {
	return Compare(strings.ToLower(a), strings.ToLower(b))
}

// Return s prefix up to sep starting from the left.
// If sep not found returns an empty string.
func FetchLeft(s, sep string) string {
	var i = strings.Index(s, sep)
	if i < 0 {
		return ""
	}
	return s[0:i]
}

// Return s prefix up to sep starting from the left.
// Separator search is case-insesitive, result casing is not modified.
// If sep not found returns an empty string.
func FetchLeftFold(s, sep string) string {
	return s[0:strings.Index(strings.ToLower(s), strings.ToLower(sep))]
}

// Return s suffix up to sep starting from the right.
// If sep not found returns an empty string.
func FetchRight(s, sep string) string {
	var i = -1
	for j := strings.Index(s, sep); j > -1; j = strings.Index(s[j+1:], sep) {
		i = j
	}
	return s[i+len(sep):]
}

// Return s suffix up to sep starting from the right.
// Separator search is case-insesitive, result casing is not modified.
// If sep not found returns an empty string.
func FetchRightFold(s, sep string) string {
	return FetchRight(strings.ToLower(s), strings.ToLower(sep))
}

// HasPrefixFold tests whether the string s begins with prefix.
// Case-insensitive.
func HasPrefixFold(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffixFold tests whether the string s ends with suffix
// Case-insensitive.
func HasSuffixFold(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// IndexFold returns the index of the first instance of substr in s, or -1 if
// substr is not present in s. Search is case-insensitive.
func IndexFold(s, substr string) int {
	return strings.Index(strings.ToLower(s), strings.ToLower(substr))
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashstr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashstr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])

	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s".
func Indexes(s, sep string) (r []int) {
	n := len(sep)
	switch {
	case n == 0:
		return
	case n == 1:
		c := sep[0]
		// special case worth making fast
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				r = append(r, i)
			}
		}
		return
	case n == len(s):
		if sep == s {
			r = append(r, 0)
			return
		}
	case n > len(s):
		return
	}
	// Hash sep.
	hashsep, pow := hashstr(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && s[:n] == sep {
		r = append(r, 0)
		return
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && s[i-n:i] == sep {
			r = append(r, i-n)
		}
	}
	return r
}

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s". Case-insensitive.
func IndexesFold(s, sep string) []int {
	return Indexes(strings.ToLower(s), strings.ToLower(sep))
}

// Unwrap unpacks the string by removing prefix and suffix.
// If both prefix and suffix were found result is an unpacked string and true
// else result is s and false.
// Both prefix and suffix are optional and can be empty in which case their
// removal is not performed.
func Unwrap(s, prefix, suffix string) (string, bool) {
	if !strings.HasPrefix(s, prefix) {
		return s, false
	}
	if !strings.HasSuffix(s, suffix) {
		return s, false
	}
	return s[len(prefix) : len(s)-len(suffix)], true
}

// UnwrapFold is the case-insensitive version of Unpack.
func UnwrapFold(s, prefix, suffix string) (string, bool) {
	if !HasPrefixFold(s, prefix) {
		return s, false
	}
	if !HasSuffixFold(s, suffix) {
		return s, false
	}
	return s[len(prefix) : len(s)-len(suffix)], true
}

// UnquoteSingle removes single quotes around s and returns it and true on
// success. If either leading or trailing quote is not found result is s, false.
func UnquoteSingle(s string) (string, bool) { return Unwrap(s, "'", "'") }

// UnquoteDouble removes double quotes around s and returns it and true on
// success. If either leading or trailing quote is not found result is s, false.
func UnquoteDouble(s string) (string, bool) { return Unwrap(s, "\"", "\"") }

// Wrap wraps s within prefix and suffix.
func Wrap(s, prefix, suffix string) string { return prefix + s + suffix }

// QuoteSingle wraps s with single quotes.
func QuoteSingle(s string) string { return Wrap(s, "'", "'") }

// QuoteDouble wraps s with double quotes.
func QuoteDouble(s string) string { return Wrap(s, "\"", "\"") }

// Matches "text" against "pattern". Case insensitive. Returns truth.
// * matches any number of characters.
// ? matches one character.
func MatchesWildcard(text, pattern string) bool {
	if text == "" || pattern == "" {
		return false
	}

	var (
		t, w   = []rune(text), []rune(pattern)
		it, iw = 0, 0
	)
	for it < len(t) && iw < len(w) {
		if w[iw] == '*' {
			break
		}
		if w[iw] != '?' && !strings.EqualFold(string(t[it]), string(w[iw])) {
			return false
		}
		it++
		iw++
	}

	var sw, st = 0, -1
	for it < len(t) && iw < len(w) {
		if w[iw] == '*' {
			iw++
			if iw >= len(w) {
				return true
			}
			sw = iw
			st = it
		} else {
			if w[iw] == '?' || strings.EqualFold(string(t[it]), string(w[iw])) {
				it++
				iw++
			} else {
				it = st
				st++
				iw = sw
			}
		}
	}

	for iw < len(w) && w[iw] == '*' {
		iw++
	}

	return iw == len(w)
}

// Segment returns the word up to the next separator and index of s just after
// the separator. Scan of s starts at start index.
//
// If sep was not found returns s starting at start and -1.
// If the end of s was reached resulting next will be -1.
// Returns an empty string and -1 if start is out of range or sep is empty.
func Segment(s, sep string, start int) (segment string, next int) {
	if len(s) == 0 || start < 0 || start > len(s)-1 {
		return "", -1
	}
	var end = strings.Index(s[start:], sep)
	if end == -1 {
		return s[start:], -1
	}
	if end == 0 {
		start = start + len(sep)
		end = start + end
		if start+len(sep) > len(s) {
			return s[start : start+len(sep)], -1
		}
		return s[start : start+len(sep)], start + len(sep)
	}
	end = start + end
	return s[start:end], end + len(sep)
}

// SegmentFold is the case-insensitive version of Segment.
func SegmentFold(s, sep string, start int) (segment string, next int) {
	if len(s) == 0 || start < 0 || start > len(s)-1 {
		return "", -1
	}
	var end = IndexFold(s[start:], sep)
	if end == -1 {
		return s[start:], -1
	}
	if end == 0 {
		start = start + len(sep)
		end = start + end
		if start+len(sep) > len(s) {
			return s[start : start+len(sep)], -1
		}
		return s[start : start+len(sep)], start + len(sep)
	}
	end = start + end
	return s[start:end], end + len(sep)
}

// WrapText wraps text into multiple lines at first whitespace before or exactly
// at cols.
//
// If a word is longer than cols and force is true it is split at cols length
// regardless of white space. If force is false, the word is not split and
// placed into its own line at first next whitespace.
//
// Allocates runes of text and appends to out without preallocation.
func WrapText(text string, cols int, force bool) (out []string) {

	var (
		idx   = 0            // scan index
		start = 0            // copy offset
		space = -1           // last space position
		col   = 0            // column counter
		runes = []rune(text) // alloc
	)

	for l := len(runes); idx < l; {

		// Wrap at newlines.
		if runes[idx] == '\n' {
			out = append(out, text[start:idx])
			col = 0
			start = idx + 1
			space = -1
			idx++
			continue
		}

		var isSpace = runes[idx] == ' '
		if isSpace && col == 0 {
			idx++
			start = idx
			continue
		}

		if col == cols-1 {

			// Last col is space.
			if isSpace {
				out = append(out, text[start:idx])
				col = 0
				start = idx + 1
				space = -1
				idx++
				continue
			}

			// Wrap at last space.
			if space > -1 {
				out = append(out, text[start:space])
				col = cols - (space - start) - 1
				start = space + 1
				space = -1
				idx++
				continue
			}

			// Wrap if forced or next rune is a space.
			if force || idx < l-1 && runes[idx+1] == ' ' {
				out = append(out, text[start:idx+1])
				start = idx + 1
				idx++
				col = 0
				space = -1
				continue
			}

		}

		if isSpace {

			if col > cols-1 {
				// Word is longer than cols.
				out = append(out, text[start:idx])
				col = 0
				start = idx + 1
				space = -1
				idx++
				continue
			}

			space = idx
		}

		col++
		idx++
	}

	if start < len(runes) {
		out = append(out, text[start:])
	}

	return
}

// Unique returns unique strings from in.
//
// Returned strings are in order as passed in.
func Unique(in ...string) (out []string) {

	var (
		m = make(map[string]int)
		r = make(map[int]string)
	)

	for i, s := range in {
		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}

	var n = make([]int, 0, len(m))
	for k, v := range m {
		r[v] = k
		n = append(n, v)
	}
	sort.Ints(n)

	out = make([]string, 0, len(m))

	for _, i := range n {
		out = append(out, r[i])
	}

	return
}
