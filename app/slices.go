package main
import "fmt"

func main(){
	//var fslice []int
	slice := []byte {'a','b','c','d'}
	fmt.Println(slice)

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(ar)

	var aSlice, bSlice []byte

	aSlice = ar[:3] 
	fmt.Println(aSlice)
	aSlice = ar[5:] 
	fmt.Println(aSlice)
	aSlice = ar[:] 
	fmt.Println(aSlice)
	fmt.Println(bSlice)
}

