package heap

import "fmt"
import "hong/sort"
import "math/rand"

// 堆数据结构
type Heap struct {
	data       []interface{}
	comparator sort.Comparator
	size       int
}

func (this *Heap) up(a []interface{}, i int) {

	for i > 0 {
		p := i / 2
		if this.comparator(a[p], a[i]) > 0 {
			a[p], a[i] = a[i], a[p]
			i = p
		} else {
			break
		}
	}

}

func (this *Heap) Insert(v interface{}) {
	m := append(this.data, v)
	this.up(m, len(m)-1)
	this.data = m
}

func (this *Heap) Remove() interface{} {

	t := this.data[0]
	this.data[0] = this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	this.down(this.data, 0)
	return t
}

func (this *Heap) down(data []interface{}, i int) {

	for i < len(data) {
		left := 2*i + 1
		right := 2*i + 2

		largest := i
		if left < len(data) && this.comparator(data[left], data[largest]) < 0 {
			largest = left
		}
		if right < len(data) && this.comparator(data[right], data[largest]) < 0 {
			largest = right
		}

		if largest == i {
			break
		}
		data[i], data[largest] = data[largest], data[i]
		i = largest
	}
}

func NewHeap(c sort.Comparator) *Heap {
	return &Heap{make([]interface{}, 0), c, 0}
}

func Int(a, b interface{}) int {
	aa, _ := a.(int)
	bb, _ := b.(int)
	return -(aa - bb)
}

func Test_Heap() {
	a := NewHeap(Int)
	for i := 0; i < 20; i++ {
		a.Insert(rand.Intn(100))
	}
	fmt.Println(a.data)
	for i := 0; i < 10; i++ {
		fmt.Println(a.Remove())
	}

}
