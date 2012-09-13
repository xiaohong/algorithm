package main

import "clj"
import "fmt"

//import "unicode/utf8"

func main() {
	s := clj.NewStringReader("单独d")
	t := s.Read()
	fmt.Println(111)
	fmt.Println(rune(t))

	print(string(int('单')))

}
