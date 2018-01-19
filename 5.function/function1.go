package main

import (
	"fmt"
)

// type Salutation struct {
// 	name     string
// 	greeting string
// }

// func info(value Salutation) {
// 	fmt.Println(value.name)
// 	fmt.Println(value.greeting)
// }
// func main() {
// 	var infoinvidual = Salutation{"Hi", "Bob"}
// 	info(infoinvidual)
// }
type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}
func main() {
	var s = Salutation{"Bob", "Hi"} // ở đây phải ghi giá trị trong dấu ngoặc nhọn {}
	Greet(s)
}
