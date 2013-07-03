package main

import ("net"

	"regexp"
	"strconv"
	"fmt"
)

type DataType int

const {
	Val = iota
	IntVal
	ListVal 
	MapVal 

}

type Key struct{
	name []byte
	valType DataType
}

type Value interface{}


var data map[interface{}]interface{} =  make(map[interface{}]interface{},0)

func readline(conn net.Conn) string{
	c := make([]byte,1)
	_,err := conn.Read(c)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte,0)
	s := ""
	for  {
		if c[0] != '\n' {
			data = append(data,c...)
			_,err = conn.Read(c)
			if err != nil {
				fmt.Println(err)
			}
		}else {
			s = string(data[0:len(data)-1])
			break
		}
	}
	fmt.Println(s)
	return s
}

func findNum(h string) int{
	r, _ := regexp.Compile("\\d+")
	s := r.FindStringSubmatch(h)
	num, _ := strconv.Atoi(s[0])
	return num
}

func handeConnection(conn net.Conn){

	for {
			h := readline(conn)
	fmt.Println("con...")
	num := findNum(h)
	cmd := make([]string, 0)
	for i := 0; i < num; i++{
		byteNum := readline(conn)
		 findNum(byteNum)
		cmd = append(cmd, readline(conn))
	}

	fmt.Println("接受到命令", cmd)

	ret := execCmd(cmd)
	conn.Write(ret)

	}

}

var crlf string = "\r\n"

func ok() []byte{
	return []byte("+OK\r\n")
}

func err(s string) []byte{
	return append([]byte("-"+s), crlf...)
}

func integerReply(i int) []byte{
	return append([]byte(":"+strconv.Itoa(i)), crlf...)
}

func bulkReply(s string) []byte{
	d := []byte(s)
	return []byte("$"+strconv.Itoa(len(d))+crlf+s+crlf)
}

type cmdFunc func (args []string) []byte

func ping(args []string) []byte{
	return bulkReply("PONG")
}

var cmds map[string]cmdFunc

func init(){
	cmds = make(map[string]cmdFunc, 10)
	cmds["Ping"] = ping
}

func execCmd(cmd []string) []byte {
	f,ok := cmds[cmd[0]]
	if ok {
		return f(cmd)
	}
	return err("error cmd " + fmt.Sprint( cmd))
}



func multiBulkReply(data [][]byte) []byte{
	ret := make([]byte,0)

	ret = append(ret, "*"+strconv.Itoa(len(data))+crlf...)
	for _, v := range data{
		if v == nil {
ret = append(ret, "$-1"+crlf...)
			}else{
		ret = append(ret, "$"+strconv.Itoa(len(v))+crlf...)
		ret = append(ret, v...)
		ret = append(ret, crlf...)
		}
	}
	return ret
}

func main(){
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {

	}

	for {
		conn, err := ln.Accept()
		if err != nil {

		}

		go handeConnection(conn)
	}
}