package strutils

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestRandom(t *testing.T) {
	set := []byte("abc")
	result := Random(set)
	if len(result) != 1 {
		t.Errorf("Expected length 1, got %d", len(result))
	}
	if !strings.Contains(string(set), result) {
		t.Errorf("Result not in set")
	}
}

func BenchmarkRandom(b *testing.B) {
	set := []byte("abc")
	for i := 0; i < b.N; i++ {
		Random(set)
	}
}

func TestRandoms(t *testing.T) {
	set := []byte("abc")
	length := 5
	result := Randoms(set, length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	for _, r := range result {
		if !strings.Contains(string(set), string(r)) {
			t.Errorf("Result not in set")
		}
	}

	length = 0
	result = Randoms(set, length)
	if len(result) != 0 {
		t.Errorf("Expected empty string, got %s", result)
	}

	length = -1
	result = Randoms(set, length)
	if len(result) != 0 {
		t.Errorf("Expected empty string, got %s", result)
	}

}

func BenchmarkRandoms(b *testing.B) {
	set := []byte("abc")
	length := 5
	for i := 0; i < b.N; i++ {
		Randoms(set, length)
	}
}

func TestRandomNum(t *testing.T) {
	result := RandomNum()
	if len(result) != 1 {
		t.Errorf("Expected length 1, got %d", len(result))
	}
	if !strings.Contains(string(Numerals), result) {
		t.Errorf("Result not in numerals")
	}
}

func BenchmarkRandomNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNum()
	}
}

func TestRandomNums(t *testing.T) {
	length := 5
	result := RandomNums(length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	for _, r := range result {
		if !strings.Contains(string(Numerals), string(r)) {
			t.Errorf("Result not in numerals")
		}
	}
}

func BenchmarkRandomNums(b *testing.B) {
	length := 5
	for i := 0; i < b.N; i++ {
		RandomNums(length)
	}
}

func TestRandomUpper(t *testing.T) {
	result := RandomUpper()
	if len(result) != 1 {
		t.Errorf("Expected length 1, got %d", len(result))
	}
	if !strings.Contains(string(AlphaUpper), result) {
		t.Errorf("Result not in alphaUpper")
	}
}

func BenchmarkRandomUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomUpper()
	}
}

func TestRandomUppers(t *testing.T) {
	length := 5
	result := RandomUppers(length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	for _, r := range result {
		if !strings.Contains(string(AlphaUpper), string(r)) {
			t.Errorf("Result not in alphaUpper")
		}
	}
}

func BenchmarkRandomUppers(b *testing.B) {
	length := 5
	for i := 0; i < b.N; i++ {
		RandomUppers(length)
	}
}

func TestRandomLower(t *testing.T) {
	result := RandomLower()
	if len(result) != 1 {
		t.Errorf("Expected length 1, got %d", len(result))
	}
	if !strings.Contains(string(AlphaLower), result) {
		t.Errorf("Result not in alphaLower")
	}
}

func BenchmarkRandomLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomLower()
	}
}

func TestRandomLowers(t *testing.T) {
	length := 5
	result := RandomLowers(length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	for _, r := range result {
		if !strings.Contains(string(AlphaLower), string(r)) {
			t.Errorf("Result not in alphaLower")
		}
	}
}

func BenchmarkRandomLowers(b *testing.B) {
	length := 5
	for i := 0; i < b.N; i++ {
		RandomLowers(length)
	}
}

func TestRandomSpecial(t *testing.T) {
	result := RandomSpecial()
	if len(result) != 1 {
		t.Errorf("Expected length 1, got %d", len(result))
	}
	if !strings.Contains(string(SpecialChars), result) {
		t.Errorf("Result not in specialChars")
	}
}

func BenchmarkRandomSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomSpecial()
	}
}

func TestRandomSpecials(t *testing.T) {
	length := 5
	result := RandomSpecials(length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	for _, r := range result {
		if !strings.Contains(string(SpecialChars), string(r)) {
			t.Errorf("Result not in specialChars")
		}
	}
}

func BenchmarkRandomSpecials(b *testing.B) {
	length := 5
	for i := 0; i < b.N; i++ {
		RandomSpecials(length)
	}
}

func TestRandomString(t *testing.T) {
	length := 5
	lo := true
	up := true
	nums := true
	result := RandomString(lo, up, nums, length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}

	length = 5
	lo = false
	up = false
	nums = false
	result = RandomString(lo, up, nums, length)
	if len(result) != 0 {
		t.Errorf("Expected empty string")
	}

	length = 0
	lo = true
	up = true
	nums = true
	result = RandomString(lo, up, nums, length)
	if len(result) != 0 {
		t.Errorf("Expected empty string")
	}

	length = -1
	lo = true
	up = true
	nums = true
	result = RandomString(lo, up, nums, length)
	if len(result) != 0 {
		t.Errorf("Expected empty string")
	}

}

func BenchmarkRandomString(b *testing.B) {
	length := 10
	lo := true
	up := true
	nums := true
	for i := 0; i < b.N; i++ {
		RandomString(lo, up, nums, length)
	}
}

func TestFoo_word(t *testing.T) {
	foo := NewFoo()
	length := 5
	result := foo.word(length)
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
}

func BenchmarkFoo_word(b *testing.B) {
	foo := NewFoo()
	length := 5
	for i := 0; i < b.N; i++ {
		foo.word(length)
	}
}

func TestFoo_intRng(t *testing.T) {
	foo := NewFoo()
	lo := 1
	hi := 10
	result := foo.intRng(lo, hi)
	if result < lo || result >= hi {
		t.Errorf("Result out of range")
	}
}

func BenchmarkFoo_intRng(b *testing.B) {
	foo := NewFoo()
	lo := 1
	hi := 10
	for i := 0; i < b.N; i++ {
		foo.intRng(lo, hi)
	}
}

func TestFoo_Bool(t *testing.T) {
	foo := NewFoo()
	result := foo.Bool()
	_ = result // Just call function, cannot test randomness
}

func BenchmarkFoo_Bool(b *testing.B) {
	foo := NewFoo()
	for i := 0; i < b.N; i++ {
		foo.Bool()
	}
}

func TestFoo_Name(t *testing.T) {
	foo := NewFoo()
	result := foo.Name()
	if len(result) < foo.MinName || len(result) > foo.MaxName {
		t.Errorf("Result length out of range")
	}
}

func BenchmarkFoo_Name(b *testing.B) {
	foo := NewFoo()
	for i := 0; i < b.N; i++ {
		foo.Name()
	}
}

func TestFoo_EMail(t *testing.T) {
	foo := NewFoo()
	result := foo.EMail()
	if !strings.Contains(result, "@") {
		t.Errorf("Result does not contain @")
	}
	domainFound := false
	for _, domain := range foo.Domains {
		if strings.Contains(result, domain) {
			domainFound = true
			break
		}
	}
	if !domainFound {
		t.Errorf("Result does not contain configured domain")
	}
}

func BenchmarkFoo_EMail(b *testing.B) {
	foo := NewFoo()
	for i := 0; i < b.N; i++ {
		foo.EMail()
	}
}

func TestNewFoo(t *testing.T) {
	foo := NewFoo()
	if foo.MinName != 2 {
		t.Errorf("Expected MinName 2, got %d", foo.MinName)
	}
	if foo.MaxName != 10 {
		t.Errorf("Expected MaxName 10, got %d", foo.MaxName)
	}
	if len(foo.Domains) != 3 {
		t.Errorf("Expected Domains length 3, got %d", len(foo.Domains))
	}
}

func BenchmarkNewFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFoo()
	}
}
