package concurrency

import "fmt"

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

    testNum := 3
    fmt.Println("[main] sent testNum to squareChan")

    // square command

    fmt.Println("[main] resuming")
    fmt.Println("[main] sent testNum to cubeChan")

    // cube command 

    fmt.Println("[main] resuming")
    fmt.Println("[main] reading from channels")

    // аштфд part here

    fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
    fmt.Println("[main] main() stopped")
}