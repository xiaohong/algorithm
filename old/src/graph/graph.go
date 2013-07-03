package graph

import "fmt"


type DirectedGraph struct{
	// 所有顶点点列表
	nodes []interface{}
	// 每个顶点到相应顶点的连接关系
	neighbors map[interface{}][]interface{}
}

func (this *DirectedGraph) Neighbors(v interface{}) []interface{}{
	return this.neighbors[v]
}

func (this *DirectedGraph) reverse() *DirectedGraph{
	a := make(map[interface{}][]interface{},10)

	for _, v := range this.nodes {
		ns := this.Neighbors(v)
		for _, n := range ns{
			val, ok := a[n]
			if ok {
				a[n] = append(val, v)
			}else{
				a[n] = []interface{}{v}
			}
		}
	}

	return &DirectedGraph{this.nodes, a}
}

func (this *DirectedGraph) String() string{
	s := "{"
	s += fmt.Sprint(this.nodes) 
	s += "}"

	for k,v := range this.neighbors {
		s += fmt.Sprint(k, "=", v, ", ")
	}
	return s
}

func Test_Reverse(){
	a := make(map[interface{}][]interface{},10)
	a[1] = []interface{}{2,3}
	a[2] = []interface{}{3}

	n := []interface{}{1,2,3}	

	g := &DirectedGraph{n, a}
	fmt.Println(g)

	fmt.Println(	g.reverse())
}