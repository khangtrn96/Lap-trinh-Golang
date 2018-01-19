package main

import (
	"fmt"
)

type Salutation struct {
	Nation string
	City   string
}
type funccustom func(datacustom string)

func showresult(datainput Salutation, resultaftercustom funccustom) {
	valuewhereyoulive1, valuewhereyoulive2 := controldatawhereyoulive(datainput.Nation, datainput.City)
	resultaftercustom(valuewhereyoulive1)
	resultaftercustom(valuewhereyoulive2)
}
func controldatawhereyoulive(datainput1, datainput2 string) (valueWYL1, valueWYL2 string) {
	valueWYL1 = "Quốc gia bạn đang sống là " + datainput1
	valueWYL2 = "Thành phố bạn đang sống là " + datainput2
	return
}
func Printcustom(valuecustom string) funccustom {
	return func(textVLWYL string) {
		fmt.Println(textVLWYL + valuecustom)
	}
}
func main() {
	whereyoulive := Salutation{"Việt Nam", "Hồ Chí Minh"}
	showresult(whereyoulive, Printcustom(" @_@"))
}
