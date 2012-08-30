package binarySearchTree

import "fmt"

type Node struct{
	Value interface{}
	left, right, parent *Node
}

func (this *Node) ToString() string{
	return fmt.Sprint(this.Value)
}

type Tree struct{
	Root *Node
	Size int 
	Compare func (a, b interface{}) int
}

func compare(a, b interface{}) int {
	switch v := a.(type) {
	case int :
		switch m := b.(type) {
		case int :
			return int(v) - int(m)
		}
	}

	return 0
}


func (this *Tree) leftSubTree() *Tree{
	if this.Root.left == nil {
		return nil
	}
	return &Tree{Root: this.Root.left}
}

func (this *Tree) rightSubTree() *Tree{
	if this.Root.right == nil {
		return nil
	}
	return &Tree{Root : this.Root.right}
}

func NewTree() *Tree{
	return &Tree{Compare:compare}
}

func (this *Tree) Insert(v interface{}) {
	if this.Root == nil {
		this.Root = &Node{Value: v}
		return 
	}


	t := this.Root
	p := t 

	for ; t != nil ; {
		p = t
		n := this.Compare(t.Value, v)
		if n > 0 {
			t = t.left
		}else if  n < 0 {
			t = t.right
		}else {
			break;
		}
	}

	n := this.Compare(p.Value, v)
	if n > 0 {
		p.left = &Node{Value : v}
	}else if n < 0 {
		p.right = &Node{Value : v}
	}
}
func (this *Tree) Inorder(){
	left := this.leftSubTree()
	if left != nil {
		left.Inorder()
	}
	fmt.Print(this.Root.Value, ",")
	right := this.rightSubTree()
	if right != nil {
		right.Inorder()
	}
}
