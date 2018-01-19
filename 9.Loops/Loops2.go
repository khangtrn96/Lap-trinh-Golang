package main

/*
khi chạy vòng lặp for, kết quả trả về "imported and not used" tức là chưa sử dụng thư viện fmt,
nên ta phải thêm dấu _ ở chỗ import
*/
import (
	_ "fmt" //thêm dấu gạch ngang
)

//cách 1 để dùng vòng lặp for
func for_basic(numbers_1 int) {
	for i := 0; i < numbers_1; i++ {
		println("Tôi tên là Trần An Khang")
	}
}

//cách 2 để dùng vòng lặp for
func for_while(numbers_2 int) {
	j := 0
	for j < numbers_2 {
		println("Bạn gái tôi tên là:")
		j++
	}
}

//cách 3 để dùng vòng lặp for
func for_break(numbers_3 int) {
	k := 0
	for {
		if k >= numbers_3 {
			break //nếu k>=numbers_3 thì thoát ra khỏi vòng lặp
		}
		println("Tôi là ai")
		k++
	}
}

//cách 4 để dùng vòng lặp for VÔ HẠN LẦN
// func for_infinite() {
// 	for {
// 		println("Tôi biết rồi")
// 	}
// }
func for_continue(numbers_4 int) {
	for q := 0; q < numbers_4; q++ {
		if q%2 == 0 {

			continue //khi q chia hết cho 2 thì nó sẽ bỏ qua và không thực hiện lệnh,
		}
		println(q) //lệnh

	}
}
func main() {
	for_basic(5)
	for_while(4)
	for_break(2)
	//for_infinite()
	for_continue(7)
}
