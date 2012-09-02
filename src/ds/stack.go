package ds

import "fmt"
import "strings"

type Stack interface{
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
}

// 栈指针可以指向当前栈顶元素，也可以指向下一个要压栈的位置
// 这2者之间的差异主要是栈顶元素的表示
// 通过size来表示当前元素的数量，然后通过减1来获取当前栈顶元素
type arrayStack struct{
	data []interface{}
	size int 
}

func (this *arrayStack) init(){
	this.data = make([]interface{}, 10)
	this.size = 0
}

func (this *arrayStack) resize(newSize int){
	t := make([]interface{}, newSize)
	copy(t, this.data)
	this.data = t
}

func (this *arrayStack) Push(v interface{}){
	if this.size == len(this.data){
		this.resize(this.size*2)
	}
	this.data[this.size] = v
	this.size++
}

func (this *arrayStack) Pop() interface{}{
	if this.size <= 0{
		panic("empty stack")
	}
	this.size--
	t := this.data[this.size]
	this.data[this.size] = nil
	return t
}

func (this *arrayStack) Peek() interface{}{
	if this.size <= 0{
		panic("empty stack")
	}
	return this.data[this.size-1]	
}

func (this *arrayStack) ToString() string{
	s := []string{}
	v := this.data[0:this.size]
	for _, n := range v {
		s = append(s, fmt.Sprint(n))
	}
	return "[" +strings.Join(s,",") +"]"
}

func (this *arrayStack) Size() int{
	return this.size
}

func NewArrayStack() Stack{
	t := new(arrayStack)
	t.init()
	return t
}