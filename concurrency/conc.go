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


/*
	Сымитировать результат работы функция вычисления time.Sleep(time.Millisecond_)
	в Функции получай и отправляй результаты 
	все одновременно должны работать 3 такие функции 
	куда отправляются значенияб там размер 10 (откуда приходят тоже )
	отправь 5 результатов 
	получи 5 ответов 

*/

func Conc2_pull_workers() {
	
	sqrWorker := func(id int) {
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
	}
	_ = sqrWorker


	for i := 0; i < 3; i++ {
		
	}

	for i := 0; i < 5; i++ {
		//i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	for i := 0; i < 5; i++ {
	}
}
