package main
import "fmt"

func main() {

    var arr = [5]int {1,2,3,4,5}
    for index := range arr {
	
	if(index == 3) {
			
            break
	}
	fmt.Println(arr[index])
    }
}
