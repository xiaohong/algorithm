package disjointset


type node struct{
	head *node
	next *node
	value interface{]}
}

type Set {
	head *node
	tail *node
}

