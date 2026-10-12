package concurrency

import (
	"fmt"
	"time"
)

/*
В канал ch отправить 3 числа 111, 222, 333 и зкарыть канал в отдельной горутине
Вывести все числа из канала в главной горутине
*/
func Ch1() {
	ch := make(chan int)
	_ = ch
}

/*
Вывести все значения отправленные канал в главной горутине в цикле for
Простой способ без range
*/
func Ch2() {
	c := make(chan int)

	squares := func(c chan int) {
		for i := 0; i <= 9; i++ {
			c <- i * i
		}
		close(c) // close channel
	}

	go squares(c) // start goroutine
}

/* Вывести данные из закрытого канала в цикле */
func Ch3() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

/**
в функциях square и cube посчитать значения,
- затем отправить их обратно через тот же канал
- в функции main получить оба значения и просуммировать
- вывести результат

*/

func Ch4() {
	square := func(c chan int) {
		fmt.Println("[square] reading")
	}

	cube := func(c chan int) {
		fmt.Println("[cube] reading")
	}

	fmt.Println(" started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	// testNum := 3
	fmt.Println("[main] sent testNum to squareChan")

	// square command

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	// cube command

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	// аштфд part here

	// fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
	fmt.Println("[main] main() stopped")
}

/*
в функции green сделать ограничить канал, сделав его только для чтения
*/
func Ch5_stricter_chan_type() {
	greet := func(roc chan string) {
		fmt.Println("Hello " + <-roc + "!")
		roc <- "asd"
	}

	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}

/*
  Сделать так, чтобы обрабатывалась только первая (самая быстрая) отправка из 2 горутин
*/
func Ch6_read_from_several_chans() {
	var start time.Time

	service1 := func(c chan string) {
		time.Sleep(7 * time.Second)
		c <- "Hello from service 1"
	}

	service2 := func(c chan string) {
		time.Sleep(1 * time.Second)
		c <- "Hello from service 2"
	}

	fmt.Println("main() started", time.Since(start).Milliseconds()%1000)

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	res1 := <-chan1
	fmt.Println("Response from service 1", res1, time.Since(start).Milliseconds()%1000)
	res2 := <-chan2
	fmt.Println("Response from service 2", res2, time.Since(start).Milliseconds()%1000)

	fmt.Println("main() stopped", time.Since(start).Milliseconds()%1000)
}

/* 
	Сделать так, чтобы из селекта возвращалось значение по умолчанию по умолчанию, 
	если в горутинах нет значения

*/
func Ch7_read_from_several_chans() {
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
	}
	fmt.Println("main() stopped", time.Since(start).Milliseconds() % 1000)
}