package main

// 二叉树是结合有序数组(查找)和链表(插入和删除)优点的数据结构，
// 树是一种特殊的图，二叉树表示最多有2个节点的树

import "math/rand"
import "fmt"

type ArrayTree struct {
	nodes []interface{}
}

type node struct {
	val   interface{}
	left  *node
	right *node
}

// 二叉树数据必须是可比较的对象
type BinaryTree struct {
	root *node
	cmp  func(a, b interface{}) int
	size int
}

// 查找节点是否存在
func (this *BinaryTree) Find(v interface{}) bool {
	if this == nil {
		return false
	}

	// 开始指向根节点
	r := this.root
	for r != nil {
		c := this.cmp(r.val, v)
		if c == 0 {
			return true
		} else if c > 0 {
			r = r.left
		} else {
			r = r.right
		}
	}

	return false
}

func (this *BinaryTree) Add(v interface{}) {
	n := &node{val: v}

	if this.root == nil {
		this.root = n
		return
	}

	r := this.root
	var pre *node = nil
	var c int

	for r != nil {
		pre = r
		c = this.cmp(r.val, v)
		if c == 0 {
			return
		} else if c > 0 {
			r = r.left
		} else {
			r = r.right
		}
	}

	if pre == nil {
		this.root = n
	} else if c > 0 {
		pre.left = n
	} else {
		pre.right = n
	}

	this.size++
}

func (this *BinaryTree) String() string {

	ret := ""

	data := this.InOrder_iter()
	for _, v := range data {
		ret = fmt.Sprint(ret, v.val, ", ")
	}

	return ret
}

// 中序遍历

func (this *BinaryTree) InOrder_iter() []*node {
	// 把根节点及其所有的左子节点都压入栈
	// 初始条件 r == root
	// 当 r不为null时，入栈然后r = r.left再循环
	// 如果r == nil, 弹出栈顶元素并访问，然后 r = 栈顶元素的右节点

	stack := []*node{}

	ret := []*node{}
	r := this.root
	for ; r != nil; r = r.left {
		stack = append(stack, r)
	}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		ret = append(ret, pop)
		stack = stack[:len(stack)-1]

		r = pop.right
		for ; r != nil; r = r.left {
			stack = append(stack, r)
		}
	}

	return ret
}

func (this *BinaryTree) InOrder_iter_1() []*node {
	stack := []*node{}
	ret := []*node{}

	r := this.root
	for r != nil || len(stack) > 0 {
		if r != nil {
			stack = append(stack, r)
			r = r.left
		} else {
			v := stack[len(stack)-1]
			ret = append(ret, v)
			stack = stack[:len(stack)-1]
			r = v.right
		}
	}

	return ret
}

func (this *BinaryTree) InOrder() []*node {
	d := []*node{}

	if this.root == nil {
		return d
	}

	return _inorder(this.root)
}

func _inorder(n *node) []*node {
	d := []*node{}
	if n.left != nil {
		d = append(d, _inorder(n.left)...)
	}
	d = append(d, n)
	if n.right != nil {
		d = append(d, _inorder(n.right)...)
	}

	return d
}

// 先序遍历
func (this *BinaryTree) PreOrder_iter() []*node {
	data := []*node{}
	stack := []*node{}

	r := this.root
	for r != nil || len(stack) > 0 {
		if r != nil {

			data = append(data, r)
			if r.right != nil {
				stack = append(stack, r.right)
			}

			r = r.left
		} else {
			r = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return data
}

// 后续遍历
func (this *BinaryTree) PostOrder_iter() []*node {
	data := []*node{}

	r := this.root
	if r == nil {
		return data
	}

	stack := []*node{r}
	var prevNode *node

	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		if prevNode == nil || prevNode.left == currNode || prevNode.right == currNode {
			if currNode.left != nil {
				stack = append(stack, currNode.left)
			} else if currNode.right != nil {
				stack = append(stack, currNode.right)
			} else {
				// 达到一个叶子节点
			}
		} else if currNode.left == prevNode {
			if currNode.right != nil {
				stack = append(stack, currNode.right)
			}
		} else {
			data = append(data, currNode)
			stack = stack[:len(stack)-1]
		}

		prevNode = currNode
	}

	return data
}

// 后序遍历
func (this *BinaryTree) PostOrder() []*node {
	if this.root == nil {
		return []*node{}
	}

	return _postOrder(this.root)
}

func _postOrder(n *node) []*node {
	data := []*node{}
	if n.left != nil {
		data = append(data, _postOrder(n.left)...)
	}
	if n.right != nil {
		data = append(data, _postOrder(n.right)...)
	}

	data = append(data, n)

	return data
}

func NewBinaryTree(f func(a, b interface{}) int) *BinaryTree {
	return &BinaryTree{cmp: f}
}

func ToVals(ns []*node) []interface{} {
	l := len(ns)
	data := make([]interface{}, l)

	for i, v := range ns {
		data[i] = v.val
	}

	return data
}

func (this *BinaryTree) Min() *node {
	cur := this.root

	for cur != nil {
		cur = cur.left
	}

	return cur
}

func (this *BinaryTree) Delete(v interface{}) {
	currNode := this.root
	var preNode *node
	var findNode *node

	for currNode != nil {
		c := this.cmp(currNode.val, v)
		if c == 0 {
			findNode = currNode
			break
		} else if c > 0 {
			preNode = currNode
			currNode = currNode.left
		} else {
			preNode = currNode
			currNode = currNode.right
		}

	}

	if findNode == nil {
		return
	}

	if findNode.left == nil && findNode.right == nil {
		if preNode == nil {
			this.root = nil
		} else {
			if preNode.left == findNode {
				preNode.left = nil
			} else {
				preNode.right = nil
			}
		}
	} else if findNode.left == nil {
		if preNode == nil {
			this.root = findNode.right
		} else {
			if preNode.left == findNode {
				preNode.left = findNode.right
			} else {
				preNode.right = findNode.right
			}
		}
	} else if findNode.right == nil {
		if preNode == nil {
			this.root = findNode.left
		} else {
			if preNode.left == findNode {
				preNode.left = findNode.left
			} else {
				preNode.right = findNode.left
			}
		}
	} else {
		// 2个子节点，把右子节点挂在左节点的最大节点的右孩子
		// 或者把左子树挂在右节点的最小节点的左孩子
		/*	l := findNode.left
			for l != nil {
				l = l.right
			}
			l.right = findNode.right

			if preNode == nil {
				this.root = findNode.left
			} else {
				if preNode.left == findNode {
					preNode.left = findNode.left
				} else {
					preNode.right = findNode.left
				}
			}
		*/

		// 或者把删除的后继者找到，然后用后继者替换要被删除的节点。
		// 后继者的左子节点肯定是空的
		// 如果后继者是被删除节点的右节点，那么设置后继者的左节点为要被删除节点的左节点

		succ := Successor(findNode)

		if preNode == nil {
			this.root = succ
		} else {
			if preNode.left == findNode {
				preNode.left = succ
			} else {
				preNode.right = succ
			}
		}

	}
}

func Successor(n *node) *node {
	current := n.right
	parent := n
	successor := n.right

	for current.left != nil {
		parent = current
		current = current.left
		successor = current
	}

	if successor == n.right {
		successor.left = n.left
	} else {
		parent.left = successor.right
		successor.right = n.right
		successor.left = n.left
	}

	return successor
}

func main() {
	fn := func(a, b interface{}) int {
		aint, _ := a.(int)
		bint, _ := b.(int)

		return aint - bint
	}

	tree := NewBinaryTree(fn)

	for i := 0; i < 100; i++ {
		tree.Add(Rand())
	}

	fmt.Println(tree)

	k := ToVals(tree.InOrder_iter())

	fmt.Println(ToVals(tree.InOrder_iter()))
	fmt.Println(ToVals(tree.InOrder_iter_1()))
	fmt.Println(ToVals(tree.PreOrder_iter()))
	fmt.Println(ToVals(tree.PostOrder_iter()))

	for _, v := range k[0:(len(k) / 2)] {
		tree.Delete(v)
	}

	fmt.Println(ToVals(tree.InOrder_iter()))
	fmt.Println(ToVals(tree.InOrder_iter_1()))
	fmt.Println(ToVals(tree.PreOrder_iter()))
	fmt.Println(ToVals(tree.PostOrder_iter()))
}

func Rand() int {
	return rand.Int() % 100
}
