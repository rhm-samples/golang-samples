package main
import "fmt"

func main() {

	var mapNumberToWord = make(map[int]string)
	mapNumberToWord[1]="one"
	mapNumberToWord[2]="two"
	mapNumberToWord[3]="three"

	for key,value := range mapNumberToWord {

		  fmt.Println(key,value)
	}

	//check if key is present
	value,isPresent := mapNumberToWord[4]
	if(isPresent == true) {

		fmt.Println("Word corresponding to 4 is - "+value)
	}else {
		
		fmt.Println("4 is not present")
	}
}
