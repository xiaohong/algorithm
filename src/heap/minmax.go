package heap

import "fmt"

type MinMaxHeap struct{
	size int
	data []int
}


func (this *MinMaxHeap) Add(v int){
	if this.size < len(this.data){
		this.data[this.size] = v	
	}else{
		this.data = append(this.data, v) 
	}
	
	this.size++

	this.percolateUp(this.size-1)
}

func (this *MinMaxHeap) percolateUp(n int){
	parent := (n -1) /2

	level := 0
	for  i := 2 ; i <= n+1; i =  i*2{
		level ++
	}
	

	if level % 2 == 0{
		// 如果是偶数行，说明是最小层，父节点是大于这个节点的节点
		// 如果当前节点比父节点大，那么应该交换，然后逐层上升最大节点
		// 如果比父节点小，那么应该在最小堆中的父节点中进行上升
		if parent >= 0 && this.data[n] > this.data[parent]{
			this.data[n], this.data[parent] = this.data[parent], this.data[n]
			this.percolateMax(parent)

		}else{
			this.percolateMin(n)

		}
	}else{
		// 如果是偶数行，那么是最大堆所在行，首先判断新插入的元素是不是比最小堆中的父节点
		// 要小，如果小的话，交换位置，在最小堆中进行上升
		// 如果大，
		if parent >= 0 && this.data[n] < this.data[parent]{
			this.data[n], this.data[parent] = this.data[parent], this.data[n]
			this.percolateMin(parent)
		}else{
			this.percolateMax(n)
		}
	}
}

func (this *MinMaxHeap) percolateMin(n int){

	parent := ( n + 1) /4-1

	
	if parent >= 0 && this.data[n] < this.data[parent] {
		this.data[n], this.data[parent] = this.data[parent], this.data[n]
		this.percolateMin(parent)

		fmt.Println(this.data)
	}
}

func (this *MinMaxHeap) percolateMax(n int){
	parent := (n + 1)/4-1
	if parent >= 0 && this.data[n] > this.data[parent] {
		this.data[n], this.data[parent] = this.data[parent], this.data[n]
		this.percolateMax(parent)
	}
}

func (this *MinMaxHeap) Max() int{
	if this.size <= 0{
		return 0
	}
	if this.size == 1 {
		return this.data[0]
	}

	if this.size == 2{
		return this.data[1]
	}


	if this.data[1] > this.data[2]{
		return this.data[1]
	}else{
		return this.data[2]
	}
	return 0
}

func (this *MinMaxHeap) Min() int{
	if this.size == 0{
		return 0
	}
	return this.data[0]
}

func (this *MinMaxHeap) RemoveMax() int{
	maxIndex := 0
	if this.size == 0{
		maxIndex = 0
	}else if this.size == 1{
		maxIndex = 1
	}else {
		if this.data[1] > this.data[2]{
			maxIndex = 1
		}else{
			maxIndex = 2
		}
	}
	t := this.data[maxIndex]
	this.size--
	this.data[maxIndex] = this.data[this.size]
	
	this.down(maxIndex)

	return t
}

func (this *MinMaxHeap) RemoveMin() int{
	t := this.data[0]
	
	this.size--
	this.data[0] = this.data[this.size]


	this.down(0)
	return t
}

// 把当前元素下降
// 如果当前元素在最小层，那么当前元素可能比子节点要大，和最小的子节点进行交换
// 如果当前元素在最大层，那么当前节点可能小，然后选取最大子节点和其交换
func (this *MinMaxHeap) down(n int){
	level := 0

	// 因为n是从0开始计数的，所以要加1, n >= 2^level and n < 2^(level+1)
	for i := 2; i <= n+1; i++{
		level ++ 
	}

	if level %2 == 0{
		this.downMin(n)
	}else{
		this.downMax(n)
	}
}

func (this *MinMaxHeap) downMin(n int){

}

func (this *MinMaxHeap) downMax(n int){
	
}
func NewMinMaxHeap() *MinMaxHeap{
	return &MinMaxHeap{0, make([]int,1)}
}

