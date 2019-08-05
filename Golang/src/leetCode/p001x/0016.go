package p001x

import "learnAlgo"

func threeSumClosest(nums []int, target int) int {
    learnAlgo.QuickSortInt(nums)
    var minDiff int = 1<<31 - 1
    var theSum, theDiff int
    var twoTarget int
    var l, r int
    for i, le := 0, len(nums); i < le; i++ {
        twoTarget = target - nums[i]
        l, r = i+1, le-1
        for l < r {
            theDiff = twoTarget - nums[l] - nums[r]
            if abs(theDiff) < minDiff {
                minDiff, theSum = abs(theDiff), nums[i]+nums[l]+nums[r]
            }
            if minDiff == 0 {
                return theSum
            }
            if theDiff > 0 {
                l++
            } else {
                r--
            }
        }
    }
    return theSum
}

func threeSumClosest1(nums []int, target int) int {
   const minInt = - 1 << 31
   const maxInt = 1<<31 - 1
   // 题目没有说错误处理，姑且假设len(nums)>=3
   learnAlgo.QuickSortInt(nums)
   var iNum int = minInt
   var minDiff = maxInt
   var l, r, sum, theSum int
   for i, le := 0, len(nums); i < le && iNum < target; i++ {
       if nums[i] == iNum {
           continue
       }
       iNum = nums[i]
       l, r = i+1, le-1
       for l < r {
           if nums[l] == nums[l+1] {
               sum = iNum + 2*nums[l]
               if abs(target-sum) < minDiff {
                   minDiff, theSum = abs(target-sum), sum
               }
               for l++;l < r && nums[l] == nums[l+1] ; l++ {
               }
           }
           if nums[r] == nums[r-1] {
               sum = iNum + 2*nums[r]
               if abs(target-sum) < minDiff {
                   minDiff, theSum = abs(target-sum), sum
               }
               for r--;l < r && nums[r] == nums[r-1]; r-- {
               }
           }
           sum = iNum + nums[l] + nums[r]
           if abs(target-sum) < minDiff {
               minDiff, theSum = abs(target-sum), sum
           }
           if minDiff == 0 {
               goto STOPLOOP
           }
           if sum > target {
               r--
           } else {
               l++
           }
       }
   }
STOPLOOP:
   return theSum
}

func abs(x int) int {
    if x >= 0 {
        return x
    }
    return -x
}
