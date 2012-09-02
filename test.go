package main

import tree "binarySearchTree"
import "fmt"

func main(){
	t := tree.NewTree()
	fmt.Println(t)
	t.Insert(5)
	t.Insert(10)
	t.Insert(1)
	t.Insert(3)
	t.Insert(7)

	t.Inorder()

	t.Inorder_with_stack()
	t.Inorder_with_stack_2()

	t.Preorder()
  	fmt.Println( t.Search(3))

  	fmt.Println(t.Search_iterative(3))

  	fmt.Println(t.Maximum())
  	fmt.Println(t.Minimum())

  	fmt.Println(t.Search_iterative(3).Successor())
  	fmt.Println(t.Search(3).Predecessor())

  	fmt.Println(t.Minimum_recusive())

  	t.Inorder_minimum_successor()
}