package p001x

import "unsafe"

var RomanTable = [7]byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
func intToRoman(num int) string {
	// 题目规定 1<= num <=3999
	var r int
	var result = make([]byte, 16)
	var p int = 15
	for i := 0; i < 4; i++ {
		r = i*2
		switch num % 10 {
		case 1:
			result[p] = RomanTable[r]
			p--
		case 2:
			result[p-1], result[p] = RomanTable[r], RomanTable[r]
			p -= 2
		case 3:
			result[p-2], result[p-1], result[p] = RomanTable[r], RomanTable[r], RomanTable[r]
			p -= 3
		case 4:
			result[p-1], result[p] = RomanTable[r], RomanTable[r+1]
			p -= 2
		case 5:
			result[p] = RomanTable[r+1]
			p--
		case 6:
			result[p-1], result[p] = RomanTable[r+1], RomanTable[r]
			p -= 2
		case 7:
			result[p-2], result[p-1], result[p] = RomanTable[r+1], RomanTable[r], RomanTable[r]
			p -= 3
		case 8:
			result[p-3], result[p-2], result[p-1], result[p] = RomanTable[r+1], RomanTable[r], RomanTable[r], RomanTable[r]
			p -= 4
		case 9:
			result[p-1], result[p] = RomanTable[r], RomanTable[r+2]
			p -= 2
		}
		num /= 10
	}
	return (*(*string)(unsafe.Pointer(&(result))))[p+1:]
}


//var RomanTable = [7]byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
//func intToRoman(num int) string {
//	// 题目规定 1<= num <=3999
//	var romanTable []byte
//	var result = make([]byte, 16)
//	var p int = 15
//	for i := 0; i < 4; i++ {
//		romanTable = RomanTable[i*2:]
//		switch num % 10 {
//		case 1:
//			result[p] = romanTable[0]
//			p--
//		case 2:
//			result[p-1], result[p] = romanTable[0], romanTable[0]
//			p -= 2
//		case 3:
//			result[p-2], result[p-1], result[p] = romanTable[0], romanTable[0], romanTable[0]
//			p -= 3
//		case 4:
//			result[p-1], result[p] = romanTable[0], romanTable[1]
//			p -= 2
//		case 5:
//			result[p] = romanTable[1]
//			p--
//		case 6:
//			result[p-1], result[p] = romanTable[1], romanTable[0]
//			p -= 2
//		case 7:
//			result[p-2], result[p-1], result[p] = romanTable[1], romanTable[0], romanTable[0]
//			p -= 3
//		case 8:
//			result[p-3], result[p-2], result[p-1], result[p] = romanTable[1], romanTable[0], romanTable[0], romanTable[0]
//			p -= 4
//		case 9:
//			result[p-1], result[p] = romanTable[0], romanTable[2]
//			p -= 2
//		}
//		num /= 10
//	}
//	return (*(*string)(unsafe.Pointer(&(result))))[p+1:]
//}



//var RomanTable = [7]string{"I", "V", "X", "L", "C", "D", "M"}
//func intToRoman(num int) string {
//	// 题目规定 1<= num <=3999
//	var romanTable []string
//	var result = make([]string, 4)
//	for i := 0; i < 4; i++ {
//		romanTable = RomanTable[i*2:]
//		switch num % 10 {
//		case 1:
//			result[3-i] = fmt.Sprint(romanTable[0])
//		case 2:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0])
//		case 3:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[0], romanTable[0])
//		case 4:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[1])
//		case 5:
//			result[3-i] = fmt.Sprint(romanTable[1])
//		case 6:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0])
//		case 7:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0])
//		case 8:
//			result[3-i] = fmt.Sprint(romanTable[1], romanTable[0], romanTable[0], romanTable[0])
//		case 9:
//			result[3-i] = fmt.Sprint(romanTable[0], romanTable[2])
//		}
//		num /= 10
//	}
//	return strings.Join(result, "")
//}
