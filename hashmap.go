package main

import "fmt"

type Hashable interface{
	Hash() int
}

type Map interface{
	Put(key Hashable, v interface{})
	Get(key Hashable) interface{}
	Size() int
}

type Iterator interface{
	HasNext() bool
	Next() interface{}
}



type Entry struct{
	key Hashable
	val interface{}
	hash int 
	next *Entry
}

// hashmap实现
type HashMap struct{
	tables []*Entry
	size   int
}


func NewHashMap() *HashMap{
	t := new(HashMap)
	t.tables = make([]*Entry, 10)
	t.size = 0

	return t
}

func (m *HashMap) Put(key Hashable, v interface{}){
	index := key.Hash()%(len(m.tables))


	var t *Entry = nil
	if m.tables[index] != nil {
		t = m.tables[index]
	}
	e := Entry{key, v, key.Hash(), t}
	m.tables[index] = &e
	m.size++
}

func (m *HashMap) Get(key Hashable) interface{}{
	index := key.Hash()%(len(m.tables))

	entry := m.tables[index]

	for ; entry != nil ;{
		if entry.key == key {
			return entry.val
		}
		entry = entry.next
	}
	return nil
}

func (m *HashMap) Size() int{
	return m.size
}

func (m *HashMap) String() string{
	if m.size == 0{
		return "{}"
	}
	s := "{"
	for _, v := range m.tables {
		t := v 
		for ;t != nil; {
			s = s + fmt.Sprint( t.key , "=" , t.val , ", ")
			t = t.next
		}
	}
	return s +"}"
}

func (m *HashMap) resize(size int){
	newTable := make([]*Entry,size)

	oldTable := m.tables
	for _, v := range oldTable {
		for ; v != nil; {
			index := v.hash & (size-1)
			old := v.next 
			v.next = newTable[index]
			newTable[index] = v 
			v = old
		}
	}
	m.tables = newTable
}

type hashMapIterator struct{
	current *Entry
	next    *Entry
	index int 
	hashmap *HashMap
}

func (m *HashMap) Iterator() Iterator{
	i := 0;
	var begin *Entry = nil
	for ; i < len(m.tables); i++{
		if m.tables[i] != nil {
			begin = m.tables[i]
			break
		}
	}
	return &hashMapIterator{nil, begin, i, m}
}

func (this *hashMapIterator) HasNext() bool{
	return this.next != nil
}

func (this *hashMapIterator) Next() interface{}{
	t := this.next 
	this.next = this.next.next 
	if this.next == nil{
		for this.index++ ; this.index < len(this.hashmap.tables); {
			if this.hashmap.tables[this.index] != nil {
				this.next = this.hashmap.tables[this.index]
				break
			}
			this.index++ 
		}
	}
	this.current = t
	return t
}

type S string
func (s S) Hash() int{
	return len(s)
}

type linkedEntry struct{
	Entry
	before, after *linkedEntry
}

type linkedHashMap struct{
	*HashMap
	header *linkedEntry
}

func NewLinkedHashMap() Map{
	return &linkedHashMap{NewHashMap(),nil}
}

 
func main() {
	m := NewHashMap()

	m.Put(S("ss"),1)

	m.Put(S("ssss"),1)

	m.Put(S("ssss"),22)

	m.resize(30)

	fmt.Println(m)

	i := m.Iterator()

	fmt.Printf("%+v", i, "\n")

	for ; i.HasNext(); {
		fmt.Println(i.Next())
	}

	a := NewLinkedHashMap()
	fmt.Println(a)

}


