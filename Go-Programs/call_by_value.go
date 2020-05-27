package main
import "fmt"

func increment(a int) {
	
    a++
}

func main() {

    a := 5
    increment(a)
    fmt.Println(a)
}
