package strutils

import (
	"errors"
	"strconv"
	"strings"
)

// TagKey is a recognized key in a set of key=value pairs parsed
// from a raw tag value.
//
// In struct tags:  `TagName:"TagKey=somevalue"`
// In doc comments: //TagName:"TagKey=somevalue"
type TagKey = string

// Values is a parsed map of key=value pairs from a tag value.
// An entry under some [TagKey] can have multiple values, stored as a slice.
//
// Given:
//
// 	const TagName = "mytag"
//
// both the
//
//  structTag = `json:"omitempty" MyTag:"key1,key2=value1,key2=value2,key3"`
//
// and
//
//  goDocLine = //myTag:"key1,key2=value1,key2=value2,key3"
//
// Results in the following:
//
// 	Values{
//		"key1": nil,
//		"key2": []string{
//			"value1",
//			"value2",
//      }
// 		"key3": nil,
// 	}
//
// 
// Contents of the quoted string following "MyTag:" are parsed as input which
// has following rules:
// 
// Keys may appear without values or in key=value format. Multiple keys or pairs
// are separated by a comma. Values may not contain commas or  double quotes. 
// Space is trimmed from any pair values. Specifying a pair with the same key
// multiple times adds values to an entry under key in parsed [Values].
//
// See [Tag.ParseStructTag] and [Tag.ParseDocs].
type Values map[TagKey][]string

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

// Tag is a tag parser.
//
// It parses comma delimited set of keys or key=value pairs.
// Internal representation is a keyvalue map.
type Tag struct {
	// TagName is the name of the tag to parse.
	TagName string

	// Keys is a set of recognised Values keys to parse.
	//
	// If keys is an empty slice all key= value pairs will be parsed.
	// If keys is not an empty slice, unrecognised keys will be ignored or
	// an error will be thrown if one is encountered if
	// [Tag.ErrorOnUnknownKey] is true.
	Keys []TagKey

	// ErrorOnUnknownKey, if true will make parse functions throw an error if
	// an unrecognised tag is found and [Tag.Keys] is not an empty slice.
	//
	// Default: false
	ErrorOnUnknownKey bool

	// Values are the parsed values.
	// Nil until parsed.
	Values
}

func (self *Tag) init() error {
	if self.TagName == "" {
		return errors.New("tag name not specified")
	}
	self.Values = make(Values)
	return nil
}

// ErrTagNotFound is returned when the defined tag was not found.
var ErrTagNotFound = errors.New("tag not found")

// validKey returns true if key is in [Config.Keys] or it is empty,
// false otherwise.
func (self *Tag) validKey(key string) (valid bool) {
	if len(self.Keys) == 0 {
		return true
	}
	for _, k := range self.Keys {
		if k == key {
			return true
		}
	}
	return false
}

// ParseStructTag parses a raw struct tag into [Values].
//
// rawTag must be a raw struct tag string, possibly quoted with (``) and 
// containing other tags such as "json", "db", etc. 
//
// It looks for a value under a key specified by [Tag.TagName] and parses its 
// value into [Values]. See [Values] on how the tag value is parsed.
func (self *Tag) ParseStructTag(rawTag string) (err error) {

	if err = self.init(); err != nil {
		return
	}

	rawTag, _ = Unwrap(rawTag, "`", "`")
	var exists bool
	if rawTag, exists = LookupTag(rawTag, self.TagName); !exists {
		return ErrTagNotFound
	}

	for key, i := Segment(rawTag, ",", 0); i > -1 || key != ""; key, i = Segment(rawTag, ",", i) {
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

// ParseDocs parses go doc comments into [Values].
//
// docs must be a slice of raw lines from a declaration doc comment from which 
// lines in the following form are parsed:
//  //mytag:"key1,key2=value2,key3"
//
// I.e., the standard build tag way of a tag immediately following a double 
// slash, a colon denoting the value that follows and a value inside double 
// quotes that is parsed into [Tag.Values].
//
// Tag to parse is defined by [Tag.TagName]. See [Values] for details on how 
// the tag value is parsed.
func (self *Tag) ParseDocs(docs []string) (err error) {

	if err = self.init(); err != nil {
		return
	}

	var tagPrefix = self.TagName + ":"
	for _, line := range docs {
		line = strings.TrimSpace(strings.TrimPrefix(line, "//"))
		if !strings.HasPrefix(line, tagPrefix) {
			continue
		}
		line = strings.TrimPrefix(line, tagPrefix)
		line, _ = UnquoteDouble(line)
		for _, token := range strings.Split(line, ",") {
			if token = strings.TrimSpace(token); token == "" {
				continue
			}
			var key, val, pair = strings.Cut(token, "=")
			if !pair {
				self.Add(key, "")
				continue
			}
			key = strings.TrimSpace(key)
			val = strings.TrimSpace(val)
			self.Values.Add(key, strings.Split(val, ",")...)
		}
	}
	return nil
}

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
