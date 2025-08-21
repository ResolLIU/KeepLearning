package golang

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 来自2025/08/20 的一场面试
// step 1 : call a func for 1000 times by using goroutine,ensure that every goroutine will run before the main func end
// step 2 : what if some error exit and we should close all goroutine which are still alive

func call() error {
	n := rand.Intn(1000)
	if n > 990 {
		err := errors.New("some err occur and the program will exit")
		fmt.Println("fail , n is ", n)
		return err
	}
	<-time.After(3 * time.Second)
	fmt.Println("success , n is ", n)
	return nil
}

// 如果要实现1000个线程并发管理
// 可以直接使用waitGroup

func Step1() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			call()
			defer wg.Done()
		}()
	}
	wg.Wait()
}

// Step 2 : 我当时想的是暴力，只要err != nil 就直接os退出，面试官说不够优雅
// 正解应该是使用 context, 这里对call函数进行一点改写
func call2(ctx context.Context) error {
	n := rand.Intn(1000)
	if n > 990 {
		err := errors.New("some err occur and the program will exit")
		fmt.Println("fail , n is ", n)
		return err
	}
	select {
	case <-ctx.Done():
		fmt.Println("[cancel]")
		return nil
	case <-time.After(3 * time.Second):
		fmt.Println("success , n is ", n)
		return nil
	}
}
func Step2() {
	// 默认初始化，通过暴露的withCancel接口获取到终止函数
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			err := call2(ctx)
			if err != nil {
				cancel()
			}
		}()
	}
}

// 这里想起了之前的一道题，有n个协程，请对其进行排序,比如轮询输出abcabcabc
func printN(n int) {
	fmt.Println("num is", n)
}
func PrintForNTurns(n int) {
	que := make([]chan int, n)
	wg := sync.WaitGroup{}
	for i := range que {
		que[i] = make(chan int, 1)
	}
	defer func() {
		for i := range que {
			close(que[i])
		}
	}()
	que[0] <- 1
	for j := 0; j < n; j++ {
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				defer func() {
					que[(i+1)%n] <- 1
					wg.Done()
				}()
				<-que[i]
				printN(i)
			}()
		}
	}
	wg.Wait()
}
