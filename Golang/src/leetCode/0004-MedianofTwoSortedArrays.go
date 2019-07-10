/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func Main0004() {
	var k [2][]int
	var v float64

	k, v = [2][]int{[]int{1, 3}, []int{2}}, 2
	fmt.Printf("Input: '%v', Answer: %v, Yours: %v \n", k, v, findMedianSortedArrays(k[0], k[1]))
	k, v = [2][]int{[]int{1, 2}, []int{3, 4}}, 2.5
	fmt.Printf("Input: '%v', Answer: %v, Yours: %v \n", k, v, findMedianSortedArrays(k[0], k[1]))
	k, v = [2][]int{[]int{1, 2, 3, 4}, []int{0, 5, 7, 9}}, 3.5
	fmt.Printf("Input: '%v', Answer: %v, Yours: %v \n", k, v, findMedianSortedArrays(k[0], k[1]))
	k, v = [2][]int{[]int{1, 2}, []int{1, 1}}, 1
	fmt.Printf("Input: '%v', Answer: %v, Yours: %v \n", k, v, findMedianSortedArrays(k[0], k[1]))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	lo, hi := 0, len1*2
	var Lmax1, Rmin1, Lmax2, Rmin2, c1, c2 int
	for lo <= hi {
		c1 = (lo + hi) / 2
		c2 = len1 + len2 - c1

		if c1 <= 0 {
			Lmax1 = intsets.MinInt
		} else {
			Lmax1 = nums1[(c1-1)/2]
		}
		if c1 >= len1*2 {
			Rmin1 = intsets.MaxInt
		} else {
			Rmin1 = nums1[c1/2]
		}
		if c2 <= 0 {
			Lmax2 = intsets.MinInt
		} else {
			Lmax2 = nums2[(c2-1)/2]
		}
		if c2 >= len2*2 {
			Rmin2 = intsets.MaxInt
		} else {
			Rmin2 = nums2[c2/2]
		}

		if (Rmin2 >= Lmax1) && (Rmin1 >= Lmax2) {
			break
		}
		if Lmax1 > Rmin2 {
			hi = c1
		} else {
			lo = c1 + 1
		}
	}
	return float64(mymax(Lmax1, Lmax2)+mymin(Rmin1, Rmin2)) / 2
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func mymin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
小结：

三个很重要的思路：

1. 把任意**奇偶长的数组**转换为**偶数长的数组**：
	在每个元素前面插入一个虚拟的null值，新数组的元素下标为[0,...2n]，使用list[x/2]即可访问(新下标x)对应的老数组的元素。

2. 无穷小/无穷大的应用
	当数组执行到边界（0或len(list)）的时候，使用无穷小/无穷大作为返回值，参与比较，这样可以大大的简化代码。

3. 二分法的核心思想：
	var p int
	while lo <= hi {

		// do something

		if ...{
			hi = p
		} else {
			lo = p+1
		}
	}

*/
