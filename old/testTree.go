package main

//import tree "binarySearchTree"
//import rb  "redblack"
import dp "dyamicProgramming"
import "fmt"

func main() {
	//tree.Test_Insert()
	//	rb.Test_RedBlackTree()

	dp.Test_revenue()
	dp.Test_lcs()

	fmt.Println("00000000")
	m := a(1)
	n := &m
	n.add(10)
	fmt.Println(*n)
}

type a int

func (this *a) add(n int) {
	t := int(*this)
	g := t + n
	fmt.Println(g)
	m := a(g)
	this = &m
	fmt.Println(*this)
}
