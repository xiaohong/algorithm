package ds

//import "math/rand"
import "fmt"


// 链表可以使单向的，双向的，循环的，有序的
// 其中的节点存储数据和到其他数据的指针
// 头指针有2种方式，明确的使用front, back指针，或者使用一个哑节点

type Node struct{
	pre,next *Node
	Value interface{}
}

type LinkedList struct{
	header *Node
	size int
}

// 使用2个单独的节点表示头，尾指针
//type LinkedList struct{
	//head, tail *Node
	//size int
//}

func (this *LinkedList) Insert(v interface{}) {
	addFirst(v, this.header)
	this.size ++
}

// 添加到p节点之后
func  addFirst(v interface{}, p *Node){
	n := &Node{Value :v}
	p.next.pre = n 
	n.next = p.next
	n.pre = p
	p.next = n
}

// 添加到p节点之前
func addLast(v interface{}, p *Node){
	n := &Node{Value: v}
	p.pre.next = n
	n.pre = p.pre

	n.next = p
	p.pre = n
}

func (this *LinkedList) RemoveFirst() interface{}{
	if this.header.next == this.header {
		return nil
	}
	t := this.header.next

	this.header.next = this.header.next.next
	this.header.next.pre = this.header
	return t.Value
}

func (this *LinkedList) String() string{
	t := this.header.next
	s := "["
	for t != this.header{
		s =  fmt.Sprint(s, t.Value, ", ")
		t = t.next
	}
	return s + "]"
}

func NewLinkedList() *LinkedList{
	l := new(LinkedList)
	l.header = new(Node)
	l.header.next , l.header.pre = l.header, l.header

	return l
}


func Test_NewLinkedList(){
	l := NewLinkedList()

	for i := 0; i < 10; i++{
		l.Insert(i)	
	}
	
	fmt.Println(l)
}