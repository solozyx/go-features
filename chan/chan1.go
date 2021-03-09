package main

func main() {
	var (
		ch = make(chan *int, 1)
		n  = 1
		p  = &n
	)
	go func() {
		gp := <-ch
		println("*gp=",*gp) // read
	}()
	ch <- p
	*p += 1  // write
	println("*p=",*p)
}
