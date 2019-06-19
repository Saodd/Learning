package learnTour

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Prefix is a Markov chain prefix of one or more words.

// Prefix 为拥有一个或多个单词的链马尔可夫链的前缀。
type Prefix []string

// String returns the Prefix as a string (for use as a map key).

// String 将 Prefix 作为一个（用作映射键的）字符串返回。
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.

// Shift 从 Prefix 中移除第一个单词并追加上给定的单词。
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.

// Chain 包含一个从前缀到一个后缀列表的映射（“chain”）。
// 一个前缀就是一个加入了空格的，拥有 prefixLen 个单词的字符串。
// 一个后缀就是一个单词。一个前缀可拥有多个后缀。
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.

// NewChain 返回一个拥有 prefixLen 个单词前缀的 Chain。
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.

// Build 从提供的 Reader 中读取文本，并将它解析为存储了前缀和后缀的 Chain。
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

// Generate returns a string of at most n words generated from Chain.

// Generate 返回一个从 Chain 生成的，最多有 n 个单词的字符串。
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

//func Main0015() {
//	// 寄存命令行标记。
//	numWords := flag.Int("words", 100, "maximum number of words to print")
//	prefixLen := flag.Int("prefix", 2, "prefix length in words")
//
//	flag.Parse()                     // 解析命令行标记。
//	rand.Seed(time.Now().UnixNano()) // 设置随机数生成器的种子。
//
//	c := NewChain(*prefixLen)     // 初始化一个新的 Chain。
//	c.Build(os.Stdin)             // 从标准输入中构建链。
//	text := c.Generate(*numWords) // 生成文本。
//	fmt.Println(text)             // 将文本写入标准输出。
//}

// -----------------------------------------------------------
// -----------------------------------------------------------
type Prefix_zh []string

func (p Prefix_zh) String() string {
	return strings.Join(p, "")
}
func (p Prefix_zh) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain_zh struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain_zh(prefixLen int) *Chain_zh {
	return &Chain_zh{make(map[string][]string), prefixLen}
}
func (c *Chain_zh) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix_zh, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		for _, b := range s {
			key := p.String()
			c.chain[key] = append(c.chain[key], string(b))
			p.Shift(string(b))
		}
	}
}
func (c *Chain_zh) Generate(n int) string {
	p := make(Prefix_zh, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, "")
}

func Main0015() {
	numWords := flag.Int("words", 200, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")

	flag.Parse()                     // 解析命令行标记。
	rand.Seed(time.Now().UnixNano()) // 设置随机数生成器的种子。

	c := NewChain_zh(*prefixLen)  // 初始化一个新的 Chain。
	c.Build(os.Stdin)             // 从标准输入中构建链。
	text := c.Generate(*numWords) // 生成文本。
	fmt.Println(text)             // 将文本写入标准输出。
}
