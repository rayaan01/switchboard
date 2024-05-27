package db

type StorageEngine interface {
	set(key string, value string) ([]byte, error)
	get(key string) ([]byte, error)
	del(key string) ([]byte, error)
	visualizeAVLTree(node *AVLTreeNode)
	visualizeHashTable()
	getStore() *AVLTreeNode
}

var StoreMapper = map[string]StorageEngine{}
