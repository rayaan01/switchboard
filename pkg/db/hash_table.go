package db

import (
	"fmt"
	"sort"
	"strings"
)

type HashTable struct {
	store map[string]string
}

func (ht *HashTable) get(key string) ([]byte, error) {
	value, ok := ht.store[key]
	if !ok {
		return []byte("(nil)"), nil
	}
	return []byte(value), nil
}

func (ht *HashTable) set(key string, value string) ([]byte, error) {
	ht.store[key] = value
	return []byte("OK"), nil
}

func (ht *HashTable) del(key string) ([]byte, error) {
	delete(ht.store, key)
	return []byte("OK"), nil
}

func (ht *HashTable) get_range(low string, high string) ([]byte, int, error) {
	keys := []string{}
	var results *[]string = &[]string{}
	for k := range ht.store {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		if k >= low && k <= high {
			result := fmt.Sprintf("[%s, %s]", k, ht.store[k])
			*results = append(*results, result)
		}
	}

	keyValuePairs := "[" + strings.Join(*results, ",") + "]"
	return []byte(keyValuePairs), len(*results), nil
}
