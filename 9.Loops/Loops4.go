package main

//File này nói về cách sử dụng goto trong vòng lặp Loop
//Có thể hiểu 1 cách đơn giản là cách mà chúng ta chuyển vòng lặp
import (
	"fmt"
)

func main() {
	/*local variable definition*/
	var a int = 10
	/* do loop execution*/
	//LOOP_1 và LOOP_2 đóng vai trò là tên của lable, nó được dùng cho lệnh goto
LOOP_1:
	for a < 20 {
		if a == 15 {
			/* skip the iterntion*/
			a = a + 1

			/*Sau khi so sánh giá trị a đúng bằng 15 thì lập tức cộng thêm 1 đơn vị và chuyển sang LOOP_2*/
			goto LOOP_2
		}
		fmt.Printf("Value of a :%d\n", a)
		a++

	LOOP_2:
		for a%2 == 0 {
			fmt.Printf("Gía trị của a chia hết cho 2 là: %d\n", a)

			/*Sau khi thực hiện xong dòng lệnh Printf ta sẽ chuyển lại vòng lập LOOP_1 để tiếp tục thực hiện */
			goto LOOP_1
		}
	}

}
