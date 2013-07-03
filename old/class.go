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

func readU1(file *os.File) byte{
	a := make([]byte,1)
	file.Read(a)
	return a[0]
}

func readN(file *os.File, n int) []byte{
	a := make([]byte,n)
	file.Read(a)
	return a
}

type JavaClass struct{
	magic string
	minor, major int
	constantPool []CPInfo
//	uint16 thisClass
}

type CPInfo interface{
	Tag() int
}

type ClassInfo struct{
	tag int 
	nameIndex uint16
}

func main(){
		file, err := os.Open("HoolaiApplication.class")
	if err != nil {
		print(err)
	}


	fmt.Printf("%x\n",readU4(file))
	fmt.Println(readU2(file))
	fmt.Println(readU2(file))
	
	count := int( readU2(file))



	tag := make([]byte,1)
	for i := 0; i <  count-2; i++{
		file.Read(tag)
		//fmt.Println("tag",tag)
		typeTag := int(tag[0])
		switch typeTag {
		case 7 :
			readU2(file)
		case 9,10,11:
			readU2(file)
			readU2(file)
		case 8:
			readU2(file)
		case 3,4:
			readU4(file)
		case 5,6:
			readU4(file)
				readU4(file)
		case 12:
			readU2(file)
			readU2(file)
		case 1:
			l := readU2(file)
			data := make([]byte, l)
			file.Read(data)
			fmt.Println(string(data))
		case 15:
			readU1(file)
			readU2(file)
		case 16:
			readU2(file)
		case 18:
			readU2(file)
			readU2(file)
		default:
			//panic("dd")
			fmt.Println("ddd",tag)
		}
	}

	fmt.Println(readU2(file))
	fmt.Println(readU2(file))
	fmt.Println(readU2(file))

	iCount := int(readU2(file))
	for i := 0; i < iCount; i++{
		readU2(file)
	}

	fieldCount := int(readU2(file))
	for i := 0; i < fieldCount; i++{
		readU2(file)
		readU2(file)
		readU2(file)

		ac := int( readU2(file))
		for j := 0 ; j < ac; j++{
			parseAttribute(file)
		}
	}
	fmt.Println()


	fmt.Println(count)
}

func parseAttribute(file *os.File) {
	i := readU2(file)
	l := readU4(file)
	data := readN(file, int(l))
	fmt.Println(i, "==", data)
}