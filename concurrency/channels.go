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

func Ch2() {
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
		fmt.Println(val, ok)
	}
	fmt.Println("main() stopped")

}