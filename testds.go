package main

import  "ds"
import "fmt"

func main() {
	d  := ds.NewArrayStack()
	d.Push(1)
	d.Push(2)
	d.Push(3)

	fmt.Println(d.Pop())
	fmt.Println(d.Peek())

	fmt.Println(d)

	fmt.Println(U+000A)
}