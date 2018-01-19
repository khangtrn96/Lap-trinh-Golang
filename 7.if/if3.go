package main

//if3 giúp ta biết thêm về cách dặt điều kiện của if
//Chú ý dòng 17
import (
	"fmt"
)

type Salutation struct {
	choice1 string
	choice2 string
	choice3 string
}
type format func(string)

func showresult(choicedatainput Salutation, custom format, isFormal int) {
	valueafterchoice1, valueafterchoice2, valueafterchoice3 := controlvalue(choicedatainput.choice1, choicedatainput.choice2, choicedatainput.choice3, ". Bạn đã quyết định sáng suốt!!")
	if comment := "Bạn đã lực chọn cửa "; isFormal < 2 { //isFormal ở đây là điều kiện để so sánh
		custom(comment + valueafterchoice1)
	} else if isFormal < 3 {
		custom(comment + valueafterchoice2)
	} else {
		custom(comment + valueafterchoice3)
	}
}
func controlvalue(value1, value2 string, value3 ...string) (aftersolve1, aftersolve2, aftersolve3 string) {
	aftersolve1 = value1 + ". Hay lắm"
	aftersolve2 = value2 + ". Bạn chọn nhầm cửa rồi"
	aftersolve3 = value3[0] + value3[1]
	return
}
func custom(symbol string) string {
	//return func(text string) { //text thay thế cho "custom+valueafterchoice1"
	fmt.Println(text + " " + symbol)
	//}
	return
}
func main() {
	yourchoice := Salutation{"Số 1", "Số 2", "Số 3"}
	showresult(yourchoice, custom("^_^"), 1)
}
