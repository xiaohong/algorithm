package hashTable

import "fmt"
// 直接地址表格使用key作为数组的索引，而且key不能重复的条件下

type DirectAddressTable []interface{}

func (this DirectAddressTable) Search(i int) interface{}{
	return []interface{}(this)[i]
}

func (this DirectAddressTable) Insert(key int, v interface{}){
	[]interface{}(this)[key]=v
}

func (this DirectAddressTable) Delete(key int){
	[]interface{}(this)[key]=nil	
}

func NewDirectAddressTable() DirectAddressTable{
	return DirectAddressTable(make([]interface{},100))
}

func Test_DirectAddressTable(){
	d := NewDirectAddressTable()
	d.Insert(1,44)
	d.Insert(5, "ddd")
	fmt.Println(d)
	fmt.Println(d.Search(5))
	d.Delete(5)
	fmt.Println(d)
}

// 11.1-1
// 遍历数组所有元素，最糟糕的情况是O(n)

// 11.1-2
// 位为1表示有值，位为0表示没有值