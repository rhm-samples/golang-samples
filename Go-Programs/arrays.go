package main
import "fmt"

func main() {

	var arr = [5]int {1,2,3,4,5}
	var sum int = 0
	for i:=0;i<len(arr);i++ {
		
		  sum = sum + arr[i]
	}
	fmt.Println(sum)
}
