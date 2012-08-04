package main

import "fmt"

func select_sort(a []int) []int{	
	for i := 1; i < len(a); i++{
		j := i
		for k := i; k < len(a); k++{
			if a[j] > a[k]{
				j = k
			}
		}
		a[i],a[j] = a[j],a[i]
	}
	return a;
}

func merge(a []int, p int, q int, r int) []int{
	n1 := q - p + 1
	n2 := r -q
	l := make([]int, n1)
	right := make([]int, n2)
	for i :=0; i< n1; i++{
		l[i] = a[p+i]
	}
	for i := 0; i < n2; i++{
		right[i] = a[q+i+1]
	}
	fmt.Println(l, right)
	i := 0
	j := 0
	for k := p ; k <= r; k++{
		fmt.Println(i,j,p,k)
		if  i < n1 && j < n2 {
			if( l[i] < right[j]){
			a[k] = l[i]
			i = i +1
			}else{
			a[k] = right[j]
			j = j + 1
			}
		}else if i < n1{
			a[k] = l[i]
			i = i +1
		}else{
			
			a[k] = right[j]
			j = j + 1
		}
	}
//	fmt.Println(a, p, q,r)
	return a
}

func merge_sort_(a []int, p int , q int) []int{
	if p < q {
		fmt.Println(p,q)
		l := (p+q)/2
		merge_sort_(a, p, l)
		merge_sort_(a, l+1 , q)
		return merge(a,p,l,q)
	}
	fmt.Println(p,q)
	return a
}

func merge_sort(a []int) []int{
	return merge_sort_(a,0,len(a)-1)
}


func main(){
	fmt.Println(merge_sort([]int{1,22,3,44,5,44,5,34}))
	fmt.Println(1);
}