package p003x

import (
	"learnAlgo"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列
重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，
则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var top int
	for top = len(nums) - 1; top > 0; top-- {
		if nums[top-1] < nums[top] {
			break
		}
	}
	if top == 0 {
		swapSortInt(nums)
		return
	}
	exchId := indexOfMinGreater(nums[top:], nums[top-1]) + top
	nums[top-1], nums[exchId] = nums[exchId], nums[top-1]
	swapSortInt(nums[top:])
	return
}

// 二分查找
func indexOfMinGreater(nums []int, target int) int {
	var gt, lt int
	for gt, lt = 0, len(nums); lt-gt > 1; {
		mid := (gt + lt) / 2
		if nums[mid] > target {
			gt = mid
		} else {
			lt = mid
		}
	}
	return gt
}

// 交换排序
func swapSortInt(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation1(nums []int) {
	if len(nums) < 2 {
		return
	}
	var top int
	// 找到顶点，即降序排列的头部
	for top = len(nums) - 1; top > 0; top-- {
		if nums[top-1] < nums[top] {
			break
		}
	}
	if top == 0 {
		learnAlgo.QuickSortInt(nums) // 快速排序函数
		return
	}
	// 在右边数组中找到一个值去跟左边交换
	var targetNum, updateIndex int = nums[top-1], top
	for i := len(nums) - 1; i > top; i-- {
		if nums[i] > targetNum && nums[i] < nums[updateIndex] {
			updateIndex = i
		}
	}
	nums[top-1], nums[updateIndex] = nums[updateIndex], nums[top-1]
	learnAlgo.QuickSortInt(nums[top:]) // 快速排序函数
	return
}
