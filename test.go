package main

//import c "chapter2"
import "fmt"
import d "chapter4"
import "hong/sort"
import "heap"
import "graph"
import "list"

func main() {
	//	c.TestInsertSort()

	//	c.Test_find()

	//c.Test_hanoi()

	d.Test_a()

	print(sort.Int()(3, 2))

	heap.Test_Heap()

	graph.Test_Reverse()

	heap.Test_MinMaxHeap()
//	heap.Test_a()
//tt()


list.Test_Insert()
}


func tt() {

	t := func(n int) int{
		level := 0
		for  i := 2 ; i < n+1; i =  i*2{
			level ++
		}
		return level
	}
	for i := 0; i < 20; i++{
		fmt.Println(i, t(i))	
	}
	
}