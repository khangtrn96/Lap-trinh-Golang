package main

import (
	"fmt"
)

type Salutation struct {
	whoareyou string
	yourage   string
	yourjob   string
}

func identifyourinfo(infologin Salutation) {
	//tencuaban, tuoicuaban, congvieccuaban := controlyourinfo(infologin.whoareyou, infologin.yourage, infologin.yourjob)
	tencuaban, tuoicuaban, congvieccuaban, monhocbanday := controlyourinfo(infologin.whoareyou, infologin.yourage, infologin.yourjob, "Toán")
	fmt.Println(tencuaban)
	fmt.Println(tuoicuaban)
	fmt.Println("Hai kết quả xong đã ứng dụng Variadic function: ")
	fmt.Println(congvieccuaban)
	fmt.Println(monhocbanday)
}

//func controlyourinfo(value1, value2, value3 string) (tencuaban1, tuoicuaban1, congvieccuaban1 string) {
//Muốn dùng "variadic function". ta sẽ thêm ký hiệu dấu "..." trước gia trị biến khai báo
//Các biến còn lại là value1, value2 phải khai báo kiểu dữ liệu riêng(mà ở đây là kiểu string) với biến "Variadic function"
//Lúc này, mà cụ thể là "value3" sẽ là 1 "variadic function"
//value3[0]=infologin.yourage
//value3[1]="Toán"
func controlyourinfo(value1, value2 string, value3 ...string) (tencuaban1, tuoicuaban1, congvieccuaban1, monhocbanday1 string) {
	tencuaban1 = "Ten cua ban la " + value1
	tuoicuaban1 = "Tuoi cua ban la " + value2
	congvieccuaban1 = "Cong viec cua ban la " + value3[0]
	monhocbanday1 = "Mon hoc ban day la " + value3[1]
	//Để xem 1 "Variaidc function" có bao nhiêu phần tử, ta sẽ dung hàm "len()"
	fmt.Println("đỗ dài của biến Variadic function là", len(value3))
	return
}
func main() {
	aboutyou := Salutation{"Khang", "19", "giao vien"}
	identifyourinfo(aboutyou)
}
