package learnTour

//import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

/*
通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。

package main

import "fmt"

type IPAddr [4]byte

// td: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
*/

import "fmt"

type IPAddr [4]byte

func (self IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", self[0], self[1], self[2], self[3])
}

func Main0006() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
