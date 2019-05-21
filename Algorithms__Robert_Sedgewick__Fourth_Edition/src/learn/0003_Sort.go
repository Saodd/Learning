package learn

import (
	"fmt"
	"mytools"
	"time"
)

func is_Sorted(li []int) bool {
	for i := 0; i < len(li)-1; i++ {
		if li[i] > li[i+1] {
			return false
		}
	}
	return true
}

// 选择排序算法,10万个int用时：5.4108552s,5.3938423s,5.4068515s
// 20万个int用时:21.5873776s,21.6013872s
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

func Main0003() {
	data := mytools.Gen_ints_list(200000)
	starttime := time.Now()
	sort_Select(data)
	t := time.Since(starttime)
	fmt.Printf("Success: %v,   used time: %v seconds.", is_Sorted(data), t)
}
