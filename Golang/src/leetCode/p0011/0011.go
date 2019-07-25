package p0011

func maxArea(height []int) int {
	// 题目规定：参数数组长度最小为2
	var maxA int = 0 // 假设面积只用int就可以表示
	for l, r, thisA, lh, rh := 0, len(height)-1, 0, height[0], height[len(height)-1]; l < r; {
		if lh > rh {
			thisA = rh * (r - l)
			if thisA > maxA {
				maxA = thisA
			}
			r--
			rh = height[r]
		} else {
			thisA = lh * (r - l)
			if thisA > maxA {
				maxA = thisA
			}
			l++
			lh = height[l]
		}
	}
	return maxA
}
