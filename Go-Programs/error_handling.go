package main

import "fmt"
import "errors"

func divideNumber(number1 int,number2 int) (int,error) {

	if number2 == 0 {

		return 0,errors.New("The divisor is zero!!")
	}
	return number1/number2,nil
}

func main() {

	number1 := 20
	number2 := 0
	answer,err := divideNumber(number1,number2)

	if err!=nil {

		fmt.Println(err)
	}else{
	
		fmt.Println(answer)
	}
}
