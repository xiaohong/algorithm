package heap

/*

import "hong/sort"
// A fibonacci heap is a collection of rooted trees that are min-heap ordered
type fibNode struct{
	degree int
	mark bool

	prev,next *fibNode

	parent *fibNode
	child *fibNode

	key int
	value interface{}

	comp sort.Comparator
}

type FibHeap struct{
	min *fibNode
	size int
}

func NewFibHeap() *FibHeap{
	return &FibHeap{}
}

func (this *FibHeap) Insert(v interface{}, key int){
	x := &fibNode{}
	x.degree = 0
	x.key = key

	if this.min == nil {
		this.min = x
		this.min.pre = x
		this.min.next = x
		return
	}

	// insert x into the root list
	x.pre = this.min.pre 
	x.next = this.min 
	this.min.pre.next = x
	this.min.pre = x

	if x.key < this.min.key{
		this.min = x
	}
	this.size++
}

func (this *FibHeap) Minimim() interface{}{
	if this.min == nil {
		return nil
	}

	return this.min.value
}

func (this *FibHeap) Union(other *FibHeap){
	// 取2个之中最小的元素的作为min
	// concat the root list
}
*/