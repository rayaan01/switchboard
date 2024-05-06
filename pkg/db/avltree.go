package db

type AVLTreeNode struct {
	key    string
	value  string
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int16
}

type AVLTree struct {
	store AVLTreeNode
}

func (at *AVLTree) get(key string) ([]byte, error) {
	return []byte("OK"), nil
}

func (at *AVLTree) set(key string, value string) ([]byte, error) {
	return []byte("OK"), nil
}

func (at *AVLTree) del(key string) ([]byte, error) {
	return []byte("OK"), nil
}
