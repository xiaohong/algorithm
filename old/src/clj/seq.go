package main 

import "fmt"

type Seqable interface{
	Seq() ISeq
}

type ISeq interface{
	Seqable
	First() interface{}
	Next() ISeq
	More() ISeq
	Cons(o interface{}) ISeq
}

type Cons struct{
	first interface{}
	more ISeq
}

func NewCons(f interface{}, m ISeq) *Cons{
	return &Cons{f,m}
}

func (c *Cons) First() interface{}{
	return c.first
}

func (c *Cons) Next() ISeq{
	return c.more.Seq()
}

func (c *Cons) More() ISeq{
	if c.more == nil {
		return nil
	}
	return c.more
}

func (c *Cons) Seq() ISeq{
	return c
}

func (c *Cons) Cons(f interface{}) ISeq{
	return NewCons(f,c)
}

func (c *Cons) String() string{
	s := fmt.Sprint("(", c.first)
	t := c.more
	for t != nil {
		s = fmt.Sprint(s, ",", t.First())
		t = t.More()
	}
	s = fmt.Sprint(s,")")
	return s
}

type Range struct{
	start, end int
}

func (r *Range) 

func main(){
	c := NewCons(1,nil)
	d := NewCons(2,c)
	f := d.Cons(3).Cons(4).Cons(5)
	fmt.Println(f)
}
