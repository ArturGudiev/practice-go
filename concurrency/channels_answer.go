package concurrency

import "fmt"

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

func Ch2_answer() {
	c := make(chan int)

	squares := func(c chan int) {
		for i := 0; i <= 9; i++ {
			c <- i * i
		}
		close(c) // close channel
	}

	go squares(c) // start goroutine
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}
	fmt.Println("main() stopped")

}

func Ch3_answer() {
	c := make(chan int, 3)
    c <- 1
    c <- 2
    c <- 3
    close(c)

    // iteration terminates after receiving 3 values
    for elem := range c {
        fmt.Println(elem)
    }
}
