package p003x

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4

示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func search(nums []int, target int) int {
    switch len(nums) {
    case 0:
        return -1
    case 1:
        if target == nums[0] {
            return 0
        }
        return -1
    }

    var minId int = indexOfMin(nums)
    var maxId int
    if minId == 0 {
        maxId = len(nums) - 1
    } else {
        maxId = minId - 1
    }

    //fmt.Println(maxId)
    switch {
    case target > nums[maxId] || target < nums[minId]:
        return -1
    case target >= nums[0]:
        return binarySearch(nums[:maxId+1], target)
    default:
        if got := binarySearch(nums[minId:], target); got == -1 {
            return -1
        } else {
            return got + minId
        }
    }
}

func indexOfMin(nums []int) int {
    // 假设1：nums长度大于1
    // 假设2：nums没有重复元素
    var firstElem = nums[0]
    for left, right := 0, len(nums); ; {
        mid := (left + right) / 2
        if mid == left {
            if left == len(nums)-1{
                return 0
            }
            return left + 1
        }
        if nums[mid] > firstElem {
            left = mid
        } else {
            right = mid
        }
    }
}

func binarySearch(nums []int, target int) int {
    // 假设1：nums已经升序排列
    // 假设2：nums没有重复元素
    if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
        return -1
    }
    for left, right := 0, len(nums); ; {
        mid := (left + right) / 2
        switch {
        case nums[mid] < target:
            if left == mid {
                return -1
            }
            left = mid
        case nums[mid] > target:
            right = mid
        default:
            return mid
        }
    }
}
