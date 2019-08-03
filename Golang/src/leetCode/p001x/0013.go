package p001x

var rtiTable = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int
	for i, l, left, right:=0, len(s)-1, rtiTable[s[0]], 0; i<l; i++{
		right = rtiTable[s[i+1]]
		if  left < right{
			result -= left
		}else {
			result += left
		}
		left = right
	}
	return result + rtiTable[s[len(s)-1]]
}

//func romanToInt(s string) int {
//	if len(s)==1{
//		return rtiTable[s[0]]
//	}
//	var result int
//	for i, l:=1, len(s); i<l; i++{
//		temp := rtiTable[s[i-1]]
//		if  temp < rtiTable[s[i]]{
//			result -= temp
//		}else {
//			result += temp
//		}
//	}
//	return result + rtiTable[s[len(s)-1]]
//}