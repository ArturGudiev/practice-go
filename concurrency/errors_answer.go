package concurrency

import "fmt"

func Err1_answer() {
	greet := func(c chan string) {
		<-c // for John
		<-c // for Mike
	}

	c := make(chan string, 1)
	go greet(c)
	c <- "John"
	c <- "Mike"
	fmt.Println("Err1 finished successfully")
}
