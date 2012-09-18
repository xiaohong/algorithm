package heap



type MinMaxHeap struct{
	size int
	data []int
}


func (this *MinMaxHeap) Add(v int){
	this.data[this.size] = v
	this.size++

	this.percolateUp(this.size-1)
}

func (this *MinMaxHeap) percolateUp(n int){
	parent := (n -1) /2

	level := 0
	for  i := 1 ; i < n; i = i*2{
		level ++
	}

	if level % 2 == 0{
		if parent >= 0 && this.data[n] > this.data[parent]{
			this.data[n], this.data[parent] = this.data[parent], this.data[n]
			this.percolateMax(parent)
		}else{
			this.percolateMin(n)
		}
	}else{
		if parent >= 0 && this.data[n] < this.data[parent]{
			this.data[n], this.data[parent] = this.data[parent], this.data[n]
			this.percolateMin(parent)
		}else{
			this.percolateMax(n)
		}
	}
}

func (this *MinMaxHeap) percolateMin(n int){
	parent := ( n - 1) /4
	if parent >= 0 && this.data[n] < this.data[parent] {
		this.data[n], this.data[parent] = this.data[parent], this.data[n]
		this.percolateMin(parent)
	}
}

func (this *MinMaxHeap) percolateMax(n int){
	parent := (n - 1)/4
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

func NewMinMaxHeap() *MinMaxHeap{
	return &MinMaxHeap{0, make([]int,100)}
}

