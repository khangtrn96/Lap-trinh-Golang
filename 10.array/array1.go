package main

//file này nói về cách khai báo mảng
/*cú pháp :
*var variable_name [SIZE] variable_type
 */
/*
* Cách cài đặt 1 mảng
* var variable_name=[SIZE] variable_type{số lương phần tử phụ thuộc vào kích thước SIZE }
* hoặc
* var variable_name=[] variable_type{số lương phần tử không cần quan tâm đến kích thước SIZE}
*Lý do: vì SIZE để trống
 */

import "fmt"

func main() {
	//Cách 1: khai báo mảng
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	// cách 2: khai báo mảng
	var phan_tu = [7]int{1, 4, 5, 9, 6, 7, 8}

	//thực hiện tính trung bình của mảng x
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total = total + x[i]
	}
	//Kết quả trung bình của mảng x
	fmt.Println(total / 5)
	fmt.Printf("\n")
	//in từng giá trị có trong mảng phan_tu
	for j := 0; j < 7; j++ {
		fmt.Printf("Gía trị phan_tu[%d] là %d", j, phan_tu[j])
		fmt.Printf("\n")
	}
}
