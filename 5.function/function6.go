package main

import (
	"fmt"
)

type Salutation struct {
	nameobject1 string
	nameobject2 string
	nameobject3 string
	nameobject4 string
}
type xapxepketqua func(string)

func showresult(contentobject Salutation, procedure xapxepketqua) {
	result1object, result2object, result3object, result4object := controlresultobject(contentobject.nameobject1, contentobject.nameobject2, contentobject.nameobject3, contentobject.nameobject4)
	procedure(result1object)
	procedure(result2object)
	procedure(result3object)
	procedure(result4object)
}
func controlresultobject(infoobject1, infoobject2, infoobject3, infoobject4 string) (assignobject1, assignobject2, assignobject3, assignobject4 string) {
	assignobject1 = "Môn học thứ nhất là " + infoobject1
	assignobject2 = "Môn học thứ hai la " + infoobject2
	assignobject3 = "Môn học thứ ba là " + infoobject3
	assignobject4 = "Môn học thứ tư là " + infoobject4
	return
}
func hamxuongdong(text string) {
	fmt.Println(text)
}
func hamtrendong(text string) {
	fmt.Print(text)
}
func main() {
	yourobject := Salutation{"Toán", "Văn", "Lý", "Hoá"}
	showresult(yourobject, hamtrendong)
}
