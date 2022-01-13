package main

import "fmt"

// func add(x int, y int) int {
// 	return x+y
// }

func add(x,y int) (sum int) {
	sum = x+y
	return
}

func swap(x,y int)(int,int){
	return y,x
}

// func main() {
//     fmt.Println(add(46, 9))
// }

func main() {
		fmt.Println(add(46,9))
}