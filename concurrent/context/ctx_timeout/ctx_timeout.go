package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r *http.Response
	err error
}

func process() {
	// 2秒超时控制 返回子context 和 1个取消函数
	// context内部维护1个timer timer.C
	// 任何的子协程执行任务都可以用context控制 [通用超时控制]
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	// defer当process函数执行完退出前执行 取消函数 取消上下文
	// 2秒的定时器调用cancel后被取消 做清理工作
	defer cancel()

	// http底层使用TCP
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// channel 用于存储本次http请求的结果
	c := make(chan Result, 1)

	// http GET 请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}

	// 子goroutine做具体事情
	go func() {
		// 发出http请求 请求正常结果放在resp 请求错误结果放在err
		// 但是 client.Do(req) 多长时间返回不能确定
		resp, err := client.Do(req)

		// 子goroutine做什么任务不重要 访问后端 读写文件 排序 都可以用context控制超时
		// time.Sleep(3 * time.Second)

		pack := Result{r: resp, err: err}
		// Client.Do()请求结束 把结果写入channel
		// 如果channel能读取到数据了说明请求结束了
		c <- pack
	}()

	// context 如何处理超时
	// 后端服务没问题 走正常分支
	// redis mysql 负载高了 挂了 走context超时分支
	// 没有超时控制 很多子协程在执行任务 服务不响应 程序的负载也会高 所以要有超时控制
	// etcd客户端使用context做超时控制
	select {
	case <-ctx.Done():
		// Client.Do() 超过2秒没有返回 则请求超时
		// context内部Done方法也维护了1个channel 如果该channel可读了 说明超时了
		// context.Done() 方法当context内部的timer到时了 该方法可执行
		tr.CancelRequest(req) // 超时则取消本次正在执行的http请求
		// 取消请求后悔返回err
		res := <-c
		fmt.Println("Server Timeout! err:", res.err)
	case res := <-c: // 能从channel读出数据了 说明 Client.Do() 请求结束了
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
}

func main() {
	process()
}