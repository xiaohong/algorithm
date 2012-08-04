package main

import "fmt"

// 3 8 4 2 9 1
// i current index of already sorted array
// j current iterator index

func partition(a []int, p,r int) int{
	x := a[r]
	i := p -1 
	for j := p; j < r; j++{ // // 当前已经排序的索引，这个值之前的数据都是比中间值要大或者要小的所有值
		if a[j] <= x {
			i = i + 1
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1],a[r] = a[r],a[i+1]
	return i+1
}
// 1-2
// 根据计数次数
func partition_(a []int, p, r int) int{
	x := a[r]
	i := p -1 
	n := 0
	for j := p; j < r; j++{ // 
		if a[j] < x {
			i = i + 1
			a[j], a[i] = a[i], a[j]
		}else if a[j] == x && n%2 == 0{
			i = i + 1
			a[j], a[i] = a[i], a[j]
			n = n +1
		}
	}
	a[i+1],a[r] = a[r],a[i+1]
	return i+1
}


func  quick_sort_(a []int, p,r int){
	if (p < r) {
		q := partition_(a, p, r)
		quick_sort_(a, p, q-1)
		quick_sort_(a, q+1, r)
	}
}

func quick_sort(a []int){
	quick_sort_(a,0,len(a)-1)
}
func main(){
	a := []int{3,1,4,6,4,4,33,2,11,4}
	quick_sort(a)
	fmt.Println(a)
	
	fmt.Println(partition_([]int{1,1,1,1,1},0,4))
}