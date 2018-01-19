package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		//Khai báo từ khóa go để chạy hàm f()
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
