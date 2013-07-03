package main


import "os"
import "fmt"


func readU4(file *os.File) uint32 {
	a := make([]byte,4)
	file.Read(a)
	return (uint32(a[0]) << 24 ) | (uint32(a[1]) << 16 ) | (uint32(a[2]) << 8 ) | (uint32(a[3]) )
}

func readU2(file *os.File) uint16{
	a := make([]byte,2)
	file.Read(a)
	return (uint16(a[0]) << 8 ) | (uint16(a[1]) ) 
}

type Frame struct{
	operand []interface{}
	local []interface{}
}

type Runtime struct{
	pc int 
	stack []*Frame
}


func main(){
	file, err := os.Open("HoolaiApplication.class")
	if err != nil {
		print(err)
	}


	fmt.Printf("%x\n",readU4(file))
	fmt.Println(readU2(file))
	fmt.Println(readU2(file))
	fmt.Println(readU2(file))

}