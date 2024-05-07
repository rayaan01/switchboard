package db

type StorageEngine interface {
	set(key string, value string) ([]byte, error)
	get(key string) ([]byte, error)
	del(key string) ([]byte, error)
	visualize(node *AVLTreeNode)
	getStore() *AVLTreeNode
}

var StoreMapper = map[string]StorageEngine{}
