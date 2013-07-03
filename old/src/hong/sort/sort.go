package sort

type Comparable interface{
	Compare(other interface{}) int
}

type Comparator func(a, b interface{}) int


func Int() Comparator{
	return func (a, b interface{}) int{
		bb, _ := b.(int)
		switch v := a.(type){
		default:
			return 0
		case int :
			return v -bb
		} 
			return 0
	}

}