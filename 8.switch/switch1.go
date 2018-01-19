package main

//switch1 cho biết cách sữ dụng hàm case cách 2 bằng cách khai báo nhiều giá trị so sánh
//Chú ý dòng 29, 31 ,33
import (
	"fmt"
)

type Salutation struct {
	nation string
	place  string
}

func showresult(datawhereinfo Salutation, Isformal bool, ten string) (result3, b, c string) {
	result1, result2 := controlvalueinput(datawhereinfo.nation, datawhereinfo.place, "Thanks for your support!!!!")
	if prefix := Getprefix(ten); Isformal {
		fmt.Println(prefix + result1)
	} else {
		fmt.Println(prefix + result2)
	}
	return
}
func controlvalueinput(datacontrol_1 string, datacontrol_2 ...string) (datareturncontrol_1, datareturncontrol_2 string) {
	datareturncontrol_1 = ". " + datacontrol_2[0] + " " + datacontrol_2[1] + ". Chào mừng bạn đến quê tôi"
	datareturncontrol_2 = ". " + datacontrol_2[1]
	return
}
func Getprefix(text string) (custom string) {
	switch text {
	case "Khang", "Giao": // Gía trị so sánh là Khang và Giao
		custom = "Xin chao ban tre"
	case "Khiêm", "Vinh": //Gia tri so sánh là Khiêm và Vinh
		custom = "Xin chao mày"
	case "Huy", "Hoàng": //Gía trị so sánh là Huy và Hoàng
		custom = "Biến đi nhé"
	}
	return
}
func main() {
	whereareyou := Salutation{"Viet Nam", "Nam Dinh"}
	showresult(whereareyou, false, "Khiêm")
}
