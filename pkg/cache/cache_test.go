package cache

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCacheNewLRUCache(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		val  int
		want *LRU
	}{
		"simple": {val: 3, want: &LRU{
			capacity: 3,
			items:    make(map[string]*list.Element, 3),
			queue:    list.New(),
		}},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := NewLRUCache(testCase.val)
			req.Equal(testCase.want, res)
		})
	}

}

func TestCacheAdd(t *testing.T) {
	tests := map[string]struct {
		firstVal  string
		secondVal string
		want      bool
	}{
		"simple": {firstVal: "TestKey1", secondVal: "TestValue1", want: true},
	}
	req := require.New(t)
	cache := NewLRUCache(5)
	req.NotEmpty(cache)
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := cache.Add(testCase.firstVal, testCase.secondVal)
			req.Equal(testCase.want, res)
			resCache := cache.Add(testCase.firstVal, testCase.secondVal)
			req.NotEqual(testCase.want, resCache)
		})
	}

}

func TestCacheRemove(t *testing.T) {
	req := require.New(t)
	cache := NewLRUCache(5)
	tests := map[string]struct {
		firstVal  string
		secondVal string
		want      bool
	}{
		"simple": {firstVal: "TestKey1", secondVal: "TestValue1", want: true},
	}
	for name, testCase := range tests {
		cache.Add(testCase.firstVal, testCase.secondVal)
		t.Run(name, func(t *testing.T) {
			res := cache.Remove(testCase.firstVal)
			req.Equal(res, testCase.want)
		})
	}
}

func TestCacheGet(t *testing.T) {
	req := require.New(t)
	cache := NewLRUCache(5)
	tests := map[string]struct {
		firstVal  string
		secondVal string
		want      bool
	}{
		"simple": {firstVal: "TestKey1", secondVal: "TestValue1", want: true},
	}
	for name, testCase := range tests {
		cache.Add(testCase.firstVal, testCase.secondVal)
		t.Run(name, func(t *testing.T) {
			res, ok := cache.Get(testCase.firstVal)
			req.NotEqual("", res)
			req.Equal(ok, testCase.want)
		})
	}
}
