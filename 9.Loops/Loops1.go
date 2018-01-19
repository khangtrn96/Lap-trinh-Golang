package main

/*File này nói về 3 khai báo vòng lặp for
* Cách 1: for [condition]{}
* Cách 2: for [init; condition: increment]{}
* Cách 3: for [range]
 */
import (
	"fmt"
)

func main() {
	var b int = 15
	var a int

	// Tạo 1 mảng có 6 phần tử, riêng phần tử thứ 5 và 6 thì mang giá trị 0
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution*/

	/*Cách 1 :for [condition]{}*/
	for a < b {
		a++
		fmt.Printf("Value of a : %d\n", a)
	}

	//Cách 2: for [init; condition: increment]{}
	for a = 0; a < 10; a++ {
		fmt.Printf("Value of a:%d\n", a)
	}

	//Cách 3: for [range]
	for i, x := range numbers {
		fmt.Printf("Value of x: %d at %d\n", x, i)
	}
}
