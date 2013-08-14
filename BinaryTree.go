package cslib

type BinaryTreeNode struct {
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
	Value  int
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (tree *BinaryTree) Insert(value int) {
	node := &BinaryTreeNode{Value: value}

	if tree.Root == nil {
		tree.Root = node

	} else {
		current := tree.Root
		for {
			if current.Value >= value {
				if current.Left == nil {
					current.Left = node
					node.Parent = current
					break
				} else {
					current = current.Left
				}
			} else {
				if current.Right == nil {
					current.Right = node
					node.Parent = current
					break
				} else {
					current = current.Right
				}
			}
		}
	}
}

func (tree *BinaryTree) Search(value int) *BinaryTreeNode {
	current := tree.Root
	for current != nil {
		if current.Value == value {
			break
		} else if current.Value > value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return current
}

func (tree *BinaryTree) Traverse(f func(int)) {
	if tree.Root != nil {
		tree.Root.traverse(f)
	}
}

func (node *BinaryTreeNode) traverse(f func(int)) {
	if node.Left != nil {
		node.Left.traverse(f)
	}
	f(node.Value)
	if node.Right != nil {
		node.Right.traverse(f)
	}
}
