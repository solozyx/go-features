// deadlock 死锁demo

package main

func writeCloseChanDeadlock(){
	intChan := make(chan int,5)
	intChan <- 1
	intChan <- 2
	close(intChan)
	intChan <- 3
}

func writeChanData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i
	}
	close(intChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1); close(exitChan)
	go writeChanData(intChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	// writeCloseChanDeadlock()
	go writeCloseChanDeadlock()
}