package main

import (
	"fmt"
)

type salutation struct {
	name     string
	greeting string
}

func showresult(infoinvi salutation, isformal bool) (result1, result2 string) {
	result1, result2 = controlvalue(infoinvi.name, infoinvi.greeting, ". Rat vui duoc gap ban")
	if prefix := Getprefix(infoinvi.name); isformal {
		fmt.Println(prefix + result1)
	} else {
		fmt.Println(result2 + " " + prefix + result1)
	}
	return
}
func controlvalue(dainfoname string, datainfogreeting ...string) (answer1, answer2 string) {
	answer1 = dainfoname + datainfogreeting[1] //Khang hoặc Giao+ Rat vui duoc gap ban
	answer2 = datainfogreeting[0]              //Xin chao
	return
}

func Getprefix(text string) (custom string) {
	switch text {
	case "Khang":
		custom = "Mr "
	case "Giao":
		custom = "Mrs "
	default: //nếu kết quả không có ở case trước đó thì ta chọn giá trị mặc định "Dude"
		custom = "Dude "
	}
	return
}

func main() {
	greetingstatement := salutation{"Giao", "Xin chao"}
	showresult(greetingstatement, false)
}
