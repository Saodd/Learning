package learnTour

import (
	"golang.org/x/tour/tree"
)

/*
练习：等价二叉查找树
1. 实现 Walk 函数。

2. 测试 Walk 函数。

函数 tree.New(k) 用于构造一个随机结构的已排序二叉查找树，它保存了值 k, 2k, 3k, ..., 10k。

创建一个新的信道 ch 并且对其进行步进：

go Walk(tree.New(1), ch)
然后从信道中读取并打印 10 个值。应当是数字 1, 2, 3, ..., 10。

3. 用 Walk 实现 Same 函数来检测 t1 和 t2 是否存储了相同的值。

4. 测试 Same 函数。

Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1), tree.New(2)) 应当返回 false。

Tree 的文档可在这里https://godoc.org/golang.org/x/tour/tree#Tree找到。
-------------------------------------------------------------------
import "golang.org/x/tour/tree"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int)

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool

func main() {
}
-------------------------------------------------------------------
*/

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk2(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk2(t.Right, ch)
	}
	close(ch)
}
func Walk2(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk2(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk2(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	//var v1, v2 int
	//var ok2 bool
	for v1 := range ch1 {
		v2, ok2 := <-ch2
		println(v1, v2)
		if !ok2 {
			//defer close(ch1)
			return false
		}
		if v1 != v2 {
			//defer close(ch1) 不关闭了，让GC自动回收吧
			//defer close(ch2)
			return false
		}
	}
	return true
}

func Main0012() bool {
	t1, t2 := tree.New(1), tree.New(1)
	println(t1.String())
	println(t2.String())
	r := Same(t1, t2)
	println("Compare result: ", r)
	return r

}

//妈呀不懂二叉树，先暂停一下，先去学算法。
