package avltree

import (
	"fmt"
	bst "labtree/bynarysearchtree"
	"math"

	"golang.org/x/exp/constraints"
)

type BalanceFactor int

const (
	UNBALANCED_RIGHT          BalanceFactor = 1
	SLIGHTLY_UMBALANCED_RIGHT BalanceFactor = 2
	BALANCED                  BalanceFactor = 3
	SLIGHTLY_UMBALANCED_LEFT  BalanceFactor = 4
	UNBALANCED_LEFT           BalanceFactor = 5
)

type AVLTree[T constraints.Ordered] struct {
	bst.Bynarysearchtree[T]
}

func NewTree[T constraints.Ordered](compareFn func(node1, node2 T) int) *AVLTree[T] {
	return &AVLTree[T]{
		Bynarysearchtree: *bst.NewTree[T](compareFn),
	}
}

func (t *AVLTree[T]) Remove(key T) {
	t.Root = t.RemoveNode(t.Root, key)

	if t.Root == nil {
		return
	}

	balanceFactor := t.getBalance(t.Root)
	if balanceFactor == UNBALANCED_LEFT {
		balanceFactorLeft := t.getBalance(t.Root.Left)
		if balanceFactorLeft == BALANCED || balanceFactorLeft == SLIGHTLY_UMBALANCED_LEFT {
			t.rotationLL(t.Root)
			return
		}
		if balanceFactorLeft == SLIGHTLY_UMBALANCED_RIGHT {
			t.rotationLR(t.Root.Left)
		}
	}

	if balanceFactor == UNBALANCED_RIGHT {
		balanceFactorRight := t.getBalance(t.Root.Right)
		if balanceFactorRight == BALANCED || balanceFactorRight == SLIGHTLY_UMBALANCED_RIGHT {
			t.rotationRR(t.Root)
			return
		}
		if balanceFactorRight == SLIGHTLY_UMBALANCED_LEFT {
			t.rotationRL(t.Root.Right)
		}
	}
}

func (t *AVLTree[T]) Insert(key T) {
	t.Root = t.insertNode(t.Root, key)
}

func (t *AVLTree[T]) insertNode(node *bst.Node[T], key T) *bst.Node[T] {
	if node == nil {
		return bst.NewNode(key)
	} else if t.CompareFn(key, node.Key) == -1 {
		node.Left = t.insertNode(node.Left, key)
	} else if t.CompareFn(key, node.Key) == 1 {
		node.Right = t.insertNode(node.Right, key)
	} else {
		return node
	}

	balanceFactor := t.getBalance(node)
	if balanceFactor == UNBALANCED_LEFT {
		if t.CompareFn(key, node.Left.Key) == -1 {
			node = t.rotationLL(node)
		} else {
			return t.rotationLR(node)
		}
	}

	if balanceFactor == UNBALANCED_RIGHT {
		if t.CompareFn(key, node.Right.Key) == 1 {
			node = t.rotationRR(node)
		} else {
			return t.rotationRL(node)
		}
	}

	return node
}

func (t *AVLTree[T]) getNodeHeight(node *bst.Node[T]) int {
	if node == nil {
		return -1
	}

	return int(math.Max(
		float64(t.getNodeHeight(node.Left)), float64(t.getNodeHeight(node.Right)),
	)) + 1
}

func (t *AVLTree[T]) getBalance(node *bst.Node[T]) BalanceFactor {
	heightDifference := t.getNodeHeight(node.Left) - t.getNodeHeight(node.Right)
	switch heightDifference {
	case -2:
		return UNBALANCED_RIGHT
	case -1:
		return SLIGHTLY_UMBALANCED_RIGHT
	case 1:
		return SLIGHTLY_UMBALANCED_LEFT
	case 2:
		return UNBALANCED_LEFT
	default:
		return BALANCED
	}
}

func (t *AVLTree[T]) rotationLL(node *bst.Node[T]) *bst.Node[T] {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	return tmp
}

func (t *AVLTree[T]) rotationRR(node *bst.Node[T]) *bst.Node[T] {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	return tmp
}

func (t *AVLTree[T]) rotationLR(node *bst.Node[T]) *bst.Node[T] {
	node.Left = t.rotationRR(node.Left)
	return t.rotationLL(node)
}

func (t *AVLTree[T]) rotationRL(node *bst.Node[T]) *bst.Node[T] {
	node.Right = t.rotationLL(node.Right)
	return t.rotationRR(node)
}

func (t *AVLTree[T]) TreeDebug() {
	t.treeDebug(t.Root)
}

func (t *AVLTree[T]) treeDebug(node *bst.Node[T]) {
	if node == nil {
		return
	}

	t.treeDebug(node.Left)

	if node.Left != nil {
		fmt.Println("node.left: ", node.Left.Key)
	} else {
		fmt.Println("node.left: ")
	}

	fmt.Println("node.Key: ", node.Key)

	if node.Right != nil {
		fmt.Println("node.Right: ", node.Right.Key)
	} else {
		fmt.Println("node.Right: ")
	}
	fmt.Println()

	t.treeDebug(node.Right)
}
