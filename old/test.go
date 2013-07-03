package main

//import c "chapter2"
import "fmt"
import "reflect"


func main() {
	 s := "ddf的地方大放送"

	 for i, v := range s {
	 	fmt.Println(i,  rune(v),reflect.TypeOf(v))
	 	fmt.Printf("%s", v)
	 
	 }
}


