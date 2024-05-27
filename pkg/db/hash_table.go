package db

import "fmt"

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

func (ht *HashTable) getStore() *AVLTreeNode {
	return nil
}

func (ht *HashTable) visualizeAVLTree(node *AVLTreeNode) {
	fmt.Println("Could not visualize AVL Tree")
}

func (ht *HashTable) visualizeHashTable() {
	for key, value := range ht.store {
		fmt.Printf("%s:%s\n", key, value)
	}
}
