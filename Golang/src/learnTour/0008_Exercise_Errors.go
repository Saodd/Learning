package learnTour

/*
从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

type ErrNegativeSqrt float64
并为其实现

func (e ErrNegativeSqrt) Error() string
方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回
"cannot Sqrt negative number: -2"。

注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：
fmt.Sprint(float64(e))。这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。

package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

*/
import (
	"fmt"
)

type ErrNegativeSqrt float64

func (self ErrNegativeSqrt) Error() string {
	// 将self转换为string会进入死循环。
	// 因为转换的时候就是在调用Error()方法。
	return fmt.Sprintf("cannot Sqrt negative number: %f", self)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		// copied
		z := 1.0
		for y := 0.0; (y-z > 1e-8) || (z-y > 1e-8); {
			y = z
			z = y - (y*y-x)/(2*y)
		}
		return z, nil
	}
}

func Main0008() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
