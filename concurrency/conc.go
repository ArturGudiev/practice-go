package concurrency

import (
	"fmt"
	"time"
)

/*
Вызови сервис в 3 горутинах и выведи строку, когда все горутины завершатся.
*/
func Conc1_() {
	service := func(instance int) {
		time.Sleep(2 * time.Second)
		fmt.Println("Service called on instance", instance)
	}

	fmt.Println("main() started")

	for i := 1; i <= 3; i++ {

		go service(i)
	}

	fmt.Println("main() stopped")
}
