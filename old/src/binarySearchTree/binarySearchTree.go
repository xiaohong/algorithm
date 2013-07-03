package binarySearchTree

import "fmt"

type Node struct {
	Value               interface{}
	left, right, parent *Node
}

func (this *Node) String() string {
	return fmt.Sprint(this.Value)
}



type Tree struct {
	Root    *Node
	Size    int
	Compare func(a, b interface{}) int
}

func compare(a, b interface{}) int {
	switch v := a.(type) {
	case int:
		switch m := b.(type) {
		case int:
			return int(v) - int(m)
		}
	}

	return 0
}

func (this *Tree) leftSubTree() *Tree {
	if this.Root.left == nil {
		return nil
	}
	return &Tree{Root: this.Root.left, Compare: this.Compare}
}

func (this *Tree) rightSubTree() *Tree {
	if this.Root.right == nil {
		return nil
	}
	return &Tree{Root: this.Root.right, Compare: this.Compare}
}

func NewTree() *Tree {
	return &Tree{Compare: compare}
}

// 
func (this *Tree) Insert(v interface{}) {
	if this.Root == nil {
		this.Root = &Node{Value: v}
		return
	}

	t := this.Root
	p := t
	for t != nil {
		p = t
		n := this.Compare(t.Value, v)
		if n > 0 {
			t = t.left
		} else if n < 0 {
			t = t.right
		} else {
			break
		}
	}

	n := this.Compare(p.Value, v)
	if n > 0 {
		p.left = &Node{Value: v, parent: p}
	} else if n < 0 {
		p.right = &Node{Value: v, parent: p}
	}
}

// 首先需要查找插入的位置
// 通常遍历的时候，有一个指针指向正在遍历的数据
// 另外一个指针维护遍历的轨迹
func (this *Tree) Insert_1(v interface{}){
	var y *Node= nil
	x := this.Root

	for x != nil {
		y = x // 记住上次遍历的变量
		n := this.Compare(x.Value, v)
		// 如果是相等元素，等往右子树上放
		if n > 0 {
			x = x.left
		}else {
			x = x.right
		}
	}

	z := &Node{Value : v, parent: y}
	if y == nil {
		this.Root = z
	}else {
		n := this.Compare(y.Value, v) 
		if n > 0 {
			y.left = z
		}else{
			y.right = z
		}
	}
}

func (this *Tree) Inorder() {
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

func (this *Tree) search(v interface{}) *Node{
	r := this.Root 
	for r != nil {
		if r.Value == v{
			return r
		}

		if this.Compare(r.Value, v) > 0 {
			r = r.left 
		}else {
			r = r.right
		}
	}
	return nil
}

// v节点替换u节点
func (this *Tree) transplant(u, v *Node){
	if u.parent == nil {
		this.Root = v
	}else if u == u.parent.left {
		u.parent.left = v
	}else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent	
	}
}

func (this *Tree) Delete(v interface{}) {
	r := this.search(v) 
	if r == nil {
		return
	}

	if r.left == nil {
		this.transplant(r, r.right)
	}else if r.right == nil {
		this.transplant(r, r.left)
	}else {
		// 左右子树都存在的情况，取出后继者
		// 看后继者是不是他的直接右节点，如果不是则需要先删除后继者
		// 因为后继者肯定没有左子树，最多只有一个右子节点，所以直接用右子树替换
		// 然后再用后继者替换删除节点
		successor := r.Successor()
		if successor.parent != r { //右子树中还有左子树
			this.transplant(successor, successor.right)
			successor.right = r.right 
			successor.right.parent = r
		}
		// successor肯定没有左子树了
		this.transplant(r, successor)
		// 更新左子树
	//	y.left = z.left
	//	y.left.parent = y

	}
}


func Test_Insert() {
	t := NewTree()
	t.Insert_1(1)
	t.Insert_1(3)
	t.Insert_1(5)
	t.Insert_1(3)
	t.Insert_1(6)

	t.Inorder()
	fmt.Println(t)

}