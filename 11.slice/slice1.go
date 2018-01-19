package main

//File này nó về cách định nghĩa 1 slice là như thế nào
/*
* Bản chất 1 slice cũng như giống như 1 array nhưng bỏ trống giá trị SIZE
* Cụ thể:
* khai báo 1 array : variable_name=[SIZE]variable_type{giá trị phần tử}
* khai báo 1 slice(cách 1): variable_name= []variable_type{giá trị phần tử}
* khai báo 1 slice(cách 2): variable_name= make([]variable_type, len, cap)
* với len cho biết độ dài của slice, cap cho biết sức chứa của slice
 */
import (
	"fmt"
)

/* Khai báo hàm để xuất giá trị len, cap và slice*/
func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
func main() {
	// Tạo 1 slice có kiểu dữ liệu int, có độ dài lả 3, có sức chứa là 5
	var numbers = make([]int, 3, 5)

	/* In ra slice ta dùng hàm printSlice*/
	printSlice(numbers)
}
