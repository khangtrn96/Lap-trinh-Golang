package main

/*file  này giới thiệu về cách sử dụng gorountines để chạy concurrently(tức chạy nhiều hàm cùng lúc)
* Bản chất gorountine cũng là 1 hàm
* Bằng cách sử dụng từ khóa go dê chạy 1 gorountines
 */
import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	//Khai báo từ khóa go để chạy hàm f()
	go f(4)
	var input string
	fmt.Scanln(&input)
}
