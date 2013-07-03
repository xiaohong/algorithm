package hashTable

import "fmt"

type openEntry struct{
	key, value interface{}
	hash int
}

func (this *openEntry) String() string{
	return fmt.Sprint("[",this.key, this.value, this.hash, "]")
}

type OpenAdressHashMap struct{
	size int 
	data []*openEntry
}

func (this *OpenAdressHashMap) Put(k, v interface{}){
	n := &openEntry{k,v,hash(k)}

	h := hash(k) % (len(this.data)-1)
	for this.data[h] != nil {
		// 使用线性
		h = lineProbe(h, (len(this.data)-1))
	}
	this.data[h] = n 
	this.size++
	if this.size >= len(this.data){
		this.resize()
	}
}

func (this *OpenAdressHashMap) resize(){

}

// 顺序存储冲突的数据
func lineProbe(n, l int) int{
	return (n + 1)%l
}

func (this *OpenAdressHashMap) Get(k interface{}) interface{}{
	h := hash(k) % (len(this.data)-1)

	for this.data[h] != nil {
		fmt.Println(this.data[h] , h)
		if this.data[h].key == k {
			return this.data[h].value
		}
		h = lineProbe(h, (len(this.data)-1))
	}
	return nil
}

func (this *OpenAdressHashMap) Size() int{
	return this.size
}

func (this *OpenAdressHashMap) Delete(key interface{}) bool{
	h := hash(key) % (len(this.data) -1)

	for this.data[h] != nil {
		if this.data[h].key == key {
			// 直接删除貌似有问题，如果后面有和前面相同的hash数据，则会搜索不到
			// 这时要做一个假数据，说明这个地方是一个被删除的地方，如果有插入操作在此，
			// 可以经常操作，如果是搜索操作则可以跳过继续进行下面的数据
			this.data[h] = nil
			return true
		}
	}
	return false
}

func NewOpenAdressHashMap() Map{
	m := &OpenAdressHashMap{data: make([]*openEntry,20)}
	return m
}

func Test_OpenAdressHashMap(){
	m := NewOpenAdressHashMap()
	m.Put(S("aa"), 1)
	m.Put(S("CC"), 2)

	fmt.Println(m)
	fmt.Println(m.Get(S("CC")))
}

// 11.4-3
// 上界就是把所有元素都检查完了


// 完美hash table使用2级hash table实现
// 第一级hash table和正常的一样
// 第二级hash table用于存储冲突的,通过选择hash function可以避免在二级hash table中发生冲突
// h(k) = ((ak+b) mod p) mod m
// hj(k) = ((aj*k + bj) mod p) mod mj  
// mj = nj^2
