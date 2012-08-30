package main 

import "fmt"


// 二叉树实现学习

type node struct{
	val interface{}
	left, right *node
}

type tree struct{
	root *node
	size int 
	compare func (a, b interface{}) int
}

func compare(a, b interface{}) int{
	switch i := a.(type){
	case int:
		switch j := b.(type){
		case int:
			return i - j
		}
	}
	return 0
}

func NewTree() *tree{

	return &tree{nil,0,compare}
}

func (this *tree) Add(v interface{}){
	node := &node{v, nil, nil} 
	if this.root == nil {
		this.root = node
		return
	}

	t := this.root
	for ; ;{
		if this.compare(t.val, v) > 0 {
			if t.left == nil {
				t.left = node
				break;
			}else{
				t = t.left
			}
		}else {
			if t.right == nil {
				t.right = node
				break;
			}else{
				t = t.right
			}
		}
	}
}

func subtree(n *node) *tree{
	return &tree{n,0,nil}
}

func (this *tree) inorderWalk() string {
	var s = ""
	if this.root.left != nil {
		s =  subtree(this.root.left).inorderWalk()
	}
	s = s + ","+fmt.Sprint(this.root.val)
	if this.root.right != nil {
		s = s + ","+ subtree(this.root.right).inorderWalk()
	}
	return s
}

func (this *tree) preorder(){
}

func main() {

	tr := NewTree()
	tr.Add(10)
	tr.Add(1)
	tr.Add(2)
	tr.Add(3)
	tr.Add(14)

fmt.Println(	tr.inorderWalk())

	fmt.Println(compare(1,2))
}