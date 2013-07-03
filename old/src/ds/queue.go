package ds


type Queue interface{
	Enqueue(v interface{})
	Dequeue() interface{}
	Size() int
	Peek() interface{}
}

// tail指针指向下一个要进队的位置，默认为0
// head指针指向队头，默认为0，通过size来判断队列头是否有效
// 默认head, tail都为0，tail表示下一个进栈的位置，这样0符合要求
// head表示当前的队列头，默认为0，但是默认是没有数据的，所以要通过size
// 来判断head是否有效
type arrayQueue struct{
	data []interface{}
	head, tail int // putIndex, getIndex
	size int
}

func (this *arrayQueue) Enqueue(v interface{}){
	if this.size == len(this.data){
		//	panic("queue full")
		this.ensureCap(this.size*2)
	}
	this.data[this.tail] = v
	if this.tail == len(this.data) -1 {
		this.tail = 0
	}else{
		this.tail++
	}
	this.size++
}

func (this *arrayQueue) Dequeue() interface{}{
	if this.size == 0{
		panic("empty queue")
	}

	t := this.data[this.head]
	if this.head == len(this.data)-1 {
		this.head =0
	}else{
		this.head++
	}
	this.size--;
	return t
}

func (this *arrayQueue) ensureCap(newSize int){
	newData := make([]interface{}, newSize)

	t := this.head
	for i := 0; i < this.size; i++{
		newData[i] = this.data[t]
		t = (t +1)%this.size
	}
	this.head = 0
	this.tail = this.size+1
	this.data = newData
}

func (this *arrayQueue) Peek() interface{}{
	if this.size == 0{
		panic("empty queue")
	}
	return this.data[this.head]
}

func (this *arrayQueue) Size() int{
	return this.size
}

func NewArrayQueue(size int) Queue{
	q := new(arrayQueue)
	q.data = make([]interface{}, size)
	q.tail = 0
	q.head = 0
	q.size = 0
	return q
}

//
