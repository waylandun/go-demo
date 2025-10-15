// 教学内容: goroutine、channel 的创建与通信、缓冲通道、select 多路复用、同步等待。
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(100 * time.Millisecond) // 模拟工作耗时
		fmt.Printf("worker %d 处理任务 %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	// 启动一个简单的 goroutine
	go func() {
		fmt.Println("独立 goroutine 开始执行")
	}()
	time.Sleep(10 * time.Millisecond)

	// 无缓冲通道: 发送与接收必须同步
	message := make(chan string)
	go func() {
		message <- "来自 goroutine 的消息"
	}()
	fmt.Println(<-message)

	// 有缓冲通道
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println("缓冲通道读取:", <-buffered, <-buffered)

	// 使用通道池处理任务
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < 5; a++ {
		fmt.Println("结果:", <-results)
	}

	// select 等待多个 channel
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- "任务一完成"
	}()
	go func() {
		time.Sleep(80 * time.Millisecond)
		c2 <- "任务二完成"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("select 收到:", msg)
		case msg := <-c2:
			fmt.Println("select 收到:", msg)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("超时处理")
		}
	}
}
