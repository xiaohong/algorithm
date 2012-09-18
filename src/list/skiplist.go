package list

type node struct{
	x interface{}
	next []*node	
}

func NewNode(x interface{}, h int) *node{
	return &node{x, make([]*node, h+1)}
}

func (this *node) height() int{
	return  len(this.next) -1 
}