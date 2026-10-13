package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Conc1_answer() {
	service := func(wg *sync.WaitGroup, instance int) {
		time.Sleep(2 * time.Second)
		fmt.Println("Service called on instance", instance)
		wg.Done() // decrement counter
	}

	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}

func Conc2_pull_workers_answer() {
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

func Conc3_with_wg_answer() {
	sqrWorker := func(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
		for num := range tasks {
			time.Sleep(time.Millisecond)
			fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
			results <- num * num
		}
		wg.Done()
	}

	fmt.Println("[main] main() started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
	}

	go func() {
		i := 0
		for result := range results {
			fmt.Println("[main] Result", i, ":", result)
			i++;
		}
	}()

	for i := 0; i < 100; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}
	close(tasks)
	fmt.Println("[main] Wrote 5 tasks")
	wg.Wait()
	fmt.Println("[main] main() stopped")
}
