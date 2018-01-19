package main

/*
Cú pháp pointer: var var_name *var_data_type
*/
import (
	"fmt"
)

func main() {
	//gán giá trị cho message
	message := "Hello World!"
	//tạo ra biến con trỏ cho greeting
	//lúc này greeting mang địa chỉ của biến message
	//và giá trị của greeting mang kiểu string
	//do greeting là kiểu pointer nên muốn lấy giá trị của nó ta phải thêm dấu * thành *greeting
	//do message không phải là kiểu pointer nên để khai báo giá trị ta không cần *
	//cách lấy địa chỉ của một biến không phải kiểu pointer ta chỉ cần thêm dấu &
	var greeting *string = &message
	// thay đổi giá trị cho *greeting tức là thay đổi giá trị của biến message
	*greeting = "How do you do!"
	/*
		nhận xét: sayhi, greeting và &message cùng mang 1 địa chỉ(mà cụ thể ở đây là địa chỉ của message)
		nên khi 1 trong 2 pointer trên (sayhi và greeting) thay đổi giá trị bên trong nó thì giá trị của message
		cũng thay đổi y chang vậy
	*/
	var sayhi *string = &message
	*sayhi = "Hello"

	fmt.Println(message)
	fmt.Println(greeting) //0x10acc0d0
	fmt.Println(*greeting)
	fmt.Println(message)
}
