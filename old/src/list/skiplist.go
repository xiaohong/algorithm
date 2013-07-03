package list

import "hong/sort"
import "math/rand"
import "fmt"

// 跳跃表的实现
// 查找平衡的链表数据结果，类似于平衡二叉树，但是不需要在插入时做旋转操作
// 基本核心思想就是多几层链表结构，每个节点有多个指针，不同的指针指向不同的下一个
// 节点，每个节点有一级别level,可以存储Level+1个指针，第0个指针总是指向链表中的下一个元素

const (
	maxLevel int = 7
)

// 链表节点
type node struct{
	x interface{}
	next []*node	
}

func NewNode(x interface{}, h int) *node{
	return &node{x, make([]*node, h+1)}
}

func (this *node) height() int{
	return  len(this.next) -1 
}

// 
type SkipList struct{
	header *node
	level int
	cmp sort.Comparator
}

// 查找相应的节点
// 从最高级别开始，如果当前的节点小于要搜索的几点，然后用指向的节点替换当前节点
// 继续判断当前level上的下一个节点，如果大于，则下降level，知道最后一个级别
// 这样当前节点是小于key的最大节点，那么这个几点的下一个节点就是等于或大于key
// 的第一个节点
func (this *SkipList) Search(key interface{}) bool{
	x := this.header

	for i := this.level; i >=0; i--{
		for this.cmp(x.next[i].x, key) < 0{
			x = x.next[i]
		}
	}
	x = x.next[0]
	if this.cmp(x, key) == 0{
		return true
	}
	return false
}

// 1. 寻找要插入的位置，并且维护搜索路径
// 2. 生成新节点和新的随机level
// 3. 如果新的level比当前level还大，那么添加高Level的搜索路径
// 4. 插入新的节点，并更新相应路径的
func (this *SkipList) Insert(v interface{}){
	x := this.header

	update := make([]*node, this.level+1)
	for i := this.level; i >= 0; i--{
		for  x.next[i] != nil && this.cmp(x.next[i].x, v) < 0 {
			x = x.next[i]
		}
		update[i] = x
	}
	// 指向小于v的最大节点的下一个元素
	x = x.next[0]

	newLevel := randromLevel()
	if newLevel > this.level {
		for i := this.level; i <= newLevel; i++{
			update = append(update, this.header)
		}
		this.level = newLevel
	}

	y :=  NewNode(v, newLevel)
	for i := 0; i <= newLevel; i++{
		y.next[i] = update[i].next[i]
		update[i].next[i] = y
	}
}

func printUpdate(x interface{}, u []*node){
	fmt.Print(x, " << ")
	for _, v := range u{
		fmt.Print(v.x, " ")
	}
	fmt.Println()
}

func (this *SkipList) print() {
	for i := this.level; i >0; i--{
		x := this.header.next[i-1]
		for x != nil {
			fmt.Print(x.x, "--> ")
			x = x.next[i-1]
		}
		fmt.Println()
	}
}

func (this *SkipList) String() string{
	x := this.header.next[0]
	var s string = ""
	for ; x != nil; x = x.next[0]{
		if s == "" {
			s = fmt.Sprint(x.x)
		}else{
			s = fmt.Sprint(s, ", ", x.x)	
		}
		
	}
	return s
}

func randromLevel() int{
	level := 1
	for rand.Intn(2) < 1{
		level++
	}
	if level > 7{
		return 7
	}
	return level
}

func (this *SkipList) Delete(v interface{}) {
	x := this.header
	update := make([]*node, maxLevel+1)

	for i := this.level -1; i >=0 ; i--{
		for x.next[i] != nil && this.cmp(x.next[i].x, v) < 0{
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]

	// 删除指定的元素
	if x != nil && this.cmp(x.x, v) == 0{
		for i := 0; i < this.level; i++{
			// 更新那些指向x的前面的指针
			if update[i].next[i] != x{
				break
			}
			update[i].next[i] = x.next[i]
		}

		// 如果最高级别的指针被更新到nil了，这时要下降level
		for this.level > 0 && this.header.next[this.level-1] == nil{
			this.level--
		}
	}
}

func NewSkipList(cmp sort.Comparator) *SkipList{
	return &SkipList{header: NewNode(nil, maxLevel), level:0, cmp:cmp}
}