package main

import "ds"

//import "fmt"

func main() {
	d := ds.NewArrayStack()
	d.Push(1)
	d.Push(2)
	d.Push(3)

	ds.Test_PrintTree()
}
