package augmentDs

// 如果扩展数据基础数据结构

const{
	red = 1
	black = 2
}

type Interval struct{
	low, hight int
}

type node struct {
	value Interval
	max int

	left, right, parent *node
	color int
}

type IntervalTree struct{
	root *node
	size int
}

func (this Interval) overlap(other Interval) bool {
	return false
}

func (this *IntervalTree) Search(i Interval) Interval {
	x := this.root
	for x != nil && !i.overlap(x.value){
		if x.left != nil and x.left.max >= i.low{
			x = x.left
		}else{
			x = x.right
		}
	}
	return x
}

func (this *IntervalTree) Insert(v Interval) {
	n := &node{value: v, max : v.hight, color:red}
	
	var pre *node = nil
	t := this.root
	var c int
	for t != nil {
		pre = t
		c = t.value.low > v.low
		if c > 0 {
			t = t.left
		}else{
			t = t.right
		}
	}

	if pre == nil {
		n.color = black
		this.root = n 
		return
	}

	if c > 0 {
		pre.left = n
	}else{
		pre.right = n
	}
	n.parent = pre
	insertFixup(this, n)
}

func leftRotate(this *IntervalTree, x *node){
	y := x.right

	// y的左子树替换y的右子树
	x.right = y.left 
	if y.left != nil {
		y.left.parent = x
	}

	// y替换x
	y.parent = x.parent
	if x.parent == nil {
		this.root = y
	}else if x == x.parent.left {
		y.parent.left = y
	}else {
		y.parent.right = y
	}

	// x替换y的左子树
	y.left = x
	x.parent = y

	updateMax(x)
	updateMax(y)
}

func updateMax(x *node){
	x.max = x.value.hight
	if x.left != nil && x.left.max > x.max {
		x.max = x.left.max
	}
	if x.right != nil && x.right.max > x.max {
		x.max = x.right.max
	}
}

func rightRotate(this *IntervalTree, x *node){
	y := x.left 

	x.left = y.right 
	if y.right != nil {
		x.left.parent = x
	}

	// 替换y用x
	y.parent = x.parent
	if x.parent == nil {
		this.root = y
	}else x == x.parent.left {
		y.parent.left = y
	}else{
		y.parent.right = y
	}

	// x替换y的右子树
	y.right = x
	x.parent = y 

	updateMax(x)
	updateMax(y)
}

func insertFixUp(this *IntervalTree, x *node){
	// 当前节点和父节点都是红节点
	for x != nil && x != this.root && x.parent.color == red {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y.color == red { // 伯父节点是红节点，那么爷爷节点肯定是黑节点，这时把父节点
			   x.parent.color = black                 // 和伯父节点都设置为黑节点，爷爷节点设置为红色节点，x指向a
			   y.color = black
			   x.parent.parent = red
			   x = x.parent.parent
			}else{ // 伯父节点是黑色的,使2个红节点都在左子树上，为了
				if x == x.parent.right {
					x = x.parent
					leftRotate(this, x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				rightRotate(this, x.parent.parent)
			}
		}else {
			y := x.parent.parent.left
			if y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			}else{
				if x == x.parent.left{
					x = x.parent
					rightRotate(this, x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				leftRotate(this, x.parent.parent)
			}
		}
	}
	this.root.color = black
}

func tranplant(this *IntervalTree, x y *node){
	if x.parent == nil {
		this.root = y
	}else if x == x.parent.left{
		x.parent.left = y
	}else {
		x.parent.right = y
	}

	if y != nil {
		y.parent= x.parent
	}
}

func (this *IntervalTree) Delete(x *node){
	// 保存被删除或者移动的节点
	y := x 
	yColor := x.color
	if x.left == nil {
		tranplant(this, x, x.right)
	}else if x.right == nil {
		tranplant(this, x, x.right)
	}
}