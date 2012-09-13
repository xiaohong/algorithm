package chapter2

import "fmt"

func InsertSort(data []int) []int{
	for i := 0; i < len(data); i++{
		t := data[i]
		j := i-1
		for ; j >= 0; j--{
			if data[j] > t{
				data[j+1]=data[j]
			}else{
				break
			}
		}
		data[j+1]=t
		fmt.Println(data)

	}
	return data
}

func InsertSort_(data []int) []int{
	for i := 0; i < len(data); i++{
		t := data[i]
		j := i - 1
		for j >= 0 && data[j] > t {
			data[j+1]=data[j]
			j--
		}
		data[j+1] = t
	}
	return data
}


func TestInsertSort(){
	fmt.Println(InsertSort([]int{10,1,3,2,54,5,3,2,4}))

	fmt.Println()
	insertSort_chan([]int{10,1,3,2,54,5,3,2,4})

fmt.Println(mergeSort([]int{10,1,3,2,54,5,3,2,4}))	
}


func insertSort_chan(data []int){
	a := make(chan int)

	go func(){
		for _, v := range data {
			a <- v
		}
		close(a)
	}()

	t := make([]int, 0, len(data)+1)
	for {
		n, ok := <- a
		if ok {
			t = t[0:len(t)+1]
			i := len(t)-2
			for ; i >=0; i--{
				if t[i] > n {
					t[i+1] = t[i]
				}else{
					break
				}

			}
			
			t[i+1] = n	
		}else{
			break
		}
		
	}

	fmt.Println(t)

}

// 模拟2个二进制数据相加
// 1. 初始化一个进位标识
// 2. 取最大长度为循环次数变量
// 3. 一次取相应的数据还有进位相加，如果大于等于2，等结果
// 为2的余数设置进位标识为true
// 4. 一直循环完
// 5. 如果进位标识为true,则设置最高位为1，否则为0

type bit int

func max(a, b int) int{
	return a
}

func add(a, b []int) []int{
	l := max(len(a), len(b))

	c := make([]int, l)

	m := false
	for i := 0; i < l; i++{
		t := a[i]+b[i]
		if t == 0{
			if m {
				c[i] = 1
				m = false
			}else{
				c[i] = 0
			}
		}else if t == 1{
			if m {
				c[i] = 0
				m = true
			}else{
				c[i] = 1
			}
		}else if t == 2{
			if m {
				c[i] = 1
				m = true
			}else{
				c[i] = 0
				m = true
			}
		}
	}
	if m {
		c[len(c)-1]=1
	}

	return c
}

// 1. 初始化时已排序的为空
// 2. 每次迭代时选取最小的元素追加
func select_sort(data []int) []int{

	for i :=0; i < len(data); i++{
		min := i
		j := i+1
		for ; j < len(data); j++{
			if data[j] < data[min] {
				min = j
			}
		}
		data[i],data[min] = data[min],data[j]
	}
	return data
}


func merge(a, b []int) []int{
	t := make([]int, len(a)+len(b))

	k,j := 0,0
	for i := 0; i < len(a)+len(b); i++{
		if k < len(a) && j < len(b){
			if a[k] < b[j]{
				t[i] = a[k]
				k++
			}else{
				t[i] = b[j]
				j++
			}
		}else if k < len(a){
			t = append(t, a[k:]...)
		}else if j < len(b){
			t = append(t, b[j:]...)
		}
	}
	return t
}

func mergeSort(data []int) []int{
	if len(data) == 1{
		return data
	}
	i := len(data)/2
	a := mergeSort(data[0:i])
	b := mergeSort(data[i:])
	return merge(a,b)
}

func find(data []int, a int) []int{
	m := len(data)
	for i := 0; i < m ; i++{
		n := a - data[i]
		for k := m-1; k > i ; k--{
			if data[k] > n {
				m--
				continue
			}
			if data[k] == n {
				fmt.Println(data[i],data[k])
				break;
			}
			break;
		}
	} 
	return nil
}

func Test_find(){
	find([]int{1,2,5,8,33,78},9)
}

func hanoi(n int, src, dst, tmp []int){
	if n <= 1 {
		dst[0] = src[0]
		return
	}
	hanoi(n-1, src, tmp, dst)
	dst[n-1] = src[n-1]
	fmt.Println(dst[n-1] , "->", src[n-1])
	hanoi(n-1, tmp, dst, src)
}

func Test_hanoi(){
	a := []int{1,2,3,4,5,6}
	c := make([]int, 6)
	hanoi(6,a, c,make([]int, 6))
	fmt.Println(c)
}