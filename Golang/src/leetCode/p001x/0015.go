package p001x

import (
	"learnAlgo"
)

func threeSum(nums []int) [][]int {
	var result = [][]int{}
	learnAlgo.QuickSortInt(nums)
	var lp, rp, sum int
	for i, le, iNum := 0, len(nums)-1, -1<<31; i <= le && iNum <= 0; i++ {
		if nums[i] == iNum {
			continue
		}
		iNum = nums[i]
		lp, rp = i+1, le
		for lp < rp {
			sum = iNum + nums[lp] + nums[rp]
			switch {
			case sum == 0:
				result = append(result, []int{iNum, nums[lp], nums[rp]})
				for lp < rp && nums[lp] == nums[lp+1] {
					lp++
				}
				for lp < rp && nums[rp] == nums[rp-1] {
					rp--
				}
				lp++
				rp--
			case sum < 0:
				lp++
			case sum > 0:
				rp--
			}
		}
	}
	return result
}

//func threeSum(nums []int) [][]int {
//	var result = [][]int{}
//	sort.Ints(nums)
//	var lp, rp, lNum, rTarget int
//	for i, le, iNum := 0, len(nums), math.MinInt32; i < le && iNum<=0; i++ {
//		if nums[i] == iNum {
//			continue
//		}
//		iNum = nums[i]
//		lp, rp, lNum = i+1, le-1,math.MinInt32
//		for ; lp < rp; lp++ {
//			if nums[lp] == lNum {
//				continue
//			}
//			lNum = nums[lp]
//			rTarget = - iNum - lNum
//			for rTarget < nums[rp] && lp < rp{
//				rp--
//			}
//			if lp < rp && rTarget ==  nums[rp] {
//				result = append(result, []int{iNum, nums[lp], nums[rp]})
//			}
//
//		}
//
//	}
//	return result
//}

//func threeSum(nums []int) [][]int {
//	var twoSumCache = map[int]map[int]bool{}
//	var numCount = map[int]int{}
//	var result = [][]int{}
//	// Calc twoSum
//	for i, le, lnum, rnum := 0, len(nums), 0, 0; i < le; i++ {
//		lnum = nums[i]
//		numCount[lnum] += 1
//		for j := i + 1; j < le; j++ {
//			rnum = nums[j]
//			if lnum <= rnum {
//				if _, ok := twoSumCache[lnum+rnum]; ok {
//					twoSumCache[lnum+rnum][lnum] = true
//				} else {
//					twoSumCache[lnum+rnum] = map[int]bool{lnum: true}
//				}
//			} else {
//				if _, ok := twoSumCache[lnum+rnum]; ok {
//					twoSumCache[lnum+rnum][rnum] = true
//				} else {
//					twoSumCache[lnum+rnum] = map[int]bool{rnum: true}
//				}
//			}
//
//		}
//	}
//
//	for _, theNum := range nums {
//		if numSet, ok := twoSumCache[-theNum]; ok {
//			for k, _ := range numSet {
//				if theNum < k {
//					result = append(result, []int{theNum, k, -theNum - k})
//					delete(numSet, k)
//				} else if theNum == k {
//					if theNum == 0 && numCount[theNum] > 2 {
//						result = append(result, []int{theNum, k, -theNum - k})
//						delete(numSet, k)
//					} else if theNum != 0 && numCount[theNum] > 1 {
//						result = append(result, []int{theNum, k, -theNum - k})
//						delete(numSet, k)
//					}
//				}
//			}
//			delete(twoSumCache, -theNum)
//		}
//	}
//
//	return result
//}
