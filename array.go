package main

import "fmt"

func brute_force(a []int) []int{
	i, j, k := 0,0,0;
	sum := 0;
	for m :=0 ; m < len(a); m++{
		sum = 0;
		for n := m; n < len(a); n++{
			sum = sum + a[n]
			if(sum > k){
				k = sum
				i = m
				j = n
			}
		}
	}
	return []int{i,j,k};
}

func brute_force_crossing(a []int) []int{
	n := -100000
	var m []int ;
	for i := 0; i < len(a); i++{
		t := find_max_crossing_subarray(a,0,i,len(a)-1)
		fmt.Println(t)
		if(t[2] > n){
			n = t[2]
			m = t
		}
	}
	return m;
}

func find_max_crossing_subarray(a []int, low int, mid int, hight int) []int{
	left := -10000
	left_index := 0;
	sum := 0;
	for i := mid; i>= low ; i--{
		sum = sum + a[i]
		if(sum > left){
			left = sum
			left_index = i
		}
	}
	
	right := -10000
	right_index := 0
	sum = 0
	for i := mid+1; i <= hight; i++{
		sum = sum+ a[i]
		if sum > right {
			right = sum
			right_index = i
		}
	}
	
	return []int{left_index, right_index, left + right}
}

func max_sub_array(a []int, low int, high int) []int{
	
	if low == high {
		return []int{low,high,a[low]}
	}
	mid := (low + high)/2
	left := max_sub_array(a, low, mid)
	right := max_sub_array(a, mid+1, high)
	middle := find_max_crossing_subarray(a,low,mid,high)
	fmt.Println(low,mid,high,left,right,middle)
	if (middle[2] > left[2] && middle[2] > right[2]) {
		return middle
	}else if( left[2] > right[2]){
		return left
	}else{
		return right
	}
	return []int{}
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
func linear_sub_array(a []int) int{
	maxSoFor := 0
	maxEndingHere := 0
	for i := 0 ; i < len(a); i++{
		maxEndingHere = max(maxEndingHere+a[i], 0)
		maxSoFor = max(maxSoFor, maxEndingHere)
	}
	return maxSoFor
}

func linear(a []int) int{
	current := 0;
	//current_start_index := 0
	start_index := 0
	end_index := 0
	sum := 0
	for i :=0; i < len(a);  i ++{
		value := current + a[i]
		if (value > 0){
			if (current == 0){
				start_index = i	
			}
			current = value
		}else{
			current = 0
		}
		
		if sum < current {
			sum = current
			end_index = i
		}
	}
	fmt.Println(start_index, end_index)
	return sum
}

func main(){
	fmt.Println(brute_force([]int{1,2,-10,4,5}))
	a := []int{13,-3,-25,20,-3,-616,-23,18,20,-7,12,-5,-22,15,-4,7}
	fmt.Println(brute_force(a))
	fmt.Println(max_sub_array(a, 0, len(a)-1))
	
	fmt.Println(brute_force_crossing(a))
	
	fmt.Println(linear_sub_array(a))
	
	fmt.Println(linear(a))

}