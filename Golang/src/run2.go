package main

import (
    "fmt"
    "time"
)

func main() {
    var ch = make(chan string, 1)
    go func() {
        for i := range ch {
            ch <- i
        }
    }()
    time.Sleep(time.Second)

    {
        start := time.Now()
        for i := 0; i < 1000000; i++ {
            ch <- "1"
            <- ch
        }
        totalTime := time.Since(start)
        close(ch)
        fmt.Println(float64(totalTime.Seconds()) / 1000000)
    }
}
