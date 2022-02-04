package main
import "fmt"
// func main(){
// 	var i [2]int
// 	i[0] = 1
// 	i[1] = 2
// 	fmt.Println(i)
// }

func main() {
	var arr [10] int
	arr[0] = 42
	arr[1] = 13
	fmt.Println("The first element is %d\n" , arr[0])
	fmt.Println("The last element is %d\n" , arr[9])
	sub()
	doublearray()
}

func sub() {
	a := [3]int{1,2,3}
	b := [10]int{1,2,3}
	c := [...]int{4,5,6}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func doublearray() {
	doublearray := [2][4]int{[4]int{1,2,3,4},[4]int{5,7,6,8}}
	easyArray := [2][4]int{{1,2,4,5},{5,6,7,8}}
	fmt.Println(doublearray)
	fmt.Println(easyArray)
}