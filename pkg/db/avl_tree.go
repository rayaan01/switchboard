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
