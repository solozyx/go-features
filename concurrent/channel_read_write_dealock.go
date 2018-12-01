// deadlock 死锁demo

package main
func writeChanData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i
	}
	close(intChan)
}
func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeChanData(intChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}