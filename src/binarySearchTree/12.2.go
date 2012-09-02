package binarySearchTree

func (this *Tree) Search(k interface{}) *Node {
	if this.Root == nil || this.Root.Value == k {
		return this.Root
	}

	c := this.Compare(this.Root.Value, k)
	if c > 0 {
		left := this.leftSubTree()
		if left != nil {
			return left.Search(k)	
		}
		
	}else {
		right := this.rightSubTree()
		if right != nil {
			return right.Search(k)
		}
	}
	return nil
}

func (this *Tree) Search_iterative(k interface{}) *Node{
	t := this.Root 

	for ; t != nil && t.Value != k; {
		if this.Compare(t.Value, k) > 0{
			t = t.left
		}else {
			t = t.right
		}
	}
	return t
}

func (this *Tree) Minimum() *Node{
	t := this.Root 
	if t == nil {
		return nil
	}

	for ; t.left != nil ;{
		t = t.left
	}
	return t
}

func (this *Tree) Maximum() *Node{
	if this.Root == nil {
		return nil
	}

	t := this.Root
	for ; t.right != nil; {
		t = t.right
	}

	return t
}

func (this *Node) Successor() *Node{
	if this.right != nil {
		return (&Tree{Root: this.right}).Minimum()
	}

	t := this
	p := t.parent
	for ; p != nil && p.right == t ; {
		t = p
		p = p.parent
	}
	return p
}

func (this *Node) Predecessor() *Node{
	if this.left != nil {
	  return (&Tree{Root : this.left}).Maximum()
	}

	t := this
	p := t.parent

	for ; p != nil && p.left == t;{
		t = p 
		p = p.parent
	}

	return p
}