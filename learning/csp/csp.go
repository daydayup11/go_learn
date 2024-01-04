package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // 模拟耗时的工作
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动三个goroutine，模拟三个工作线程
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送五个工作任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// 收集工作结果
	for a := 1; a <= numJobs; a++ {
		r := <-results
		fmt.Println(r)
	}
}
