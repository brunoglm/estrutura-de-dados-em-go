package avltree

import (
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
