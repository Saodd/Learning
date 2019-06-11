package p004_structure

import "errors"

// 二叉堆实现优先队列 --------------------------------------------------
type MyPriorityQueueInt struct {
	items []int
}

func (self *MyPriorityQueueInt) swim(x int) {
	// 从下向上浮，只需要跟父节点比较，一般用在插入
	for f := x / 2; (x > 1) && (self.items[x] > self.items[f]); {
		self.items[x], self.items[f] = self.items[f], self.items[x]
		x = f
		f = x / 2
	}
}

func (self *MyPriorityQueueInt) sink(x int) {
	// 从上向下沉，
	for l, m := len(self.items)-1, 2*x; m <= l; m = 2 * x {
		//在它和它的两个子节点中，寻找最大的那个跟它本身交换
		// 先跟左边的比,m指向更大的
		if self.items[x] >= self.items[m] {
			m = x
		}
		// 再跟右边的比,m指向更大的
		if n := x*2 + 1; n <= l {
			if self.items[n] > self.items[m] {
				m = n
			}
		}
		if m == x { //最大值是本身的话，就停止下沉了
			break
		} else { // 最大值是子节点，那就继续下沉
			self.items[x], self.items[m] = self.items[m], self.items[x]
			x = m
		}
	}
}

func (self *MyPriorityQueueInt) Push(x int) {
	if len(self.items) == 0 {
		self.items = append(self.items, 0)
	}
	self.items = append(self.items, x)
	self.swim(len(self.items) - 1)
}

func (self *MyPriorityQueueInt) Pop() (int, error) {
	if len(self.items) <= 1 {
		return 0, errors.New("Empty queue!")
	}
	x := self.items[1]
	self.items[1] = self.items[len(self.items)-1]
	self.items = self.items[:len(self.items)-1]
	self.sink(1)
	return x, nil
}

func (self *MyPriorityQueueInt) Length() int {
	return len(self.items) - 1
}

// 三叉堆？不能实现，因为元素不能在正整数轴上完美铺开 --------------------------------------------------
//type MyPriorityQueue3Int struct {
//	items []int
//}
//
//func (self *MyPriorityQueue3Int) swim(x int) {
//	// 从下向上浮，只需要跟父节点比较，一般用在插入
//	for f := x / 3; (x > 1) && (self.items[x] > self.items[f]); {
//		self.items[x], self.items[f] = self.items[f], self.items[x]
//		x = f
//		f = x / 3
//	}
//}
//
//func (self *MyPriorityQueue3Int) sink(x int) {
//	// 从上向下沉，
//
//	for l:= len(self.items)-1; x*3 <= l; {
//		//在它和它的4个子节点中，寻找最大的那个跟它本身交换
//		m :=x
//		if n := x*3 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+1 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+2 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+3 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+4 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+5 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+6 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+7 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//		if n := x*3+8 ; n <= l {
//			if self.items[n] > self.items[m] {
//				m = n
//			}
//		}
//
//		if m == x { //最大值是本身的话，就停止下沉了
//			break
//		} else { // 最大值是子节点，那就继续下沉
//			self.items[x], self.items[m] = self.items[m], self.items[x]
//			x = m
//		}
//	}
//}
//
//func (self *MyPriorityQueue3Int) Push(x int) {
//	if len(self.items) == 0 {
//		self.items = append(self.items, 0)
//	}
//	self.items = append(self.items, x)
//	self.swim(len(self.items) - 1)
//}
//
//func (self *MyPriorityQueue3Int) Pop() (int, error) {
//	if len(self.items) <= 1 {
//		return 0, errors.New("Empty queue!")
//	}
//	x := self.items[1]
//	self.items[1] = self.items[len(self.items)-1]
//	self.items = self.items[:len(self.items)-1]
//	self.sink(1)
//	return x, nil
//}
//
//func (self *MyPriorityQueue3Int) Length() int {
//	return len(self.items) - 1
//}
