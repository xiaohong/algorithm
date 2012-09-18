package heap

import "math/rand"
import "fmt"


func Test_MinMaxHeap(){
	a := NewMinMaxHeap()
	for i := 0; i < 20; i++{
		a.Add(rand.Intn(100))	
	}

	fmt.Println(a.Min())
	fmt.Println(a.Max())
	fmt.Println(a.data)
}