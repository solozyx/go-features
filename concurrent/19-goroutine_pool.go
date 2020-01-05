package main

import (
	"fmt"
	"time"
)

// 定义一个任务类型 Task
type Task struct {
	// Task中有1个具体业务,业务名称f
	f func() error
	// 这里可以添加任务的优先级参数
}

// 创建一个Task任务
func NewTask(fn func() error) *Task {
	t := Task{
		f: fn,
	}
	return &t
}

// Task执行业务的方法
func (t *Task) Execute() {
	// 调用任务中已经绑定好的业务方法
	t.f()
}

//----------有关协程池 Pool角色的功能
// 定义一个Pool协程池 类型
type Pool struct {
	// 对外的Task入口 EntryChannel
	EntryChannel chan *Task
	// 内部的Task队列 jobsChannel
	jobsChannel chan *Task
	// 协程池中最大worker的数量
	workerNum int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		jobsChannel:  make(chan *Task),
		workerNum:    cap,
	}
	return &p
}

// 协程池创建一个Worker 并且让这个Worker去工作
func (p *Pool) worker(workerID int) {
	// 1 永久的从 jobsChannel 去取任务
	for task := range p.jobsChannel {
		// task 就是当前 worker 从 jobsChannel 中拿到的任务
		// 2 一旦取到任务 执行这个任务 这里可以做任务优先级的封装
		task.Execute()
		fmt.Println("workerID=", workerID, " 执行完1个任务")
	}
}

// 让协程池开始真正的工作 协程池一个启动方法
func (p *Pool) run() {
	// 1 根据 workerNum 来创建worker去工作
	for i := 0; i < p.workerNum; i++ {
		// 每个worker都应该是起一个 goroutine
		go p.worker(i)
	}

	// 2 从 EntryChannel 中取任务 将取到的任务 发送给 jobsChannel
	for task := range p.EntryChannel {
		// 一旦有task读到
		p.jobsChannel <- task
	}
}

func main() {
	// 1 创建一些任务
	t := NewTask(func() error {
		// 任务逻辑 ...
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
		return nil
	})

	// 2 创建一个 Pool 协程池
	p := NewPool(4)

	// 3 将这些任务 交给协程池Pool
	taskNum := 0 //统计任务的数量的初始值

	go func() {
		for {
			// 不断的向 Pool 中写入任务t,每个任务就是打印当前的时间
			p.EntryChannel <- t
			// 统计任务的数量
			taskNum += 1
			fmt.Println("当前一共执行了 ", taskNum, "个任务")
		}
	}()

	// 4 启动Pool 此时Pool会创建worker 让worker工作
	p.run()
}
