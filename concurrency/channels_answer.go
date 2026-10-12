package concurrency

import (
	"fmt"
	"time"
)

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

func Ch4_answer() {
	square := func(c chan int) {
		fmt.Println("[square] reading")
		num := <-c
		c <- num * num
	}

	cube := func(c chan int) {
		fmt.Println("[cube] reading")
		num := <-c
		c <- num * num * num
	}

	fmt.Println(" started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
	fmt.Println("[main] main() stopped")
}

func Ch5_stricter_chan_type_answer() {
	greet := func(roc <-chan string) {
		fmt.Println("Hello " + <-roc + "!")
	}

	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}

func Ch6_read_from_several_chans_answer() {
	var start time.Time

	service1 := func(c chan string) {
		time.Sleep(7 * time.Second)
		c <- "Hello from service 1"
	}

	service2 := func(c chan string) {
		time.Sleep(1 * time.Second)
		c <- "Hello from service 2"
	}

	fmt.Println("main() started", time.Since(start).Milliseconds() % 1000)

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start).Milliseconds() % 1000)
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start).Milliseconds() % 1000)
	}

	fmt.Println("main() stopped", time.Since(start).Milliseconds() % 1000)
}

func Ch7_read_from_several_chans_answer() {
	start := time.Now()
	service1 := func(c chan string) {
		time.Sleep(7 * time.Second)
		c <- "Hello from service 1"
	}
	
	service2 := func(c chan string) {
		time.Sleep(2 * time.Second)
		c <- "Hello from service 2"
	}
	
	fmt.Println("main() started")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start).Milliseconds() % 1000)
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start).Milliseconds() % 1000)
	default:
		fmt.Println("No response received", time.Since(start).Milliseconds() % 1000)
	}
	fmt.Println("main() stopped", time.Since(start).Milliseconds() % 1000)
}