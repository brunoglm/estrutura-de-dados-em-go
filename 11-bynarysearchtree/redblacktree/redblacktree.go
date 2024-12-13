package bynarysearchtree

import "golang.org/x/exp/constraints"

type Color int

const (
	RED Color = iota
	BLACK
)

type Node[T constraints.Ordered] struct {
	Key    T
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]
	Color  Color
}

func NewNode[T constraints.Ordered](key T) *Node[T] {
	return &Node[T]{
		Key:   key,
		Color: RED,
	}
}

func (t *Node[T]) IsRed() bool {
	return t.Color == RED
}

type Redblacktree[T constraints.Ordered] struct {
	CompareFn func(node1, node2 T) int
	Root      *Node[T]
}

func NewTree[T constraints.Ordered](compareFn func(node1, node2 T) int) *Redblacktree[T] {
	if compareFn == nil {
		compareFn = defaultCompare[T]
	}

	return &Redblacktree[T]{
		CompareFn: compareFn,
	}
}

func (t *Redblacktree[T]) Insert(key T) {
	if t.Root == nil {
		t.Root = NewNode(key)
		t.Root.Color = BLACK
		return
	}

	newNode := t.insertNode(t.Root, key)
	t.fixTreeProperties(newNode)
}

func (t *Redblacktree[T]) insertNode(node *Node[T], key T) *Node[T] {
	if t.CompareFn(key, node.Key) == -1 {
		if node.Left == nil {
			node.Left = NewNode(key)
			node.Left.Parent = node
			return node.Left
		} else {
			return t.insertNode(node.Left, key)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(key)
			node.Right.Parent = node
			return node.Right
		} else {
			return t.insertNode(node.Right, key)
		}
	}
}

func (t *Redblacktree[T]) fixTreeProperties(node *Node[T]) {
	for node != nil && node.Parent != nil && node.Parent.IsRed() && node.Color != BLACK {
		parent := node.Parent
		grandParent := parent.Parent

		// caso 1A: o tio do nó também é vermelho - basta recolorir
		if grandParent != nil && grandParent.Left == parent {
			uncle := grandParent.Right
			if uncle != nil && uncle.IsRed() {
				grandParent.Color = RED
				parent.Color = BLACK
				uncle.Color = BLACK
				node = grandParent
			} else {
				// caso 2A: o nó é o filho à direita - Rotação à esquerda
				if node == parent.Right {
					t.rotationRR(parent)
					node = parent
					parent = node.Parent
				}
				// caso 3A: o nó é o filho à esquerda - Rotação à direita
				t.rotationLL(grandParent)
				parent.Color = BLACK
				grandParent.Color = RED
				node = parent
			}
		} else { // Caso B: o pai é o filho à direita
			uncle := grandParent.Left
			// caso 1B: o tio é vermelho - basta recolorir
			if uncle != nil && uncle.IsRed() {
				grandParent.Color = RED
				parent.Color = BLACK
				uncle.Color = BLACK
				node = grandParent
			} else {
				// caso 2B: o nó é o filho à esquerda - Rotação à direita
				if node == parent.Left {
					t.rotationLL(parent)
					node = parent
					parent = node.Parent
				}
				// caso 3B: o nó é o filho à direita - Rotação à esquerda
				t.rotationRR(grandParent)
				parent.Color = BLACK
				grandParent.Color = RED
				node = parent
			}

		}
	}
	t.Root.Color = BLACK
}

func (t *Redblacktree[T]) InOrderTraverse(callback func(key T)) {
	t.inOrderTraverseNode(t.Root, callback)
}

func (t *Redblacktree[T]) inOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.inOrderTraverseNode(node.Left, callback)
	callback(node.Key)
	t.inOrderTraverseNode(node.Right, callback)
}

func (t *Redblacktree[T]) PreOrderTraverse(callback func(key T)) {
	t.preOrderTraverseNode(t.Root, callback)
}

func (t *Redblacktree[T]) preOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	callback(node.Key)
	t.preOrderTraverseNode(node.Left, callback)
	t.preOrderTraverseNode(node.Right, callback)
}

func (t *Redblacktree[T]) PostOrderTraverse(callback func(key T)) {
	t.postOrderTraverseNode(t.Root, callback)
}

func (t *Redblacktree[T]) postOrderTraverseNode(node *Node[T], callback func(key T)) {
	if node == nil {
		return
	}

	t.postOrderTraverseNode(node.Left, callback)
	t.postOrderTraverseNode(node.Right, callback)
	callback(node.Key)
}

func (t *Redblacktree[T]) Min() T {
	return t.minNode(t.Root).Key
}

func (t *Redblacktree[T]) minNode(node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *Redblacktree[T]) Max() T {
	return t.maxNode(t.Root).Key
}

func (t *Redblacktree[T]) maxNode(node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Right != nil {
		current = current.Right
	}
	return current
}

func (t *Redblacktree[T]) Search(key T) bool {
	return t.searchNode(t.Root, key)
}

func (t *Redblacktree[T]) searchNode(node *Node[T], key T) bool {
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

func (t *Redblacktree[T]) Remove(key T) {
	t.Root = t.RemoveNode(t.Root, key)
}

func (t *Redblacktree[T]) RemoveNode(node *Node[T], key T) *Node[T] {
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

func (t *Redblacktree[T]) rotationLL(node *Node[T]) *Node[T] {
	tmp := node.Left
	node.Left = tmp.Right

	if tmp.Right != nil {
		tmp.Right.Parent = node
	}
	tmp.Parent = node.Parent
	if node.Parent == nil {
		t.Root = tmp
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = tmp
		} else {
			node.Parent.Right = tmp
		}
	}

	tmp.Right = node
	node.Parent = tmp
	return tmp
}

func (t *Redblacktree[T]) rotationRR(node *Node[T]) *Node[T] {
	tmp := node.Right
	node.Right = tmp.Left

	if tmp.Left != nil {
		tmp.Left.Parent = node
	}

	tmp.Parent = node.Parent

	if node.Parent == nil {
		t.Root = tmp
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = tmp
		} else {
			node.Parent.Right = tmp
		}
	}

	tmp.Left = node
	node.Parent = tmp
	return tmp
}
