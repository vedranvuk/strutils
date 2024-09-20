package strutils

import "testing"

func TestTag(t *testing.T) {

	const tag = `json:"omitempty" tag:"key1,key2=value1,key2=value2,key3" db:"name=foo"`

	var config = &Tag{
		TagKey: "tag",
		KnownPairKeys:    []string{"key1", "key2", "key3"},
	}

	if err := config.Parse(tag); err != nil {
		t.Fatal(err)
	}

	if !config.Values.Exists("key1") {
		t.Fatal("key1 not found")
	}
	if !config.Values.Exists("key2") {
		t.Fatal("key2 not found")
	}
	if !config.Values.Exists("key3") {
		t.Fatal("key3 not found")
	}
	if len(config.Values["key1"]) != 0 {
		t.Fatal("Invalid length for key1")
	}
	if len(config.Values["key2"]) != 2 {
		t.Fatal("Invalid length for key2")
	}
	if len(config.Values["key1"]) != 0 {
		t.Fatal("Invalid length for key3")
	}
	if config.Values.First("key2") != "value1" {
		t.Fatal("First failed")
	}
}