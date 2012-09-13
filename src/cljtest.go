package main 

import "clj"

func main() {
	s := clj.NewStringReader("我是的")
	print(s.Read())

}