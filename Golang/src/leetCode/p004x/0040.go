package p004x

import "sort"

/*
给定一个数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum2(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    temp0040 = make([]int, len(candidates))

    // 对原数组计数
    counter0040 = map[int]int{}
    for _, num := range candidates {
        counter0040[num]++
    }

    // 得到一个无重复元素的数组，并排序
    count :=0
    for key, _ := range counter0040 {
        candidates[count] = key
        count++
    }
    len0040 = count
    cd = candidates[:count]
    sort.Ints(cd)

    // 初始化剩下的全局变量
    result0040 = [][]int{}

    recurCombinationSum2(0, 0, target)
    return result0040
}

var cd []int
var len0040 int
var temp0040 []int
var result0040 [][]int
var counter0040 map[int]int

func recurCombinationSum2(cdPos, tempPos int, target int) {
    // 在无重复元素的数组上前进，跳过的值就是相当于是0个的
    for ; cdPos < len0040; cdPos++ {
        num := cd[cdPos]
        tg := target
        // 根据个数，分别求1个，2个，3个……的情况，并递归
        for dupl := 0; dupl < counter0040[num]; dupl++ {
            temp0040[tempPos+dupl] = num
            tg -= num

            if tg == 0 {
                newSolution := make([]int, tempPos+dupl+1)
                copy(newSolution, temp0040)
                result0040 = append(result0040, newSolution)
            } else if tg > 0 {
                recurCombinationSum2(cdPos+1, tempPos+dupl+1, tg)
            } else {
                break
            }
        }
    }
}
