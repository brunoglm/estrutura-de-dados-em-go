package bynarysearchtree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	key   T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T constraints.Ordered](key T) *Node[T] {
	return &Node[T]{
		key: key,
	}
}

type Bynarysearchtree[T constraints.Ordered] struct {
	compareFn func(node1, node2 T) int
	root      *Node[T]
}

func NewBynarysearchtree[T constraints.Ordered](compareFn func(node1, node2 T) int) *Bynarysearchtree[T] {
	if compareFn == nil {
		compareFn = defaultCompare[T]
	}

	return &Bynarysearchtree[T]{
		compareFn: compareFn,
	}
}

func (t *Bynarysearchtree[T]) Insert(key T) {
	if t.root == nil {
		t.root = NewNode(key)
		return
	}

	t.insertNode(t.root, key)
}

func (t *Bynarysearchtree[T]) insertNode(node *Node[T], key T) {
	if t.compareFn(key, node.key) == -1 {
		if node.left == nil {
			node.left = NewNode(key)
		} else {
			t.insertNode(node.left, key)
		}
	} else {
		if node.right == nil {
			node.right = NewNode(key)
		} else {
			t.insertNode(node.right, key)
		}
	}
}

func (t *Bynarysearchtree[T]) InOrderTraverse(callback func(key T)) {
	t.inOrderTraverseNode(t.root, callback)
}

func (t *Bynarysearchtree[T]) inOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.inOrderTraverseNode(node.left, callback)
	callback(node.key)
	t.inOrderTraverseNode(node.right, callback)
}

func (t *Bynarysearchtree[T]) PreOrderTraverse(callback func(key T)) {
	t.preOrderTraverseNode(t.root, callback)
}

func (t *Bynarysearchtree[T]) preOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	callback(node.key)
	t.preOrderTraverseNode(node.left, callback)
	t.preOrderTraverseNode(node.right, callback)
}

func (t *Bynarysearchtree[T]) PostOrderTraverse(callback func(key T)) {
	t.postOrderTraverseNode(t.root, callback)
}

func (t *Bynarysearchtree[T]) postOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.postOrderTraverseNode(node.left, callback)
	t.postOrderTraverseNode(node.right, callback)
	callback(node.key)
}

func (t *Bynarysearchtree[T]) Min() T {
	return t.minNode(t.root)
}

func (t *Bynarysearchtree[T]) minNode(node *Node[T]) T {
	current := node
	for current != nil && current.left != nil {
		current = current.left
	}
	return current.key
}

func (t *Bynarysearchtree[T]) Max() T {
	return t.maxNode(t.root)
}

func (t *Bynarysearchtree[T]) maxNode(node *Node[T]) T {
	current := node
	for current != nil && current.right != nil {
		current = current.right
	}
	return current.key
}

func (t *Bynarysearchtree[T]) Search(key T) bool {
	return t.searchNode(t.root, key)
}

func (t *Bynarysearchtree[T]) searchNode(node *Node[T], key T) bool {
	if node == nil {
		return false
	}

	if t.compareFn(key, node.key) == -1 {
		return t.searchNode(node.left, key)
	} else if t.compareFn(key, node.key) == 1 {
		return t.searchNode(node.right, key)
	} else {
		return true
	}
}

func defaultCompare[T constraints.Ordered](node1, node2 T) int {
	if node1 == node2 {
		return 0
	}

	if node1 < node2 {
		return -1
	}

	return 1
}
