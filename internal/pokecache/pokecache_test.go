package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
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
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	testCase := struct {
		key string
		val []byte
	}{
		key: "https://example.com",
		val: []byte("testdata"),
	}

	cache := NewCache(baseTime)
	cache.Add(testCase.key, testCase.val)

	if _, ok := cache.Get(testCase.key); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get(testCase.key); ok {
		t.Errorf("expected key to be reaped")
		return
	}
}
