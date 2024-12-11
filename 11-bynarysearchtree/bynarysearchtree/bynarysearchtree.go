package bynarysearchtree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Key   T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T constraints.Ordered](key T) *Node[T] {
	return &Node[T]{
		Key: key,
	}
}

type Bynarysearchtree[T constraints.Ordered] struct {
	CompareFn func(node1, node2 T) int
	Root      *Node[T]
}

func NewTree[T constraints.Ordered](compareFn func(node1, node2 T) int) *Bynarysearchtree[T] {
	if compareFn == nil {
		compareFn = defaultCompare[T]
	}

	return &Bynarysearchtree[T]{
		CompareFn: compareFn,
	}
}

func (t *Bynarysearchtree[T]) Insert(key T) {
	if t.Root == nil {
		t.Root = NewNode(key)
		return
	}

	t.insertNode(t.Root, key)
}

func (t *Bynarysearchtree[T]) insertNode(node *Node[T], key T) {
	if t.CompareFn(key, node.Key) == -1 {
		if node.Left == nil {
			node.Left = NewNode(key)
		} else {
			t.insertNode(node.Left, key)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(key)
		} else {
			t.insertNode(node.Right, key)
		}
	}
}

func (t *Bynarysearchtree[T]) InOrderTraverse(callback func(key T)) {
	t.inOrderTraverseNode(t.Root, callback)
}

func (t *Bynarysearchtree[T]) inOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.inOrderTraverseNode(node.Left, callback)
	callback(node.Key)
	t.inOrderTraverseNode(node.Right, callback)
}

func (t *Bynarysearchtree[T]) PreOrderTraverse(callback func(key T)) {
	t.preOrderTraverseNode(t.Root, callback)
}

func (t *Bynarysearchtree[T]) preOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	callback(node.Key)
	t.preOrderTraverseNode(node.Left, callback)
	t.preOrderTraverseNode(node.Right, callback)
}

func (t *Bynarysearchtree[T]) PostOrderTraverse(callback func(key T)) {
	t.postOrderTraverseNode(t.Root, callback)
}

func (t *Bynarysearchtree[T]) postOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.postOrderTraverseNode(node.Left, callback)
	t.postOrderTraverseNode(node.Right, callback)
	callback(node.Key)
}

func (t *Bynarysearchtree[T]) Min() T {
	return t.minNode(t.Root).Key
}

func (t *Bynarysearchtree[T]) minNode(node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *Bynarysearchtree[T]) Max() T {
	return t.maxNode(t.Root).Key
}

func (t *Bynarysearchtree[T]) maxNode(node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Right != nil {
		current = current.Right
	}
	return current
}

func (t *Bynarysearchtree[T]) Search(key T) bool {
	return t.searchNode(t.Root, key)
}

func (t *Bynarysearchtree[T]) searchNode(node *Node[T], key T) bool {
	if node == nil {
		return false
	}

	if t.CompareFn(key, node.Key) == -1 {
		return t.searchNode(node.Left, key)
	} else if t.CompareFn(key, node.Key) == 1 {
		return t.searchNode(node.Right, key)
	} else {
		return true
	}
}

func (t *Bynarysearchtree[T]) Remove(key T) {
	t.Root = t.RemoveNode(t.Root, key)
}

func (t *Bynarysearchtree[T]) RemoveNode(node *Node[T], key T) *Node[T] {
	if node == nil {
		return nil
	}

	if t.CompareFn(key, node.Key) == -1 {
		node.Left = t.RemoveNode(node.Left, key)
		return node
	} else if t.CompareFn(key, node.Key) == 1 {
		node.Right = t.RemoveNode(node.Right, key)
		return node
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
			return node
		}

		if node.Left == nil {
			node = node.Right
			return node
		}

		if node.Right == nil {
			node = node.Left
			return node
		}

		aux := t.minNode(node.Right)
		node.Key = aux.Key
		node.Right = t.RemoveNode(node.Right, aux.Key)
		return node
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
