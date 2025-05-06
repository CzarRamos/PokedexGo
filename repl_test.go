package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/CzarRamos/pokedexcli/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("The strings expected:%d and actual:%d are not the same length!", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words expected:%s and actual:%s does not match", expectedWord, word)
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte("testdata"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			val: []byte("moretestdata"),
		},
	}

	client := pokeapi.NewClient()

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := client.Config.CachedInfo
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	client := pokeapi.NewClient()
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := client.Config.CachedInfo
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
