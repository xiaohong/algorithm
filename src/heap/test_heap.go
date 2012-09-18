package heap

import "math/rand"
import "fmt"


func Test_MinMaxHeap(){
	a := NewMinMaxHeap()
	for i := 0; i < 20; i++{
		t := rand.Intn(100)
		a.Add(t)	
		fmt.Println(t, i,  a.data)
	}

	fmt.Println(a.Min())
	fmt.Println(a.Max())
	
}

func Test_a(){
	t := []int{28,95,66,58}
	a := &MinMaxHeap{ len(t), t}
	fmt.Println(a.data)
	a.Add(47)
	a.Add(43)

	fmt.Println(a.data)
}