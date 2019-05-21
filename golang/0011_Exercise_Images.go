package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

/*
还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

ColorModel 应当返回 color.RGBAModel。

At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
*/

type Image struct {
	dx   int
	dy   int
	data [][]uint8
}

func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.dx, self.dy)
}
func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (self *Image) At(x, y int) color.Color {
	return color.RGBA{self.data[y][x], self.data[y][x], 255, 255}
}
func (self *Image) Generate() { // 要修改原来的对象(struct)，就必须要*传入指针
	self.dx, self.dy = 256,256
	z := make([][]uint8, self.dy)
	for y := range z {
		z[y] = make([]uint8, self.dx)
	}
	var x int
	for y := range z {
		for x = 0; x < self.dx; x++ {
			z[y][x] = uint8(x * y)
		}
	}
	self.data = z
}

func main() {
	m := &Image{}  // 这里也必须要&取指针，否则后面没法运行
	m.Generate()
	ShowImage(m)
}

/*
感触：
一开始觉得教程上的提示也太少了吧，不过写着写着还是明朗了。
关键是明确目标函数ShowImage(m image.Image)，然后去找image.Image这个接口，然后逐个实现就好。
#论IDE的重要性#
 */