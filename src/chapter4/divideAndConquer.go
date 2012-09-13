package chapter4

import "fmt"


// 分治法的基本思路就是找到一个切入点
// 然后把原问题分解成多个子问题，然后再递归的解决子问题
// 所有的子问题最终都会演绎到基本情况，而基本情况是直接
// 给出答案的。
// 分解到子问题时，通常还会伴随的其他非原问题的问题，这
// 些是另外的问题，通常可以直接解决的，例如最大子数组
// 中的包括中间元素的最大子数组的解法
// 计算子问题后，需要对子问题的解进行合并得出原问题的解
// 分治法分解的子问题通常是不会重复的，这个动态规划方法
// 最不同的地方。
func maxSubArray(data []int) []int{
	m, n := 0,0
	max := 0
	for i := 0; i < len(data); i++{
		t := 0
		for j := i;j < len(data); j++{
			t = t + data[j]
				if max < t {
					fmt.Println(m,n,max)
					m = i 
					n = j
					max = t
				}		
		}
	}
	return []int{m,n, max}
}

func Test_a(){
	fmt.Println(maxSubArray([]int{11,2,3,4,-800,4,5,6}))
}

// 最大子数组问题的分治法的思想是根据中间元素把数组分为2部分，
// 最大子数组要么在左边，要么在右边，要么包括中间元素
// 然后递归的计算出左右子数组的最大子数组，还要包括中间元素
// 的最大子数组，然后3者进行比较
// 只有一个元素的数组的最大子数组就是本身


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

// 如果累加的和小于0后，则需要扔掉
// 然后重新开始计算
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

// 取中间的数据，然后把数据划分成2部分，
// 小于中间数的，大于中间数的，等于中间数的
// 然后比较k与小于中间数的size,如果小，则在
// 等于中间数的找，如果还小，则大大于中间数的
func Knth(data []int) int{
	return 1
}

func nqueuens(data []int, r int){
	if r == len(data){
		print(data)
		return
	}

	for j := 0; j < len(data); j++{
		//legal := true
		for i := 0; i < r; i++{
			//
		}
	}

}