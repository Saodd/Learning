package learnTour

import (
	"io"
	"os"
	"strings"
)

/*
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader
的 *gzip.Reader（解压后的数据流）。

编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，
通过应用 rot13 代换密码对数据流进行修改。

rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

*/

// 先实现rot13
func rot13(b byte) byte {
	// A=65 ,Z=90,a=97 z=122
	// 注意，byte加减操作之后会变成数字（实质上就是数字，按照ascii翻译）
	// 要转换为字符串就需要string()，但是这里只需要返回byte即可，后面的函数会处理成字符串
	switch {
	case (b > 122) || (b < 65):
		return b
	case (90 < b) && (b < 97):
		return b
	case (b > 109) || ((b > 77) && (b < 91)):
		return b - 13
	default:
		return b + 13
	}
}

type rot13Reader struct {
	r io.Reader
}

func (self rot13Reader) Read(x []byte) (int, error) {
	num, err := self.r.Read(x)
	if err != nil {
		return 0, err
	}
	for i := 0; i < num; i++ {
		x[i] = rot13(x[i])
	}
	return num, nil
	//return self.r.Read(x)
}

func Main0010() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
