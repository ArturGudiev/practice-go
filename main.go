package main

import (
	"fmt"
	"go-practice/concurrency"
	"os"
)

func Default() {
	println(" ---- default option ---- ")
	concurrency.Ch2_answer()
	concurrency.Err1_answer()
}

func runPracticeItem(itemType, itemNumber string) {
	println(" ---- item option ---- ")
	println("itemType:", itemType)
	println("itemNumber:", itemNumber)

	switch [2]string{itemType, itemNumber} {
	case [2]string{"ch", "1"}:
		concurrency.Ch1()
	case [2]string{"ch", "1a"}:
		concurrency.Ch1_answer()
	case [2]string{"err", "1"}:
		concurrency.Err1()
	case [2]string{"err", "1a"}:
		concurrency.Err1_answer()
	}
}

func main() {
	// fmt.Println(os.Args)
	switch len(os.Args) {
	case 1, 2:
		Default()
	case 3:
		itemType, itemNumber := os.Args[1], os.Args[2]
		runPracticeItem(itemType, itemNumber)
	default:
		fmt.Println("usage: go run . [arg1] [arg2]")
	}
}
