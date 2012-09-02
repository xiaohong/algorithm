package ds


import "fmt"

// 10.2-1
// 都头尾插入数据是O(1), 删除数据时需要查找数据，而查找数据是线性的O(n)
// 所以删除数据是O(n)

// 10.2-2
// 在链表头上添加和删除就可以实现stack

// 10.2-3
// 使用单链表实现栈需要头指针和尾指针

// 10.2-4
// 消除nil， 使用size进行循环，这个就不需要通过指针比较了，
// 这个地方需要的循环一遍的终止条件，任何等价的终止条件都行

// 10.2-5
// 单链表循环
// 查询O(1), deleteO(1), search O(n)因为

type ringNode struct{
	Value interface{}
	next *ringNode
}

type ring struct{
	header *ringNode
	size int
}

func (this *ring) Insert(v interface{}){
	r := &ringNode{Value: v, next : this.header.next}
	this.header.next = r
}

func (this *ring) Search(v interface{}) bool {
	t := this.header.next
	for t != this.header {
		if t.Value == v {
			return true
		}
	}
	return false
}

func (this *ring) String() string{
	s := "["
	t := this.header.next

	for t != this.header {
		s = fmt.Sprint(s, t.Value, ", ")
		t = t.next
		
	}

	return s +"]"
}

func NewRing() *ring{
	r := &ring{}
	h := &ringNode{}
	h.next = h
	r.header = h
	return r
}

func TestRing(){
	r := NewRing()
	r.Insert(2)
	r.Insert(3)
	r.Insert(4)
	r.Insert(5)

	fmt.Println(r)
	fmt.Println(r.Search(5))

	r.reverse()
	fmt.Println(r)
	fmt.Println(r.Search(5))
}

// 12.2-6
// 使用头尾指针可以在O(1)时间内完成
// list1.tail.next = list2.head 
// list2.head.pre = list1.tail 
// list1.head.pre = list2.tail
// list2.tail.pre = list1.head

// 12.2-7
// 翻转单链表

// 使用2个指针，一个指针维护当前遍历到的对象，另一个指针
// 维护已经反转好的数据
func (this *ring) reverse(){
	t := this.header.next 

	a := this.header
	for t != this.header && t != nil {
		m := t.next
		t.next = a
		a = t 
		t = m
	}
	this.header.next =a
	
}