package strutils

import (
	"errors"
	"strconv"
	"strings"
)

// TagKey names a key inside a tag string literal whose value is to be parsed
// into [Values].
//
// Given:
//
//	tagKey := "foo"
//	rawTag := `json:"omitempty" foo:"key1,key2=value1,key2=value2,key3"`
//
// TagKey specifies the "foo" key inside a tag string literal.
type TagKey = string

// PairKey is a recognized key in a set of key=value pairs parsed from a tag
// value.
//
// Given:
//
//	tagKey  := "foo"
//	pairKey := "key1"
//	rawTag  := `json:"omitempty" foo:"key1,key2=value1,key2=value2,key3"`
//
// pairKey specifies the "key1" inside a value keyed by tagKey.
type PairKey = string

// Tag parses value of a key inside a tag string literal into [Values] map.
//
// Given:
//
//	tag := `json:"omitempty" MyTag:"key1,key2=value1,key2=value2,key3"`
//
// Results in following Values structure:
//
//		Values{
//			"key1": nil,
//			"key2": []string{
//				"value1",
//				"value2",
//	     }
//			"key3": nil,
//		}
//
// The following rules, equal to how Go parses tags apply:
//
// Keys may appear without values or in key=value format. Multiple keys or pairs
// are separated by a comma. Values may not contain commas or double quotes.
//
// Leading and trailing space is trimmed from pair values.
// Specifying a pair with the same key multiple times adds values to an entry
// under key in parsed [Values].
//
// Specifying a PairKey without a value adds an entry without values in the
// [Values] map. Specifying a PairKey with an empty value multiple times is a
// no-op.
//
// See [Tag.Parse] for details.
type Tag struct {
	// TagKey is the name of the tag whose value is to be parsed into [Values].
	TagKey

	// Separator is the pair separator.
	//
	// Defaults to comma ",".
	Separator string

	// KnownPairKeys is a set of recognised pair keys found inside a value of
	// a tag.
	//
	// If it is an empty slice all keys or key=value pairs will be parsed.
	// If it is not an empty slice, unrecognised keys will be skipped silently
	// or an error will be thrown if [Tag.ErrorOnUnknownKey] is true.
	KnownPairKeys []PairKey

	// ErrorOnUnknownKey, if true will make parse functions throw an error if
	// an unrecognised tag is found and [Tag.Keys] is not an empty slice.
	//
	// Default: false
	ErrorOnUnknownKey bool

	// Raw is the raw tag value that was parsed.
	// Set after [Tag.Parse].
	Raw string

	// Values are the parsed values. Values are nil until [Tag.ParseDocs] or
	// [Tag.ParseStructTag] is called.
	Values
}

// ErrTagNotFound is returned when tag named [Tag.TagKey] was not found in a
// tag string literal.
var ErrTagNotFound = errors.New("tag not found")

// Parse parses a tag string literal into [Values].
//
// tag may be a backquoted string (for instance extracted from struct field
// type using reflect) in which case it is unquoted before parsing.
//
// See [Tag] on details how the tag string is parsed.
func (self *Tag) Parse(tag string) (err error) {

	if self.TagKey == "" {
		return errors.New("tag name not specified")
	}
	if self.Separator == "" {
		self.Separator = ","
	}

	if self.Values == nil {
		self.Values = make(Values)
	}

	tag, _ = Unwrap(tag, "`", "`")

	var exists bool
	if tag, exists = LookupTag(tag, self.TagKey); !exists {
		return ErrTagNotFound
	}
	_, self.Raw, _ = strings.Cut(tag, "=")

	for key, i := Segment(tag, self.Separator, 0); i > -1 || key != ""; key, i = Segment(tag, self.Separator, i) {
		var k, v, pair = strings.Cut(key, "=")
		if !self.validKey(k) {
			return errors.New("invalid key: " + k)
		}
		if pair {
			self.Values.Add(k, v)
		} else {
			self.Values.Add(k)
		}
	}
	return nil
}

// validKey returns true if key is in [Config.Keys] or it is empty,
// false otherwise.
func (self *Tag) validKey(key string) (valid bool) {
	if len(self.KnownPairKeys) == 0 {
		return true
	}
	for _, k := range self.KnownPairKeys {
		if k == key {
			return true
		}
	}
	return false
}

// Values is a map of parsed key=value pairs from a tag value.
type Values map[PairKey][]string

// Add appends values to value slice under key.
//
// If no values exist under key a new entry is added. If no values were
// specified target slice is unmodified. Initial value of an empty entry is
// a nil slice.
func (self Values) Add(key string, values ...string) {
	if s, exists := self[key]; exists {
		s = append(s, values...)
		self[key] = s
		return
	}
	self[key] = values
	return
}

// Exists returns true if entry under key exists.
func (self Values) Exists(key string) (exists bool) {
	_, exists = self[key]
	return
}

// Exists returns true if entry under key exists and has at least one value
// which is not empty.
func (self Values) ExistsNonEmpty(key string) (exists bool) {
	var val []string
	if val, exists = self[key]; !exists {
		return
	}
	return len(val) > 0 && val[0] != ""
}

// First returns the first value keyed under key.
// If entry not found or no values for entry found returns an empty string.
// Use [Values.Exists] to check if an entry exists.
func (self Values) First(key string) (s string) {
	if !self.Exists(key) {
		return
	}
	if a := self[key]; len(a) > 0 {
		return a[0]
	}
	return
}

// Set sets value of out to value under specified key and returns true.
// If key was not found or has no value out is not set and false is returned.
func (self Values) Set(key string, out *string) (exists bool) {
	var val []string
	if val, exists = self[key]; !exists {
		return
	}
	if len(val) == 0 || val[0] == "" {
		return false
	}
	*out = val[0]
	return true
}

// Clear clears any loaded values.
func (self Values) Clear() { clear(self) }

// LookupTag returns the value associated with key in the tag string.
// If the key is present in the tag the value (which may be empty)
// is returned. Otherwise the returned value will be the empty string.
// The ok return value reports whether the value was explicitly set in
// the tag string. If the tag does not have the conventional format,
// the value returned by LookupTag is unspecified.
//
// LookupTag is a copy of (reflect.Lookup) to skip the reflect include.
func LookupTag(rawStructTag, key string) (value string, ok bool) {
	// When modifying this code, also update the validateStructTag code
	// in cmd/vet/structtag.go.

	for rawStructTag != "" {
		// Skip leading space.
		i := 0
		for i < len(rawStructTag) && rawStructTag[i] == ' ' {
			i++
		}
		rawStructTag = rawStructTag[i:]
		if rawStructTag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(rawStructTag) && rawStructTag[i] > ' ' && rawStructTag[i] != ':' && rawStructTag[i] != '"' && rawStructTag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(rawStructTag) || rawStructTag[i] != ':' || rawStructTag[i+1] != '"' {
			break
		}
		name := string(rawStructTag[:i])
		rawStructTag = rawStructTag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(rawStructTag) && rawStructTag[i] != '"' {
			if rawStructTag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(rawStructTag) {
			break
		}
		qvalue := string(rawStructTag[:i+1])
		rawStructTag = rawStructTag[i+1:]

		if key == name {
			value, err := strconv.Unquote(qvalue)
			if err != nil {
				break
			}
			return value, true
		}
	}
	return "", false
}
