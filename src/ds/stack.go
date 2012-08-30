package ds

import "fmt"
import "strings"

type Stack interface{
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
}

type arrayStack struct{
	data []interface{}
	size int 
}

func (this *arrayStack) init(){
	this.data = make([]interface{}, 10)
	this.size = 0
}

func (this *arrayStack) resize(newSize int){
	//t := make([]interface{}, newSize)

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
	return this.data[this.size]
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