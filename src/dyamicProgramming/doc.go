package dyamicProgramming

import "fmt"
// optimal substructure
// overlaping subproblems

// a problem exhibits optimal substructure if an optimal solution to
// the problem contains within it optimal solutions to subproblems

// 最长

// 计算c[i,j]的公共字符串的最大长度
func lcs(x, y []int) string{
	m := len(x)
	n := len(y)
	c := make([][]int, m+1)
	for i := 0; i < m+1; i++{
		c[i] = make([]int, n+1)
	}

	//如果x[i]与y[j]相同，那么[i,j]的最长子序列就是[i-1,j-1]+
	// 如果不相同的话，那么[i,j]的最长子序列就是
	for i := 1; i <= m ; i++{
		for j := 1; j <= n; j++{
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1]+1
			}else{
				if c[i][j-1] > c[i-1][j] {
					c[i][j] = c[i][j-1] 
				}else{
					c[i][j] = c[i-1][j]
				}
			}
		}
	}

	s := ""
	xx := len(x)
	yy := len(y)
	for xx >= 1 && yy >=1{
		if x[xx-1] == y[yy-1]{
			s = fmt.Sprint(s, x[xx-1])
			xx--
			yy--
		}else {
			if c[xx-1][yy] < c[xx][yy-1]{
				yy--
			}else{
				xx--
			}
		}
	}
	return s
}




func Test_lcs(){
	fmt.Println(lcs([]int{1,1}, []int{1,1}))
}