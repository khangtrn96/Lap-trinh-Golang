package main

// dùng hàm hằng giúp ta gắn giá trị "cố định" cho biến
// và biến này không thể thay đổi giá trị được
import (
	"fmt"
)

//Khai bao const kiểu 0
const (
	PI       = 3.14
	Language = "Golang"
)

// Khai bao const kieu 1
// có thể ghi thêm kiểu dữ liệu "int" nhưng điều này là không cần thiết
const (
	studentcode int = 1511130
)

func main() {
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(studentcode)
}
