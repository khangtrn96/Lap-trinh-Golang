package main

//switch2 muốn nói về cách sử dụng fallthrough để bỏ qua giá trị trong switch
//Chú ý dòng
//switch2 còn đề cập đến việc sử dụng cách gán giá trị của switch
//Chú ý dòng
// switch2 còn đề cập đến việc sử dụng nhiều cách so sánh để trả về dữ liệu. Ở đây đã dùng hàm len dể tính dộ dài
import (
	"fmt"
)

type Salutation struct {
	schoolname string
	place      string
}

func showresult(datainputstudy Salutation, level string, isFormal bool) (result3, result4 string) {
	result1, result2 := controlvalue(datainputstudy.schoolname, datainputstudy.place, ". Cam on da lua chon ngoi truong nay")
	if loichao := doanloichao(level); isFormal {
		fmt.Println(loichao + result1)
	} else {
		fmt.Println(loichao + result2)
	}
	return
}
func controlvalue(datacontrol_1 string, datacontrol_2 ...string) (datasolve_1, datasolve_2 string) {
	datasolve_1 = datacontrol_1 + " " + datacontrol_2[0] + datacontrol_2[1]
	datasolve_2 = datacontrol_2[1]
	return
}
func doanloichao(text string) (custom string) {
	switch {
	case text == "hoc sinh", text == "phu huynh": //text mang tính chất so sánh
		custom = "Chao ban da den ngoi truong cua chúng tôi ."
		fallthrough // fallthrough có chúc năng chuyển đến dòng case tiếp theo
		// fallthrough chỉ có tác dụng trong case mà nó ở
	case text == "người lạ", len(text) == 10:
		// text mang tính chát so sánh
		// nếu giá trị text có độ dài chuỗi là 10 thì sẽ chọn case này
		custom = "Chao mung ban. "
	case text == "Hiệu trưởng", len(text) == 12:
		//text mang tính chất so sánh
		//nếu giá trị text có độ dài chuỗi là 12 thì sẽ chọn case này
		custom = "Xin chao thay đến trường ."
	default:
		custom = "Biến đi nhé!!! "
	}
	return
}

func main() {
	yourstudy := Salutation{"KHTN", "Thu duc"}
	showresult(yourstudy, "1234567890", true)
}
