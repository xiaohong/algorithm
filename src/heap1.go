package main

import "fmt"

type Heap []int

type Interface interface{
	Insert(v int)
	Pop() int
}

func up(a []int) {
	i := len(a)-1
	p := (i - 1)/2
	for ; i > 0 && a[p] < a [i]; {
		a[p],a[i] = a[i],a[p]
		i = p
	}
}

func (a *Heap) Insert(v int){
	t := []int(a)
	m := append(t,v)
	up(m)
}

func (a *Heap) Pop() int {
	return 1
}

func main(){
	var a Heap = Heap([]int{})
	a.Insert(1)
	a.Insert(2)
	a.Insert(3)
	a.Insert(4)
	a.Insert(5)
	
	fmt.Println(a)
}
