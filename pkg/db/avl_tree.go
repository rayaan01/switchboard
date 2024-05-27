package db

import "fmt"

type AVLTreeNode struct {
	key    string
	value  string
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int16
}

type AVLTree struct {
	store *AVLTreeNode
}

func (at *AVLTree) get(key string) ([]byte, error) {
	value := retrieve(at.store, key)
	return []byte(value), nil
}

func (at *AVLTree) set(key string, value string) ([]byte, error) {
	at.store = insert(at.store, key, value)
	return []byte("OK"), nil
}

func (at *AVLTree) del(key string) ([]byte, error) {
	at.store = remove(at.store, key)
	return []byte("OK"), nil
}

func (at *AVLTree) getStore() *AVLTreeNode {
	return at.store
}

func (at *AVLTree) visualizeAVLTree(node *AVLTreeNode) {
	if node == nil {
		return
	}
	at.visualizeAVLTree(node.left)
	fmt.Println("Node", node)
	fmt.Println("Left", node.left)
	fmt.Println("Right", node.right)
	at.visualizeAVLTree(node.right)
}

func (at *AVLTree) visualizeHashTable() {
	fmt.Println("Could not visualize Hash Table")
}

func insert(node *AVLTreeNode, key string, value string) *AVLTreeNode {
	if node == nil {
		return createNode(key, value)
	}

	if key < node.key {
		node.left = insert(node.left, key, value)
	} else {
		node.right = insert(node.right, key, value)
	}

	node.height = updateHeight(node)
	bf := getBalanceFactor(node)

	if bf < -1 && key > node.right.key {
		return rotateLeft(node)
	}

	if bf > 1 && key < node.left.key {
		return rotateRight(node)
	}

	if bf < -1 && key < node.right.key {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	if bf > 1 && key > node.left.key {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	return node
}

func remove(node *AVLTreeNode, key string) *AVLTreeNode {
	if node == nil {
		return node
	}

	if key < node.key {
		node.left = remove(node.left, key)
	} else if key > node.key {
		node.right = remove(node.right, key)
	} else {
		if node.left == nil {
			temp := node.right
			node = nil
			return temp
		} else if node.right == nil {
			temp := node.left
			node = nil
			return temp
		}

		temp := getMinValueNode(node.right)
		node.value = temp.value
		node.right = remove(node.right, temp.value)
	}

	node.height = updateHeight(node)
	bf := getBalanceFactor(node)

	if bf < -1 && getBalanceFactor(node.right) <= 0 {
		return rotateLeft(node)
	}

	if bf > 1 && getBalanceFactor(node.left) >= 0 {
		return rotateRight(node)
	}

	if bf < -1 && getBalanceFactor(node.right) > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	if bf > 1 && getBalanceFactor(node.left) < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	return node
}

func getBalanceFactor(node *AVLTreeNode) int16 {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

func retrieve(node *AVLTreeNode, key string) string {
	if node == nil {
		return "(nil)"
	}

	if node.key == key {
		return node.value
	}

	if key < node.key {
		return retrieve(node.left, key)
	} else {
		return retrieve(node.right, key)
	}
}

func createNode(key string, value string) *AVLTreeNode {
	return &AVLTreeNode{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func getHeight(node *AVLTreeNode) int16 {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

func updateHeight(node *AVLTreeNode) int16 {
	var (
		heightLeftSubtree  int16
		heightRightSubtree int16
	)

	if node.left == nil {
		heightLeftSubtree = 0
	} else {
		heightLeftSubtree = node.left.height
	}

	if node.right == nil {
		heightRightSubtree = 0
	} else {
		heightRightSubtree = node.right.height
	}

	return 1 + max(heightLeftSubtree, heightRightSubtree)
}

func getMinValueNode(node *AVLTreeNode) *AVLTreeNode {
	if node == nil || node.left == nil {
		return node
	}
	return getMinValueNode(node.left)
}

func rotateLeft(node *AVLTreeNode) *AVLTreeNode {
	root := node.right
	rootLeftSubtree := root.left
	root.left = node
	node.right = rootLeftSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))

	return root
}

func rotateRight(node *AVLTreeNode) *AVLTreeNode {
	root := node.left
	rootRightSubtree := root.right
	root.right = node
	node.left = rootRightSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))

	return root
}
