package main

//ở các ngôn ngữ lập trình khác, ta không thể khai báo biến cho dữ liệu trả về nhưng Go cho phép diều đó
//Chú ý dòng 18 dê thấy được diều đó
import (
	"fmt"
)

type Salutation struct {
	nation  string
	capital string
}

func showinfo(infoinvidual Salutation) {
	info1, info2 := idplace(infoinvidual.nation, infoinvidual.capital)
	//khi nhận giá trị trả về từ hàm "idplace", với cách xắp xếp thứ tự các biến trong hàm "idplace" thì info1=info3, info2=info4
	fmt.Println(info1)
	fmt.Println(info2)

}
func idplace(nationname, capitalname string) (info3 string, info4 string) {
	info3 = "Nước của tôi là " + nationname
	info4 = "Thủ đô của nước tôi là " + capitalname
	return info3, info4
	//không cần ghi info3, info4 thì compile vẫn trả về 2 giá trị info3, info4
}
func main() {
	whereyoulive := Salutation{"Viet  Nam", "Ha Noi"}
	showinfo(whereyoulive)
}
