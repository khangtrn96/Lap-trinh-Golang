package main

import (
	"fmt"
	//Khai báo thư viện math để sử dụng các hàm toán có sẵn trong ngôn ngữ go
	"math"
)

/*
define a circle
*/
type Circle struct {
	x, y, radius float64
}

/*
define a method for circle
Khi khai báo thư viện math ta mới có thể sử dụng được giá trị Pi trong toán thông qua biến math.Pi
*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func main() {
	hinhtron := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area:%f", hinhtron.area())
}
