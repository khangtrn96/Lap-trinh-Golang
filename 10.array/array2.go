package main

//File này nói về cách sử dụng mảng đa chiều
import (
	"fmt"
)

func main() {
	//panic("Có lỗi xảy ra")

	/*
	 * Tạo 1 mảng có 4 hàng 3 cột
	 */
	var mang_da_chieu = [4][2]int{{0, 1}, {2, 4}, {3, 6}, {5, 7}}
	var i, j int
	/*
	 *Xuất giá trị của mang_da_chieu
	 */
	for i = 0; i < len(mang_da_chieu); i++ {
		for j = 0; j < len(mang_da_chieu[1]); i++ {
			fmt.Printf("mang_da_chieu[1][%d]=%d", j, mang_da_chieu[1][j])
			fmt.Println("\n")
		}
	}
}
