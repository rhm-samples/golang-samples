package main
import "fmt"

func main() {

	var arr = [5]int {10,20,30,40,50}
	for index := range arr {
	
		fmt.Println(arr[index])
	}
}
