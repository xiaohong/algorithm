package main

import "fmt"

func union(a, b []int) []int{
	for t := range b {
		a = append(a, t)
	}
	return a
}

func activity_selector(s, f []int, k, n int) []int{
	m := k + 1
	for; m < n && s[m]<f[k] ; {
		m = m + 1
	}
	if m < n {
		return union([]int{m}, activity_selector(s,f,m,n))
	}else{
		return []int{}
	}
	return []int{}
}

func exchange(a, b int) int, int{
	a = a+b
	b = a - b
	a = a -b
	return a, b
}

func main(){
	fmt.Println(3,4)
	fmt.Println(exchange(3,4))
	fmt.Println(1)
}