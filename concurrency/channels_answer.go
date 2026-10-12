package concurrency


func Ch1_answer() {
	
	ch := make(chan int)

	go func() {
		ch <- 111
		ch <- 222
		ch <- 333
		close(ch)
	}()

	for num := range ch {
		println(num)
	}
}