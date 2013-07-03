package dyamicProgramming

import "fmt"

// divede-and-conquer 把原问题分解成相互独立的子问题
// dynamic programming applies when the subproblems overlap
// 1. Characterize the structure of an optimal solution
// 2. Recursively define the value of an optimal solution
// 3. Compute the value of an optimal solution, typically in a bottom-up fashion
// 4. Construct an optimal solution from computed infomation

// 把第0次切当做一个整体
func revenue(data []int, n int) int {
	
	max := 0
	for i := 0; i < n ; i++{
		t := 0
		if i == 0 {
			t = data[n-1]
		}else{
			t = data[i-1] + revenue(data, n - i)	
		}
		if t > max {
			max = t
		}
	}
	return max
}

func Max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

// 从第一刀切，然后第n刀表示没有切
func revenue_1(data []int, n int) int{
	if n == 0 {
		return 0
	}

	max := 0
	for i := 1; i <= n; i++{
		max = Max(max,  data[i-1]+revenue_1(data, n-i))
	}

	return max
}

func memoized_revenue(data []int, n int) int{
	cache := make(map[int]int, n)
	s := make([]int, n+1)
	m := f(data,n, cache,s)
    fmt.Println(s)
	return m
}

func f (data []int, m int, cache map[int]int, s []int) int{
		if v,ok := cache[m]; ok {
			return v
		}
		if m == 0 {
			return 0
		}
		max := 0
		for i := 1; i <= m; i++{
			t := data[i-1] +f(data, m - i, cache, s)
			if max < t {
				max = t 
				s[m] = i
			}
		}
		cache[m] = max
		return max
	}

// 自低向上计算
// 首先设置基本条件0为0
// 然后依赖前面已经计算好的值，从1往上计算
func bottom_up_cut_rod(p []int, n int) (int, []int) {
	r := make([]int, n+1)
	s := make([]int, n+1)
	r[0] = 0

	for i := 1; i <=n ; i++{
		q := 0
		for j := 1; j <=i; j++{
			t := p[j-1]+r[i-j]
			if q < t{
				q = t
				s[i]=j
			}
		}
		r[i] = q
	}
	printSolution(s, n)
	return r[n],s
}

// 每次打印第一刀切的位置，
// 直到切完
func printSolution(s []int, n int){
	for n > 0{
		fmt.Print(s[n],",")
		n = n - s[n]
	}
	fmt.Println("")
}

// f(x) = f(x-1)+f(x-2)
// f(0) = 1, f(1) =1

func fib(n int) int{
	if n == 0 || n == 1{
		return 1
	}
	return fib(n-1)+fib(n-2)
}

func fib_it(n int) int{
	if n == 0 || n == 1{
		return 1
	}
	f1, f2 := 1,1
	for i := 2; i <= n; i++ {
		f := f1 + f2
		f1 = f2
		f2 = f
	}
	return f2
}

func fib_c(n int) int{
	c := make([]int, n +1)
	c[0] = 1
	c[1] = 1
	for i := 2; i <=n ; i++{
		c[i] = c[i-1]+c[i-2]
	}
	return c[n]
}

func Test_revenue(){
	d := []int{1,5,8,9,10,17,17,20,24,30}

	fmt.Println(revenue(d,10))
	fmt.Println(revenue_1(d,10))

	fmt.Println(memoized_revenue(d, 10))
	fmt.Println(bottom_up_cut_rod(d, 10))

	fmt.Println(fib(10))
	fmt.Println(fib_c(10))
	fmt.Println(fib_it(10))

}
