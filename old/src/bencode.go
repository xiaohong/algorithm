package main 

// string len:string byte data
// int    i12e
// list   l<bencode value>e
// dictionary(map)  d<bencoded string><bencoded element>e 

import "fmt"
import "reflect"

type  Decoder struct{
	Data []byte
	Index    int
}

func (d *Decoder) current() byte{
	return d.Data[d.Index]
}

func NewDecoder(data []byte) *Decoder{
	a := new(Decoder)
	a.Data = data 
	return a
}

func EncodeInteger(i int) []byte{
	s := fmt.Sprintf("i%de", i)
	return []byte(s)
}

func DecodeInteger(data *Decoder) int{
	if data.Data[data.Index] != 'i' {
		panic("error integer")
	}

	data.Index++
	var i int = 0;
	k := data.Index
	for  ; k < len(data.Data); k++{
		if data.Data[k] == 'e'{
			break;
		}
		i = i * 10 + int(data.Data[k])-48
	}
	k++
	data.Index = k
	return i
}

func EncodeString(s string) []byte{
	return []byte(fmt.Sprintf("%d:%s", len([]byte(s)), s))
}


func DecodeString(data *Decoder) string{
	if data.current() < '0' || data.current() > '9' {
		panic("error string")
	}

	var l int = 0
	i := data.Index
	for ; i < len(data.Data); i++{
		if data.Data[i] == ':' {
			break
		}
		l = l * 10 +  int( data.Data[i]) - 48
	}

	data.Index = i+l+1

	return string(data.Data[i+1:i+1+l])
}

func Encode(v interface{}) []byte{
	switch reflect.TypeOf(v).Kind(){
	case reflect.String:
		return EncodeString(v.(string))
	case reflect.Int:
		return EncodeInteger(v.(int))
	default:
		panic("error ")
	}
	return nil
}

func EncodeList(data []interface{}) []byte{
	r := []byte{'l'}
	for _, v := range data {
		r = append(r, Encode(v)...)
	}
	return append(r, 'e')
}

func DecodeList(data *Decoder) []interface{}{
	if data.current() != 'l' {
		panic("error")
	}

	t := []interface{}{}
	data.Index++
	for{
		if data.current() == 'i' {
			t = append(t, DecodeInteger(data))
		}else {
			t = append(t, DecodeString(data))
		}
		if data.current() == 'e'{
			break
		}
	}
	return t
}

func EncodeMap(data map[string]interface{}) byte[]{
	r := []byte{'d'}

	keys := []string{}

	for k,_ := range data {
		keys = append(keys, k)
	}
	
}

func main() {
	a := EncodeInteger(100334)
	fmt.Println(DecodeInteger(NewDecoder( a)))

	s := EncodeString("hello, ping")
	fmt.Println(DecodeString(NewDecoder( s)))

	fmt.Println(EncodeList([]interface{}{1,"rest,",3}))

	fmt.Println(DecodeList(NewDecoder(EncodeList([]interface{}{1,"rest,",3}))))
}