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

var practiceItems = map[string]func(){
	"ch1": concurrency.Ch1, "ch1a": concurrency.Ch1_answer,
	"ch2": concurrency.Ch2, "ch2a": concurrency.Ch2_answer,
	"ch3": concurrency.Ch3, "ch3a": concurrency.Ch3_answer,
	"ch4": concurrency.Ch4, "ch4a": concurrency.Ch4_answer,
	"ch5": concurrency.Ch5_stricter_chan_type, "ch5a": concurrency.Ch5_stricter_chan_type_answer,
	"ch6": concurrency.Ch6_read_from_several_chans, "ch6a": concurrency.Ch6_read_from_several_chans_answer,
	"ch7": concurrency.Ch7_read_from_several_chans, "ch7a": concurrency.Ch7_read_from_several_chans_answer,
	"ch8": concurrency.Ch8_wait_in_select_some_time_then_default, "ch8a": concurrency.Ch8_wait_in_select_some_time_then_default_answer,

	"conc1": concurrency.Conc1_, "conc1a": concurrency.Conc1_answer,
	"conc2": concurrency.Conc2_pull_workers_answer, "conc2a": concurrency.Conc2_pull_workers_answer,
	
	"conc3": concurrency.Conc3_with_wg,
	"conc3a": concurrency.Conc3_with_wg_answer,

	"conc4": concurrency.Conc4,
	"conc4a": concurrency.Conc4_answer,

	"conc5a": concurrency.Conc5_fib_answer,

	"err1": concurrency.Err1, "err1a": concurrency.Err1_answer,
}

func runPracticeItem(key string) {
	fn, ok := practiceItems[key]
	if !ok {
		fmt.Println("unknown:", key)
		return
	}
	fn()
}

func main() {
	// fmt.Println(os.Args)
	switch len(os.Args) {
	case 1:
		Default()
	case 2:
		practiceItem := os.Args[1]
		runPracticeItem(practiceItem)
	default:
		fmt.Println("usage: go run . [arg1] [arg2]")
	}
}
