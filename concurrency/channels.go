package concurrency


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