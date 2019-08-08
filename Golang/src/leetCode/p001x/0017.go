package p001x

var dialDigitMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

var result0017 []string
var input0017 string
var length0017 int
var temp0017 []byte

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result0017 = []string{}
	input0017 = digits
	length0017 = len(digits)
	temp0017 = make([]byte, length0017)
	recur0017(0)
	return result0017
}

func recur0017(pos int) {
	if pos == length0017-1 {
		for _, c := range dialDigitMap[input0017[pos]] {
			temp0017[pos] = c
			result0017 = append(result0017, string(temp0017))
		}

	} else {
		for _, c := range dialDigitMap[input0017[pos]] {
			temp0017[pos] = c
			recur0017(pos + 1)
		}
	}
}
