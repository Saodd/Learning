package main

/*
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)
*/
import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x1 := 0
	x2 := 1
	var temp int
	return func() int {
		temp = x1
		x1, x2 = x2, x1+x2
		return temp
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
