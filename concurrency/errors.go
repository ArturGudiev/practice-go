package concurrency

import "fmt"

func Err1() {
	greet := func(c chan string) {
		<-c // for John
		<-c // for Mike
	}

	c := make(chan string, 1)
	go greet(c)
	c <- "John"
	close(c)
	c <- "Mike"
	fmt.Println("Err1 Finished successfully")

}
