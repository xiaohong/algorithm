package main

import ("net"
	"fmt"
	"bufio"
	)


func main(){

	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {

	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	bf := bufio.NewReader(conn)
	status, err := bf.ReadString('\n') 

	for err == nil {
		fmt.Println(status)	
		status, err = bf.ReadString('\n') 
	}


	

	
}