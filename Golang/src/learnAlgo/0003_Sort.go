package learnAlgo

import (
	"fmt"
	"learnAlgo/p004_structure"
	"learnAlgo/tools"
	"time"
)

func is_Sorted(li []int) bool { // 检测是否递增，即前面的数必须小于等于后面的
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
// 20万个int用时：16ms, 15ms, 16ms, 单位是毫秒，比选择排序快1000倍……
// 200万个int用时：167ms, 171ms, 168ms, 单位还是毫秒……
// 2000万个int用时：1.94s, 1.94s, 1.94s
// 2亿个int用时：21.99s, 22.02s, 内存1578M
// 几乎是线性的时间复杂度。太强了！的确是把选择排序不可能完成的任务完成了。
type sort_Quick struct {
}

func (self sort_Quick) sort(li []int) {
	self._sort(li, 0, len(li)-1)
}
func (self sort_Quick) _sort(li []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := self.partition(li, lo, hi)
	self._sort(li, lo, mid-1)
	self._sort(li, mid+1, hi)
}
func (self sort_Quick) partition(li []int, lo, hi int) int { // 自己琢磨的算法，想了2个小时
	// mid和hi去循环
	mid := lo
	midv := li[lo]
	for mid < hi {
		for mid < hi { // 寻找左边更大的元素
			if li[mid] > midv {
				break
			}
			mid++
		}
		for mid < hi { // 寻找右边更小的元素
			if li[hi] < midv {
				break
			}
			hi--
		}
		if mid < hi { // 没有相遇就交换二者
			li[mid], li[hi] = li[hi], li[mid]
		} else if li[mid] < li[lo] { // 相遇了那就再检查一下相遇的位置，找到更小的
			li[mid], li[lo] = li[lo], li[mid]
			lo = mid
		} else { // 相遇了那就再检查一下相遇的位置，找到更小的
			li[mid-1], li[lo] = li[lo], li[mid-1]
			lo = mid - 1
		}
	}
	return lo
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
	//fmt.Println("Checking input for Quick-sort:", is_Sorted(data2))
	//starttime = time.Now()
	//sort_Quick{}.sort(data2)
	//t = time.Since(starttime)
	//fmt.Printf("Quick-sort: %v,   used time: %v seconds.\n", is_Sorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for Quick-sort-Book:", is_Sorted(data2))
	starttime = time.Now()
	sort_Quick_Book{}.sort(data2)
	t = time.Since(starttime)
	fmt.Printf("Quick-sort-Book: %v,   used time: %v seconds.\n", is_Sorted(data2), t)
	// ---------------------------
	//for i := range data{
	//	data2[i] = data[i]
	//}
	//fmt.Println("Checking input for Select-sort:", is_Sorted(data2))
	//starttime = time.Now()
	//sort_Select(data2)
	//t = time.Since(starttime)
	//fmt.Printf("Select-sort: %v,   used time: %v seconds.\n", is_Sorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for sort_Quick_Book_Insert(5):", is_Sorted(data2))
	starttime = time.Now()
	sort_Quick_Book_Insert{}.sort(data2)
	t = time.Since(starttime)
	fmt.Printf("sort_Quick_Book_Insert(5): %v,   used time: %v seconds.\n", is_Sorted(data2), t)
	// ---------------------------
	for i := range data {
		data2[i] = data[i]
	}
	fmt.Println("Checking input for PriorityQueue:", is_Sorted(data2))
	starttime = time.Now()
	q := &p004_structure.MyPriorityQueueInt{}
	for i := range data2 {
		q.Push(data2[i])
	}
	t = time.Since(starttime)
	for i := q.Length() - 1; i >= 0; i-- {
		data2[i], _ = q.Pop()
	}
	fmt.Printf("PriorityQueue: %v,   used time: %v seconds.\n", is_Sorted(data2), t)

	// ---------------------------三叉堆是不能实现的
	//for i := range data {
	//	data2[i] = data[i]
	//}
	//fmt.Println("Checking input for PriorityQueue3:", is_Sorted(data2))
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
	//fmt.Printf("PriorityQueue3: %v,   used time: %v seconds.\n", is_Sorted(data2), t)

}
