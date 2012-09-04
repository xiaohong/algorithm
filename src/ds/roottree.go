package ds

// root tree, 不同的树有不同的子节点，例如二叉树有2个子节点
// 没有边界分支的树，通过表示左节点和右兄弟节点，来表示无穷的子节点

// 如果是父节点的最右边节点，则rightSibling为nil
// 如果有子节点的话，left不为空
type node struct{
	parent *node
	left *node
	rightSibling *node
	Value interface{}
}

type RootTree struct{
	root *node
}

func (this *RootTree) Insert(v interface{}){
	n := &node{Value: v}
	if this.root == nil {
		this.root = n 
		return
	}
	if this.root.left == nil {
		this.root.left = n
		n.parent = this.root
	}else {
		pre := this.root.left
		cur := pre.rightSibling
		for cur != nil {
			pre = cur
			cur = cur.rightSibling
		}
		pre.rightSibling = n 
		n.parent = this.root
	}
}

// 可以通过数组模拟树，通过index表示父子关系