package main

import "ds"
import "fmt"
import "math/rand"

func main() {
	q := ds.NewArrayQueue(10)

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())
	fmt.Println(q.Size())

	for i := 0; i < 100; i++{
		q.Enqueue(rand.Int())
	}

	fmt.Println(q.Size())



	twoStack := ds.NewTwoStack(5)
	twoStack.PushHead(1)
	twoStack.PushHead(1)
	twoStack.PushTail(2)
	twoStack.PushTail(2)
	twoStack.PushTail(3)
	fmt.Println(twoStack)

//	twoStack.PushTail(2)

	ds.Test_10_1_3()


	d := ds.NewDeque(2)
	d.PushHead(11)
	fmt.Println(d)
	d.PushTail(12)
	fmt.Println(d)
	d.PushHead(3)
	fmt.Println(d)

	d.PushHead(3)
	fmt.Println(d)
	d.PushTail(3)
	d.PushTail(3)

	d.PushTail(29)

	d.PushTail(30)

	d.PushTail(31)

	d.PushTail(32)
	d.PushTail(32)
	fmt.Println(d)
	t := d.PopTail()
	for ;t != nil; {
		fmt.Print(t, ", ")
		t = d.PopTail()
	}
		fmt.Println(d)

	fmt.Println(5%3)


	a := ds.NewTwoStackQueue()
	a.Enqueue(1)
	a.Enqueue(2)
	fmt.Println(a)
	fmt.Println(a.Dequeue())
		fmt.Println(a.Dequeue())

}

// 10.1-7
// 使用2个队列实现一个栈
// 一个主队列，一个辅助队列，入栈时进入主队列，
// 出栈时，首先把主队列所有元素出对列，一次加入到辅助队列，但主队列为空时
// 当前元素不入辅助队列，把返回，然后再把辅助队列的数据转移到主队列中