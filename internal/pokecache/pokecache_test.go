package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestReap(t *testing.T) {
	interval := 5 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("short-lived", []byte("value"))

	_, exists := cache.Get("short-lived")
	if !exists {
		t.Fatal("expected key to exist immediately after adding")
	}

	time.Sleep(interval * 2)

	_, exists = cache.Get("short-lived")
	if exists {
		t.Fatal("expected key to be reaped after interval passed")
	}
}

func TestCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("greeting", []byte("hello"))
	got, exists := cache.Get("greeting")
	if !exists {
		t.Fatal("expected the key to exist")
	}

	if string(got) != "hello" {
		t.Fatal("wrong")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://pokeapi.co",
			value: []byte("bulbasaur data"),
		},
		{
			key:   "https://pokeapi.co",
			value: []byte("charmander data"),
		},
	}

	for _, tc := range cases {
		cache.Add(tc.key, tc.value)

		actual, exists := cache.Get(tc.key)
		if !exists {
			t.Errorf("expected to find key %q", tc.key)
			continue
		}

		if !bytes.Equal(actual, tc.value) {
			t.Errorf("expected value %q, got %q", string(tc.value), string(actual))
		}
	}
}

func TestMissingKey(t *testing.T) {
	cache := NewCache(5 * time.Second)

	actual, exists := cache.Get("https://pokeapi.co")
	if exists {
		t.Error("expected key to not exist")
	}

	if actual != nil {
		t.Errorf("expected returned value to be nil, got %v", actual)
	}
}

func TestReplaceValue(t *testing.T) {
	cache := NewCache(5 * time.Second)
	key := "https://pokeapi.co"

	cache.Add(key, []byte("old data"))
	cache.Add(key, []byte("new data"))

	actual, exists := cache.Get(key)
	if !exists {
		t.Errorf("expected to find key %q", key)
	}

	if !bytes.Equal(actual, []byte("new data")) {
		t.Errorf("expected value to be replaced with %q, got %q", "new data", string(actual))
	}
}