package ds


import "fmt"

// 10.1_1
// 通过size来表示下一个进栈的位置，所以默认为0，第一次进栈的位置在数组的
// 第一个位置，进展后增加size的大小
// [4] size = 1
// [4,1] 
// [4,1,3] 
// [4,1] 
// [4,1,8] 
// [4,1] 


// 10.2
// 一个栈从0开始向上增加，另外一个栈从n开始往下减，当两个指针相同时表示还有一个空
// 位置，但从n开始的指针等于从0开始的指针-1时，说明栈满，

type TwoStack struct{
	data []interface{}
	head, tail int
}

func (this *TwoStack) full() bool{
	if this.tail <= this.head - 1{
		return true
	}
	return false
}

func (this *TwoStack) PushHead(v interface{}){
	if this.full(){
		panic("stack full")
	}
	this.data[this.head] = v
	this.head++
}

func (this *TwoStack) PushTail(v interface{}){

	if this.full(){
		panic("stack full")
	}
	this.data[this.tail] = v
	this.tail--
}

func NewTwoStack(size int) *TwoStack{
	t :=new(TwoStack)
	t.data = make([]interface{}, size)
	t.head = 0
	t.tail = size-1
	return t
}


// 10.1-3
func (this *arrayQueue) Trace() {
	s := "["
	t := this.head
	for i := 0; i < this.Size(); i++{
		s = s + fmt.Sprint(this.data[t])+" ,"
		t = (t +1)%len(this.data)
	}
	s = s + "] head=" + fmt.Sprint(this.head) +", tail="+ fmt.Sprint(this.tail)
	fmt.Println(s)
}

func Test_10_1_3(){
	t ,_:= (NewArrayQueue(6)).(*arrayQueue)
	t.Enqueue(4)
	t.Trace()
	t.Enqueue(1)
	t.Trace()
	t.Enqueue(3)
	t.Trace()
	t.Dequeue()
	t.Trace()
	t.Enqueue(8)
	t.Trace()
	t.Dequeue()
	t.Trace()
}

//10.1-4
// 进队列时要检查栈是否满，主要通过size是否已经达到数组的容量
// 出队列时检查是否为空

// 10.1-5

// 头，尾都从0开始，头指针表示下一个要插入的位置，尾指针表示当前已经插入的位置，
// 这样当头、尾指针相同时说明满了// 再加入元素后检查是否已经满了，然后扩容，
// 而不是在要添加的时候再去检查容量
type Deque struct{
	data []interface{}
	head, tail int
}

func (this *Deque) PushHead(v interface{}){
	this.data[this.head] = v
	this.head = (this.head+1)&(len(this.data)-1)

	if this.head == this.tail{
		this.ensureCap()
	}
}

func (this *Deque) PushTail(v interface{}){
	this.tail = (this.tail -1) & (len(this.data)-1)
	// 因为head指针表示即将要插入的位置，所以不会冲突
	this.data[this.tail] = v
	if this.head == this.tail {
		this.ensureCap()
	}
}

// 通过设置不存在的数据为null,然后假定数据存在取数据，
// 如果数据不为空，则走正常流程，如果为空，说明队列空了
// 然后直接返回null,这样就隐含的假定不能存储null数据(java的做法)
// 或者可以维护一个size变量，这样就可以判断是否为空，如果
// 在空的情况下依然pop的话，则抛出异常
func (this *Deque) PopHead() interface{}{
	t := this.head - 1
	if this.data[t] == nil {
		return nil
	}
	n := this.data[t]
	this.data[t] = nil
	this.head = t
	return n
}

func (this *Deque) PopTail() interface{}{
	t := this.data[this.tail]
	if t == nil {
		return nil
	}
	this.data[this.tail] = nil
	this.tail = (this.tail + 1)%len(this.data)
	return t
}

// 主要是复制已存在的数据，然后重新设置tail,head指针位置
func (this *Deque) ensureCap(){
	newData := make([]interface{}, 2*len(this.data))

	tailLen := len(this.data)-this.tail
	copy(newData[0:tailLen], this.data[this.tail:])
	copy(newData[tailLen:], this.data[0:this.head])
	this.data = newData
	this.head = len(this.data)/2
	this.tail = 0
}

func NewDeque(size int) *Deque{
	t := new(Deque)
	t.data = make([]interface{}, size)
	return t
}

// 双端队列主要是同时要维护2个指针，一个指向下个元素的位置，一个指向当前元素的位置

// 10.1-6
// 用2个栈实现队列

type twoStackQueue struct{
	left, right Stack
}

func (this *twoStackQueue) Enqueue(v interface{}){
	this.left.Push(v)
}

func (this *twoStackQueue) Dequeue() interface{}{
	for this.left.Size() > 0 {
		this.right.Push(this.left.Pop())
	}
	t := this.right.Pop()
	for this.right.Size() > 0{
		this.left.Push(this.right.Pop())
	}
	return t
}

func (this *twoStackQueue) Size() int{
	return this.left.Size()
}

func (this *twoStackQueue) Peek() interface{}{
	for this.left.Size() > 0 {
		this.right.Push(this.left.Pop())
	}
	t := this.right.Peek()
	for this.right.Size() > 0{
		this.left.Push(this.right.Pop())
	}	
	return t
}

func NewTwoStackQueue() Queue{
	t := &twoStackQueue{NewArrayStack(), NewArrayStack()}
	return t
}