package main

import (
	_ "fmt"
	_ "reflect"
)

type Salutation struct {
	name     string
	greeting string
}

//type numbers string

// func showresult_1(data_Range_1 []Salutation) {
// 	for _, s := range data_Range_1 {
// 		if s.name == "Khang" {
// 			println("Tôi tên là " + s.name + " " + s.greeting)
// 		} else if s.name == "Michale" {
// 			println("I am " + s.name + " " + s.greeting)
// 		} else {
// 			println("OMG!! " + s.name + " " + s.greeting)
// 		}
// 	}
// }

func showresult_2(data_Range_2 []string) {
	for i, s := range data_Range_2 {
		println("Số " + i + " trong tiếng anh là " + s)
	}
}

func main() {
	// slice_1 := []Salutation{
	// 	{"Khang", "Xin chao"},
	// 	{"Michale", "Hello"},
	// 	{"Mark Zurkenburg", "Bonjour"},
	// }
	slice_2 := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	//showresult_1(slice_1)
	showresult_2(slice_2)
}
