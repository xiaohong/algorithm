package binarySearchTree

import "ds"
import "fmt"

// 12.1-1
// 深度为k的二叉树，最多有2^k-1个节点，也就是说树的高度是层数，
// 只有一个根节点的树的高度为1
//        10
//      /    \
//     4     16
//    / \      \
//    1  5     17
//               \
//               21  

//         10
//      /     \
//     4       17
//    / \     /  \
//    1  5    16 21

// 12.1-2
// 二叉搜索树严格满足根节点大于所有的左子树节点，小于所有右子树节点
// 最小堆值要求父节点小于左右节点，但是没有对左右节点的大小进行规定

// 最小堆不能够在O(n)的时间内打印出有序数据，因为当前我们只能保证打印到
//最小数据，等这个数据从堆中删除我们需要重新选择出最小的数据放在根
//上，而这个查找的时间复杂度是lg(n)???

// 12.1-3

func (this *Tree) Inorder_with_stack() {
	stack := ds.NewArrayStack()
	root := this.Root

	// 首先把根节点下面的所有左节点全部进栈
	for root != nil {
		stack.Push(root)
		root = root.left
	}

	// 首先打印最小数据，然后打印这个节点的右子树
	// 如果这个右子树有左节点时，再进栈
	for stack.Size() > 0 {
		t, ok := (stack.Pop()).(*Node)
		if ok {
			fmt.Println(t.Value)
			t = t.right
			for t != nil {
				stack.Push(t)
				t = t.left
			}
		}
	}
}

func (this *Tree) Inorder_with_stack_2() {
	n := this.Root
	stack := ds.NewArrayStack()

	for stack.Size() > 0 || n != nil {
		if n != nil {
			stack.Push(n)
			n = n.left
		} else {
			t, _ := stack.Pop().(*Node)
			n = t.right
		}
	}
}

//12.1-4
func (this *Tree) Preorder() {
	fmt.Print(this.Root.Value, ", ")
	if this.leftSubTree() != nil {
		this.leftSubTree().Preorder()
	}
	if this.rightSubTree() != nil {
		this.rightSubTree().Preorder()
	}
}
