package main

//cách khai báo type cũng là cách mình khai báo 1 kiểu dữ liệu mới(vd int, string).
//Nhưng khai báo type giúp ta gắn thêm điều kiện cho biến type đó
import (
	"fmt"
)

// Khai báo type loại 1
type Salutation string

//Khai báo type loại 2
type Salutation2 struct {
	name     string //điều kiện
	greeting string //điều kiện
}
type Salutation3 struct {
	studentcode int    //điều kiện
	studentname string //điều kiện
}
type Salutation4 struct {
	nation  string //điều kiện
	capital string //điều kiện
}
type Salutation5 struct {
	subject string //điều kiện
	lesson  string //điều kiện
}

func main() {
	//cách chạy code loại 0
	//Nếu không gán giá trị gì cho biến trong {} thì sẽ xuất ra giá trị 0
	var message0 = Salutation3{}
	fmt.Println(message0.studentcode)
	fmt.Println(message0.studentname)
	//cách chạy type kiểu 1

	var message Salutation = "Hello World!"
	// thay vì ghi string cho biến message ta ghi Salutation,
	// nhưng trong trường hợp khác, Salutation được định nghĩa với nhiều biến hơn thì Salutatuion không thể thay thế cho String
	fmt.Println(message)
	//cách chạy code kiểu 2

	var message2 = Salutation2{"Hi", "Everybody"}
	fmt.Println(message2.name)
	fmt.Println(message2.greeting)

	var message3 = Salutation3{1511130, "Khang"}
	fmt.Println(message3.studentcode)
	fmt.Println(message3.studentname)
	//cách chạy code kiểu 3
	//nhớ la phải dùng dấu ":" khi gán cho giá trị cho biến "nation" va "capital"

	var message4 = Salutation4{nation: "Viet Nam", capital: "Ha Noi"}
	fmt.Println(message4.nation)
	fmt.Println(message4.capital)
	// cách chạy code kiểu 4
	// khác với cách chay code kiểu 0, tuy ta để rỗng trong{}, nhưng vẫn có thể gán giá trị cho biến ở bên ngoai Salutation5{}
	var message5 = Salutation5{}
	message5.subject = "Math"
	message5.lesson = "linear"
	fmt.Println(message5.subject)
	fmt.Println(message5.lesson)
}
