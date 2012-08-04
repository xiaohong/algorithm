package main

import "fmt"

func max_heapify(a []int, i int) []int{
	l := 2*i + 1
	r := 2 * i + 2
	largest := i
	if l < len(a) && a[l] > a[i] {
		largest = l
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	
	if largest != i {
		a[largest], a[i] = a[i],a[largest]
		max_heapify(a, largest)
	}
	return a
}

func max_heapify_iterator(a []int, i int) []int{
	for k := i; k < len(a); {
		l := 2*i + 1
		r := 2 * i + 2
		largest := i
		if l < len(a) && a[l] > a[i] {
			largest = l
		}
		if r < len(a) && a[r] > a[largest] {
			largest = r
		}
		
		if largest != i {
			a[largest], a[i] = a[i],a[largest]
			k = largest
		}else{
			break
		}
	}
	return a
}

func max_heapify_all(a []int){
	for i := len(a)-1; i >=0 ; i--{
		max_heapify(a,i)
	}
}

func heap_sort(a []int) []int{
	var t []int;
	for i := len(a); i >0; i--{
		t = a[0:i]
		max_heapify_all(t)
		a[0],a[len(t)-1] = a[len(t)-1],a[0]
	}
	return a
}

func main(){

	a := []int{1,11,2,22,3,4,5}
	
	for i := len(a)-1; i >=0 ; i--{
		max_heapify(a,i)
		fmt.Println(a)
	}
	
	fmt.Println(heap_sort(a))
}