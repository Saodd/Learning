package main

import (
	"learnAlgo"
	"learnAlgo/tools"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		li := tools.Gen_ints_list(2000000)
		learnAlgo.QuickSortInt(li)
	}()
	//li := tools.Gen_ints_list(2000000)
	//learnAlgo.QuickSortInt(li)
	http.ListenAndServe("0.0.0.0:6060", nil)
}
