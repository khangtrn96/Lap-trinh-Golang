package main

import (
	"fmt"
)

type Salutation struct {
	ability1 string
	ability2 string
	ability3 string
}
type format func(datacustom string)

func solvedatainput(datainput Salutation, custom format, isFormal bool) { //isFormal mang giá trị "true", "false"
	valueaftersolve1, valueaftersolve2, valueaftersolve3 := controlvalue(datainput.ability1, datainput.ability2, datainput.ability3, "amazing")
	if status := "Tôi có "; isFormal { //isFormal để không thì nó mang ý nghĩa là true
		custom(status + valueaftersolve1)
	}
	custom(valueaftersolve2)
	custom(valueaftersolve3)
}

func controlvalue(value1, value2 string, value3 ...string) (aftersolve1, aftersolve2, aftersolve3 string) {
	aftersolve1 = "khả năng thứ nhất là " + value1
	aftersolve2 = "Khả năng thứ hai của tôi là " + value2
	aftersolve3 = value3[0] + " is an " + value3[1] + " sport"
	return
}

func custom(sign string) format {
	return func(text string) {
		fmt.Println(text + " " + sign)
	}
}
func main() {
	yourabilities := Salutation{"coding", "sleeping", "swimming"}
	solvedatainput(yourabilities, custom("!!!"), false)
}
