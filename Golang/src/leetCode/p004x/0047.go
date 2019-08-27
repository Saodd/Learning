package p004x

import "sort"

/*
47. 全排列 II

给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permuteUnique(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }

    sort.Ints(nums)
    len0047 = len(nums)
    nums0047 = nums
    temp0047 = make([]int, len(nums))
    used0047 = make([]bool, len(nums))
    result0047 = make([][]int, 0, 6)

    recurPermuteUnique(0)
    return result0047
}

var nums0047 []int
var temp0047 []int
var len0047 int
var result0047 [][]int
var used0047 []bool

func recurPermuteUnique(pos int) {
    if pos == len0047 {
        newSolution := make([]int, len(temp0047))
        copy(newSolution, temp0047)
        result0047 = append(result0047, newSolution)
        return
    }
    for i := 0; i < len0047; i++ {
        if !used0047[i] {
            if i > 0 && !used0047[i-1] && nums0047[i] == nums0047[i-1] {
                continue
            }
            temp0047[pos] = nums0047[i]
            used0047[i] = true
            recurPermuteUnique(pos + 1)
            used0047[i] = false
        }
    }
}

// // ---------------------------------------
//func permuteUnique(nums []int) [][]int {
//    if len(nums)==0{
//        return [][]int{}
//    }
//
//    len0047 = len(nums) - 1
//    temp0047 = nums
//    sort.Ints(temp0047)
//
//    result0047 = make([][]int, 0, 5)
//
//    recurPermuteUnique(0)
//    return result0047
//}
//
//var temp0047 []int
//var len0047 int
//var result0047 [][]int
//
//func recurPermuteUnique(pos int) {
//    if pos == len0047 {
//        newSolution := make([]int, len(temp0047))
//        copy(newSolution, temp0047)
//        result0047 = append(result0047, newSolution)
//        return
//    }
//
//    recurPermuteUnique(pos + 1)
//    for i := pos + 1; i <= len0047; i++ {
//        for j:=pos; j<i; j++{
//            if temp0047[j]==temp0047[i]{
//                goto NEXTI
//            }
//        }
//        temp0047[pos], temp0047[i] = temp0047[i], temp0047[pos]
//        recurPermuteUnique(pos + 1)
//        temp0047[pos], temp0047[i] = temp0047[i], temp0047[pos]
//        NEXTI:
//    }
//}
