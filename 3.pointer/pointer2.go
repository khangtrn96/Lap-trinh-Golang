package main

//File này nói về khái niệm Nil Pointer in Go ( tức lá địa chỉ của biến pointer là 0 )
import (
	"fmt"
)

func main() {
	//Gọi biến Pointer ptr có kiểu dữ liệu là int
	var ptr *int

	//Nếu lúc đầu ta không nhập giá trị gì cho ptr thì địa chỉ của ptr sẽ là 0
	//Dùng %x để gán giá trị địa chỉ
	fmt.Printf("The value of ptr: %x\n", ptr)

}
