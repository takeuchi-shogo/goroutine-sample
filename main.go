package main

import (
	"fmt"
	"sync"
	"time"
)

func syncWaitGroup() {
	ch := make(chan struct{}, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-ch
				time.Sleep(2 * time.Second)
				wg.Done()
			}()

			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channel() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine(s, c)
	go goroutine(s, c)
	// 1回目のgoroutineの値
	x := <-c
	fmt.Println(x)
	// 2回目のgoroutineの値
	y := <-c
	fmt.Println(y)
}

func closeChannel() {
	// チャネルのバッファーを二つ用意する
	ch := make(chan int, 2)
	ch <- 100
	// 1
	fmt.Println(len(ch))
	ch <- 200
	// 2
	fmt.Println(len(ch))
	// chanをcloseさせないとfor文で3つ目を取りに行ってしまう
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

func main() {
	syncWaitGroup()
	channel()
	closeChannel()
	// Goroutine終了後の2秒後に呼ばれる
	fmt.Println("goroutine success!")
}
