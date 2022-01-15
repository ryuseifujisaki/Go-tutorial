package main
import (
	"fmt"
)

// func main(){
// 	defer fmt.Println("world")
// 	fmt.Println("Hello")
// }

func main() {
	fmt.Println("Countdown...")
	defer fmt.Println("GO!!")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}