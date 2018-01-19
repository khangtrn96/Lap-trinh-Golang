package main

import (
	"fmt"
)

type Salutation struct {
	nation string
	place  string
}

func showresult(datawhere Salutation, level string, isFormal bool) (result3, result4 string) {
	result1, result2 := controlvalue(datawhere.nation, datawhere.place, "Chao mung ban đến trường")
	if loichao := doanloichao(level); isFormal {
		fmt.Println(loichao + result1)
	} else {
		fmt.Println(loichao + result2)
	}
	return
}
func controlvalue(datacontrol_1 string, datacontrol_2 ...string) (data_1, data_2 string) {
	data_1 = datacontrol_1 + " " + datacontrol_2[0] + " " + datacontrol_2[2]
	data_2 = datacontrol_2[1]
	return
}
func doanloichao(text string) (custom string) {
	switch {
	case text == "học sinh", text == "phu huynh":
		custom = "Xin chao ban đến."
	case text == "hiệu trưởng":
		custom = "Xin chào thầy dến trường"
	default:
		custom = "Biến đi nhé"
	}
	return
}

func TypeSwitchText() {
	var text interface{}
	text = 123
	switch text.(type) {
	case int:
		fmt.Println("Kiểu int")
	case string:
		fmt.Println("Kiểu string")
	case Salutation:
		fmt.Println("Kiểu Salutation")
	default:
		fmt.Println("unknown")
	}
}
func main() {
	whereyoulive := Salutation{"KHTN", "Thu Duc"}
	showresult(whereyoulive, "học sinh", true)
	TypeSwitchText()
}
