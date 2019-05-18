package main

import (
	"fmt"
)

func run(pool, out chan int) {
	var sum int
	for n := range pool {
		sum = 0
		for i := 0; i < n; i++ {
			sum += i
		}
		out <- sum
	}

}

func main() {
	pool := make(chan int, 1000)
	out := make(chan int, 10)
	for i := 0; i < 1000; i++ {
		pool <- 1000000000
	}
	close(pool)
	for i := 0; i < 8; i++ {
		go run(pool, out)
	}
	// 这里out通道没人close，会死锁报错，不过暂时先不管了
	// 8个CPU全部跑满，感觉真high啊。比python爽多了…………
	for sum := range out {
		fmt.Println(sum)
	}
}
