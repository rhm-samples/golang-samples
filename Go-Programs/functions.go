package main
import "fmt"

func solve(number1 int,number2 int) (int,int) {
    
    sum := number1 + number2
    difference := number1 - number2
    return sum,difference
}

func add(number1 int,number2 int) {

    result := number1 + number2
    fmt.Printf("Sum of 2 numbers in add() function = %d\n",result)
}

func main() {

    a := 10
    b := 20
    add(a,b)

    sum,difference := solve(a,b)
    fmt.Printf("The sum of 2 numbers = %d\n",sum)
    fmt.Printf("The difference between 2 numbers = %d\n",difference)
}
