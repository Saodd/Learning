package learnTour

import (
	"fmt"
	"io"
	"os"
)

/*
实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TD: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func main() {
	reader.Validate(MyReader{})
}

*/

type MyReader struct{}

func (self MyReader) Read(x []byte) (int, error) {
	for i := range x {
		x[i] = 'A' // 单引号是byte，双引号是string
	}
	return len(x), nil
}

func Main0009() {
	Validate(MyReader{})
}

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}
