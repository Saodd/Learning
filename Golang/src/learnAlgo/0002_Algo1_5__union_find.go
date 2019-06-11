package learnAlgo

import (
	"fmt"
)

/*
书P136页：
union-find算法（动态连通性问题）。

--输入：触点总数N
--输入：代表所有连接的整数对数组[(1,2),(1,4)...]
--求：联通分量的数量（即所有触点被隔离成几个部分了？）（类似于：至少需要再联多少次才能把所有的触点连接？）

--举例：9个数，5个连接[[0,3],[3,4], [4,7],[7,8],[5,8]]。结果应当是3个分量。
*/

type UF struct {
	nodes []int
}

func (self *UF) union(x, y int) {
	// 连接的时候必须把根节点相连，不能只连当前节点。
	// 思路与书P142的quick-union思路相同。
	self.nodes[self.find(y)] = self.find(x)

}

func (self *UF) find(x int) int {
	if self.nodes[x] != x {
		// 递归找到根节点，如果找到的根节点与自身储存的根节点不同，那就保存新的
		// 这种解法在书中称为“路径压缩算法”，即每个节点在寻路的过程中直接保存最新的根节点
		if self.nodes[x] != self.find(self.nodes[x]) {
			self.nodes[x] = self.find(self.nodes[x])
			// 上面find()了两次，其实可以缓存的，但是考虑到并不一定每次都要改，缓存也许更慢呢？
			// 而且由于路径压缩，find()的效率本身就非常高。
		}
		return self.nodes[x]
	} else {
		return x
	}
}
func (self *UF) count() int {
	set := []int{}
	sum := 0
	cache := 0
	var connected bool
	for x := range self.nodes {
		cache = self.find(x)
		connected = false
		for _, s := range set {
			if s == cache {
				connected = true
			}
		}
		if !connected {
			sum++
			set = append(set, cache)
		}
	}
	return sum
}

func Main0002() {
	uf := &UF{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	mapping := [][]int{{0, 3}, {3, 4}, {4, 7}, {7, 8}, {5, 8}}
	for _, m := range mapping {
		uf.union(m[0], m[1])
	}
	for i := range uf.nodes {
		fmt.Println(uf.find(i))
	}
	//for _,v := range uf.nodes {
	//	fmt.Println(v)
	//}
	fmt.Print("group num: ", uf.count())
}

/*
小结：
花了2个小时左右，然后在回家的路上顿悟了这个最优解。
但不可否认的是，在开始构建的时候受到书的启发，把核心逻辑分为find()和union()两部分，这让我事半功倍。
*/
