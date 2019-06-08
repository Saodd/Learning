package p004_structure

import (
	"errors"
	"sync"
)

// 入门版 ------------------------------------------------------------
func NewMyFixedStackInt(cap int) MyFixedStackInt {
	return MyFixedStackInt{cap, 0, make([]int, cap)}
}

type MyFixedStackInt struct {
	cap    int
	length int
	box    []int
}

func (self *MyFixedStackInt) len() int {
	return self.length
}

func (self *MyFixedStackInt) isempty() bool {
	if self.length == 0 {
		return true
	}
	return false
}

func (self *MyFixedStackInt) push(x int) error {
	if self.length < self.cap {
		self.box[self.length] = x
		self.length += 1
		return nil
	} else {
		return errors.New("超出容量")
	}
}

func (self *MyFixedStackInt) pop() (int, error) {
	if self.length > 0 {
		self.length--
		return self.box[self.length], nil
	} else {
		return 0, errors.New("没有元素可以取出")
	}
}

// 升级版 ------------------------------------------------------------
//func NewMyStack() *MyStackInt {
//	return &MyStackInt{}
//}

type MyStackInt struct {
	items []int
	lock  sync.Mutex
}

func (self *MyStackInt) Len() int {
	return len(self.items)
}

func (self *MyStackInt) Push(obj int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.items = append(self.items, obj) //append好像不会返回error??
}
func (self *MyStackInt) Pop() (int, error) {
	self.lock.Lock()
	defer self.lock.Unlock()
	if len(self.items) == 0 {
		return 0, errors.New("没有元素可以取出")
	}
	obj := self.items[len(self.items)-1]
	self.items = self.items[:len(self.items)-1]
	return obj, nil
}
