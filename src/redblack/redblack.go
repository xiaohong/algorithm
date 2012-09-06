package redblack

import "ds"
import "fmt"
import "math/rand"

// 红黑树确保没有任何一条path是其他path的2倍长
// 1. 每个节点要么红，要么黑
// 2. 根节点是黑
// 3. 叶子节点(nil)是黑
// 4. 如果一个节点是红的，那么两个孩子节点都是黑的
// 5. For each node, all simple paths from the node to descendant
// leaves contain the same number of black nodes.

// a red-black tree with n internal nodes has height at most 2log(n+1)

// 因为从根几点到所有叶子节点的黑高度都相同，而每个红节点的孩子都必须是黑节点

type rbNode struct {
	value       interface{}
	left, right *rbNode
	parent      *rbNode
	color       int
}

const (
	red   int = 0
	black int = 1
)

// 保持当前节点与即将要替换的节点引用
// 1. 右子节点的左子树作为当前节点的右子树
// 2. 右子节点替换当前节点
// 3. 当前节点当做右子节点的左节点
func leftRotate(t *RedBlackTree, x *rbNode) {
	// 替换右孩子的左节点到旋转节点的右节点
	y := x.right       // 保存要旋转节点的右孩子

	x.right = y.left   // 右孩子的左节点放到旋转节点的右子节点
	if y.left != Nil { // 更新右孩子的左节点的父指针
		y.left.parent = x
	}

	// 右节点替换旋转节点
	y.parent = x.parent
	if x.parent == Nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	// 旋转节点作为右节点的左节点
	y.left = x
	x.parent = y
}

// 
func rightRotate(this *RedBlackTree, y *rbNode) {
	x := y.left // x替换y的位置

	// 
	y.left = x.right
	if y.left != Nil {
		y.left.parent = y
	}

	// 用x替换y节点
	x.parent = y.parent
	if y.parent == Nil {
		this.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	// 把y作为x的右节点
	x.right = y
	y.parent = x
}

func compare(a, b interface{}) int {
	a1, aok := a.(int)
	b1, bok := b.(int)
	if aok && bok {
		return a1 - b1
	}
	return 1
}

var Nil *rbNode = &rbNode{color: black}

// 首先查找到要插入的位置
// 
func (this *RedBlackTree) Insert(v interface{}) {
	n := this.search(v)
	if n != Nil && n != nil {
		return
	}

	z := &rbNode{value: v, color: red, left: Nil, right: Nil}
	var y *rbNode = Nil
	x := this.root

	for x != Nil && x != nil {
		y = x
		c := compare(x.value, v)
		if  c > 0 {
			x = x.left
		} else if c < 0 {
			x = x.right
		}else{
			return 
		}
	}

	if y == Nil {
		this.root = z
	} else if compare(y.value, v) > 0 {
		y.left = z
	} else {
		y.right = z
	}
	z.parent = y
	this.size++
	this.insertFixUp(z)
}

// 插入的节点默认是红色，所有可能出现新插入的节点和父节点都是红色的
// 1. 如果伯父节点是红色，那把父节点和伯父节点都设置为黑色，把父节点的父节点设置
// 为红色，当前节点指向这个节点。
// 2. 如果伯父节点是黑色，分2中情况
//    根据当前节点是父节点的左节点还是右节点分两种情况，通过旋转把当前节点，父节点，
//    父父节点转化成一条线的树，然后改变当前节点，父父节点为黑色，父节点有红色然后旋转
//   父节点变成父父节点和当前节点的父节点
func (this *RedBlackTree) insertFixUp(z *rbNode) {
	// 如果不适用Nil节点，那么这个地方需要检测z != this.root && z != nil
	for z.parent.color == red && z != this.root && z!= Nil {
		
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == red { // 伯父节点为红色
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else { // 伯父节点是黑节点
				if z == z.parent.right { // 当前节点是右子树，然后左旋转，当前节点
					z = z.parent // 变成父节点
					leftRotate(this, z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				rightRotate(this, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rightRotate(this, z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				leftRotate(this, z.parent.parent)
			}
		}
	}
	this.root.color = black
}

type RedBlackTree struct {
	root *rbNode
	size int
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (this *RedBlackTree) String() string {
	s := ds.NewArrayStack()

	v := this.root
	for v != nil && v != Nil {
		s.Push(v)
		v = v.left
	}

	str := "["
	for s.Size() > 0 {
		pre := s.Pop()
		a, _ := pre.(*rbNode)
		str = fmt.Sprint(str, a.value, ",")

		rNode := a.right
		for rNode != nil && rNode != Nil {
			s.Push(rNode)
			rNode = rNode.left
		}
	}
	return str
}

// v节点替换u节点, u的另外一个节点的信息并没有处理
func transplant(this *RedBlackTree, u, v *rbNode) {
	if u.parent == Nil {
		this.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// 返回后续者
func minimum(this *rbNode) *rbNode {
	t := this
	for t.left != Nil {
		t = t.left
	}
	return t
}

func (this *RedBlackTree) search(v interface{}) *rbNode {
	r := this.root
	for r != nil && r != Nil {
		n := compare(r.value, v)
		if n > 0 {
			r = r.left
		} else if n < 0 {
			r = r.right
		} else {
			return r
		}
	}
	return r
}

func (this *RedBlackTree) Delete(v interface{}) {

	z := this.search(v)
	if z == nil || z == Nil {
		return
	}

	y := z // 被删除或移动的节点
	yColor := y.color
	var x *rbNode = nil // 新插入的节点

	if z.left == Nil {
		// 左节点为空，用右节点填充
		x = z.right
		transplant(this, z, z.right)
	} else if z.right == Nil {
		x = z.left
		transplant(this, z, z.left)
	} else {
		// 有2个子节点，去后继者
		y = minimum(z.right)
		yColor = y.color
		x = y.right
		if y.parent == z { // 右孩子只有一个节点
			x.parent = y
		} else { // 用y的右子树替换y
	
			transplant(this, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		// 用y替换z
		transplant(this, z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yColor == black {
		deleteFixUp(this, x)
	}
}

// 移动额外的黑节点到
// 1. x 指向一个红节点，然后直接设置为黑色节点
// 2. x 指向根节点，然后移除额外的黑
// 3. 持续重着色和旋转，达到条件1、2

func deleteFixUp(this *RedBlackTree, x *rbNode) {
	for x != this.root && x.color == black { // x指向双黑非根节点
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				leftRotate(this, x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				// x, w都去掉一个黑，然后x指向x.parent,使x.parent多一个黑节点
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					// w的左孩子是红节点
					w.color = red
					w.left.color = black
					rightRotate(this, w.left)
					w = x.parent.right
				}
				// x 为黑色，w为x的颜色，w右孩子变成黑色，然后左旋转x.parent
				// 就是左边少一个黑节点，右边少一个黑节点，此时不违反相应的数据
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				leftRotate(this, x.parent)
				x = this.root
			}

		} else {
			w := x.parent.left

			if w.color == red {
				w.color = black
				x.parent.color = red
				rightRotate(this, x.parent)
				w = x.parent.left
			}

			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == red {
					w.color = red
					w.right.color = black
					leftRotate(this, w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				leftRotate(this, x.parent)
				w.right.color = black
				x = this.root
			}
		}

	}
}

func (this *RedBlackTree) debug(){
	fmt.Println(this.root.value, "[]", this.root.color)
	fmt.Println(this.root.left.value, "[]", this.root.left.color)
	fmt.Println(this.root.right.value, "[]", this.root.right.color)
	
}

// 如果被删除或移动的节点是黑色节点，那么
// 1. 如果y是根节点，x是红色子节点，那么这是根节点变成红色
// 2. 如果x和x.parent都是红节点，违反红节点的子节点必须是黑色节点
// 3. 移动黑节点y促使原来包含y的任何path的黑高度少1，

// 解决问题3的办法是假装x节点有一个额外的黑色，用来表示被移动的黑色节点，
// 这是节点x可能是2个黑或者一红一黑

// 为什么不选择新插入的节点为黑色呢？这样将就不会违反属性4，而是违反属性5
// 因为属性5不好操作，属性4我们可以递归的处理当前节点的父树

func Test_RedBlackTree() {
	r := NewRedBlackTree()
	r.Insert(41)
	fmt.Println(r)

	r.Insert(388)
	fmt.Println(r)

	r.Insert(31)
	fmt.Println(r)

	r.Insert(12)
	fmt.Println(r)

	r.Insert(19)
	r.Insert(8)

	fmt.Println(r)

	for i := 0; i < 100; i++{
		r.Insert(rand.Int()%1000)
	}

	fmt.Println(r)	
	fmt.Println("----------")
	r.Delete(388)
	fmt.Println(r)	
}
