package main

import (
	"fmt"
)

type Salutation struct {
	firstsport  string
	secondsport string
	thirdsport  string
}

//ý nghĩa việc dùng hàm procedure là giúp đưa ra cách sắp xếp kết quả theo ý muốn
//Ở đây có 2 sự lựa chọn là "textsport1"(dong 30) và "textsport2"(dòng 33)
// nếu dùng hàm textsport1 thì  sẽ không xuống dòng
// dùng hàm textsport2 sẽ xuống dòng
// Ta có thể dùng dòng 22 fmt.Println là 1 cách để trả về dử liệu
func Expresssportinvilike(Sportinvi Salutation, procedure func(string)) {
	info1, info2, info3, info4 := showinfo(Sportinvi.firstsport, Sportinvi.secondsport, Sportinvi.thirdsport, "tennis", "badminton")
	procedure(info1)
	procedure(info2)
	fmt.Println(info3)
	procedure(info4)
}
func showinfo(sport1, sport2 string, sport3 ...string) (info5, info6, info7, info8 string) {
	info5 = "Mon the thao thu nhat la " + sport1
	info6 = "Mon the thao thu hai la " + sport2
	info7 = "Mon the thao thu ba la " + sport3[0]
	info8 = "Mon the thao thu tu va thu nam la " + sport3[1] + " và " + sport3[2]
	return
}
func textsport1(text string) {
	fmt.Print(text)
}
func textsport2(text string) {
	fmt.Println(text)
}
func main() {
	sportyoulike := Salutation{"Football", "Swimming", "Baseball"}
	fmt.Println("Kết quả khi dùng textsport1:")
	Expresssportinvilike(sportyoulike, textsport1)
	fmt.Println("\n")
	fmt.Println("Kết quả khi dùng textsport2:")
	Expresssportinvilike(sportyoulike, textsport2)
}
