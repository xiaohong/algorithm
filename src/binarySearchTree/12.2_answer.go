package binarySearchTree

import "fmt"

//12.2-1
// 可以根据规律如果第一个元素比第二个元素小，那么第一个元素
// 要比后面所有的元素都小，如果大的花，要比后面的所有元素都大
// 根据这个规律检查所有俩俩相挨的数据
// 因为每次选在都是从根节点开始，如果选择左子树，那么后面所有的
// 数据都比他小，如果选择右节点，那么随后的数据都比他大

func check(data []int) bool {
	for i := 0; i < len(data)-1; i++{
		if data[i] > data[i+1]{
			for k := i+2; k < len(data); k++{
				if data[i] < data[k] {
					return false
				}
			}
		}else if data[i] < data[i+1] {
			for k := i+2; k < len(data); k++{
				if data[i] > data[k] {
					return false
				}
			}
		}
	}
	return true
}

func test_check(){
	fmt.Println(check([]int{935,278,347,621,299,391,392,358,363}))
}

// 12.2-2
func (this *Tree) Minimum_recusive() *Node {
	if this.Root.left == nil {
		return this.Root
	}

	return (&Tree{Root: this.Root.left}).Minimum_recusive()
}


// 12.2-3
// 前继者查找算法
//  1. 如果有左子树，左子树的最大节点就是前继者
//  2. 没有左子树，则寻找父代节点，最近的当前节点在父节点的左子树上。


// 12.2-4
// 对于单个节点来说是满足的，但是对于a节点的由子树，于b节点的左子树，而且a比b高2个级别时
// 先走由子树，然后左子树，然后右子树，这时第二个节点的左子树节点就比search path的第一个值要大b<a
// 貌似任意的a肯定比任意的c小

// 12.2-5
// 如果一个节点有2个子节点，那么他的后续者没有左孩子，如果有左孩子的话，这个左孩子应该比他小，应该是后续者
// 前继者没有右孩子，如果有右孩子的话，这个右孩子还小于当前节点的最大节点，应该成为其前继者

// 12.6
// 因为后继者是大于当前节点的最小节点，如果有右子树的话，则是右子树的最小节点，没有右子树，只能从父节点中
// 查找，如果当前节点在父节点的右子树的话，那么当前节点比这个父节点要大，只有遇到第一个父节点，并且在这个
// 父节点的左子树中，这个父节点比当前节点要大，而且是最近的，所以是其后继者

//12.2-7

// 需要调用n次Successor函数，而每次调用都是lg(n), 那最终结果应该是n*lg(n)

func (this *Tree) Inorder_minimum_successor() {

	min := this.Minimum()
	fmt.Print(min.Value, ", ")
	for s := min.Successor(); s != nil; {
		fmt.Print(s.Value, ", ")
		s = s.Successor()
	}
}

// 12.2-6 总有一次要回溯到根节点O(k + h)