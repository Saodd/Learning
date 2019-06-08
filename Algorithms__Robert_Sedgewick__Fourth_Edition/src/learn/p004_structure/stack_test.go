package p004_structure

import (
	"fmt"
	"testing"
)

// 入门版 ------------------------------------------------------------
func TestNewMyFixedStackInt(t *testing.T) {
	s := NewMyFixedStackInt(1)
	var e error
	var r int
	if e = s.push(100); e != nil {
		t.Error("push报错: " + fmt.Sprint(e))
	}
	if e = s.push(100); e == nil {
		t.Error("push超过容量时，没有报错。")
	}
	if r = s.length; r != 1 {
		t.Error("length错误，应当为1，实际为: " + fmt.Sprint(r))
	}
	if r, e = s.pop(); e != nil || r != 100 {
		t.Error(fmt.Sprintf("pop错误: 返回值(应为100): %d, 错误提示:%s", r, e))
	}
	if r = s.length; r != 0 {
		t.Error("length错误，应当为0，实际为: " + fmt.Sprint(r))
	}
	if r, e = s.pop(); e == nil || r != 0 {
		t.Error(fmt.Sprintf("空栈pop没有正确报错: 返回值(应为0): %d, 错误提示:%s", r, e))
	}
	s = NewMyFixedStackInt(2)
	s.push(1)
	s.push(2)
	if r, e = s.pop(); e != nil || r != 2 {
		t.Error(fmt.Sprintf("pop错误: 返回值(应为2): %d, 错误提示:%s", r, e))
	}
	if r, e = s.pop(); e != nil || r != 1 {
		t.Error(fmt.Sprintf("pop错误: 返回值(应为1): %d, 错误提示:%s", r, e))
	}

}

// 升级版 ------------------------------------------------------------
func TestMyStack(t *testing.T) {
	s := &MyStackInt{}
	//s.Len()
	s.Push(1)
	s.Push(2)
	if r := s.Len(); r != 2 {
		t.Error("length错误，应当为2，实际为: " + fmt.Sprint(r))
	}
	if r, e := s.Pop(); e != nil || r != 2 {
		t.Error(fmt.Sprintf("pop错误: 返回值(应为2): %d, 错误提示:%s", r, e))
	}
	if r, e := s.Pop(); e != nil || r != 1 {
		t.Error(fmt.Sprintf("pop错误: 返回值(应为1): %d, 错误提示:%s", r, e))
	}
	if r, e := s.Pop(); e == nil || r != 0 {
		t.Error(fmt.Sprintf("空栈pop没有正确报错: 返回值(应为0): %d, 错误提示:%s", r, e))
	}
}
