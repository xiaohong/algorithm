package main

//import "os"
import "fmt"
import "io"
import "unsafe"

type StringReader struct{
	data []byte
	i int
}

func New(s string) io.ByteReader{
	return &StringReader{[]byte(s),0}
}

func (this *StringReader) ReadByte() (byte, error){
	this.i++
	return this.data[this.i-1], nil
}

type A struct {
	a int
}

func (a *A) add(b int) int{
	a.a = a.a+b
	return a.a
}

type B struct{
	*A
	b int
}

func (b *B) dis(a int){
	b.b = b.A.add(a)+a
}

type TT string
type T2 []TT

func main(){

	s := "的地方的的"
	fmt.Printf("%T", len(s))

	var a [10]int
	a[1] =2
	fmt.Println(a, len(a))

	for v,i := range s {
		fmt.Println(v, string(i))
	}

	c := make(chan int)

	go func(){
		for i := int64(0); i < int64( 100000); i++{

		}
		c <- 1
		fmt.Println("111")
		}()

	k := <- c 
	fmt.Println(k)

	var t5 T2 = []TT{"dd"}
	fmt.Println(t5)


	var yy int = 1
	fmt.Println( unsafe.Pointer(&yy))

	fmt.Println(unsafe.Sizeof([]string{"123"}))

}