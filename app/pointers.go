package main
import (
	"fmt"
)

func main() {
	i := 46
	p := &i //iへのポインタ
	fmt.Println(*p)//ポインタpを通してiの値を読む
	*p =64 // ポインタpを通してiの値変更
	fmt.Println(i)
}