package list

import "fmt"
import "hong/sort"
import "math/rand"

func Test_Insert(){
	a := NewSkipList(sort.Int())

	for i := 0; i < 100; i++{
		a.Insert(rand.Intn(100))
	
	}
	
	fmt.Println(a.level, " == ",a)
	a.print()

}