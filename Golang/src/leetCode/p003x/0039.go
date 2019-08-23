package p003x

import "sort"

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 

示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]

示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    cd = candidates
    len0039 = len(cd)
    sort.Ints(cd)
    temp0039 = make([]int, target/cd[0]+1)
    result0039 = [][]int{}

    recurCombinationSum(0, 0, target)
    return result0039
}

var cd []int
var len0039 int
var temp0039 []int
var result0039 [][]int

func recurCombinationSum(cdPos, tempPos int, target int) {
    for ; cdPos < len0039; cdPos++ {
        if cd[cdPos] > target {
            return
        }
        temp0039[tempPos] = cd[cdPos]
        if cd[cdPos] == target {
            newSolution := make([]int, tempPos+1)
            copy(newSolution, temp0039)
            result0039 = append(result0039, newSolution)
            return
        } else {
            recurCombinationSum(cdPos, tempPos+1, target-cd[cdPos])
        }
    }
}
