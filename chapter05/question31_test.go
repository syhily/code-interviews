package chapter05

import (
	"reflect"
	"testing"
)

func Test_lruCache_get(t *testing.T) {
	cache1 := newLRUCache[string, int](10, -1)
	cache2 := newLRUCache[string, int](10, -1)
	cache2.put("key", 1000)

	type args[K comparable] struct {
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		l    LRUCache[K, V]
		args args[K]
		want V
	}
	tests := []testCase[string, int]{
		{
			name: "Get a default value with a non-existed key",
			l:    cache1,
			args: args[string]{k: "key1"},
			want: -1,
		},
		{
			name: "Get a default value with an existing key",
			l:    cache2,
			args: args[string]{k: "key"},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lruCache_put(t *testing.T) {
	cache1 := newLRUCache[string, int](10, -1)
	t.Run("the put key can be queried by get", func(t *testing.T) {
		cache1.put("key", 10)
		val := cache1.get("key")
		if val != 10 {
			t.Errorf("Invalid value from put(), expect %v", 10)
		}
	})

	cache2 := newLRUCache[string, int](2, -1)
	cache2.put("key1", 1)
	cache2.put("key2", 2)
	t.Run(
		"if the cache size exceed the capacity, we will remove oldest first",
		func(t *testing.T) {
			val1 := cache2.get("key1")
			if val1 != 1 {
				t.Errorf("Expected the key1 exists in cache")
			}

			cache2.put("key3", 3)
			val1 = cache2.get("key1")
			val3 := cache2.get("key3")
			if val1 != -1 {
				t.Errorf("Expected the key1 should be removed in cache")
			}
			if val3 != 3 {
				t.Errorf("Expected the key3 exists in cache")
			}
		})
}
