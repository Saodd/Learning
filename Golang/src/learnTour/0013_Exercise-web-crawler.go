package learnTour

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
练习：Web 爬虫
在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

-------------------------------------------------------------------------

-------------------------------------------------------------------------

*/

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, chParent chan bool) {
	defer func() { chParent <- true }()
	if depth <= 0 {
		return
	}
	if urlfetched.IsFetched(url) {
		return
	}
	<-goPool
	body, urls, err := fetcher.Fetch(url)
	goPool <- true
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	chChild := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, chChild)
	}
	for range urls {
		<-chChild
	}
	return
}

var goPool chan bool

var urlfetched *urlFetched = &urlFetched{urls: map[string]bool{}}

type urlFetched struct {
	urls map[string]bool
	lock sync.Mutex
}

func (self *urlFetched) IsFetched(url string) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	if b, _ := self.urls[url]; !b {
		self.urls[url] = true
		return false
	}
	return true
}

func Main0013() {
	goPool = make(chan bool, 1)
	for i := 0; i < 1; i++ {
		goPool <- true
	}
	ch := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Worker机制的Crawler -----------------------------------------
func Main0013_2() {
	c := NewCrawler(100000, 10, "https://golang.org/", fetcher)
	c.Run()
}

type task struct {
	depth int
	url   string
}

func NewCrawler(numGo int, depth int, startURL string, fetcher Fetcher) *Crawler {
	ch := make(chan bool, numGo)
	for i := 0; i < numGo; i++ {
		ch <- true
	}
	return &Crawler{passport: ch, tasks: []task{{depth, startURL}}, fetcher: fetcher}
}

type Crawler struct {
	passport chan bool
	tasks    []task
	fetcher  Fetcher
	lock     sync.Mutex
}

func (self *Crawler) Run() {
	numGo := cap(self.passport)

	for i := 0; i < numGo; i++ {
		go self.work()
	}
	time.Sleep(time.Second)
	for len(self.passport) != numGo {
		time.Sleep(100 * time.Millisecond)
	}
}

func (self *Crawler) getTask() (t task) {
	self.lock.Lock()
	defer self.lock.Unlock()
	if len(self.tasks) != 0 {
		t = self.tasks[0]
		self.tasks = self.tasks[1:]
	} else {
		t = task{}
	}

	return t
}
func (self *Crawler) putTasks(ts []task) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.tasks = append(self.tasks, ts...)
}

func (self *Crawler) work() {
	var t task
	for {
		<-self.passport
		t = self.getTask()
		if t.depth <= 0 {
			self.passport <- true
			time.Sleep(time.Second)
			continue
		}
		if urlfetched.IsFetched(t.url) {
			self.passport <- true
			continue
		}
		body, urls, err := fetcher.Fetch(t.url)
		if err != nil {
			fmt.Println(err)
			self.passport <- true
			continue
		}
		fmt.Printf("found: %s %q\n", t.url, body)
		if d := t.depth - 1; d > 0 {
			ts := make([]task, len(urls))
			for i, u := range urls {
				ts[i] = task{d, u}
			}
			self.putTasks(ts)
		}
		self.passport <- true
	}
}
