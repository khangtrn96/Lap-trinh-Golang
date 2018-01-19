package main

// File này nói về Array of Pointer
import (
	"fmt"
)

//Đặt biến MAX là hằng số
const (
	MAX int = 3
)

func main() {
	//Tạo một biến slice a có 3 phần tử
	a := []int{10, 100, 200}
	var i int

	/*Tạo 1 mảng 3 phần tử (thực chất đó là MAX=3) pointer với tên biến là ptr có
	kiểu dữ liệu của pointer là int
	 * Tức từng phần tử trong mảng ptr là 1 cái địa chỉ
	*/
	var ptr [MAX]*int

	/*Dùng vòng lặp để gán địa chỉ của từng phần tử có trong slice a cho biến từng phần tử có trong mảng ptr*/
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /*assign the address of integer for ptr*/
	}

	/*Dùng vòng lặp in ra giá trị của từng phần tử có trong mảng ptr
	 * Lưu ý: Do các phần tử trong mảng ptr là kiểu pointer nên khi muốn sử dụng chúng ở dạng giá trị
	 * trị thì ta dùng dấu *, cụ thể là *ptr[i]
	 */
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}
