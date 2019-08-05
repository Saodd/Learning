package learnAlgo

import (
	"fmt"
	"learnAlgo/p004_structure"
	"learnAlgo/tools"
	"time"
)

func isSorted(li []int) bool { // 检测是否递增，即前面的数必须小于等于后面的
	for i := 0; i < len(li)-1; i++ {
		if li[i] > li[i+1] {
			return false
		}
	}
	return true
}

// 选择排序算法 ------------------------------------------------------------
// 10万个int用时：5.4108552s,5.3938423s,5.4068515s
// 20万个int用时:21.5873776s,21.6013872s,内存总共2.6M
func sort_Select(li []int) {
	minindex := 0 // 健壮性先不考虑
	length := len(li)
	var i, ii int
	for i = range li {
		minindex = i
		// 在i之后的剩余部分中寻找最小值对应的位置
		for ii = i + 1; ii < length; ii++ {
			if li[ii] < li[minindex] {
				minindex = ii
			}
		}
		// 遍历完成后把最小值交换到i的位置
		li[i], li[minindex] = li[minindex], li[i]
	}
}

// 快速排序 ---------------------------------------------------------------
func QuickSortInt(li []int) {
	if len(li)==0{
		return
	}
	quickSortInt(li, 0, len(li)-1)
}

func quickSortInt(li []int, lo, hi int) {
	stack := quickSortStackInt{}
	stack.Push(lo, hi)
	for stack.Len() > 0 {
		x, y, _ := stack.Pop()
		mid := quickSortIntPartition(li, x, y)
		if mid-x > 15 {
			stack.Push(x, mid-1)
		} else if mid-x > 1 {
			quickSortIntSelectSort(li, x, mid-1)
		}
		if y-mid > 15 {
			stack.Push(mid+1, y)
		} else if y-mid > 1 {
			quickSortIntSelectSort(li, mid+1, y)
		}
	}
}

func quickSortIntPartition(li []int, lo, hi int) (mid int) {
	l, r := lo, hi
	//li[lo], li[(lo+hi)/2] = li[(lo+hi)/2], li[lo]
	midValue := li[lo]
	for l < r {
		for l <= hi { // 找大的
			if li[l] > midValue {
				break
			}
			l++
		}
		for r >= lo {
			if li[r] <= midValue {
				break
			}
			r--
		}
		if l < r {
			li[l], li[r] = li[r], li[l]
		} else {
			break
		}
	}
	li[lo], li[r] = li[r], li[lo]
	return r
}

func quickSortIntSelectSort(li []int, lo, hi int) {
	var min int
	for ; lo < hi; lo++ {
		min = lo
		for i := lo + 1; i <= hi; i++ {
			if li[i] < li[min] {
				min = i
			}
		}
		if lo != min {
			li[lo], li[min] = li[min], li[lo]
		}
	}
}

type quickSortStackInt struct {
	items []int
}

func (self *quickSortStackInt) Len() int {
	return len(self.items)
}

func (self *quickSortStackInt) Push(x, y int) {
	self.items = append(self.items, x, y)
}
func (self *quickSortStackInt) Pop() (int, int, error) {
	l := len(self.items)
	//if l < 2 {
	//	return 0, 0, errors.New("没有元素可以取出")
	//}
	x, y := self.items[l-2], self.items[l-1]
	self.items = self.items[:l-2]
	return x, y, nil
}

// 书上的算法，只有partition不同
// 我vs书 时间：1.94/1.81, 1.94/1.81, 1.95/1.81，我的比书上的慢8%……内存1578M完全相同
type sort_Quick_Book struct {
}

func (self sort_Quick_Book) sort(li []int) {
	self._sort(li, 0, len(li)-1)
}
func (self sort_Quick_Book) _sort(li []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := self.partition(li, lo, hi)
	self._sort(li, lo, mid-1)
	self._sort(li, mid+1, hi)
}
func (self sort_Quick_Book) partition(li []int, lo, hi int) (mid int) {
	i, j := lo, hi+1 // 从结果看来，在递归函数中定义变量，并不会带来额外的内存开销
	v := li[lo]
	for {
		for i++; li[i] < v; i++ {
			if i == hi {
				break
			}
		}
		for j--; v < li[j]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		li[i], li[j] = li[j], li[i]
	}
	li[lo], li[j] = li[j], li[lo]
	return j
}

// 基于书上的算法，加上插入排序进行优化
type sort_Quick_Book_Insert struct {
	sort_Quick_Book
	criterion int
}

func (self sort_Quick_Book_Insert) _sort(li []int, lo, hi int) {
	if lo >= hi {
		return
	}
	if hi-lo >= self.criterion {
		self.insert(li, lo, hi)
		return
	}
	mid := self.partition(li, lo, hi)
	self._sort(li, lo, mid-1)
	self._sort(li, mid+1, hi)
}
func (self sort_Quick_Book_Insert) insert(li []int, lo, hi int) {
	var temp, j, k int
	for i := lo + 1; i <= hi; i++ {
		temp = li[i]
		if temp < li[i-1] {
			// 更小的话就往前挪
			for j = lo; j < i; j++ {
				// j是这个更小的数字需要往前插入的位置
				if temp < li[j] {
					break
				}
			}
			// [....j.....i]把i移动到j
			for k = i; k > j; k-- {
				li[k] = li[k-1]
			}
			li[j] = temp
		}
	}
}

// 对比运行 ---------------------------------------------------------------
func Main0003() {
	data := tools.Gen_ints_list(2000000)
	data2 := make([]int, len(data))
	starttime := time.Now()
	t := time.Microsecond
	// ---------------------------
	//for i := range data {
	//	data2[i] = data[i]
	//}
	//fmt.Println("Checking input for Quick-sort:", isSorted(data2))
	//starttime = time.Now()
	//sort_Quick{}.sort(data2)
	//t = time.Since(starttime)
	//fmt.Printf("Quick-sort: %v,   used time: %v seconds.\n", isSorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for Quick-sort-Book:", isSorted(data2))
	starttime = time.Now()
	sort_Quick_Book{}.sort(data2)
	t = time.Since(starttime)
	fmt.Printf("Quick-sort-Book: %v,   used time: %v seconds.\n", isSorted(data2), t)
	// ---------------------------
	//for i := range data{
	//	data2[i] = data[i]
	//}
	//fmt.Println("Checking input for Select-sort:", isSorted(data2))
	//starttime = time.Now()
	//sort_Select(data2)
	//t = time.Since(starttime)
	//fmt.Printf("Select-sort: %v,   used time: %v seconds.\n", isSorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for sort_Quick_Book_Insert(5):", isSorted(data2))
	starttime = time.Now()
	sort_Quick_Book_Insert{}.sort(data2)
	t = time.Since(starttime)
	fmt.Printf("sort_Quick_Book_Insert(5): %v,   used time: %v seconds.\n", isSorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for PriorityQueue:", isSorted(data2))
	starttime = time.Now()
	q := &p004_structure.MyPriorityQueueInt{}
	for i := range data2 {
		q.Push(data2[i])
	}
	t = time.Since(starttime)
	for i := q.Length() - 1; i >= 0; i-- {
		data2[i], _ = q.Pop()
	}
	fmt.Printf("PriorityQueue: %v,   used time: %v seconds.\n", isSorted(data2), t)

	// ---------------------------三叉堆是不能实现的
	//for i := range data {
	//	data2[i] = data[i]
	//}
	//fmt.Println("Checking input for PriorityQueue3:", isSorted(data2))
	//
	//starttime = time.Now()
	//q3 := &p004_structure.MyPriorityQueue3Int{}
	//for i := range data2 {
	//	q3.Push(data2[i])
	//}
	//t = time.Since(starttime)
	//for i := q3.Length()-1; i >=0; i-- {
	//	data2[i], _ = q3.Pop()
	//}
	//fmt.Printf("PriorityQueue3: %v,   used time: %v seconds.\n", isSorted(data2), t)

}
