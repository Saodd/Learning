package learnAlgo

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func Main0001() {
	fmt.Println(gcd(10, 5))
	fmt.Println(gcd(5, 10))
	fmt.Println(gcd(99, 3))
}
