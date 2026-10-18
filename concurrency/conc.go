package concurrency

import (
	"fmt"
	"sync"
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

/* 
  Добавь сюда WaitGroup, сделай 100 задач так чтобы каналы не блокировались
*/
func Conc3_with_wg() {
	// worker than make squares
	sqrWorker := func(tasks <-chan int, results chan<- int, id int) {
		for num := range tasks {
			time.Sleep(time.Millisecond) // simulating blocking task
			fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
			results <- num * num
		}
	}

	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}

/* 
   Сделай так чтобы возвращалось значение 1000 
*/
func Conc4() {
	i := 0
	worker := func(wg *sync.WaitGroup) {
		i = i + 1
		wg.Done()
	}

    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go worker(&wg)
    }

    wg.Wait()

    fmt.Println("value of i after 1000 operations is", i)
}


/*
Сделай генератор чисел фибоначчи
*/
func Conc5_fib() {
	fib := func(length int) {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			print(i)
		}		
	}
	_ = fib

    // for fn := range fib(10) {
    //     fmt.Println("Current fibonacci number is", fn)
    // }
}

/*
  Смержить выходные каналы в один канал
*/ 
func Conc6_fan_in() {

	getInputChan := func() <-chan int {
		input := make(chan int, 100)
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		go func() {
			for num := range numbers {
				input <- num
			}
			close(input)
		}()
		return input
	}

	getSquareChan := func(input <-chan int) <-chan int {
		output := make(chan int, 100)

		go func() {
			for num := range input {
				output <- num * num
			}

			close(output)
		}()
		return output
	}

	merge := func(outputsChan ...<-chan int) <-chan int {
		merged := make(chan int, 100)
		return merged
	}

	chanInputNums := getInputChan()
	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)

	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)
	sqrSum := 0

	for num := range chanMergedSqr {
		sqrSum += num
	}

	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}
