package main

import "fmt"

// 用要排序的数作为临时数组的索引，然后计算每个数的排序位置
// 用数在临时数组中查找相应的位置，在最终位置的数组上设置相应的值
// 要求所有的数都小于k,
func countingSort(a []int, k int) []int{
	c := make([]int, k)
	for i := 0; i < len(a); i++{
		c[a[i]] = c[a[i]]+1
	}
	
	
	for i := 1; i < k; i++{
		c[i]= c[i]+c[i-1]
	}
	b := make([]int, len(a)+1)
//	for j := len(a)-1; j>=0; j--{
	for j := 0; j < len(a); j++{
		b[c[a[j]]] = a[j]
		c[a[j]] = c[a[j]] - 1
	}

	return b
}

func main(){
	a := []int{1,2,3,6,4,5,4,6,7,8}
	t := countingSort(a,10)
	fmt.Println(t)

	fmt.Println(countingSort([]int{6,0,2,1,3,4,6,1,3,2},7))
}
