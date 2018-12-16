package Search

import "fmt"

// # 二分探索木
// メモリ上の配列に対する二分探索は効率的であるが、探索対象の集合の要素が頻繁に変化するようでは、せっかくの整列済み配列に対する二分探索の
// 効率の良さが台無しになってしまうほど効率が落ちる。要素が動的に変化する集合に対する高速な検索には配列というデータ構造は適していない。
// ハッシュを使えばいいではないかという声はもっともであるが、大規模な集合に対するハッシュを用いた探索アルゴリズムは空間計算量を膨大に
// 必要とするので、ベストではない。それなりの量のデータを高速に探索するためにはどのようなデータ構造を用いればいいのだろうか。
// そこで登場するのが探索木というデータ構造である。もっとも使われるのは二分探索木で、一人の親に対して左右の子がいるという木構造を成す。
// 二分探索木が有効な場面は例えば次のような場合がある。
// - データを昇順もしくは降順で走査する必要がある時
// - データセットのサイズが未知の時。実装の段階では最大でメモリに格納できるサイズまで対応させる必要がある。
// - データセットが動的に変化する時。多数の挿入および削除が起こることが予想される。
// 二分探索木は常に平衡を取れているとは限らない。つまり常に左に子供がいる（右には子供がいない）というものも二分探索木であるということである。
// しかし常に左に子供がいるような二分探索木はもはや木ではなく配列と同じである。平衡が取れていないような木構造に変化することを「退行」という。
// 二分探索木を用いる場合には要素の挿入削除ごとに左右が平衡を取るように実装することが重要である。

type Node struct {
	value       int
	left, right *Node
	height      int
}

func NewNode(val int) *Node {
	return &Node{value: val, left: nil, right: nil, height: 0}
}

func (node *Node) Add(val int) *Node {
	// 素朴な二分探索木の実装
	if val <= node.value {
		if node.left != nil {
			node.left.Add(val)
		} else {
			node.left = NewNode(val)
		}
	} else {
		if node.right != nil {
			node.right.Add(val)
		} else {
			node.right = NewNode(val)
		}
	}
	return node
	// TODO: AVL木になるように実装
	// newRoot := node
	// if val <= node.value {
	// 	node.left = node.addToSubtree(node.left, val)
	// 	if node.calcHeightDifference() == 2 {
	// 		if val <= node.left.value {
	// 			newRoot = node.rotateRight()
	// 		} else {
	// 			newRoot = node.rotateLeftRight()
	// 		}
	// 	}
	// } else {
	// 	node.right = node.addToSubtree(node.right, val)
	// 	if node.calcHeightDifference() == -2 {
	// 		if node.right.value < val {
	// 			newRoot = node.rotateLeft()
	// 		} else {
	// 			newRoot = node.rotateRightLeft()
	// 		}
	// 	}
	// }
	// newRoot.calcHeight()
	// return newRoot
}

func (node *Node) calcHeight() {
	height := -1
	if node.left != nil {
		height = max(height, node.left.height)
	}
	if node.right != nil {
		height = max(height, node.right.height)
	}
	node.height = height + 1
}

func (node *Node) calcHeightDifference() int {
	leftHeight := 0
	rightHeight := 0
	if node.left != nil {
		leftHeight = node.left.height + 1
	}
	if node.right != nil {
		rightHeight = node.right.height + 1
	}
	return leftHeight - rightHeight
}

func (node *Node) addToSubtree(parent *Node, val int) *Node {
	if parent == nil {
		return NewNode(val)
	}
	parent = parent.Add(val)
	return parent
}

func (node *Node) rotateRight() *Node {
	newRoot := node.left
	grandson := newRoot.right
	node.left = grandson
	newRoot.right = node
	node.calcHeight()
	return newRoot
}

func (node *Node) rotateRightLeft() *Node {
	child := node.right
	newRoot := child.left
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.left = grand2
	child.right = grand1
	newRoot.left = node
	newRoot.right = child
	child.calcHeight()
	node.calcHeight()
	return newRoot
}

func (node *Node) rotateLeft() *Node {
	newRoot := node.right
	grandson := newRoot.left
	node.right = grandson
	newRoot.left = node
	node.calcHeight()
	return newRoot
}

func (node *Node) rotateLeftRight() *Node {
	child := node.left
	newRoot := child.right
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.right = grand2
	node.left = grand1
	newRoot.left = child
	newRoot.right = node
	child.calcHeight()
	node.calcHeight()
	return newRoot
}

func (node *Node) RemoveFromParent(parent *Node, val int) *Node {
	if parent != nil {
		return parent.remove(val)
	}
	return nil
}

func (node *Node) remove(val int) *Node {
	newRoot := node
	if val == node.value {
		if node.left == nil {
			return node.right
		}
		child := node.left
		for child.right != nil {
			child = child.right
		}
		childKey := child.value
		node.left = node.RemoveFromParent(node.left, childKey)
		node.value = childKey
		if node.calcHeightDifference() == -2 {
			if node.right.calcHeightDifference() <= 0 {
				newRoot = node.rotateLeft()
			} else {
				newRoot = node.rotateRightLeft()
			}
		}
	} else if val < node.value {
		node.left = node.RemoveFromParent(node.left, val)
		if node.calcHeightDifference() == -2 {
			if node.right.calcHeightDifference() <= 0 {
				newRoot = node.rotateLeft()
			} else {
				newRoot = node.rotateRightLeft()
			}
		}
	} else {
		node.right = node.RemoveFromParent(node.right, val)
		if node.calcHeightDifference() == 2 {
			if node.left.calcHeightDifference() >= 0 {
				newRoot = node.rotateRight()
			} else {
				newRoot = node.rotateLeftRight()
			}
		}
	}
	newRoot.calcHeight()
	return newRoot
}

func (node *Node) Inorder() {
	if node.left != nil {
		node.left.Inorder()
	}
	fmt.Println(node.value)
	if node.right != nil {
		node.right.Inorder()
	}
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

func (tree *Tree) Add(val int) {
	if tree.root == nil {
		tree.root = NewNode(val)
	} else {
		tree.root.Add(val)
	}
}

func (tree *Tree) Search(t int) bool {
	node := tree.root
	for node != nil {
		if t < node.value {
			node = node.left
		} else if node.value < t {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (tree *Tree) Inorder() {
	if tree.root != nil {
		tree.root.Inorder()
	} else {
		fmt.Println("There is no tree.")
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
