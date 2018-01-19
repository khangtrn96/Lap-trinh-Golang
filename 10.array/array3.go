package main

//file này nói về cách sử dụng passing array
import "fmt"

func main() {

	/*an int array with 5 elements*/
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/*pass array as an argument*/
	avg = getAverage(balance, 5)

	/*output the returned value*/
	fmt.Printf("Average value is: %f", avg)
}

/* Tạo hàm để tính giá trị trung bình dựa vào
 * 	mảng arr và biến size
 * Biến size dùng để giá trị cần chia để lấy giá trị trung bình
 */
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	/*
	 *Ta cần có float32 trước (sum/size) để định dạng đúng về kiểu dữ liệu của avg
	 *Do sum và size đều là kiểu int nên khi lấy sum/size thì vẫn ra kiểu int nên khi
	 *gán avg=sum/size sẽ báo lỗi
	 *Ta ghi float32 để biến giá trị sum/size từ kiểu int về kiểu float32
	 */
	avg = float32(sum / size)

	return avg
}
