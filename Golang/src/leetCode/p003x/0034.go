package p003x

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}
    if left := searchRangeLeft(nums, target); left != -1 {
        result[0] = left
        result[1] = searchRangeRight(nums[left:], target) + left
    }
    return result
}

func searchRangeLeft(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    var lo, hi, mid int
    for lo, hi = 0, len(nums); lo < hi; {
        mid = lo + (hi-lo)/2
        if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    if hi == len(nums) {
        return -1
    }
    if nums[hi] == target {
        return hi
    }
    return -1
}

func searchRangeRight(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    var lo, hi, mid int
    for lo, hi = 0, len(nums); lo < hi; {
        mid = lo + (hi-lo)/2
        if nums[mid] > target {
            hi = mid
        } else {
            lo = mid + 1
        }
    }

    if hi == 0 {
        return -1
    }
    if nums[hi-1] == target {
        return hi - 1
    }
    return -1
}

func searchRange1(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    var lo, hi int
    for lo, hi = 0, len(nums)-1; lo <= hi; {
        mid := (lo + hi) / 2
        switch {
        case nums[mid] < target:
            lo = mid + 1
        case nums[mid] > target:
            hi = mid - 1
        default:
            return []int{searchRangeLow(nums[lo:mid+1], target) + lo, searchRangeHigh(nums[mid:hi+1], target) + mid}
        }
    }
    return []int{-1, -1}
}

func searchRangeLow(nums []int, target int) int {
    fmt.Println(nums)
    if len(nums) == 0 {
        return 0
    }

    var lo, hi int
    for lo, hi = 0, len(nums)-1; lo < hi; {
        mid := (lo + hi) / 2
        if nums[mid] == target {
            hi = mid
        } else {
            lo = mid + 1
        }
    }

    if nums[hi] == target {
        return hi
    }
    return -1
}

func searchRangeHigh(nums []int, target int) int {
    fmt.Println(nums)
    if len(nums) == 0 {
        return 0
    }

    var lo, hi int
    for lo, hi = 0, len(nums); lo < hi; {
        mid := (lo + hi) / 2
        if nums[mid] != target {
            hi = mid
        } else {
            if lo == mid {
                break
            }
            lo = mid
        }
    }
    if nums[lo] == target {
        return lo
    }
    return -1
}
