package learnTour

/*
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
*/
import (
	"fmt"
	"strings"
)

// Test runs a test suite against f.
func Test(f func(string) map[string]int) {
	ok := true
	for _, c := range testCases {
		got := f(c.in)
		if len(c.want) != len(got) {
			ok = false
		} else {
			for k := range c.want {
				if c.want[k] != got[k] {
					ok = false
				}
			}
		}
		if !ok {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
				c.in, got, c.want)
			break
		}
		fmt.Printf("PASS\n f(%q) = \n  %#v\n", c.in, got)
	}
}

var testCases = []struct {
	in   string
	want map[string]int
}{
	{"I am learning Go!", map[string]int{
		"I": 1, "am": 1, "learning": 1, "Go!": 1,
	}},
	{"The quick brown fox jumped over the lazy dog.", map[string]int{
		"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
		"over": 1, "the": 1, "lazy": 1, "dog.": 1,
	}},
	{"I ate a donut. Then I ate another donut.", map[string]int{
		"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
	}},
	{"A man a plan a canal panama.", map[string]int{
		"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
	}},
}

func Main0004() {
	Test(WordCount)
	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//WordCount("ha haa ha")
}

func WordCount(s string) map[string]int {
	// 如果不用make，那m一开始是nil，而不是map。
	// 类似于python中的m={}
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		//fmt.Println(m[v])  //非常奇葩，nil可以取值（得到零值），但是不可以赋值……
		m[v] = m[v] + 1
	}
	return m
}
