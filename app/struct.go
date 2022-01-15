package main 
import "fmt"

type Name struct {
	A string
	B string
}

func main() {
	n1 := Name{"hoge","piyo"}
	n2 := Name{A: "hoge"}
	n3 := Name{}
	p := &Name{"hoge" , "piyo"}
	fmt.Println(n1,n2,n3,p)
}